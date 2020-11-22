/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// n > 0
// f(1) = 1 f(0) = 1
// f(2) = 2
// 1 + 1 = 2 1 + 2 = 3 2 + 3 = 5 = f(4)
func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

// @lc code=end

