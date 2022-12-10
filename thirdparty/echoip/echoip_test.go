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
			t.Logf("TestIfConfigJson IfConfigJson err: %v", err)
			return
		}
		if echoIp == nil {
			t.Logf("TestIfConfigJson echoIp is nil")
			return
		}
		if echoIp.Ip == "" {
			t.Logf("TestIfConfigJson ip is nil")
			return
		}
		marshalIndent, err := json.MarshalIndent(echoIp, "", " ")
		if err != nil {
			t.Logf("TestIfConfigJson MarshalIndent err: %v", err)
			return
		}
		t.Logf("IfConfigJson res: %s", marshalIndent)
	})
}

func TestIfConfigIp(t *testing.T) {
	t.Run("test ip", func(t *testing.T) {
		ip, err := echoip.IfConfigIp(context.Background())
		if err != nil {
			t.Logf("TestIfConfigJson IfConfigIp err: %v", err)
			return
		}
		if ip == "" {
			t.Logf("TestIfConfigIp ip is nil")
			return
		}
		t.Logf("ip: %s", ip)
	})
}
