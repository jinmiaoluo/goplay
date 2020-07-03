package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	r1 := roaring.New()
	for i := uint32(100); i < 1000; i++ {
		r1.Add(i)
	}

	r1.RunOptimize()

	if !r1.Contains(500) {
		fmt.Errorf("should contain 500")
	}

	// 测试 runContainer 的
	fmt.Println(r1.String())
}
