package comlog

import "log"

func Init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
