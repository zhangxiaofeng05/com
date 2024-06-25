package com_heartbeat_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestGin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		abort      bool
		expectCode int
		expectBody string
	}{
		{"case 1", true, http.StatusOK, "pong"},
		{"case 2", false, http.StatusOK, "pong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建一个新的 gin 引擎
			r := gin.New()

			// 添加一个中间件处理程序来检查 c.Abort() 是否被调用
			aborted := false
			r.Use(func(c *gin.Context) {
				c.Next()
				if c.IsAborted() {
					aborted = true
				}
			})

			// 将 Gin 处理程序添加到路由中
			relativePath := "/ping"
			r.GET(relativePath, com_heartbeat.Gin(tt.abort))

			// 创建一个新的 HTTP 请求
			req, _ := http.NewRequest(http.MethodGet, relativePath, nil)
			w := httptest.NewRecorder()

			// 处理请求
			r.ServeHTTP(w, req)

			// 检查状态码
			require.Equal(t, tt.expectCode, w.Code)

			// 检查响应体
			require.Equal(t, tt.expectBody, w.Body.String())

			// 检查是否调用了 c.Abort()
			require.Equal(t, tt.abort, aborted)
		})
	}
}
