package comutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhangxiaofeng05/com/comutil"
)

func TestGetEnv(t *testing.T) {
	k := "Jack"
	v := "JackSstValue"
	err := os.Setenv(k, v)
	if err != nil {
		t.Fatal("set env fail")
	}
	gv := os.Getenv(k)
	if gv != v {
		t.Fatal("get env fail")
	}
	list := []struct {
		key   string
		exist bool
	}{
		{k, true},
		{"JackNotExist", false},
	}
	for i, s := range list {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := comutil.GetEnv(s.key, "")
			if (got != "") != s.exist {
				t.Fatalf("get key:%v env wrong", s.key)
			}
		})
	}
}
