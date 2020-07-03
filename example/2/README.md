**Roaring Bitmaps**

Bitmaps 用于存放常见的数学上的集合. 比如我们有这个一个集合 {1, 3, 7}. 它对应的 Bitmaps 就是 `10001010`

**为什么 Cardinality 在大于 4096 时会触发 Array 到 Bitmaps 的转换**

因为 4096 个 short int (2 Bytes) 组成的数组容器需要的空间是 4096 * 2 Bytes = 8 KBytes. 一个 Bitmaps 用于表示一个 32 位整数的低 2^16 表示的数值, 2^16 = 65536, 因此一个 Bitmaps 需要 65536 个 bit 位. 一个 Bitmaps 需要的空间是 65536 bits / 8 = 8 KBytes. 因为一个含有 4096 个成员的数组其占用空间和一个 Bitmaps 一样, 并且, 如果超过 4096 个成员, 数组需要更多的空间才能存放新的成员. 而 Bitmaps 只要将对应的位从 0 置 1 即可, 无需更多的空间. 因此需要做一个 Array 到 Bitmaps 的转换. 这是从空间角度考虑的. 另外, 从时间角度, 因为 Bitmaps 做 union 和 intersection 本质是 Bit Or 和 Bit And. 效率更高. 这是从时间层面考虑.

```go
func (ac *arrayContainer) iaddReturnMinimized(x uint16) container {
	// Special case adding to the end of the container.
	l := len(ac.content)
	if l > 0 && l < arrayDefaultMaxSize && ac.content[l-1] < x {
		ac.content = append(ac.content, x)
		return ac
	}

	loc := binarySearch(ac.content, x)

	if loc < 0 {
		if len(ac.content) >= arrayDefaultMaxSize {
			a := ac.toBitmapContainer()
			a.iadd(x)
			return a
		}
		s := ac.content
		i := -loc - 1
		s = append(s, 0)
		copy(s[i+1:], s[i:])
		s[i] = x
		ac.content = s
	}
	return ac
}
```

**Roaring Bitmaps 如何实现数据的压缩**

压缩空间. 用 runContainer 表示 arrayContainer 中连续的成员.

**bitmapContainer 如何存储数据**

bitmapContainer 本质是一个 65536 个 bit 的大二进制(blob). 但是, 在 golang 中, 整数最大的 bit 是 64 位(即: uint64). 因为 bitmap 是通过一个 bit 置 1 来表示一个数是否存在(比如集合 [1,2,3] 用 bitmap 表示就是 00001110, 这里, 第一个 bit 位表示 0, 第二个 bit 表示 1, 第 3 个 bit 表示 2, 第 4 个 bit 表示 3, and so forth) 为了表示这个 blob, 我们将整数的值按照 64 为一份进行平分(因为 golang 中最大的 uint 类型只能有 64 bit). 由于总共有 65536 个 bit 构成一个 bitmap, 因此, 极限情况下, 需要这个数组有 1024 个 uint64 成员, 表示这 blob. 因此, bitmapContainer 实际上就是一个可能拥有 1024 个 uint64 成员的数组.

```go
func (bc *bitmapContainer) loadData(arrayContainer *arrayContainer) {
	bc.cardinality = arrayContainer.getCardinality()
	c := arrayContainer.getCardinality()
	for k := 0; k < c; k++ {
		x := arrayContainer.content[k]
		i := int(x) / 64
		bc.bitmap[i] |= (uint64(1) << uint(x%64))
	}
}
```

**参考**
- [Lemire's paper](https://arxiv.org/pdf/1402.6407.pdf)
- [高效压缩位图RoaringBitmap的原理与应用](https://www.jianshu.com/p/818ac4e90daf)
- [Roaring Bitmap更好的位图压缩算法](http://smartsi.club/better-bitmap-performance-with-roaring-bitmaps.html)
