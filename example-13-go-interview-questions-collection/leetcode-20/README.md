# LeetCode

Problem: 20

判断字符串中的符号是否成对. 比如: `(hello)[world]` 中 `(` `)` 以及 `[` `]`
就是成对出现的.

Example 1:

```
Input: "()"
Output: true

```


Example 2:

```
Input: "()[]{}"
Output: true

```

Example 3:

```
Input: "(]"
Output: false
```

Example 4:

```
Input: "([)]"
Output: false
```

Example 5:

```
Input: "{[]}"
Output: true
```

思路:

构建一个数组, 存放成对出现的符号的左部分, 如果出现右部分, 删除数组中的左边部分.
遍历完整个字符串后, 如果数组中没有剩余左边部分的字符, 即为都是成对出现的
