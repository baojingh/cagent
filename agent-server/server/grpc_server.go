package server

import (
	logger "agent-server/log"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var log = logger.New()

// grpc server entry for running
func StartGrpcServer(ip string, port int, tlsEnabled bool) {
	log.Info("GRPC Server is starting")
	ipPort := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", ipPort)
	if err != nil {
		log.Error("Failed to start GRPC Server, ", err)
		return
	}
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	log.Infof("GPRC Server starts in port %d", port)
	server.Serve(listener)
}
