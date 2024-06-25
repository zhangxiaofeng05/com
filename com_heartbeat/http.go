package com_heartbeat

import (
	"net/http"
)

// Http reference: https://github.com/go-chi/chi/blob/master/middleware/heartbeat.go
func Http(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}
