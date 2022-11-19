// recommend read: https://gobyexample.com/testing-and-benchmarking
package stringer

import (
	"fmt"
	"testing"
)

func TestEnum_String(t *testing.T) {
	var tests = []struct {
		enum Enum
		want string
	}{
		{1, "success"},
		{2, "fail"},
		{3, "unknow"},
		{4, "Enum(4)"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("enum:%d", tt.enum)
		t.Run(testname, func(t *testing.T) {
			ans := tt.enum.String()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})

	}
}

func BenchmarkEnum_String(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Enum(1).String()
		_ = Enum(2).String()
		_ = Enum(3).String()
	}
}

func TestGetMessageByCode(t *testing.T) {
	var tests = []struct {
		enum int
		want string
	}{
		{1, "success"},
		{2, "fail"},
		{3, "unknow"},
		{4, "Invalid(4)"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("enum:%d", tt.enum)
		t.Run(testname, func(t *testing.T) {
			ans := GetMessageByCode(tt.enum)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})

	}
}

func BenchmarkGetMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = GetMessageByCode(1)
		_ = GetMessageByCode(2)
		_ = GetMessageByCode(3)
	}
}
