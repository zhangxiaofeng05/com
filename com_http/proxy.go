package com_http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zhangxiaofeng05/com/com_proxy"
)

type Client struct {
	httpClient *http.Client
}

// New get client
// reference: https://gist.github.com/leafney/0beac92b784fae03c070b09983704c6f
func New(proxyUrl string) (*Client, error) {
	if proxyUrl == "" {
		return &Client{http.DefaultClient}, nil
	}
	// socks5 or http
	client, err := com_proxy.HttpClient(proxyUrl)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// Get get method. return json data
func (c *Client) Get(ctx context.Context, url string, header map[string]string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	errorBuf, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("HTTP GET url:%s response status: %v, err:%v", url, resp.Status, errorBuf)
}

// Post post method. return json data
func (c *Client) Post(ctx context.Context, url string, header map[string]string, data []byte, result any) error {
	body := bytes.NewBuffer(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	errorBuf, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("HTTP POST url:%s response status: %v, err:%v", url, resp.Status, errorBuf)
}
