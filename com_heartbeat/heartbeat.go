package com_heartbeat

import (
	"net/http"
	"strings"
)

// Heartbeat reference: https://github.com/go-chi/chi/blob/master/middleware/heartbeat.go
func Heartbeat(endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && strings.EqualFold(r.URL.Path, endpoint) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("pong"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
