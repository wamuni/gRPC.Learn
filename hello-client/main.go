package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "wamuni.cc/grpc/implementation/hello-server/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.0:9090", grpc.WithTransportCredentials(insecure.NewBundle().TransportCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSayHelloClient(conn)
	message, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "Eddie"})
	fmt.Printf("Message from gRPC server: %v", message)
}
