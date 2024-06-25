package com_heartbeat

import "github.com/gin-gonic/gin"

type gm struct {
	abort bool
}

type Option func(*gm)

func GinWithAbort() Option {
	return func(g *gm) {
		g.abort = true
	}
}

func Gin(opts ...Option) gin.HandlerFunc {
	g := &gm{
		abort: false,
	}

	for _, opt := range opts {
		opt(g)
	}
	return func(c *gin.Context) {
		Http(c.Writer, c.Request)
		if g.abort {
			c.Abort()
		}
	}
}
