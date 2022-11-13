// Package comlog provide log init
package comlog

import "log"

// Init standard library log init
func Init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
