/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	// result := make([]int, len(digits)+1)
	// result[0] = 1
	// copy(result[1:], digits)

	// result := make([]int, len(digits)+1)
	// result[0] = 1

	digits = append(digits, 0)
	digits[0] = 1

	return result
}

// @lc code=end

