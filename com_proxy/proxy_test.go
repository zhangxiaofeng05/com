package com_proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpClient(t *testing.T) {
	tests := []struct {
		proxyStr   string
		expectErr  bool
		expectDial bool
	}{
		{proxyStr: "http://localhost:8080", expectErr: false, expectDial: true},
		//{proxyStr: "socks5://localhost:8080", expectErr: false, expectDial: true}, // 不支持测试
		{proxyStr: "invalid://localhost:8080", expectErr: true, expectDial: false},
	}

	for _, test := range tests {
		client, err := HttpClient(test.proxyStr)
		if test.expectErr && err == nil {
			t.Errorf("expected an error for proxyStr: %s", test.proxyStr)
		} else if !test.expectErr && err != nil {
			t.Errorf("did not expect an error for proxyStr: %s, got: %v", test.proxyStr, err)
		}

		if client != nil {
			// Create a test server to simulate a real server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer ts.Close()

			resp, err := client.Get(ts.URL)
			if err != nil {
				defer resp.Body.Close()
			}
			if test.expectDial && err != nil {
				t.Errorf("expected to dial successfully for proxyStr: %s, got: %v", test.proxyStr, err)
			} else if !test.expectDial && err == nil {
				t.Errorf("expected not to dial for proxyStr: %s", test.proxyStr)
			}

			if resp != nil && resp.StatusCode != http.StatusOK {
				t.Errorf("expected status code 200 for proxyStr: %s, got: %d", test.proxyStr, resp.StatusCode)
			}
		}
	}
}
