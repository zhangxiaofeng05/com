package com_fmt_test

import (
	"github.com/zhangxiaofeng05/com/com_fmt"
	"testing"
)

func TestJsonPrintf(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	pp := Person{
		Name: "gg",
		Age:  18,
	}
	com_fmt.JsonPrintf("p1: %v, p2: %v\n", pp, pp)
	com_fmt.JsonPrintf("pp: %v\n", pp)
	com_fmt.JsonPrintf("pp: %+v\n", pp)
	com_fmt.JsonPrintf("pp: %#v\n", pp)
	com_fmt.JsonPrintf("pp: %#+v\n", pp)
}
