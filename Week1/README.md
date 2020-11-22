学习笔记

- 递归

  - 模板

  ```shell
  1. 退出条件
  2. 本层处理
  3. 传递到下层
  4. 如果需要，清理本层
  ```

  - 注意点

  ```shell
  1.空间：警惕堆栈溢出，代码中限制最大递归深度，超过就抛出错误
  2.时间：警惕重复计算，通过判断条件、缓存剪枝
  ```

- LRU Cache - Linked list： LRU 缓存机制（Least Recently Used 最近最少使用淘汰算法 ）

  - 缓存

  ```shell
  缓存是一种广义的概念，在计算机存储层次结构中，低一层的存储器都可以看做是高一层的缓存。比如Cache是内存的缓存，内存是硬盘的缓存，硬盘是网络的缓存（不一定网络就比磁盘访问慢）等等。
  缓存之所以有效还是因为程序和数据的局部性特点，程序的局部性指的是：
  1.一个程序在运行时被引用过一次的存储位置之后也很有可能被再次引用（时间局部性，对应程序中的循环）
  2.程序运行时，一个存储器位置被访问，它的附近位置也很有可能会被引用（空间局部性）
  所以这些将来会重复访问到的数据，就可以缓存起来，提示程序效率。
  ```

  - 缓存淘汰算法

  ```shell
  why？缓存资源有限，程序运行不可预期，势必得淘汰部分缓存内容
  what？LRU 是一种常用的缓存淘汰算法，最近最少使用的缓存内容
  how？题目：https://leetcode-cn.com/problems/lru-cache/
  ```

  

- Redis - Skip List：跳跃表、为啥 Redis 使用跳表（Skip List）而不是使用 Red-Black？

```shell
skip list 是redis针对zset（有序集合）使用的数据结构，使用它的原因如下：
1.实现简单。
2.符合功能需求，通过跳表去取得有序集合和通过平衡树的时间复杂度一致。
3.如果有序集合添加删除元素，通过跳表性能优于平衡树，索引到修改位置后，修改操作是常数级别的。
```



- Python 的 heapq

  堆是一个二叉树，它的每个父节点的值都只会小于或大于所有孩子节点（的值）

  源码解析：[Lib/heapq.py](https://github.com/python/cpython/tree/3.9/Lib/heapq.py)

```python
// TODO 这个模块提供了堆队列算法的实现，也称为优先队列算法。
```



- 高性能的 container 库

https://docs.python.org/2/library/collections.html

- 分析 Queue 和 Priority Queue 的源码

```go

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// 通过Init函数生成堆
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
    // 用于每个节点排序
		down(h, i, n)
	}
}

// 放入元素在堆尾，上移
// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

// 将堆顶元素和堆尾元素交换，抛出，将新堆顶元素下放排序
// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// 根据组合不同的sort.Interface，实现小顶堆和大顶堆

```

Queue 和 Priority Queue 都要自己实现

官网实例：https://golang.org/pkg/container/heap/#example__priorityQueue 