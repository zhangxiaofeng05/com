package com_logrus

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Init reference:https://github.com/sirupsen/logrus/issues/63
// https://github.com/sirupsen/logrus/blob/master/example_custom_caller_test.go
func Init() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				s := strings.Split(frame.Function, ".")
				funcName := s[len(s)-1]
				_, filename := path.Split(frame.File)
				return fmt.Sprintf("%s()", funcName), fmt.Sprintf(" %s:%d", filename, frame.Line)
			},
		},
	)
}
