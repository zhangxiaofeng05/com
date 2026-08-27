package ipsb_test

import (
	"context"
	"testing"

	"github.com/zhangxiaofeng05/com/third_party/ip/ipsb"
)

func TestJsonIp(t *testing.T) {
	t.Run("test jsonip", func(t *testing.T) {
		ip, err := ipsb.JsonIp(context.Background(), nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("ip.sb ip: %v", ip)
	})
}

func TestGeoIp(t *testing.T) {
	t.Run("test geoip", func(t *testing.T) {
		geoIp, err := ipsb.GeoIp(context.Background(), nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("geoIp: %#v", geoIp)
	})
}
