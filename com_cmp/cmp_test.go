package com_cmp_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/com_cmp"
)

func TestMustEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	got := Person{
		Name: "gg",
		Age:  18,
	}
	want := Person{
		Name: "gg",
		Age:  18,
	}
	err := com_cmp.MustEqual(got, want)
	if err != nil {
		t.Fatal(err)
	}
}
