package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/code-8888888888/my-grpc-go-server/internal/application/port"
	"github.com/code-8888888888/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort int
	server *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort: grpcPort,
	}
}

func unaryInterceptor(
	ctx context.Context,
	req interface{}, 
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("unary interceptor called:", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
	) error {
	log.Println("stream interceptor called:", info.FullMethod)
	return handler(srv, stream)
}

func(a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to listen on port %d: %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) stop() {
	a.server.Stop()
}