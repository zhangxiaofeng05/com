package com_net_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/zhangxiaofeng05/com/com_net"
)

// externalIP only ipv4
// https://code.google.com/p/whispering-gophers/source/browse/util/helper.go
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func TestGetOutboundIP(t *testing.T) {
	got, err := com_net.GetOutboundIP()
	if err != nil {
		t.Fatalf("GetOutboundIP err:%v", err)
	}
	ip, err := externalIP()
	if err != nil {
		t.Fatalf("externalIP err:%v", err)
	}
	if got.String() != ip {
		t.Fatalf("got: %v, want: %v", got.String(), ip)
	}
	t.Logf("got: %v, want: %v", got.String(), ip)
}

func ExampleGetOutboundIP() {
	ip, err := com_net.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.String())
	// 192.168.31.147
}

func ExampleGetPhysicalAddress() {
	physicalAddress, err := com_net.GetPhysicalAddress()
	if err != nil {
		panic(err)
	}
	fmt.Printf("physicalAddress: %v", physicalAddress)
	// mac: b1:bd:84:51:1c:bc
}
