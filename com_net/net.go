// Package com_net provide function about net
package com_net

import (
	"errors"
	"log"
	"net"
)

// GetOutboundIP need network
// Get preferred outbound ip of this machine
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("GetOutboundIP conn.Close() err: %v", err)
		}
	}()
	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, errors.New("local addr is not udp addr")
	}
	return localAddr.IP, nil
}

const (
	En0 = "en0"
)

func GetPhysicalAddress() (string, error) {
	inter, err := net.InterfaceByName(En0)
	if err != nil {
		return "", err
	}
	if inter == nil {
		return "", errors.New("inter is nil")
	}
	return inter.HardwareAddr.String(), nil
}
