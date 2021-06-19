# LeetCode

Problem: 35

给定一个排序数组和一个目标值, 在数组中找到目标值, 并返回其索引.
如果目标值不存在于数组中, 返回它将会被按顺序插入的位置.

你可以假设数组中无重复元素.

**Example 1:**

```
Input: [1,3,5,6], 5
Output: 2
```

**Example 2:**

```
Input: [1,3,5,6], 2
Output: 1
```

**Example 3:**

```
Input: [1,3,5,6], 7
Output: 4
```

**Example 4:**

```
Input: [1,3,5,6], 0
Output: 0
```

思路:

假设要找的值是 target. 遍历所有数组成员，索引为 i, 值为 v. 有两种情况：

第一种, target 小于等于 v, 返回此时的索引 i.

第二种, target 大于 v. 此时说明当前的 v 不是我们要找的.
需要继续遍历新的数组成员, 直到结束时还没找到, 此时 target 要 append 到末尾,
因此返回数组长度.

```go
func searchInsert(nums []int, target int) int {
	length := len(nums)
	for i, v := range nums {
		if target <= v {
			return i
		}
		if i == length-1 {
			return length
		}
	}
	return 0
}
```
