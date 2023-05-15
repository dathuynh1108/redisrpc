package redisrpc

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/dathuynh1108/redisrpc/rpc"

	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// redefine grpc.serverMethodHandler as it is not exposed
type serverMethodHandler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error)

type serverTransportStream struct {
	stream *serverStream
}

func (s *serverTransportStream) Method() string {
	return s.stream.method
}
func (s *serverTransportStream) SetHeader(md metadata.MD) error {
	return s.stream.SetHeader(md)
}

func (s *serverTransportStream) SendHeader(md metadata.MD) error {
	return s.stream.SendHeader(md)
}

func (s *serverTransportStream) SetTrailer(md metadata.MD) error {
	s.stream.SetTrailer(md)
	return nil
}

func serverUnaryHandler(srv interface{}, handler serverMethodHandler) handlerFunc {
	return func(s *serverStream) {
		var interceptor grpc.UnaryServerInterceptor = nil
		ctx := grpc.NewContextWithServerTransportStream(s.Context(), &serverTransportStream{stream: s})
		if s.md != nil {
			ctx = metadata.NewIncomingContext(ctx, s.md)
		}
		response, err := handler(srv, ctx, s.RecvMsg, interceptor)
		if s.ctx.Err() == nil {
			if err != nil {
				s.close(err)
				return
			}
			if s.SendMsg(response) == nil {
				s.close(err)
			}
		}
	}
}

func serverStreamHandler(srv interface{}, handler grpc.StreamHandler) handlerFunc {
	return func(s *serverStream) {
		err := handler(srv, s)
		if s.ctx.Err() == nil {
			s.close(err)
		}
	}
}

type handlerFunc func(s *serverStream)

// serviceInfo wraps information about a service. It is very similar to
// ServiceDesc and is constructed from it for internal purposes.
type serviceInfo struct {
	// Contains the implementation for the methods in this service.
	serviceImpl interface{}
	methods     map[string]*grpc.MethodDesc
	streams     map[string]*grpc.StreamDesc
	mdata       interface{}
}

// Server is the interface to gRPC over NATS
type Server struct {
	redis    *redis.Client
	ctx      context.Context
	cancel   context.CancelFunc
	log      *logrus.Logger
	handlers map[string]handlerFunc
	streams  map[string]*serverStream
	mu       sync.Mutex
	subs     map[string]*redis.PubSub
	nid      string
	services map[string]*serviceInfo // service name -> service info
}

// NewServer creates a new Proxy
func NewServer(r *redis.Client, nid string) *Server {
	s := &Server{
		redis:    r,
		handlers: make(map[string]handlerFunc),
		streams:  make(map[string]*serverStream),
		subs:     make(map[string]*redis.PubSub),
		services: make(map[string]*serviceInfo),
		log:      logrus.New(),
		nid:      nid,
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	return s
}

// Stop gracefully stops a Proxy
func (s *Server) Stop() {
	s.cancel()
	for name, sub := range s.subs {
		err := sub.Unsubscribe(s.ctx, name)
		if err != nil {
			s.log.Errorf("Unsubscribe [%v] failed %v", name, err)
		}
	}
}

func (s *Server) CloseStream(nid string) error {
	for name, st := range s.streams {
		if st.pnid == nid {
			st.done()
			s.log.Infof("CloseStream nid = %v, name = %v", nid, name)
		}
	}
	return nil
}

// RegisterService is used to register gRPC services
func (s *Server) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	prefix := fmt.Sprintf("redisrpc.%v", sd.ServiceName)
	if len(s.nid) > 0 {
		prefix = fmt.Sprintf("redisrpc.%v.%v", s.nid, sd.ServiceName)
	}
	subject := prefix
	s.log.Infof("Subscribe: subject => %v", subject)
	sub := s.redis.Subscribe(s.ctx, subject)
	ch := sub.Channel()
	go func() {
		for msg := range ch {
			s.log.Debugf("Received message: %v", msg)
			s.onMessage(msg)
		}
	}()
	s.subs[sd.ServiceName] = sub
	for _, it := range sd.Methods {
		desc := it
		path := fmt.Sprintf("%v.%v", prefix, desc.MethodName)
		s.handlers[path] = serverUnaryHandler(ss, serverMethodHandler(desc.Handler))
		s.log.Infof("RegisterService: method path => %v", path)
	}
	for _, it := range sd.Streams {
		desc := it
		path := fmt.Sprintf("%v.%v", prefix, desc.StreamName)
		s.handlers[path] = serverStreamHandler(ss, desc.Handler)
		s.log.Infof("RegisterService: stream path => %v", path)
	}

	s.register(sd, ss)
}

func (s *Server) register(sd *grpc.ServiceDesc, ss interface{}) {
	s.log.Infof("RegisterService(%q)", sd.ServiceName)

	if _, ok := s.services[sd.ServiceName]; ok {
		s.log.Fatalf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
	}
	info := &serviceInfo{
		serviceImpl: ss,
		methods:     make(map[string]*grpc.MethodDesc),
		streams:     make(map[string]*grpc.StreamDesc),
		mdata:       sd.Metadata,
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		info.methods[d.MethodName] = d
	}
	for i := range sd.Streams {
		d := &sd.Streams[i]
		info.streams[d.StreamName] = d
	}
	s.services[sd.ServiceName] = info
}

func (s *Server) GetServiceInfo() map[string]grpc.ServiceInfo {
	ret := make(map[string]grpc.ServiceInfo)
	for n, srv := range s.services {
		methods := make([]grpc.MethodInfo, 0, len(srv.methods)+len(srv.streams))
		for m := range srv.methods {
			methods = append(methods, grpc.MethodInfo{
				Name:           m,
				IsClientStream: false,
				IsServerStream: false,
			})
		}
		for m, d := range srv.streams {
			methods = append(methods, grpc.MethodInfo{
				Name:           m,
				IsClientStream: d.ClientStreams,
				IsServerStream: d.ServerStreams,
			})
		}

		ret[n] = grpc.ServiceInfo{
			Methods:  methods,
			Metadata: srv.mdata,
		}
	}
	return ret
}

func (s *Server) onMessage(msg *redis.Message) {
	//p.log.Infof("Proxy.onMessage: subject %v, replay %v, data %v", msg.Subject, msg.Reply, string(msg.Data))
	request := &rpc.Request{}
	err := proto.Unmarshal([]byte(msg.Payload), request)
	if err != nil {
		s.log.WithField("data", string(msg.Payload)).Error("unknown message")
	}

	method := request.Method
	reply := request.Reply

	log := s.log.WithField("method", method)

	s.mu.Lock()
	stream, ok := s.streams[reply]
	if !ok {
		stream = newServerStream(s, method, reply, log)
		s.streams[reply] = stream
	}
	s.mu.Unlock()
	go stream.onRequest(msg, request)
}

func (s *Server) remove(reply string) {
	s.mu.Lock()
	delete(s.streams, reply)
	s.mu.Unlock()
}

var (
	// https://github.com/grpc/grpc-go/blob/master/internal/transport/http2_server.go#L54

	// ErrIllegalHeaderWrite indicates that setting header is illegal because of
	// the stream's state.
	ErrIllegalHeaderWrite = errors.New("transport: the stream is done or WriteHeader was already called")
)

type serverStream struct {
	ctx       context.Context
	cancel    context.CancelFunc
	server    *Server
	log       *logrus.Entry
	recvRead  <-chan []byte
	recvWrite chan<- []byte
	hasBegun  bool
	md        metadata.MD // recevied metadata from client
	header    metadata.MD // send header to client
	trailer   metadata.MD // send trialer to client
	method    string
	reply     string
	pnid      string

	messageMutex sync.Mutex
}

func newServerStream(server *Server, method, reply string, log *logrus.Entry) *serverStream {
	s := &serverStream{
		server: server,
		log:    log,
		method: method,
		reply:  reply,
	}
	s.ctx, s.cancel = context.WithCancel(server.ctx)
	recv := make(chan []byte, 1)
	s.recvRead = recv
	s.recvWrite = recv
	return s
}

func (s *serverStream) done() {
	s.cancel()
	s.server.remove(s.reply)
}

func (s *serverStream) onRequest(msg *redis.Message, request *rpc.Request) {
	switch r := request.Type.(type) {
	case *rpc.Request_Call:
		//s.log.WithField("call", r.Call).Info("recv call")
		s.processCall(r.Call)
	case *rpc.Request_Data:
		//s.log.WithField("data", r.Data).Info("recv data")
		s.processData(r.Data)
	case *rpc.Request_End:
		//s.log.WithField("end", r.End).Info("recv end")
		s.processEnd(r.End)
	}
}

func (s *serverStream) processCall(call *rpc.Call) {
	s.log = s.log.WithField("method", s.method)
	handlerFunc, ok := s.server.handlers[s.method]
	if !ok {
		s.close(status.Error(codes.Unimplemented, codes.Unimplemented.String()))
		return
	}
	// save metadata to context
	if call.Metadata != nil {
		md := make(metadata.MD)
		for hdr, data := range call.Metadata.Md {
			md[hdr] = data.Values
		}
		if s.md == nil {
			s.md = md
		} else if md != nil {
			s.md = metadata.Join(s.md, md)
		}
	}
	s.pnid = call.Nid
	go handlerFunc(s)
}

func (s *serverStream) processData(data *rpc.Data) {
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()
	if s.recvWrite == nil {
		s.log.Error("data received after client closeSend")
		return
	}
	s.recvWrite <- data.Data
}

func (s *serverStream) processEnd(end *rpc.End) {
	s.log.Infof("Received End from client: %v")
	if end.Status != nil {
		s.log.WithField("status", end.Status).Info("cancel")
		s.done()
	} else {

		s.messageMutex.Lock()
		defer s.messageMutex.Unlock()

		s.log.Debug("closeSend")
		s.recvWrite <- nil
		close(s.recvWrite)
		s.recvWrite = nil
	}
}

func (s *serverStream) beginMaybe() error {
	if !s.hasBegun {
		s.hasBegun = true
		if s.header != nil {
			return s.writeBegin(&rpc.Begin{
				Header: MakeMetadata(s.header),
				Nid:    s.server.nid,
			})
		}
	}
	return nil
}

func (s *serverStream) onMessage(msg *redis.Message) {
	request := &rpc.Request{}
	err := proto.Unmarshal([]byte(msg.Payload), request)
	if err != nil {
		s.log.WithField("data", string(msg.Payload)).Error("unknown message")
	}

	go s.onRequest(msg, request)
}

func (s *serverStream) close(err error) {
	s.beginMaybe()
	s.writeEnd(&rpc.End{
		Status:  status.Convert(err).Proto(),
		Trailer: MakeMetadata(s.trailer),
	})
	s.done()
}

// Server Stream interface
func (s *serverStream) Method() string {
	return s.method
}

func (s *serverStream) SetHeader(header metadata.MD) error {
	if s.hasBegun {
		return ErrIllegalHeaderWrite
	}
	if s.header == nil {
		s.header = header
	} else if header != nil {
		s.header = metadata.Join(s.header, header)
	}
	return nil
}

func (s *serverStream) SendHeader(header metadata.MD) error {
	err := s.SetHeader(header)
	if err != nil {
		return err
	}
	return s.beginMaybe()
}

func (s *serverStream) SetTrailer(trailer metadata.MD) {
	if s.trailer == nil {
		s.trailer = trailer
	} else if trailer != nil {
		s.trailer = metadata.Join(s.trailer, trailer)
	}
}

func (s *serverStream) Context() context.Context {
	return s.ctx
}

func (s *serverStream) SendMsg(m interface{}) (err error) {
	defer func() {
		if err != nil {
			s.close(err)
		}
	}()

	err = s.beginMaybe()
	if err == nil {
		data, err := proto.Marshal(m.(proto.Message))
		if err == nil {
			s.writeData(&rpc.Data{
				Data: data,
			})
		}
	}
	return
}

func (s *serverStream) RecvMsg(m interface{}) error {
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	case bytes, ok := <-s.recvRead:
		if ok && bytes != nil {
			return proto.Unmarshal(bytes, m.(proto.Message))
		}
		return io.EOF
	}
}

func (s *serverStream) writeResponse(response *rpc.Response) error {
	//s.log.WithField("response", response).Info("send")
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	return s.server.redis.Publish(s.ctx, s.reply, data).Err()
}

func (s *serverStream) writeBegin(begin *rpc.Begin) error {
	return s.writeResponse(&rpc.Response{
		Type: &rpc.Response_Begin{
			Begin: begin,
		},
	})
}

func (s *serverStream) writeData(data *rpc.Data) error {
	return s.writeResponse(&rpc.Response{
		Type: &rpc.Response_Data{
			Data: data,
		},
	})
}

func (s *serverStream) writeEnd(end *rpc.End) error {
	return s.writeResponse(&rpc.Response{
		Type: &rpc.Response_End{
			End: end,
		},
	})
}
