package sort_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/algorithm/sort"
)

func TestShuffle(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("before n: %+v", n)
	sort.Shuffle(n)
	t.Logf("after n: %+v", n)

}
