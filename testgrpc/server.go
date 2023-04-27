package testgrpc

import (
	context "context"
	"fmt"
)

type Server struct {
	UnimplementedTestServerServer
}

func (s *Server) MakeRequest(ctx context.Context, request *Request) (*Response, error) {
	fmt.Println("Message:", request.Message)
	return &Response{
		Message: "Hello, " + request.Message,
	}, nil
}
