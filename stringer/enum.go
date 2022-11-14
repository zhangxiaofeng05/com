// Package stringer stringer(https://github.com/golang/tools/tree/master/cmd/stringer) example
package stringer

//go:generate stringer -type=Enum -linecomment

type Enum uint

const (
	Success Enum = iota + 1 // success
	Fail                    // fail
	Unknow                  // unknow
)
