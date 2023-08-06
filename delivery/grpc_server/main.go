package grpc_server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/hengkyawijaya/simple-go/delivery/grpc_server/client"
	"github.com/hengkyawijaya/simple-go/repository"
	"github.com/hengkyawijaya/simple-go/usecase"
	"google.golang.org/grpc"
)

// grpcServer is a struct that contains all dependencies for grpc server
type grpcServer struct {
	Uc *usecase.Usecase
}

// NewGRPCServer is a constructor for grpc server
func NewGRPCServer(uc *usecase.Usecase) *grpcServer {
	return &grpcServer{Uc: uc}
}

// DeliverGRPCServer is a function to deliver grpc server
func DeliverGRPCServer(
	uc *usecase.Usecase,
	wg sync.WaitGroup,
	repo *repository.Repository,
) {

	cfg := repo.ConfigRepository.ReadConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServer.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	client.RegisterGreeterServer(srv, NewGRPCServer(uc))

	log.Printf("Starting gRPC server on port %s...\n", cfg.GRPCServer.Port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	wg.Done()
}
