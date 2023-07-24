// Package com_log provide function about log standard library
package com_log

import "log"

// Init standard library log init
func Init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
