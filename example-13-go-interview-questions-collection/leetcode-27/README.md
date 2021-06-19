# LeetCode

Problem: 27

给定一个数组 nums 和一个数值 val, 将数组中所有等于 val 的元素删除,
并返回剩余的元素个数.

Example 1:

```c
Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums
being 2.

It doesn't matter what you leave beyond the returned length.
```

Example 2:

```c
Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums
containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
```

思路:

定义一个变量 k 存放不包含 val 的元素个数. 遍历数组,
假设每次遍历数组得到的数组成员的值是 v.

有两种情况:

第一种. v == val, 此时, k 保持不变.

第二种, v != val, 此时, sums[k] = v, 然后 k 自增. 因为我们要返回的 k
是经过处理的数组的元素个数(而不是索引), 所以 k++ 放在后面是合理的.

```go
func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}
```
