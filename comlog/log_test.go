package comlog_test

import (
	"log"
	"testing"

	"github.com/zhangxiaofeng05/com/comlog"
)

func TestInit(t *testing.T) {
	log.Println("default log")
	log.SetFlags(0)
	log.Println("no flag")
	comlog.Init()
	log.Println("use comlog")
}

func TestCheckError(t *testing.T) {
	comlog.CheckError(nil)
}
