/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
// 深度优先遍历-递归
func generateParenthesis(n int) []string {
	var results []string

	// left， right 左右括号出现次数
	var dfs func(left, right int, res []byte)
	dfs = func(left, right int, res []byte) {
		// terminal
		if left > n || right > n || left < right {
			return
		}
		// current handle cut branch
		if len(res) == n+n {
			results = append(results, string(res))
			return
		}
		// drop down
		dfs(left+1, right, append(res, '('))
		dfs(left, right+1, append(res, ')'))
	}
	res := make([]byte, 0, n+n)
	dfs(0, 0, res)
	return results
}

// @lc code=end

// 广度优先遍历-queue

