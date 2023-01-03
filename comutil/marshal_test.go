package comutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

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
			val, err := comutil.MarshaToString(test.p)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, test.want, val)
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
		comutil.UnmarshalAny(&p, m)
		assert.Equal(t, P{"jack", 18}, p)
	})

	t.Run("case string", func(t *testing.T) {
		comutil.UnmarshalAny(&p, `{"name":"jack","age":20}`)
		assert.Equal(t, P{"jack", 20}, p)
	})

	t.Run("case []byte", func(t *testing.T) {
		comutil.UnmarshalAny(&p, []byte(`{"name":"jack","age":22}`))
		assert.Equal(t, P{"jack", 22}, p)
	})
}
