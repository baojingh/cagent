package main

import (
	"agent-server/pb"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/**
  Author     : He Bao Jing
  Date       : 10/20/2023 5:06 PM
  Description:
*/

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAgentActionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	paramMap := make(map[string]string, 10)
	paramMap["host"] = "99.99.99"

	req := &pb.AgentServiceRequest{HopstIP: "11", Component: "vws", ParamMap: paramMap}
	r, err := c.UpdateFluentbitHost(ctx, req)
	if err != nil {
		log.Fatalf("%v, %v", err, r)
	}
	log.Printf("success")
}
