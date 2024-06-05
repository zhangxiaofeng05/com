package com_cmp

import (
	"errors"

	"github.com/google/go-cmp/cmp"
)

// MustEqual 如果x和y不相等，输出 diff 并退出
// 按照 + - 可以把 x 转为 y
func MustEqual(x, y any) error {
	eq := cmp.Equal(x, y)
	if !eq {
		diff := cmp.Diff(x, y)
		return errors.New(diff)
	}
	return nil
}
