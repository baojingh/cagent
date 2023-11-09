package server

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func GetTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("Failed to load key pair for server, %s", err)
	}
	ca := x509.NewCertPool()
	capath := "cert/server.crt"
	caByte, _ := os.ReadFile(capath)
	if ok := ca.AppendCertsFromPEM(caByte); !ok {
		log.Fatalf("Failed to parse %q", capath)
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    ca,
	}

	return tlsConfig, nil
}
