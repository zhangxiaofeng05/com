package com_math

import (
	"math/rand"
	"time"
)

var (
	Rand *rand.Rand
)

func init() {
	Rand = rand.New(rand.NewSource(time.Now().Unix()))
}
