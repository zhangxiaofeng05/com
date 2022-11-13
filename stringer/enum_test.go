package stringer

import (
	"testing"
)

func TestEnum_String(t *testing.T) {
	tests := []struct {
		enum    Enum
		want    string
		wantErr bool
	}{
		{1, "success", false},
		{2, "fail", false},
		{3, "unknow", false},
		{4, "Enum(4)", false},
		{5, "", true},
	}
	for i, tt := range tests {
		result := tt.enum.String()
		if result != tt.want {
			if tt.wantErr {
				t.Logf("case %d . result:%s want:%s", i, result, tt.want)
				return
			}
			t.Errorf("case %d run fail. result:%s != want:%s", i, result, tt.want)
			return
		}
	}
}
