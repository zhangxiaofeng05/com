package com_heartbeat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Http reference: https://github.com/go-chi/chi/blob/master/middleware/heartbeat.go
func Http(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func Gin(abort bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		Http(c.Writer, c.Request)
		if abort {
			c.Abort()
		}
	}
}
