package com_env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhangxiaofeng05/com/com_env"
)

func TestGetEnv(t *testing.T) {
	const (
		defaultValue = "defaultValue"
		k            = "Jack"
		v            = "JackSstValue"
	)

	err := os.Setenv(k, v)
	if err != nil {
		t.Fatal("set env fail")
	}
	gv := os.Getenv(k)
	if gv != v {
		t.Fatal("get env fail")
	}

	tests := []struct {
		key  string
		want string
	}{
		{k, v},
		{"JackNotExist", defaultValue},
	}
	for i, s := range tests {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := com_env.GetEnv(s.key, defaultValue)
			if got != s.want {
				t.Fatalf("get key:%v env wrong", s.key)
			}
		})
	}
}

func TestLookupEnv(t *testing.T) {
	const (
		defaultValue = "defaultValue"
		k            = "Jack"
		v            = "JackSstValue"
	)

	err := os.Setenv(k, v)
	if err != nil {
		t.Fatal("set env fail")
	}
	gv := os.Getenv(k)
	if gv != v {
		t.Fatal("get env fail")
	}
	var tests = []struct {
		key  string
		want string
	}{
		{k, v},
		{"JackNotExist", defaultValue},
	}
	for i, s := range tests {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := com_env.LookupEnv(s.key, defaultValue)
			if got != s.want {
				t.Fatalf("get key:%v env wrong", s.key)
			}
		})
	}
}
