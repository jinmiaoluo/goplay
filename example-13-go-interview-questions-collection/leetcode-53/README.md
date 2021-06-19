# LeetCode

Problem: 53

给定一个整数数组 nums, 找到一个具有最大和的连续子数组(子数组最少包含一个元素),
返回其最大和.

Example 1:

```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Example 2:

```
Input: nums = [1]
Output: 1
```

Example 3:

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Follow up**: If you have figured out the O(n) solution, try coding another
solution using the **divide and conquer approach**, which is more subtle.

思路1:

暴力破解的思路. 将所有可能的 subArray 全部都比较一遍.

思路2:

滑动窗口的思路. 这个窗口是只需要保留作为累加和的元素构成的子数组.
这里的滑动窗口有两种情况要考虑:

第一种是窗口是否扩容.

将窗口分割为左右两部分. 假设 nums 数组是 [1,3,4]. 左边是 nums[0] 右边是 nums[1],
因为 nums[0] + nums[1] >= nums[1]. 因为 nums[0] 是非负数, nums[1] 加上 nums[0],
只会让 nums[1] 变大.  所以我们可以认为, 窗口的左边部分是需要保留的.
此时就需要将窗口扩容. 扩容的操作, 就是将 nums[0]+nums[1]
相加的和作为新的窗口的左边的部分. nums[2]作为右边的部分, 再次判断.

第二种是窗口何时移动.

将窗口分割为左右两部分. 假设 nums 数组是 [-1,1]. 因为 nums[0] + nums[1] <
nums[1]. nums[1] 加上 nums[0], 只会让自己变小(因为 nums[0] 是负数).
所以我们可以认为, 窗口的左边部分是多余的了.
此时就需要将窗口移动到只有右边部分作为成员的情况.

需要两个变量, 一个记录最大值即: max. 一个记录最大区间和的值: 即 maxSum.

我们需要分为 3 个步骤:

第 1 是计算当前的区间和.

第 2 是判断当前的区间和和当前最大值哪个更大, 记录最新的最大值.

第 3 是判断区间和是否有保留的必要.

假设一个最简单的场景: [-1,1,1]

我们默认 nums[0] 是一开始最大值, 即 max = -1. 区间和 maxSum 默认为 0.

我们遍历 nums. 每次传入一个 nums 的元素.

区间是从只有一个元素开始的, 此时元素是 [-1] 时,
计算一下此时的区间和为 maxSum + (-1) = -1.  因为最大值大于等于区间和,
所以最大值不变. 因为任何数加上负数只会变得更小. 所以, 我们认为后面的 [-1,x]
这个区间是没有意义的, 直接跳到比较 x 和 maxSum 谁大谁小的问题就好了. 因此,
我们把 maxSum 置为 0.  然后来到下一个循环.

在下一个循环, 此时的元素是 1. 因为上一轮的区间和被我们置为 0 了, 所以此时的区间是 [1].
我们还是先计算区间和 maxSum 等于 maxSum + 1 等于 0 + 1 等于 1. 比较最大值,
发现新的区间和更大, 因此, 我们更新 max, 更新后 max 就等于 1 了. 因为区间和
maxSum 大于 0. 我们知道任何数加上一个正数会变得比自身更大. 所以,
我们这里的区间和就需要保留(而不是置为0). 此时 maxSum 等于 1.

在最后一个循环, 此时的元素是 1. 计算新的区间和等于 maxSum + 1 等于 2.
判断最大值可以知道, 当前新的区间和更大, 所以我们要更新最大值 max 为 2.
因为区间和 maxSum 大于 0. 所以, 我们这里的区间和就需要保留(而不是置为0).

什么是 **divide and conquer approach**: 中文指分而治之的方法,
我们常称之为分治算法.  分治算法是一个解决复杂问题的好工具,
它可以把问题分解成若干个子问题, 把子问题逐个解决, 再组合到一起形成大问题的答案.

```go
func maxSubArray(nums []int) int {
		max, subSum := nums[0], 0
		for _, v := range nums {
			// Calculate subSum
			subSum += v
			// If we should update `max' or not
			if subSum > max {
				max = subSum
			}
			// If `subSum' lower than 0, We should ignore this range
			if subSum < 0 {
				subSum = 0
			}
		}
		return max
}
```

思路3:

分治法. 我们从最小只有一个成员的区间开始. 通过遍历 nums 数组中其他成员,
不断的扩大区间的成员个数. 直到覆盖所有的 nums 中的成员.

在每次循环中, 我们会记录当前区间内的子区间最大和(这个结果不一定是对的),
存放到一个数组内, 假设这个数组叫 dp.  因为刚开始只有一个成员. 所以 dp[0] 等于
nums[0].  因为如果一个数加上负数会使当前的值变小. 所以我们有如下的关系: 如果
dp[i-1] > 0, 那么 dp[i] = dp[i-1] + nums[i]. 否则, 我们认为 dp[i] = nums[i].
这个规则成立是因为, 我们每次循环会在末尾判断 dp[i] 和 max 哪个是真正的最大值.
从而避免了小区间(最小的区间是只有一个数组元素)的和反而比大区间的和大的问题.

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, dp := nums[0], make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
			if dp[i-1] > 0 {
				dp[i] = nums[i] + dp[i-1]
			} else {
				dp[i] = nums[i]
			}
			if dp[i] > max {
				max = dp[i]
			}
	}
	return max
}
```
