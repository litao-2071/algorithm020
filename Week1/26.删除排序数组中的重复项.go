/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
// O(n) 将不重复的往前移动
func removeDuplicates(nums []int) int {
	pre := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[pre] {
			pre++
			nums[pre] = nums[i]
		}
	}
	return pre + 1
}

// @lc code=end

