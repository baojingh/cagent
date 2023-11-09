package server

import (
	"agent-server/config"
	logger "agent-server/log"
	"agent-server/pb"
	"agent-server/server/fluentbit"
	"agent-server/server/suricata"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		config, _ := GetTLSConfig()
		cre := credentials.NewTLS(config)
		opts = []grpc.ServerOption{grpc.Creds(cre)}
		log.Info("Setup TLS parameters for GRPC Server.")
	}
	grpcServer = grpc.NewServer(opts...)
	pb.RegisterAgentActionServer(grpcServer, &fluentbit.AgentServcie{})

	pb.RegisterAgentFileServiceServer(grpcServer, &suricata.AgentFileServiceServer{})

	log.Infof("GPRC Server starts in port %d", port)
	grpcServer.Serve(listener)
}

// GRPC server cannot be shutdown if there is any active conn or stream.
// Process the active worker for given seconds and then kill them forcely.
func StopGRPCServer() {
	ctxShut, cancelShut := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShut()
	go func() {
		if grpcServer != nil {
			grpcServer.GracefulStop()
		}
	}()

	select {
	case <-ctxShut.Done():
		grpcServer.Stop()
	}
	log.Warn("GRPC Server stops success.")
}
