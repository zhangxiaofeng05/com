package com_heartbeat_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_heartbeat"
)

func TestHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:5000/ping", nil)
	rr := httptest.NewRecorder()

	// Call the Heartbeat handler
	com_heartbeat.Http(rr, req)

	// Check the status code
	require.Equal(t, http.StatusOK, rr.Code)

	// Check the Content-Type header
	expectedContentType := "text/plain"
	require.Equal(t, expectedContentType, rr.Header().Get("Content-Type"))

	// Check the response body
	expectedBody := "pong"
	require.Equal(t, expectedBody, rr.Body.String())
}
