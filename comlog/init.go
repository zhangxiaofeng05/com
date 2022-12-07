// Package comlog provide log init
package comlog

import "log"

// Init standard library log init
func Init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// CheckError will fatal when err not nil
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
