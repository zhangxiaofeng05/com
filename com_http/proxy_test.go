package com_http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zhangxiaofeng05/com/com_http"
)

func ExampleNew() {
	//proxyUrl := "http://127.0.0.1:1081"
	proxyUrl := "socks5://127.0.0.1:1080"
	client, err := com_http.New(proxyUrl)
	if err != nil {
		panic(err)
	}
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
	err = client.Get(ctx, url, com_http.DefaultHeader, &res)
	if err != nil {
		panic(err)
	}
	if res.Url == "" {
		panic(fmt.Sprintf("res url can't nil. res:%v", res))
	} else {
		marshalIndent, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			panic(err)
		}
		panic(fmt.Sprintf("get res: %s", marshalIndent))
	}
}
