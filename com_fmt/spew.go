package com_fmt

import "github.com/davecgh/go-spew/spew"

func SpewPrintf(format string, a ...any) {
	spew.Printf(format, a...)
}
