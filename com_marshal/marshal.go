package com_marshal

import "encoding/json"

func MarshaToString(a any) (str string, err error) {
	switch as := a.(type) {
	case string:
		str = as
	case []byte:
		str = string(as)
	default:
		var b []byte
		b, err = json.Marshal(as)
		str = string(b)
	}
	return
}

func UnmarshalAny(str any, a any) (err error) {
	switch as := a.(type) {
	case string:
		err = json.Unmarshal([]byte(as), &str)
	case []byte:
		err = json.Unmarshal(as, &str)
	default:
		var b []byte
		b, err = json.Marshal(as)
		if err != nil {
			return
		}
		err = json.Unmarshal(b, str)
	}
	return
}
