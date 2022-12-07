package comhttp_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/zhangxiaofeng05/com/comhttp"
)

func TestGet(t *testing.T) {
	t.Run("test http get", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		url := "https://httpbin.org/get?method=get"
		type Result struct {
			Args    map[string]string `json:"args"`
			Headers map[string]string `json:"headers"`
			Origin  string            `json:"origin"`
			Url     string            `json:"url"`
		}
		var res Result
		err := comhttp.Get(ctx, url, comhttp.DefaultHeader, &res)
		if err != nil {
			t.Fatal(err)
		}
		if res.Url == "" {
			t.Fatalf("res url can't nil. res:%v", res)
		} else {
			marshalIndent, err := json.MarshalIndent(res, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("get res: %s", marshalIndent)
		}
	})
}

func TestPost(t *testing.T) {
	t.Run("test http post", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		url := "https://httpbin.org/post?method=post"

		data := struct {
			Name       string `json:"name"`
			Girlfriend string `json:"girlfriend"`
		}{
			Name:       "Jack Dawson",
			Girlfriend: "Rose Dawson",
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}

		type Result struct {
			Args    map[string]string `json:"args"`
			Data    string            `json:"data"`
			Files   any               `json:"files"`
			Form    any               `json:"form"`
			Headers map[string]string `json:"headers"`
			Json    any               `json:"json"`
			Origin  string            `json:"origin"`
			Url     string            `json:"url"`
		}
		var res Result
		err = comhttp.Post(ctx, url, comhttp.DefaultHeader, dataBytes, &res)
		if err != nil {
			t.Fatal(err)
		}
		if res.Url == "" {
			t.Fatalf("res url can't nil. res:%v", res)
		} else {
			marshalIndent, err := json.MarshalIndent(res, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("post res: %s", marshalIndent)
		}
	})
}
