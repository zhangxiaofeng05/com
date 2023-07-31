package com_fmt

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func SpewPrintf(format string, params ...any) {
	//spew.Printf(format, a...)
	wrap := make([]any, 0, len(params))
	for _, p := range params {
		formatter := spew.NewFormatter(p)
		wrap = append(wrap, formatter)
	}
	fmt.Printf(format, wrap...)
}
