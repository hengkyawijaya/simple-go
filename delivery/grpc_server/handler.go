package grpc_server

import (
	"context"

	"github.com/hengkyawijaya/simple-go/delivery/grpc_server/client"
)

// Hello is a function to say hello
func (g *grpcServer) Hello(context context.Context, request *client.HelloRequest) (*client.HelloReply, error) {
	return &client.HelloReply{Message: g.Uc.HelloUsecase.Hello()}, nil
}
