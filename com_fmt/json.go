package com_fmt

import (
	"encoding/json"
	"fmt"
)

func JsonPrintf(format string, params ...any) {
	wrap := make([]any, 0, len(params))
	for _, p := range params {
		bytes, err := json.Marshal(p)
		if err != nil {
			fmt.Printf("json marshal: %v error", p)
			return
		}
		wrap = append(wrap, string(bytes))
	}

	fmt.Printf(format, wrap...)
}
