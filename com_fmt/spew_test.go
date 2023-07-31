package com_fmt_test

import (
	"github.com/zhangxiaofeng05/com/com_fmt"
	"testing"
)

func TestSpewPrintf(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	pp := Person{
		Name: "gg",
		Age:  18,
	}
	com_fmt.SpewPrintf("p1: %v, p2: %v\n", pp, pp)
	com_fmt.SpewPrintf("pp: %v\n", pp)
	com_fmt.SpewPrintf("pp: %+v\n", pp)
	com_fmt.SpewPrintf("pp: %#v\n", pp)
	com_fmt.SpewPrintf("pp: %#+v\n", pp)
}
