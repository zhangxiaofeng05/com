package bitmap_test

import (
	"github.com/zhangxiaofeng05/com/datastruct/bitmap"
	"testing"
)

func TestBitMap(t *testing.T) {
	b := bitmap.New(1000)
	b.Set(300)
	b.Set(400)
	b.Set(600)
	exist := b.Contains(300)
	if exist != true {
		t.Fatal("want true, but get false")
	}
	exist1 := b.Contains(301)
	if exist1 != false {
		t.Fatal("want false, but get true")
	}
	b.Remove(300)
	exist2 := b.Contains(300)
	if exist2 != false {
		t.Fatal("want false, but get true")
	}
}
