/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	results := make([][]int, 0)
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
		// 到下一层
		combineCore(curN + 1)
		// 清理本层
		temp = temp[:len(temp)-1]
		combineCore(curN + 1)

	}
	combineCore(1)
	return results
}

// @lc code=end

