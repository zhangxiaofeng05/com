package sort

import "github.com/zhangxiaofeng05/com/com_math"

// Shuffle https://zh.wikipedia.org/wiki/Bogo%E6%8E%92%E5%BA%8F
func Shuffle(n []int) {
	for i := 0; i < len(n); i++ {
		swapPosition := com_math.Rand.Intn(i + 1)
		n[i], n[swapPosition] = n[swapPosition], n[i]
	}
}
