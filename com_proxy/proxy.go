package com_proxy

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

const (
	httpPrefix   = "http://"
	socks5Prefix = "socks5://"
)

func HttpClient(proxyStr string) (*http.Client, error) {
	if strings.HasPrefix(proxyStr, httpPrefix) {
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			return nil, err
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
		return client, nil
	} else if strings.HasPrefix(proxyStr, socks5Prefix) {
		// 去除前缀
		proxyStr = strings.TrimPrefix(proxyStr, socks5Prefix)
		dialer, err := proxy.SOCKS5("tcp", proxyStr, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}
		return &http.Client{
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
		}, nil
	} else {
		return nil, errors.New("proxyStr error")
	}
}
