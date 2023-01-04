package comutil_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zhangxiaofeng05/com/comutil"
	"testing"
)

func TestGetEnv(t *testing.T) {
	list := []struct {
		key   string
		exist bool
	}{
		{"GOPATH", true},
		{"GOPATH-NOT-EXIST", false},
	}
	for i, s := range list {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			value := comutil.GetEnv(s.key, "")
			assert.Equal(t, s.exist, value != "", s)
		})
	}
}
