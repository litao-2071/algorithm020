/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
// 从1-i选出1个，从i-n选出一个
func combine(n int, k int) [][]int {
	res := [][]int{}
	var helper func(start int, path []int)
	helper = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			helper(i+1, path)
			path = path[:len(path)-1]
		}
	}
	helper(1, []int{})
	return res
}

func combine1(n int, k int) [][]int {
	var results [][]int
	temp := make([]int, 0, k)
	var combineCore func(int)
	combineCore = func(curN int) {
		// 剪枝
		if len(temp)+n-curN+1 < k {
			return
		}
		// 退出条件
		if len(temp) == k {
			// 当层处理
			res := make([]int, k)
			copy(res, temp)
			results = append(results, res)
			return
		}
		temp = append(temp, curN)
		// 选择当前位置
		combineCore(curN + 1)
		// 清理本层
		temp = temp[:len(temp)-1]
		// 不选择当前位置
		combineCore(curN + 1)

	}
	combineCore(1)
	return results
}

// @lc code=end

