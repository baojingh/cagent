package main

import (
	"agent-server/server"
	"flag"

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
	ip         = flag.String("ip", "0.0.0.0", "ip")
	port       = flag.Int("port", 50051, "GRPC Server Port")
	tlsEnabled = flag.Bool("tls", false, "enable TLS")
	certPath   = flag.String("cert", "cert/", "cert directory")
	logPath    = flag.String("log", "log/", "log directory")
	configPath = flag.String("config", "config/", "log directory")
)

func main() {
	flag.Parse()
	// Values from config file will be overwrited by the command value

	server.StartGrpcServer(*ip, *port, *tlsEnabled)
}
