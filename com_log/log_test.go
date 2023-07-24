package com_log_test

import (
	"log"
	"testing"

	"github.com/zhangxiaofeng05/com/com_log"
)

func TestInit(t *testing.T) {
	log.Println("default log")
	log.SetFlags(0)
	log.Println("no flag")
	com_log.Init()
	log.Println("use com_log")
}
