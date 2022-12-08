package echoip_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/zhangxiaofeng05/com/thirdparty/echoip"
)

func TestIfConfigJson(t *testing.T) {
	t.Run("test json", func(t *testing.T) {
		echoIp, err := echoip.IfConfigJson(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if echoIp.Ip == "" {
			t.Fatalf("ip is nil")
		}
		marshalIndent, err := json.MarshalIndent(echoIp, "", " ")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("IfConfigJson res: %s", marshalIndent)
	})
}

func TestIfConfigIp(t *testing.T) {
	t.Run("test ip", func(t *testing.T) {
		ip, err := echoip.IfConfigIp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if ip == "" {
			t.Fatalf("ip is nil")
		}
		t.Logf("ip: %s", ip)
	})
}
