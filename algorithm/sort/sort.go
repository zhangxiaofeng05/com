package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle https://zh.wikipedia.org/wiki/Bogo%E6%8E%92%E5%BA%8F
func Shuffle(n []int) {
	for i := 0; i < len(n); i++ {
		swapPosition := rand.Intn(i + 1)
		n[i], n[swapPosition] = n[swapPosition], n[i]
	}
}
