package dotenv_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhangxiaofeng05/com/comutil/dotenv"
)

func TestLoad(t *testing.T) {
	type Args struct {
		Key string
	}
	tests := []struct {
		FileName string
		Args     Args
		Want     string
	}{
		{
			FileName: "../../testdata/.envExample",
			Args:     Args{"KeyMustNotExist"},
			Want:     "sir",
		},
	}
	for i, s := range tests {
		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			err := dotenv.Load(s.FileName)
			if err != nil {
				t.Fatal(err)
			}
			got, ok := os.LookupEnv(s.Args.Key)
			if !ok {
				t.Fatal("key not exist")
			}
			if got != s.Want {
				t.Fatalf("got: %v,but want: %v", got, s.Want)
			}

		})
	}

}
