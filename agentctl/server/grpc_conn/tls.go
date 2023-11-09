package grpc_conn

import (
	"crypto/tls"
	"crypto/x509"

	"os"
)

func GetTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if err != nil {
		log.Fatalf("Failed to load key pair for server, %s", err)
	}
	ca := x509.NewCertPool()
	capath := "cert/client.crt"
	caByte, _ := os.ReadFile(capath)
	if ok := ca.AppendCertsFromPEM(caByte); !ok {
		log.Fatalf("Failed to parse %q", capath)
	}

	tlsConfig := &tls.Config{
		ServerName:   "x.com",
		Certificates: []tls.Certificate{cert},
		RootCAs:      ca,
	}

	return tlsConfig, nil
}
