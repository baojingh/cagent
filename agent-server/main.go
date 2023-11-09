package main

import (
	"agent-server/server"
	"os"
	"os/signal"
	"sync"
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

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		PrepareforShutdown()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.StartGrpcServer()
	}()
	wg.Wait()
	os.Exit(0)
}

// shutdown GRPC Server gracefully
func PrepareforShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, os.Kill)
	sig := <-signalChan
	close(signalChan)
	log.Warnf("Receive signal %s and prepare for exit...", sig)
	server.StopGRPCServer()
}
