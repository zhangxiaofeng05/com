package echoip_test

//func TestIfConfigJson(t *testing.T) {
//	t.Run("test json", func(t *testing.T) {
//		echoIp, err := echoip.IfConfigJson(context.Background())
//		if err != nil {
//			t.Fatal(err)
//		}
//		if echoIp == nil {
//			t.Fatal("TestIfConfigJson echoIp is nil")
//		}
//		if echoIp.Ip == "" {
//			t.Fatal("TestIfConfigJson ip is nil")
//		}
//		marshalIndent, err := json.MarshalIndent(echoIp, "", " ")
//		if err != nil {
//			t.Fatalf("TestIfConfigJson MarshalIndent err: %v", err)
//		}
//		t.Logf("IfConfigJson res: %s", marshalIndent)
//	})
//}

//func TestIfConfigIp(t *testing.T) {
//	t.Run("test ip", func(t *testing.T) {
//		result, err := echoip.IfConfigIp(context.Background())
//		if err != nil {
//			t.Fatal(err)
//		}
//		if len(result) > 16 {
//			t.Log(len(result))
//			t.Fatal(result)
//		}
//		if result == "" {
//			t.Logf("TestIfConfigIp ip is nil")
//			return
//		}
//		t.Logf("ip: %s", result)
//	})
//}
