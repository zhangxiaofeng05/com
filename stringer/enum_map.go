package stringer

import "fmt"

const (
	SuccessMap = 1 // success
	FailMap    = 2 // fail
	UnknowMap  = 3 // unknow
)

var CodeMap = map[int]string{
	SuccessMap: "success",
	FailMap:    "fail",
	UnknowMap:  "unknow",
}

func GetMessageByCode(c int) string {
	v, ok := CodeMap[c]
	if ok {
		return v
	}
	return fmt.Sprintf("Invalid(%d)", c)
}
