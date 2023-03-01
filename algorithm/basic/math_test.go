package basic_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/algorithm/basic"
)

func TestGcd(t *testing.T) {
	got := basic.Gcd(54, 24)
	if got != 6 {
		t.Fatal("TestGcd fail")
	}
}
