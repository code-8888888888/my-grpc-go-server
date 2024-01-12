package main

import (
	"log"

	mygrpc "github.com/code-8888888888/my-grpc-go-server/internal/adapter/grpc"
	app "github.com/code-8888888888/my-grpc-go-server/internal/application"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(log.Writer())

	hs := &app.HelloService{}

	grpcAdapter := mygrpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}