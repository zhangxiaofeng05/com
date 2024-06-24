package com_grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc/credentials"
)

// MTLSServer reference: https://github.com/grpc/grpc-go/blob/master/examples/features/encryption/mTLS/server/main.go#L52
func MTLSServer(serverCertPath, serverKeyPath, clientCaCertPath string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load key pair: %w", err)
	}

	ca := x509.NewCertPool()
	caBytes, err := os.ReadFile(clientCaCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ca cert %q: %w", clientCaCertPath, err)
	}
	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		return nil, fmt.Errorf("failed to parse %q", clientCaCertPath)
	}

	tlsConfig := &tls.Config{ // nolint:gosec
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    ca,
	}
	creds := credentials.NewTLS(tlsConfig)
	return creds, nil
}

// MTLSClient reference: https://github.com/grpc/grpc-go/blob/master/examples/features/encryption/mTLS/client/main.go#L53
func MTLSClient(clientCertPath, clientKeyPath, caCertPath, serverName string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		log.Fatalf("failed to load client cert: %v", err)
	}

	ca := x509.NewCertPool()
	caBytes, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("failed to read ca cert %q: %v", caCertPath, err)
	}
	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		log.Fatalf("failed to parse %q", caCertPath)
	}

	tlsConfig := &tls.Config{ // nolint:gosec
		ServerName:   serverName,
		Certificates: []tls.Certificate{cert},
		RootCAs:      ca,
	}
	creds := credentials.NewTLS(tlsConfig)
	return creds, nil
}
