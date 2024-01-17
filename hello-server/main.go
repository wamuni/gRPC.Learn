package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "wamuni.cc/grpc/implementation/hello-server/proto"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{ResponseMsg: "Hello" + req.RequestName}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":9090")
	grpcServer := grpc.NewServer()
	pb.RegisterSayHelloServer(grpcServer, &server{})
	err := grpcServer.Serve(listen)

	if err != nil {
		fmt.Printf("failed to server: %v", err)
	}
}
