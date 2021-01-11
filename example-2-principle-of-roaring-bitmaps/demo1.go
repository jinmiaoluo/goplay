package main

import (
	"github.com/RoaringBitmap/roaring"
)

func main() {

	// 这部分的代码用于演示当基数大于等于 4096 时发生的 arrayContainer -> bitmapContainer 变换的过程
	rb1 := roaring.New()
	for i := uint32(0); i < 5000; i++ {
		rb1.Add(i)
	}
}
