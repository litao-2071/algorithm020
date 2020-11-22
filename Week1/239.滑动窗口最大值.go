/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

// time out
// 时间复杂度 O(kn)
func maxSlidingWindow2(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}
	for i := 0; i <= len(nums)-k; i++ {
		max := math.MinInt32
		for j := i; j < k+i; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k)
	window := make([]int, 0)

	for i, _ := range nums {
		// 左边界确定
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		// 确保window里的首位为最大值
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		// 添加元素如windows
		window = append(window, i)
		// 添加元素如结果集
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}

	return result
}

// @lc code=end

