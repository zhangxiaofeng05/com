package ipsb_test

import (
	"context"
	"testing"

	"github.com/zhangxiaofeng05/com/thirdparty/ipsb"
)

func TestJsonIp(t *testing.T) {
	t.Run("test jsonip", func(t *testing.T) {
		ip, err := ipsb.JsonIp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("ip.sb ip: %v", ip)
	})
}

func TestGeoIp(t *testing.T) {
	t.Run("test geoip", func(t *testing.T) {
		geoIp, err := ipsb.GeoIp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("geoIp: %#v", geoIp)
	})
}
