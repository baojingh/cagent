package grpc_conn

import (
	logger "agentctl/log"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var log = logger.New()

func GetGRPCConn(ip string, port string) (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	addr := fmt.Sprintf("%s:%s", ip, port)
	tlsConfig, err := GetTLSConfig()
	if err != nil {
		log.Errorf("Failed to laod TLS config, %v", err)
		return nil, err
	}
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Infof("Connect to server %s:%s success.", ip, port)
	return conn, nil
}
