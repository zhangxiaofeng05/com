package com_heartbeat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	defaultPath = "/ping"
)

type gm struct {
	path  string // 健康检查路径
	abort bool
}

type Option func(*gm)

func GinWithAbort() Option {
	return func(g *gm) {
		g.abort = true
	}
}

func GinWithPath(path string) Option {
	return func(g *gm) {
		g.path = path
	}
}

func Gin(opts ...Option) gin.HandlerFunc {
	g := &gm{
		path:  defaultPath,
		abort: false,
	}
	for _, opt := range opts {
		opt(g)
	}
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet && c.Request.URL.Path == g.path {
			Http(c.Writer, c.Request)
			if g.abort {
				c.Abort()
			}
			c.Next()
		}
	}
}
