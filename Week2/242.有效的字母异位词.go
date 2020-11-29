/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	var l1 [26]int32
	for _, letter := range s {
		l1[letter-'a']++
	}
	for _, letter := range t {
		l1[letter-'a']--
	}
	for i := 0; i < 26; i++ {
		if l1[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

