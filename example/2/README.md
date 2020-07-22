# Roaring Bitmaps
bitmap 用于存放常见的数学上的集合. 比如我们有这个一个集合 {1, 3, 7}. 它对应的 bitmap 就是 `10001010`. Roaring Bitmaps 是 bitmap 的实现, 并相对其他实现, 有更好的性能.

## 目录
<!-- vim-markdown-toc GFM -->

* [为什么基数大于等于 4096 时会触发 arrayContainer 到 bitmapContainer 的转换](#为什么基数大于等于-4096-时会触发-arraycontainer-到-bitmapcontainer-的转换)
* [runContainer 作用](#runcontainer-作用)
* [bitmapContainer 如何存储数据](#bitmapcontainer-如何存储数据)
* [Roaring Bitmaps 如何实现 bitmap 的压缩](#roaring-bitmaps-如何实现-bitmap-的压缩)
* [bitmapContainer 到 runContainer 转换时, 如何计算 runContainer 需要的 interval 数组的成员的个数](#bitmapcontainer-到-runcontainer-转换时-如何计算-runcontainer-需要的-interval-数组的成员的个数)
* [如何将 bitmapContainer 转换为 runContainer](#如何将-bitmapcontainer-转换为-runcontainer)
* [两个 bitmapContainer 求 And 后将结果存放在 arrayContainer 逻辑实现](#两个-bitmapcontainer-求-and-后将结果存放在-arraycontainer-逻辑实现)
* [arrayContainer 和 bitmapContainer 求 And 后将结果存放在 arrayContainer 实现](#arraycontainer-和-bitmapcontainer-求-and-后将结果存放在-arraycontainer-实现)
* [runContainer 和 bitmapContainer 求 And 实现](#runcontainer-和-bitmapcontainer-求-and-实现)
* [给 bitmapContainer 设置连续的 1 的实现](#给-bitmapcontainer-设置连续的-1-的实现)
* [求插入索引逻辑实现](#求插入索引逻辑实现)
* [runContainer 和 arrayContainer 求 And 实现](#runcontainer-和-arraycontainer-求-and-实现)
* [runContainer 和 runContainer 求 and 逻辑实现](#runcontainer-和-runcontainer-求-and-逻辑实现)
* [参考](#参考)

<!-- vim-markdown-toc -->

#### 为什么基数大于等于 4096 时会触发 arrayContainer 到 bitmapContainer 的转换

因为 4096 个 int16 (2 Bytes) 成员组成的 arrayContainer 需要的空间是 4096 * 2 Bytes = 8 KBytes. 一个 bitmap 用于表示一个 32 位整数的低 16 bit 表示的数值, 2^16 = 65536, 因此一个 bitmap 需要 65536 个 bit 位. 一个 bitmap 需要的空间转换成 KBytes 是: 65536 bits / 8 = 8 KBytes. 因为一个含有 4096 个成员的数组其占用空间和一个 bitmap 一样, 都是 8 KBytes, 并且, 如果超过 4096 个成员, 数组需要更多的空间才能存放新成员. 而 bitmap 只要将对应的位从 0 置 1 即可, 无需更多的空间. 因此需要做一个 arrayContainer 到 bitmapContainer 的转换. 这是从空间角度考虑的. 另外, 从时间角度, 因为 bitmap 做 union 和 intersection 本质是 Bit Or 和 Bit And. 效率更高. 这是从时间层面考虑

主要实现的代码如下:
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

我们可以通过 dlv(delve 项目提供的一个命令行工具, 如果没有你需要先安装) 调试的方式, 复现这个过程.

步骤如下:

```bash
# 1. 进入调试模式
dlv debug demo1.go

# 2. 我们会进入一个 shell, 下文中 `>' 开头的命令表示这是 shell 里面的命令
# 因为转换的动作只发生在 i == 4096 时, 因此我们需要用到调试中的 condition 功能
# condition 功能的作用: 当 i == 4096 时, breakpoint 1 才生效, 否则断点不生效
> break demo1.go:12
> condition 1 i == 4096
> continue

# 3. 在目标函数打断点
# 这是本次调试的目标函数, 我们打上断点并跳转后, 按步调试即可
> break github.com/RoaringBitmap/roaring.(*arrayContainer).toBitmapContainer
> continue
> step
```

#### runContainer 作用

压缩空间. 用 runContainer 表示 arrayContainer 或者 bitmapContainer. 在密集型的集合, 往往会有连续的成员, 比如, 一个集合: `{1,2,3,4,5,8}`, 这里面 `{1,2,3,4,5}` 就是连续的成员, 我们可以通过记录开始的值和连续的个数来表示这些值, 表示为 `runstart: 1, runlen: 4`, 这样, 我们就可以从存储 6 个 2 bytes (即: `{1,2,3,4,5,8}`) 值变为存储 4 个 2 bytes (即: `[{runstart: 1, runlen: 4}, {runstart: 8, runlen: 0}]`) 的值, 从而实现压缩(节省空间)

#### bitmapContainer 如何存储数据

bitmapContainer 本质是一个 65536 个 bit 的大二进制(blob). 但是, 在 golang 中, 整数最大的 bit 是 64 位(即: uint64). 因为 bitmap 是通过一个 bit 置 1 来表示一个数是否存在(比如集合 [1,2,3] 用 bitmap 表示就是 00001110, 这里, 第一个 bit 位表示 0, 第 2 个 bit 表示 1, 第 3 个 bit 表示 2, 第 4 个 bit 表示 3, and so forth) 为了表示这个 blob, 我们将整数的值按照 64 为一份进行平分(因为 golang 中最大的 uint 类型只能有 64 个 bit). 总共有 65536 个 bit 构成一个 bitmap, 因此, 极限情况下, 需要这个数组有 1024 个 uint64 成员, 表示这 blob. 因此, bitmapContainer 实际上就是一个可能拥有 1024 个 uint64 成员的数组

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

#### Roaring Bitmaps 如何实现 bitmap 的压缩

Roaring Bitmaps 实现压缩, 有很大一部分是由于 runContainer. 在调用 `runOptimize()` 时, 会判断三种数据结构所表示的容器所需要的空间, 根据最小的空间占用, 从而决定要用哪种容器, 最终达到压缩的效果

#### bitmapContainer 到 runContainer 转换时, 如何计算 runContainer 需要的 interval 数组的成员的个数

```go
func (bc *bitmapContainer) numberOfRuns() int {
	if bc.cardinality == 0 {
		return 0
	}

	var numRuns uint64
	nextWord := bc.bitmap[0]

	for i := 0; i < len(bc.bitmap)-1; i++ {
		word := nextWord
		nextWord = bc.bitmap[i+1]
		numRuns += popcount((^word)&(word<<1)) + ((word >> 63) &^ nextWord)
	}

	word := nextWord
	numRuns += popcount((^word) & (word << 1))
	if (word & 0x8000000000000000) != 0 {
		numRuns++
	}

	return int(numRuns)
}
```

思路是通过位运算 `popcount((^word)&(word<<1)) + ((word >> 63) &^ nextWord)`, 前半部分: `popcount((^word)&(word<<1))` 是将左边界为0时的情况通过二进制表示出来, 左边界为 0 时置 1, 然后我们计算位运算后的 1 的个数, 可以得到左边界是 0 的情况的个数. 后半部分 `(word >> 63) &^ nextWord`, 由于左移了一位, 因此需要单独考虑最高位的情况

最高位为 1 时有两总情况:

第一是, 下一个片段(由于 bitmap 有 65536 位, 因此我们需要进行分片段, 每个片段是 uint64, 即 64 位) 的最低位为 1, 那么当前片段可以忽略最高位, 因为下一个片段会将其作为连续 1 的一部分.

第二是, 下一个片段的最低位为 0, 那么这时候我们就需要在这一个片段计算个数时 +1 了. 因为最高位的 1 为独立的 1 或者连续的 1 的片段, 将作为一个独立的 `interval16` 存入到数组

由于我们是计算左边界为 0. 因此 bitmapContainer 的 bitmap 数组的最后一个成员需要单独分析. 这是因为, 我们再做循环时, 需要考虑数组下一个成员可能对数组最高位的影响(上一行说明的问题). 因为 bitmap 数组的最后一个成员, 是没有下一个成员的, 所以就需要单独拎出来考虑了

bitmap 数组最后一个成员的判断逻辑如下, 计算左边界为 0 的次数, 然后判断最高位是否为 1. 如果最高位为 1, 那么, 我们需要在左边界为 0 的次数上 +1, 否则, 结果就是左边界为 0 时的次数

#### 如何将 bitmapContainer 转换为 runContainer

```go
func newRunContainer16FromBitmapContainer(bc *bitmapContainer) *runContainer16 {

	rc := &runContainer16{}
	nbrRuns := bc.numberOfRuns()
	if nbrRuns == 0 {
		return rc
	}
	rc.iv = make([]interval16, nbrRuns)

	longCtr := 0            // index of current long in bitmap
	curWord := bc.bitmap[0] // its value
	runCount := 0
	for {
		// potentially multiword advance to first 1 bit
		for curWord == 0 && longCtr < len(bc.bitmap)-1 {
			longCtr++
			curWord = bc.bitmap[longCtr]
		}

		if curWord == 0 {
			// wrap up, no more runs
			return rc
		}
		localRunStart := countTrailingZeros(curWord)
		runStart := localRunStart + 64*longCtr
		// stuff 1s into number's LSBs
		curWordWith1s := curWord | (curWord - 1)

		// find the next 0, potentially in a later word
		runEnd := 0
		for curWordWith1s == maxWord && longCtr < len(bc.bitmap)-1 {
			longCtr++
			curWordWith1s = bc.bitmap[longCtr]
		}

		if curWordWith1s == maxWord {
			// a final unterminated run of 1s
			runEnd = wordSizeInBits + longCtr*64
			rc.iv[runCount].start = uint16(runStart)
			rc.iv[runCount].length = uint16(runEnd) - uint16(runStart) - 1
			return rc
		}
		localRunEnd := countTrailingZeros(^curWordWith1s)
		runEnd = localRunEnd + longCtr*64
		rc.iv[runCount].start = uint16(runStart)
		rc.iv[runCount].length = uint16(runEnd) - 1 - uint16(runStart)
		runCount++
		// now, zero out everything right of runEnd.
		curWord = curWordWith1s & (curWordWith1s + 1)
		// We've lathered and rinsed, so repeat...
	}

}
```

遍历 bitmap 数组, 如果 bitmap 数组中所有的成员都等于 0, 直接退出, 否则, longCtr 将等于该非 0 成员在 bitmap 数组中的索引, 该成员的值存放到 curWord

通过计算 curWord 二进制形式末尾的 0 个数, 确定 start 的值( start 值: 索引 + longCtr * 64, 因为我们是按照 64 为一个分片存放到 bitmap 数组的)

注意: 这一步很重要. 找到 start 值后, 我们将 curWord 二进制形式末尾的 0 全部置为 1 `curWordWith1s := curWord | (curWord - 1)`

判断 bitmap 数组中全为 1 的场景, 否则, longCtr 将等于下一个含 0 的 bitmap 数组成员的索引

注意: 这一步很重要. 将 longCtr 索引对应的 bitmap 数组成员取反 `localRunEnd := countTrailingZeros(^curWordWith1s)`, 这样经过末尾置 1(上一个注意的内容) 和取反(假设有一段尾部片段:01101100, 末尾置 1 得到: 01101111, 取反得到: 10010000, 这样, 计算末尾的 0 的个数为 4, 所以: 01101100 中, 从右往左第三个 0 的索引是 4, 从右往左第一对 11 的 Big End 的 1 的索引是 4 - 1, 为 3), 我们可以通过计算 longCtr 作为索引对应的 bitmap 数组成员末尾所有的 0 的个数, 得到 curWord 中 Big End 的 0 的索引 runEnd, `runEnd - 1` 即为连续的 1 的块对应的 Big End 的这个 1 对应的索引

注意: 这一步很重要. 由于我们在末尾置 1 操作, 现在, 我们可以将所有末尾连续的 1 置零: `curWord = curWordWith1s & (curWordWith1s + 1)` 进入下一个循环

#### 两个 bitmapContainer 求 And 后将结果存放在 arrayContainer 逻辑实现

bitmapContainer 的本质是 65536 个 bit 的 blob, 因为计算机最大的类型所能使用的 bit 是 64 位, 因此, 我们需要一个数组来表示这个 blob. 因此, bitmapContainer 是一个由 uint64 类型整数组成的包含 1024 个成员的整数数组, 数组成员默认值是 0. 假设数组成员如下 [10, 15, 20 ...], 以第一个数组成员为例. 表示成 64 位二进制是 `0...00001010`(我省略了部分0), 我们可以知道, 如果用 arrayContainer 来表示的话, 就是 `[1,3]` (因为二进制中, 从右到左, 总共有两个 1,  第一个 1 的索引是 1, 第二个 1 的索引是 3, 因此, bitmapContainer 中的第一个成员 10 表示的集合就是 `{1,3}`).

所以, 这里的算法问题是: 已知两个数组分别叫做 bitmap1 bitmap2, 数组成员是 uint64 整数, 我们需要分别将数组内所有的 uint64 成员的二进制拼接起来, 成为一个 65536 的 blob. 然后对这两个 blob 做 `&` 运算, 然后得到新的 blob, 新的 blob 我们叫它 bitmap3, bitmap3 也有 65536 位, 我们需要从右到左, 记录所有 1 的索引. 假设 bitmap3 中第一个 uint64 数字是 10, 其二进制表示 `0...1010`, 从右到左, 将所有的 1 的索引记录到一个新的数组内, 即为 `[1,3]` 如何实现?

思路如下:

```go
pos := 0
for k := 0; k < len(bitmap1); k++ {
  bitset := bitmap1[k] ^ bitmap2[k]
  for bitset != 0 {
    t := bitset & -bitset
    container[pos] = uint16((k*64 + int(popcount(t-1))))
    pos = pos + 1
    bitset ^= t
  }
}
```

我们先遍历 bitmapContainer 数组, 拿到整数 10. 然后, 求整数 10 和整数 -10 的交集(这是因为: 一个数的负数等于这个数取反+1), 可以求得, 整数 10 二进制下从最低位开始, 为 1的二进制位表示的整数 t(这里 `10 & -10 = 2`, 所以 t == 2, 2 的二进制是 `00000010`, 对照整数 10 的二进制, `00001010` 可以知道, 这一步的作用是求整数 10 的二进制从最小位开始 1 所表示的整数. 整数 2 的二进制表示的正是整数 10 的二进制中, 最小位开始的 1 所在的二进制位 ), 整数 t - 1, 将把整数 t 在二进制下 1 所在的位置 0, 1 后面的位全部置 1, 即 `00000001`(举一个很易懂的例子, 假设 t == 8, 二进制表示是 `00001000`, `t - 1 == 7`, 7 的二进制是 `00000111`, 通过计算 1 的个数为 3, 我们可以知道 t == 8 的时候, t 在二进制下, 从最小位开始, 1 出现时的索引是 3, 也就是 7 在二进制下所有的 1 的个数). 然后将其值存放到数组内并让索引自增. 最后, 我们需要将整数 10 中第一位出现的 1 置 0, 这一步可以通过将整数 10 和 t 求异或(Xor)实现. 然后我们进入下一个循环

#### arrayContainer 和 bitmapContainer 求 And 后将结果存放在 arrayContainer 实现

```go
func (bc *bitmapContainer) andArray(value2 *arrayContainer) *arrayContainer {
  answer := newArrayContainerCapacity(len(value2.content))
  answer.content = answer.content[:cap(answer.content)]
  c := value2.getCardinality()
  pos := 0
  for k := 0; k < c; k++ {
    v := value2.content[k]
    answer.content[pos] = v
    pos += int(bc.bitValue(v))
  }
  answer.content = answer.content[:pos]
  return answer
}
```

这个实现很有意思. 因为 bitmapContainer 默认会比 arrayContainer 的成员要多. 所以, bitmapContainer & arrayContainer 最终会等与 arrayContainer(因为如果一个 arrayContainer 如果成员不小于 4096时会转换为 bitmapContainer)

实现的思路是:
1. 我们遍历 arrayContainer. 得到每次遍历的成员 v
2. 将 v 的值存到目标的数组 answer.content 内, 对应的索引是 pos. pos 默认为 0
3. 将成员通过位运算, 确认对应 bitmap 的二进制位上是否为 1. 如果为 1, 返回 1. 如果不为 1, 返回 0. **将返回值累加到 pos 上**
4. 所以有两种情况
  4.1. 此时 answer.content[pos] 存放的值是无效值. 返回值为 0, 此时 pos 保持不变. 此时这个无效值需要处理
    4.1.1. 如果存在下一个循环(即 k < c 仍成立), 进入下一个循环. 由于 pos 保持不变, 因此 answer.content[pos] 的值会被覆盖
    4.1.2. 由于不存在下一个循环, 此时 answer.content[pos] 需要被删除. 因此有 `answer.content = answer.content[:pos]`
  4.2. 此时 answer.content[pos] 存放的值是有效值. 返回值为 1, 此时 pos 累加了 1
    4.2.1. 如果存在下一个循环(即 k < c 仍成立), 进入下一个循环. 由于 pos 累加了 1, 因此不存在覆盖的情况
    4.2.2. 如果不存在下一个循环, 此时, 如果 answer.content[pos] 是有效值, pos 累加 1, 退出循环, `answer.content = answer.content[:pos]` 不会删除有效值

#### runContainer 和 bitmapContainer 求 And 实现

```go
func (rc *runContainer16) andBitmapContainer(bc *bitmapContainer) container {
	bc2 := newBitmapContainerFromRun(rc)
	return bc2.andBitmap(bc)
}
```

- 先将 runContainer 转换为 bitmapContainer 并赋值给 bc2
- 调用 bc2.andBitmap() 方法求值. 从这一步开始演变成为 bitmapContainer 和 bitmapContainer 的 And 操作, 见上文

#### 给 bitmapContainer 设置连续的 1 的实现

```go
func setBitmapRange(bitmap []uint64, start int, end int) {
  if start >= end {
    return
  }
  firstword := start / 64
  endword := (end - 1) / 64
  if firstword == endword {
    bitmap[firstword] |= (^uint64(0) << uint(start%64)) & (^uint64(0) >> (uint(-end) % 64))
    return
  }
  bitmap[firstword] |= ^uint64(0) << uint(start%64)
  for i := firstword + 1; i < endword; i++ {
    bitmap[i] = ^uint64(0)
  }
  bitmap[endword] |= ^uint64(0) >> (uint(-end) % 64)
}
```

setBitmapRange() 函数的实现非常有意思. 因为这里使用了位运算的技巧. 这个函数用于将 interval 表示的连续的 1 存放到 bitmap 数组内, 传入开始值 start 和结束值 end `end = last+1`(结束值是最后一个值加 1)
- 特例判断. 如果 start >= end. 说明出错了. 不修改 bitmap 数组
- 如果 interval 表示的范围在同一个 bitmap 数组成员表示二进制范围内
  - 通过 `^uint64(0) << uint(start%64)` 实现二进制右到左置零
  - 通过 `^uint64(0) >> (uint(-end) % 64)` 实现二进制左到右置零
  - 求 `&` 实现在同一个 bitmap 数组成员下,  快速写入 start 值到 last 值的二进制(注意: 需要给结束值 last 累加 1 作为 end 值)
- 否则, 需要多个 bitmap 数组成员来表示这些值.
  - bitmap 数组在 firstword 索引下的值是从右到左置 0
  - 我们可能有 0 到 n 个连续的全为 1 的 bitmap 数组成员. 其索引是 firstword + 1
  - bitmap 数组在 endword 索引下的值是从左到右置 0

#### 求插入索引逻辑实现
假设有整数 a, 整数数组 b. b 有 length 个数组成员. 已知 a > b[pos], 求 a 作为 b 中的成员时, a 在 b 数组中的索引(要利用已知的 a > b[pos], 直接排除掉 b[:pos+1] 表示的整数)

完整的函数如下:

```go
func advanceUntil(array []uint16, pos int, length int, min uint16) int {

  lower := pos + 1

	if lower >= length || array[lower] >= min {
		return lower
	}

	spansize := 1

	for lower+spansize < length && array[lower+spansize] < min {
		spansize *= 2
	}
	var upper int
	if lower+spansize < length {
		upper = lower + spansize
	} else {
		upper = length - 1
	}

	if array[upper] == min {
		return upper
	}

	if array[upper] < min {
    // means array has no item >= min pos = array.length
		return length
	}

	// we know that the next-smallest span was too small
	lower += (spansize >> 1)

	mid := 0
	for lower+1 != upper {
		mid = (lower + upper) >> 1
		if array[mid] == min {
			return mid
		} else if array[mid] < min {
			lower = mid
		} else {
			upper = mid
		}
	}
	return upper

}
```

- 几种极限情况分析
  - b[pos+1] 已经是最后一个 b 数组的成员, 且 b[pos+1] >= a. 说明 a > b[pos] 且 a <= b[pos+1]. a 的索引是 pos+1
  - 通过 spansize 的 2 倍扩增(spansize *= 2; n += spansize), 我们快速寻找使 b[n] >= min 的 n 的值
    - 如果 b[n] = min. 则索引是 n 位
    - 如果 b[n] < min. 则索引是数组的长度 length (这意味着要将 a 作为最后一个成员加到数组 b)
    - 如果 b[n] > min. 这说明 a 的索引位于 pos + 1 + (spansize >> 1) 和 pos + 1 + spansize
      - 对这种场景. 我们对这个范围使用二分搜索. 最低的索引 `lower = pos + 1 + (spansize >> 1)`. 最高的索引是 `upper = pos + 1 + spansize`. 代码如下

      ```go
      mid := 0
      for lower+1 != upper {
        mid = (lower + upper) >> 1
        if array[mid] == min {
          return mid
        } else if array[mid] < min {
          lower = mid
        } else {
          upper = mid
        }
      }
      return upper
      ```

      - 通过二分搜索. 可以确定最终 a 在 b 中的索引

#### runContainer 和 arrayContainer 求 And 实现

代码实现如下:
```go
func (rc *runContainer16) andArray(ac *arrayContainer) container {
	if len(rc.iv) == 0 {
		return newArrayContainer()
	}

	acCardinality := ac.getCardinality()
	c := newArrayContainerCapacity(acCardinality)

	for rlePos, arrayPos := 0, 0; arrayPos < acCardinality; {
		iv := rc.iv[rlePos]
		arrayVal := ac.content[arrayPos]

		for iv.last() < arrayVal {
			rlePos++
			if rlePos == len(rc.iv) {
				return c
			}
			iv = rc.iv[rlePos]
		}

		if iv.start > arrayVal {
			arrayPos = advanceUntil(ac.content, arrayPos, len(ac.content), iv.start)
		} else {
			c.content = append(c.content, arrayVal)
			arrayPos++
		}
	}
	return c
}
```

- 通过两个循环. 外部循环遍历 arrayContainer 内的 content 数组
- 拿到初始化的 iv (表示 runContainer 数组的成员) 和 arrayVal (表示 arrayContainer 数组的成员).
- 进入内部循环, 遍历 runContainer 内的 iv 数组
  - 如果 arrayVal 比当前的 iv 的 last 成员的值大, rlePos++, 我们查询下一个 iv.
  - 如果 arrayVal 比任何的 iv 中表示的 last 成员的值还要大, 说明没有重合的部分, 退出函数
  - 否则结束循环, 此时 arrayVal 小于等于当前 iv 的 last 表示的值
- 如果 iv.start > arrayVal, 说明我们的 arrayVal 偏小了. 我们要找到 arrayContainer.content 数组中, 使 iv.start <= arrayVal 的 arrayPos
  - 如果 arrayPos = acCardinality 说明 arrayContainer 中所有的成员都不满足 arrayVal >= iv.start. 自然, 这种情况下, arrayVal 也就不会出现在任何 runContainer.iv 的数组成员内了. 直接返回结果的 arrayContainer c, 退出函数
  - 否则, arrayVal >= iv.start && arrayPos < acCardinality. 由于内部循环确保了 arrayVal <= iv.last(). 所以 arrayVal 在 runContainer 中. 将 arrayVal 加入到 arrayContainer c 中. arrayPos++. 进入下一个循环
  - 否则, 返回 arrayContainer. 退出函数

#### runContainer 和 runContainer 求 and 逻辑实现

这个逻辑的实现过程在 runcontainer.go `func (rc *runContainer16) intersec(b *runContainer16) *runContainer16` 方法内. 这个函数的实现, 由于性能的考虑, 没有采用先转为 bitmapContainer 再比较, 而是直接通过遍历所有 runContainer.iv[] 数组中的成员, 然后依次比较这些成员. 因此, 逻辑比较复杂

下面是一些变量的含义:
- a 和 b 来表示两个不同的 runContainer
- numa 和 numb 表示 a.iv[] 和 b.iv[] 数组的长度
- res 是要返回的 *runContainer16 结构体指针
- output 将赋值给 res.iv, 因此 output 是返回的 res.iv[] 数组
- acuri 和 bcuri 是 a.iv[] 和 b.iv[] 数组的索引, 用于循环遍历 a.iv[] 和 b.iv[] 的时候, 数组内的成员是 interval16 结构体
- astart 和 bstart 分别表示 a.iv[] 在 acuri 索引下的 interval16 结构体内的 start 字段的值. bstart 同理
- leftoverstart 表示, 当两个 interval16 结构体表示的数据有重合时, 非重合部分最开始的数据的值
- intersection 是一个 interval16 类型值, 表示两个 interval16 重合部分
- isOverlap 表示两个 interval16 是否重合
- isLeftoverA 表示两个 interval16 重合时, 来自 a 的 interval16 表示的范围数据有剩余(也就是说, b 的 interval16 表示的范围的数据被包含在 a 的 interval16 表示的范围内). isLeftoverB 同理
- done 表示 a.iv[acuri] 和 b.iv[acuri] 没有重合情况下(也就是 `!isOverlap` 情况下, 这很重要), 此时假设 astart < bstart, 我们需要找到下一个 a.iv[x] 成员, 使 x.start 值 < bstart 且 a.iv[x+1].start > bstart.(也就说, a.iv[] 中下一个成员一定存在)
  - 如果 acuri+1 已经是最后一个数组成员的索引, 这意味着已经不存在 x. 因为我们之前已经假设 a.iv[acuri] 和 b.iv[acuri] 不重复, 因此 done 会被赋值 true, 表示 a.iv[] 不存在任何新的 interval16 可以与 b.iv[bcurri] 做 intersect 计算
  - 如果 acuri+1 不是最后一个数组成员的索引, done 等于 false, 因为我们还有其他的 a.iv[] 成员可以继续分析. 然后, 通过 a.search() 函数找到的索引 w < acuri + 1. 我们通过 a.search() 找到的索引 w 要包含的情况有两种, 是:
    - w 索引对应的值和 b.iv[bcuri] 有重合
    - w 索引对应的值和 b.iv[bcuri] 无重合, 但是 w+1 索引对应的值和 b.iv[bcuri] 有重合. 此时 w+1 索引便是我们需要分析下一个循环的 acuri
- a.search() 方法, 返回最小的 a.iv[] 数组成员索引 w. 使 `a.iv[w].start <= b.iv[bcuri].start < a.iv[w+1].start`, 由于调用 a.search() 的前提条件是 `!isOverlap` 即: `a.iv[w]` 和 `b.iv[bcuri]` 表示的数据集是没有重复的, 所以: `a.iv[w].last < b.iv[bcuri].start < a.iv[w+1].start`

这个方法实现的逻辑全部代码加起来有 250+ lines. 不做展开. 其中的逻辑是: 用一个循环同时遍历 a.iv 和 b.iv. 然后比对 a.iv 和 b.iv 中的重合部分. 如果 a.iv[acuri] 和 b.iv[acuri] 中有重合, 记录重合和剩余未匹配的部分. 如果 a.iv[acuri] 和 b.iv[acuri] 中没有重合. astart < bstart, 那么我们就要去找下一个 a.iv 中的成员 w, 使 w.start 可以 >= bstart, 然后继续下一个循环判断有没有重合. bstart < astart 同理. 直到我们遍历完 a.iv 或者 b.iv 数组其中之一, 打破循环. 其中 a.search() 方法的作用已说明(用二分法确定最近索引). 可以节省一定的代码阅读的时间. 以供参考

一句码外话, 这是我接触的代码里面. 逻辑比较复杂一段. 我第一次勉强看完, 心里难免会觉得难受. 因为这个逻辑之缜密和细致, 是我短时间内无法做到的. 因此难免心里有落差. 但是, 想跟后来人说一声, 初级开发和高级开发的区别. 就是在写代码的时候, 能否严谨的处理所有的可能性(经验). 初级到高级是一个过程, 代码的细致化也是一个过程. 所以, 没必要感到失望. 这是时间的产物, 通过长时间的社区的协作, 逐渐改进的结果. 拥抱开源社区吧.

#### 参考
- [Lemire's paper](https://arxiv.org/pdf/1402.6407.pdf)
- [高效压缩位图RoaringBitmap的原理与应用](https://www.jianshu.com/p/818ac4e90daf)
- [Roaring Bitmap更好的位图压缩算法](http://smartsi.club/better-bitmap-performance-with-roaring-bitmaps.html)
