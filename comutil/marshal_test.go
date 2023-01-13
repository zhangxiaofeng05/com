package comutil_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/zhangxiaofeng05/com/comutil"
)

func TestMarshaToString(t *testing.T) {
	testCase := []struct {
		p    any
		want string
	}{
		{"jack", "jack"},
		{[]byte{106, 97, 99, 107}, "jack"},
		{map[string]string{"jack": "rose"}, `{"jack":"rose"}`},
	}
	for idx, test := range testCase {
		name := fmt.Sprintf("case %d", idx)
		t.Run(name, func(t *testing.T) {
			got, err := comutil.MarshaToString(test.p)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Fatalf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func TestUnmarshalAny(t *testing.T) {
	type P struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p P
	t.Run("case map", func(t *testing.T) {
		m := map[string]any{"name": "jack", "age": 18}
		err := comutil.UnmarshalAny(&p, m)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(P{"jack", 18}, p); diff != "" {
			t.Fatalf("(-want +got): \n%s", diff)
		}
	})

	t.Run("case string", func(t *testing.T) {
		err := comutil.UnmarshalAny(&p, `{"name":"jack","age":20}`)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(P{"jack", 20}, p); diff != "" {
			t.Fatalf("(-want +got): \n%s", diff)
		}
	})

	t.Run("case []byte", func(t *testing.T) {
		err := comutil.UnmarshalAny(&p, []byte(`{"name":"jack","age":22}`))
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(P{"jack", 22}, p); diff != "" {
			t.Fatalf("(-want +got): \n%s", diff)
		}
	})
}
