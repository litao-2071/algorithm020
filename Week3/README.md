学习笔记

- 分治模板

与递归模版相似，只是多了一个合并子问题的步骤

```go
// Golang
func divide_conquer(problem, param1, param2, ...){
  // recursion terminator 
	if problem == nil {
  	print_result 
		return 
	}
  // prepare data 
	data := prepare_data(problem) 
	subproblems := split_problem(problem, data) 

  // conquer subproblems 
	subresult1 := divide_conquer(subproblems[0], p1, ...) 
	subresult2 := divide_conquer(subproblems[1], p1, ...) 
	subresult3 := divide_conquer(subproblems[2], p1, ...) 

  // process and generate the final result 
  result = process_result(subresult1, subresult2, subresult3, …)
	
  // revert the current level states
}
```

- 回溯模板

几乎与递归一样

```go
// Golang
func recursion(level, param1, param2, ...){
  // recursion   
  if level > MAX_LEVEL {
    process_result 
	  return 
  } 
	// process logic in current level 
  process(level, data...) 
  // drill down 
  self.recursion(level + 1, p1, ...) 
  // reverse the current level status if needed
} 
```
```java
//java
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

```



需要注意： 在递归中传递全局变量，下层的改动会影响到上层，所以必须要清理本层；如果使用的是局部变量就不需要清理

- 题目

https://leetcode-cn.com/problems/powx-n/

https://leetcode-cn.com/problems/subsets/

https://leetcode-cn.com/problems/majority-element/description/

https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

https://leetcode-cn.com/problems/n-queens/