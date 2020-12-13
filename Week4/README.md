# 深度优先遍历和广度优先遍历

> 对于树和图的遍历实际上就是把每个结点都访问一遍，且保证每个结点都只访问一次

## DFS 代码模版

### 递归

```go
visited := make(map[int]bool)
func dfs(root *GraphNode, visited map[int]bool) {
  // terminator
  if _, ok := visited[root.Val]; ok {
    return
  }
  // current handle
  visited[root.Val] = true // records nodes accessed
  // sink
  for _, node := range root.Children { // children nodes
    dfs(node, visited)
  }
  return
}
```

### 非递归

```go
// 非递归使用stack来模拟递归栈
```

## BFS代码模板

```go
func bfs(root *GraphNode) {
  if root == nil {
    return
  }
  visited := make(map[*GraphNode]bool)
  queue := make([]*GraphNode)
  queue = append(queue, root)
  visited[root] = true
  
  for ; len(queue) != 0; {
    node := queue[0]
    // handle the node do something
    queue = queue[1:]
    for _, node := range root.Children { // children nodes
      if _, ok := visited[node]; !ok {
        queue = append(queue, node)
      }
  		}
  }
  return
}
```

对于树来说不需要visited，它的遍历就是前中后序遍历和层次遍历，图需要来防止多次访问



## 贪心算法

> 贪心算法指得是在每一步都采用当前状态下最好的选择，期望最后能够得到全局的最优解。

贪心算法常用于解决一下最优问题：比如图中的最小生成树，求哈夫曼编码等，但在实际业务中、生活中很少用的到，因为一般局部的最优解都很难得到全局的最优解。

#### 贪心算法和动态规划的相似、区别

**相同：**它们都是对重复子问题的处理，得出最终结果

**区别：**贪心算法是直接通过局部子问题最优解获取最终问题最优解，无法回溯；而动态规划会记录每一步执行的结果，并根据这些结果来选择执行下一步，可以回溯。

##### 贪心算法一般很少用得到，但只要能够使用贪心算法解决的问题，那它一般都是最优解了。由于贪心法的高效性以及其所求得的答案比较接近最优结果，贪心法也可以用作辅助算法或者直接解决一些要求结果不特别精确的问题。

##### TODO
1.使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
说明：同学们可以将自己的思路、代码写在第 4 周的学习总结中
2.https://leetcode-cn.com/problems/minesweeper/solution/cong-qi-dian-kai-shi-dfs-bfs-bian-li-yi-bian-ji-ke/
3.https://leetcode-cn.com/problems/search-a-2d-matrix/
4.https://leetcode-cn.com/problems/jump-game/
