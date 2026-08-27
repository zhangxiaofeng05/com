package com_fmt

import (
	"encoding/json"
	"log"
)

func JsonPrintf(format string, params ...any) {
	wrap := make([]any, 0, len(params))
	for _, p := range params {
		bytes, err := json.Marshal(p)
		if err != nil {
			log.Printf("json marshal: %v error", p)
			return
		}
		wrap = append(wrap, string(bytes))
	}

	log.Printf(format, wrap...)
}
