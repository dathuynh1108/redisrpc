package redisrpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"redisrpc/rpc"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	redis   *redis.Client
	ctx     context.Context
	cancel  context.CancelFunc
	log     *logrus.Logger
	streams map[string]*clientStream
	svcid   string
	nid     string
	mu      sync.Mutex
}

func NewClient(redis *redis.Client, svcid string, nid string) *Client {
	c := &Client{
		redis:   redis,
		svcid:   svcid,
		nid:     nid,
		streams: make(map[string]*clientStream),
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (p *Client) GetServiceId() string {
	return p.svcid
}

// Close gracefully stops a Client
func (p *Client) Close() error {
	p.cancel()
	for name, st := range p.streams {
		err := st.done()
		if err != nil {
			p.log.Errorf("Unsubscribe [%v] failed %v", name, err)
			return err
		}
	}
	return nil
}

func (p *Client) CloseStream(nid string) bool {
	if p.svcid == nid {
		p.Close()
		return true
	}
	return false
}

func (c *Client) remove(subj string) {
	c.mu.Lock()
	delete(c.streams, subj)
	c.mu.Unlock()
}

// Invoke performs a unary RPC and returns after the request is received
// into reply.
func (c *Client) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	prefix := "redisrpc"
	if len(c.svcid) > 0 {
		prefix = fmt.Sprintf("redisrpc.%v", c.svcid)
	}
	subj := prefix + strings.ReplaceAll(method, "/", ".")
	stream := newClientStream(ctx, c, subj, c.log, opts...)
	c.mu.Lock()
	c.streams[stream.reply] = stream
	c.mu.Unlock()
	return stream.Invoke(ctx, method, args, reply, opts...)
}

// NewStream begins a streaming RPC.
func (c *Client) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	prefix := "redisrpc"
	if len(c.svcid) > 0 {
		prefix = fmt.Sprintf("redisrpc.%v", c.svcid)
	}
	subj := prefix + strings.ReplaceAll(method, "/", ".")
	stream := newClientStream(ctx, c, subj, c.log, opts...)
	c.mu.Lock()
	c.streams[stream.reply] = stream
	c.mu.Unlock()
	return stream, nil
}

type clientStream struct {
	md      *metadata.MD
	header  *metadata.MD
	trailer *metadata.MD
	lastErr error
	ctx     context.Context
	cancel  context.CancelFunc
	log     *logrus.Logger
	client  *Client
	subject string
	reply   string

	msgCh <-chan *redis.Message
	sub   *redis.PubSub

	closed    bool
	recvRead  <-chan []byte
	recvWrite chan<- []byte
	hasBegun  bool
	pnid      string
}

func newClientStream(ctx context.Context, client *Client, subj string, log *logrus.Logger, opts ...grpc.CallOption) *clientStream {
	stream := &clientStream{
		client:  client,
		log:     log,
		subject: subj,
		reply:   uuid.NewString(),
		closed:  false,
	}
	stream.ctx, stream.cancel = context.WithCancel(ctx)

	recv := make(chan []byte, 1)
	stream.recvRead = recv
	stream.recvWrite = recv

	// stream.msgCh = make(chan *nats.Msg, 8192)
	// stream.sub, _ = client.nc.ChanSubscribe(stream.reply, stream.msgCh)

	stream.sub = client.redis.Subscribe(client.ctx, stream.reply)
	stream.msgCh = stream.sub.Channel(redis.WithChannelSize(1024))
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		//log.Printf("stream outgoing md => %v", md)
		stream.md = &md
	}

	for _, o := range opts {
		switch o := o.(type) {
		case grpc.HeaderCallOption:
			//log.Printf("o.HeaderAddr => %v", o.HeaderAddr)
			stream.header = o.HeaderAddr
		case grpc.TrailerCallOption:
			//log.Printf("o.TrailerAddr => %v", o.TrailerAddr)
			stream.trailer = o.TrailerAddr
		case grpc.PeerCallOption:
		case grpc.PerRPCCredsCallOption:
		case grpc.FailFastCallOption:
		case grpc.MaxRecvMsgSizeCallOption:
		case grpc.MaxSendMsgSizeCallOption:
		case grpc.CompressorCallOption:
		case grpc.ContentSubtypeCallOption:
		}
	}

	go stream.ReadMsg()
	return stream
}

func (c *clientStream) Header() (metadata.MD, error) {
	if c.header == nil {
		c.header = &metadata.MD{}
	}
	return *c.header, nil
}

func (c *clientStream) Trailer() metadata.MD {
	if c.trailer == nil {
		c.trailer = &metadata.MD{}
	}
	return *c.trailer
}

func (c *clientStream) CloseSend() error {
	// c.log.Infof("Client CloseSend %s", c.subject)
	c.writeEnd(&rpc.End{
		Status: status.Convert(nil).Proto(),
	})
	return c.done()
}

func (c *clientStream) close(err error) {
	c.writeEnd(&rpc.End{
		Status: status.Convert(err).Proto(),
	})
	c.done()
}

func (c *clientStream) Context() context.Context {
	return c.ctx
}

func (c *clientStream) onMessage(msg *redis.Message) error {
	response := &rpc.Response{}
	err := proto.Unmarshal([]byte(msg.Payload), response)
	if err != nil {
		c.log.WithField("data", msg.Payload).Error("unknown message")
		return err
	}

	switch r := response.Type.(type) {
	case *rpc.Response_Begin:
		//c.log.WithField("call", r.Begin).Info("recv call")
		c.processBegin(r.Begin)
	case *rpc.Response_Data:
		//c.log.WithField("data", r.Data).Info("recv data")
		c.processData(r.Data)
	case *rpc.Response_End:
		//c.log.WithField("end", r.End).Info("recv end")
		return c.processEnd(r.End)
	}
	return nil
}

func (c *clientStream) ReadMsg() error {
	for {
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		case msg, ok := <-c.msgCh:
			if ok {
				err := c.onMessage(msg)
				if err != nil {
					c.lastErr = err
					return err
				}
				break
			}
			return io.EOF
		}
	}
}

func (c *clientStream) done() error {
	if !c.closed {
		c.closed = true
		c.cancel()
		err := c.sub.Unsubscribe(c.ctx, c.reply)
		// close(c.msgCh)
		c.client.remove(c.subject)
		return err
	}
	return errors.New("Client Streaming already closed")
}

func (c *clientStream) SendMsg(m interface{}) error {
	if c.closed {
		return fmt.Errorf("client streaming closed=true")
	}

	if !c.hasBegun {
		c.hasBegun = true
		call := &rpc.Call{
			Method: c.subject,
			Nid:    c.client.nid,
		}
		if c.md != nil {
			call.Metadata = MakeMetadata(*c.md)
		}
		//write call with metatdata
		c.writeCall(call)
	}

	var data *rpc.Data
	if frame, ok := m.(*Frame); ok {
		data = &rpc.Data{
			Data: frame.Payload,
		}
	} else {
		payload, err := proto.Marshal(m.(proto.Message))
		if err != nil {
			c.log.Errorf("clientStream.SendMsg failed: %v", err)
			return err
		}
		data = &rpc.Data{
			Data: payload,
		}
	}
	//write grpc args
	return c.writeData(data)
}

func (c *clientStream) RecvMsg(m interface{}) error {
	select {
	case <-c.ctx.Done():
		if c.lastErr != nil {
			return c.lastErr
		}
		return c.ctx.Err()
	case bytes, ok := <-c.recvRead:
		if ok && bytes != nil {
			if frame, ok := m.(*Frame); ok {
				frame.Payload = bytes
				return nil
			} else {
				return proto.Unmarshal(bytes, m.(proto.Message))
			}
		}
		return io.EOF
	}
}

func (c *clientStream) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	payload, err := proto.Marshal(args.(proto.Message))
	if err != nil {
		c.log.Fatalf("%v for request", err)
		return err
	}

	//write call with metatdata
	call := &rpc.Call{
		Method: method,
	}
	if c.md != nil {
		call.Metadata = MakeMetadata(*c.md)
	}

	c.writeCall(call)

	//write grpc args
	c.writeData(&rpc.Data{
		Data: payload,
	})

	err = c.RecvMsg(reply)

	if err != nil {
		c.log.Debugf("%v for c.RecvMsg %s", err, c.subject)
	}

	c.CloseSend()

	return err
}

func (c *clientStream) writeRequest(request *rpc.Request) error {
	// c.log.WithField("request", request).Debug("send")
	data, err := proto.Marshal(request)
	if err != nil {
		return err
	}
	return c.client.redis.Publish(c.ctx, c.subject, data).Err()
}

func (c *clientStream) writeCall(call *rpc.Call) error {
	return c.writeRequest(&rpc.Request{
		Type: &rpc.Request_Call{
			Call: call,
		},
	})
}

func (c *clientStream) writeData(data *rpc.Data) error {
	return c.writeRequest(&rpc.Request{
		Type: &rpc.Request_Data{
			Data: data,
		},
	})
}

func (c *clientStream) writeEnd(end *rpc.End) error {
	return c.writeRequest(&rpc.Request{
		Type: &rpc.Request_End{
			End: end,
		},
	})
}

func (c *clientStream) processBegin(begin *rpc.Begin) error {
	c.log.Debugf("rpc.Begin: %v", begin.Header)
	if begin.Header != nil {
		if c.header == nil {
			c.header = &metadata.MD{}
		}
		for hdr, data := range begin.Header.Md {
			c.header.Append(hdr, data.Values...)
		}
	}
	c.pnid = begin.Nid
	return nil
}

func (c *clientStream) processData(data *rpc.Data) {
	if c.recvWrite == nil {
		c.log.Error("data received after client closeSend")
		return
	}
	c.recvWrite <- data.Data
}

func (c *clientStream) processEnd(end *rpc.End) error {

	if end.Trailer != nil && c.trailer != nil {
		if c.trailer == nil {
			c.trailer = &metadata.MD{}
		}
		for hdr, data := range end.Trailer.Md {
			c.trailer.Append(hdr, data.Values...)
		}
	}

	if end.Status != nil {
		c.log.WithField("status", end.Status).Info("cancel")
		c.done()
		return status.Error(codes.Code(end.Status.Code), end.Status.GetMessage())
	}
	// c.log.Infof("Server CloseSend %s", c.subject)
	c.recvWrite <- nil
	close(c.recvWrite)
	c.recvWrite = nil
	return nil
}
