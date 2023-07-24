// Package com_net provide function about net
package com_net

import (
	"fmt"
	"net"
)

// GetOutboundIP need network
// Get preferred outbound ip of this machine
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
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
		return "", fmt.Errorf("inter is nil")
	}
	return inter.HardwareAddr.String(), nil
}
