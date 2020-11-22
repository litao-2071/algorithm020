/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := numMap[need]; ok {
			return []int{numMap[need], i}
		}
		numMap[nums[i]] = i
	}
	return nil
}

// @lc code=end

