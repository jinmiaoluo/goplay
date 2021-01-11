# binary tree

这是关于二叉树的学习记录

#### 性质
- 二叉树每个节点最多有两个分支.
- 二叉树的深度从1开始. 即根节点的深度为1.
	- 第 k 层最多有:  <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2^{k-1}"> 个节点.
	- 假设深度是 i. 总共最多有: <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2^i-1"> 个节点.
- 满二叉树(Full Binary Tree): 每一层都填满所有节点.
	- 第 k 层的节点数是: <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2^{k-1}">
	- 深度为 i. 总共的节点数是: <img src="https://chart.googleapis.com/chart?cht=tx&chl=2^i-1" style="border:none;">
- 完全二叉树(Complete Binary Tree): 除最后一层, 其他层都是满的, 并且最后一层要么是满的，要么在右边缺少连续若干节点
	- 具有 n 个节点的完全二叉树, 其层数: <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=\log{_2}{n%2B1}">
	- 深度为 i. 总共的节点数:
		- 至少: <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2^{k-1}">
		- 至多: <img src="https://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2^k-1">
	- 遍历
    - 前序遍历:
    - 中序遍历:
    - 后序遍历:
	- 如何存储二叉树
		- 基于数组
      - 已知一个节点索引为 i:
        - 左边子节点索引: <img src="http://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2i%2b1">
        - 右边子节点索引: <img src="http://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=2i%2b2">
        - 父节点索引: <img src="http://chart.googleapis.com/chart?cht=tx&chco=000000&chf=bg,s,FFFFFF&chl=\mid \frac{i-1}{2} \mid">
		- 基于链表


#### 堆的特性
-

#### 参考
- [wikipedia: binary tree](https://en.wikipedia.org/wiki/Binary_tree)
- [golang 官方仓库内实现的二叉树比较](https://golang.org/doc/play/tree.go)
- [golang 官方仓库霍夫曼编码(Huffman coding)实现: HuffmanTree(二叉树在无损压缩领域的使用)](https://github.com/golang/go/blob/9eb219480e8de08d380ee052b7bff293856955f8/src/compress/bzip2/huffman.go#L15)
