package main

import (
	"agent-server/config"
	"agent-server/server"
	"flag"
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

var (
	// flag.string etc. return a pointer
	tlsEnabled = flag.Bool("tls", false, "enable TLS")
	certPath   = flag.String("cert", "", "cert directory")
	logPath    = flag.String("log", "", "log directory")
)

func main() {
	log.Info("start entry-point for agent-server")
	go PrepareforShutdown()
	flag.Parse()
	// Values from config file will be overwrited by the command value
	config.SetupConfig(*certPath, *logPath, *tlsEnabled)
	server.StartGrpcServer()
}

// shutdown GRPC Server gracefully
func PrepareforShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, os.Kill)
	sig := <-signalChan
	close(signalChan)
	log.Warnf("Receive signal %s and prepare for exitting", sig)
	server.StopGRPCServer()
	os.Exit(0)
}
