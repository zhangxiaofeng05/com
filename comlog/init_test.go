package comlog

import (
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	log.Println("default log")
	log.SetFlags(0)
	log.Println("no flag")
	Init()
	log.Println("use comlog")
}
