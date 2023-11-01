package server

import (
	"agent-server/config"
	logger "agent-server/log"
	"agent-server/pb"
	"agent-server/server/fluentbit"
	"agent-server/server/suricata"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var log = logger.New()
var grpcServer *grpc.Server

// grpc server entry for running
func StartGrpcServer() {
	ip := config.Ip
	port := config.Port
	ipPort := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", ipPort)
	if err != nil {
		log.Error("Failed to start GRPC Server, ", err)
		return
	}
	var opts []grpc.ServerOption
	if config.TLSEnabled {
		log.Info("Setup TLS parameters for GRPC Server.")
	}

	grpcServer = grpc.NewServer(opts...)
	pb.RegisterAgentActionServer(grpcServer, &fluentbit.AgentServcie{})
	pb.RegisterAgentFileServiceServer(grpcServer, &suricata.FileServiceServer{})

	log.Infof("GPRC Server starts in port %d", port)
	grpcServer.Serve(listener)
}

func StopGRPCServer() {
	if grpcServer != nil {
		grpcServer.GracefulStop()
		log.Info("GRPC Server shutdown gracefully")
	}
}
