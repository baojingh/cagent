package main

import (
	"agentctl/helloworld"
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

/**
  Author     : He Bao Jing
  Date       : 10/20/2023 5:06 PM
  Description:
*/

var (
	port = flag.Int("port", 50051, "Server Port")
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Printf("Receive: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	fmt.Printf("Server listening at %v", lis.Addr())
	s.Serve(lis)

}
