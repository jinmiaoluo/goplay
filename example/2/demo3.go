package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func createContainer(start, stop, interval uint32) *roaring.Bitmap {
	rb := roaring.New()
	for i := start; i < stop; i = i + interval {
		rb.Add(i)
	}
	return rb
}

func main() {

	// test arrayContainer vs arrayContainer And operation
	r1 := createContainer(0, 1000, 1)
	r2 := createContainer(0, 2000, 3)
	result := roaring.And(r1, r2)
	fmt.Println("ac vs ac: ", result.String())

	// test arrayContainer vs bitmapContainer And operation
	r1 = createContainer(0, 1000, 1)
	r2 = createContainer(0, 100000, 3)
	result = roaring.And(r1, r2)
	fmt.Println("ac vs bc: ", result.String())

	// test arrayContainer vs runContainer And operation
	r1 = createContainer(0, 1000, 1)
	r2 = createContainer(0, 2000, 1)
	r2.RunOptimize() // use runContainer to optimize space
	result = roaring.And(r1, r2)
	fmt.Println("ac vs rc: ", result.String())

	// test bitmapContainer vs arrayContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 1000, 3)
	result = roaring.And(r1, r2)
	fmt.Println("bc vs ac:" + result.String())

	// test bitmapContainer vs bitmapContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 100000, 3)
	result = roaring.And(r1, r2)
	fmt.Println("bc vs bc:" + result.String())

	// test bitmapContainer vs runContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 100000, 1)
	r2.RunOptimize() // use runContainer to optimize space
	result = roaring.And(r1, r2)
	fmt.Println("bc vs rc:" + result.String())

	// test runContainer vs arrayContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 1000, 1)
	r1.RunOptimize() // use runContainer to optimize space
	result = roaring.And(r1, r2)
	fmt.Println("rc vs ac:" + result.String())

	// test runContainer vs bitmapContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 6000, 1)
	r1.RunOptimize() // use runContainer to optimize space
	result = roaring.And(r1, r2)
	fmt.Println("rc vs bc:" + result.String())

	// test runContainer vs runContainer And operation
	r1 = createContainer(0, 5000, 1)
	r2 = createContainer(0, 6000, 1)
	r1.RunOptimize() // use runContainer to optimize space
	r2.RunOptimize() // use runContainer to optimize space
	result = roaring.And(r1, r2)
	fmt.Println("rc vs rc:" + result.String())

}
