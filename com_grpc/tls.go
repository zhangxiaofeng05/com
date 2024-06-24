package com_grpc

import (
	"fmt"

	"google.golang.org/grpc/credentials"
)

// TLSServer reference: https://github.com/grpc/grpc-go/blob/master/examples/features/encryption/TLS/server/main.go#L54
func TLSServer(serverCertPath, serverKeyPath string) (credentials.TransportCredentials, error) {
	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(serverCertPath, serverKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials: %v", err)
	}
	return creds, nil
}

// TLSClient reference: https://github.com/grpc/grpc-go/blob/master/examples/features/encryption/TLS/client/main.go#L50
func TLSClient(caCertPath, serverName string) (credentials.TransportCredentials, error) {
	// Create tls based credential.
	creds, err := credentials.NewClientTLSFromFile(caCertPath, serverName)
	if err != nil {
		return nil, fmt.Errorf("failed to load credentials: %v", err)
	}
	return creds, nil
}
