package server

import (
	"agentctl/pb"
	"fmt"

	"github.com/shimingyah/pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
	Create grpc connection for all commands
*/

func init() {
	p, err := pool.New("127.0.0.1:8080", pool.DefaultOptions)
	if err != nil {
		log.Fatalf("failed to new pool: %v", err)
	}
	defer p.Close()

	conn, err := p.Get()
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	defer conn.Close()

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Infof("Connect to server %s %s success.", ip, port)
	cli.client = pb.NewAgentFileServiceClient(conn)
}
