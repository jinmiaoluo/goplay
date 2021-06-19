# LeetCode

Problem: 66

给定一个由整数组成的非空数组所表示的非负整数, 在该数的基础上加一.
高位数字存放在数组的首位,  数组中每个元素只存储单个数字. 可以假设除了整数 0
之外, 这个整数不会以零开头.

Example 1:

```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

Example 2:

```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

Example 3:

```
Input: digits = [0]
Output: [1]
```

思路:

分为两种情况:

第一种, 数组最后一个成员加 1 后无进位. 在数组最后一位加 1 返回即可.

第二种, 数组最后一个成员加 1 后有进位. 从数组最后面那位元素开始遍历, 如果是进位,
我们就将当前循环对应数组元素置零, 继续循环. 否则, 循环对应数组元素加 1.
并返回, 如果结束循环后都没有返回, 说明刚好所有的位都是 9. 我们需要手动进位.
将第一位置为 1, 然后后面补 0.

```go
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		} else {
			digits[i] = 0
		}
	}
	// Carry manually
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
```
