package com_http_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_http"
)

func TestClient(t *testing.T) {
	url := "https://httpbin.org/get?method=get"
	httpProxy := "http://127.0.0.1:1081"
	sock5Proxy := "socks5://127.0.0.1:1080"

	t.Run("no proxy", func(t *testing.T) {
		client, err := com_http.New("")
		require.NoError(t, err)
		request(t, client, url)
	})
	t.Run("http proxy", func(t *testing.T) {
		client, err := com_http.New(httpProxy)
		require.NoError(t, err)
		request(t, client, url)
	})
	t.Run("sock5 proxy", func(t *testing.T) {
		client, err := com_http.New(sock5Proxy)
		require.NoError(t, err)
		request(t, client, url)
	})
}

func request(t *testing.T, client *com_http.Client, url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	type Result struct {
		Args    map[string]string `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}
	var res Result
	err := client.Get(ctx, url, com_http.DefaultHeader, &res)
	require.NoError(t, err)
	t.Logf("name: %s , origin: %s", t.Name(), res.Origin)
}
