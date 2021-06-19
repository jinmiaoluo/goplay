# LeetCode

Problem: 26

给定一个有序数组 nums, 对数组中的元素进行去重, 使得原数组中的每个元素只有一个.
最后返回去重以后数组的长度值.

Example 1:

```c
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums
being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

Example 2:

```c
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums
being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

思路:

比如 `[1,1,2]` 我们需要将重复的部分删除,
删除的做法是不改变数组所占用的内存的大小, 将重复的部分放到最末尾,
然后返回没有重复的部分的数量. 比如这里就需要返回 2 和经过调整位置的数组:
`[1,2,_]`

我们需要两个变量. 一个变量, 假设叫做 counter
用于表示无重复情况下的最后一个数组元素的索引. 另外一个变量, 假设叫做 last,
用于表示是否遍历完所有的数组元素了.

counter last 一开始都是 0, 在遍历过程中, 有两种可能.

第一种是有重复. 此时, sums[counter] 等于 sums[last + 1], 我们需要将 sums[last +
2] 的值放到 sums[last + 1] 的位置上, 这里我们可以这样做:

last 自增 1, 并让 `sums[last] = sums[last + 1]`

第二种是没有重复, 此时, sums[counter] 不等于 sums[last + 1], 说明 counter
需要加一个, 并且 sums[counter + 1] 这个位置要存放 sums[last + 1] 对应的值,
这里我们可以这样做:

counter 自增 1, 并让 `sums[counter] = sums[last + 1]`, last 自增 1

PS:

这个遍历的过程成立的条件是 last 小于 len(sums) - 1 而不是小于 len(sums),
因为我们使用了 sums[last + 1] 来读数组, 如果是小于 len(sums) 会有读索引溢出问题.

```go
func removeDuplicates(nums []int) int {
  if len(nums) == 0 {
    return 0
  }
  counter, last := 0, 0
    for last < len(nums) - 1 {
      if nums[counter] == nums[last + 1] {
        last++
        // cause we add 1 to `last' first, we have to check if `last + 1' will
        // out of range or not
        if last + 1 < len(nums) {
          nums[last] = nums[last + 1]
        }
      } else {
        counter++
        nums[counter] = nums[last + 1]
        last++
      }
    }
  return counter+1
}
```
