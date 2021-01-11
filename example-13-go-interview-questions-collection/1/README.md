# 问题
有一个大文件, 需要改一个词 a 为 b

## 关键点:
- 大文件读取(按段或者按行读取)
- 子串搜索和替换
- 如何处理子串搜索时, 按段或者按行读取刚好将待替换的词分割开, 导致替换不完整的特殊情况

## 如何判断文件大小?
```shell script
cd example/13/1/
go run demo0.go
```

## 如果是小文件如何解决?
- 小文件的读取, 一次性全部读入到内存
- 对 `[]byte` 进行子串的搜索和替换
> 下面的实现代码参考了: [Golang 超大文件读取的两个方案](https://learnku.com/articles/23559/two-schemes-for-reading-golang-super-large-files) 这篇文章

```shell script
cd example/13/1/
go run demo1.go
```

## 如果是大文件, 如何解决?

我们需要分段读取, 假设分段读取后的内容保存到 `[]byte` 类型 `frag` 变量中

特殊情况: **段切割时会刚好切割到了子串内部**

比如我们的子串是 `你好`, 表示成二进制是 `\xE4\xBD\xA0\xE5\xA5\xBD`. 我们在按大小进行分段读取 `[]byte` 类型的值的时候, 有可能出现下面的两个分段

第一段 `[]byte` 类型的值是:
```text
(省略开头的字节)...\xE4\xBD\xA0
```

第二段 `[]byte` 类型的值是:
```text
\xE5\xA5\xBD...(省略末尾的字节)
```
此时, `你好` 刚好就被分到两个不同的段里面了, 因此我们需要对这个特殊情况进行如下特殊处理:

假设我们的待修改的词是由 `targetByteNum` 个字节组成, 如果 `frag` 末尾 `targetByteNum-1` 个字节跟子串 a 有重叠, 记录重叠的字节个数 `overlay`, 我们需要再读取 `targetByteNum-overlay` 个新的字节. 将这些字节 `append` 到 `frag` 中, 然后再进行相同的判断, 确保 `frag` 末尾 `targetByteNum-1` 个字节跟子串没有重叠. 这是对第一种特殊情况的处理思路

然后对 `frag` 这段内容进行子串 a 的搜索和将 a 替换为 b 的子串替换操作, 将替换后得到的结果写入文件. 开始下一个循环, 处理下一段的数据. 直到 `io.EOF` 时退出循环

最终我们就得到经过处理的大文件

```shell script
cd example/13/1/
go run demo2.go
```

## 参考
- https://colobu.com/2016/10/12/go-file-operations/#%E4%BD%BF%E7%94%A8%E7%BC%93%E5%AD%98%E5%86%99
- https://www.devdungeon.com/content/working-files-go#write_bytes
