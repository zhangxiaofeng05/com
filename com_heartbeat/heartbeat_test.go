package com_heartbeat_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_heartbeat"
)

func TestHeartbeat(t *testing.T) {
	endpoint := "/ping"
	handler := com_heartbeat.Heartbeat(endpoint)

	tests := []struct {
		name         string
		method       string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"case 1", "GET", endpoint, http.StatusOK, "pong"},
		{"case 2", "HEAD", endpoint, http.StatusNotFound, ""},
		{"case 3", "POST", endpoint, http.StatusNotFound, ""},
		{"case 4", "GET", "/other", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			require.Equal(t, tt.expectedCode, rr.Code)
			require.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
