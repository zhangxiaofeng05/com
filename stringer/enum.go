package stringer

//go:generate stringer -type=Enum -linecomment

type Enum uint

const (
	Success Enum = iota + 1 // success
	Fail                    // fail
	Unknow                  // unknow
)
