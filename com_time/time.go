package com_time

import "time"

const (
	defaultLayout = "2006-01-02 15:04:05"
)

func Format(t time.Time) string {
	return t.Format(defaultLayout)
}
