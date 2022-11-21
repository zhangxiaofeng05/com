package comip

import (
	"errors"
	"net"
	"testing"
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
	outboundIP, err := GetOutboundIP()
	if err != nil {
		t.Fatalf("GetOutboundIP err:%v", err)
	}
	ip, err := externalIP()
	if err != nil {
		t.Fatalf("externalIP err:%v", err)
	}
	if outboundIP.String() != ip {
		t.Fatalf("got: %v, want: %v", outboundIP.String(), ip)
	}
}

func TestGetLocalIP(t *testing.T) {
	localIP, err := GetLocalIP()
	if err != nil {
		t.Fatalf("GetLocalIP err:%v", err)
	}

	ip, err := externalIP()
	if err != nil {
		t.Fatalf("externalIP err:%v", err)
	}

	if localIP != ip {
		t.Fatalf("got: %v, want: %v", localIP, ip)
	}
}
