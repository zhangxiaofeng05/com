// Package comip provide function about ip
package comip

import (
	"net"
	"os"
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

// GetLocalIP localhost ip.such as 192.168.31.147
func GetLocalIP() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil && ipv4.String() != "127.0.0.1" {
			return ipv4.String(), nil
		}
	}
	return "", nil
}
