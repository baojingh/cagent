package main

import (
	"agent-server/server"
	"os"
	"os/signal"
	"syscall"

	logger "agent-server/log"
)

/**
  Author     : He Bao Jing
  Date       : 10/20/2023 5:06 PM
  Description:
*/

var log = logger.New()

func main() {
	log.Info("Prepare for GRPC Server startup...")
	go PrepareforShutdown()
	server.StartGrpcServer()
}

// shutdown GRPC Server gracefully
func PrepareforShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, os.Kill)
	sig := <-signalChan
	close(signalChan)
	log.Warnf("Receive signal %s and prepare for exit...", sig)
	server.StopGRPCServer()
	os.Exit(0)
}
