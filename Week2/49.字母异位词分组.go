/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
// O(nklogk)
// @lc code=start
// func groupAnagrams(strs []string) [][]string {
// 	listMap := make(map[string][]string)
// 	for _, str := range strs {
// 		s := []byte(str)
// 		// sort.Strings impl func
// 		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
// 		key := string(s)
// 		if _, ok := listMap[key]; !ok {
// 			listMap[key] = []string{str}
// 			continue
// 		}
// 		listMap[key] = append(listMap[key], str)
// 	}
// 	res := make([][]string, 0, len(listMap))
// 	for _, list := range listMap {
// 		res = append(res, list)
// 	}
// 	return res
// }

// 时间复杂度：O(NK) 线性，k=26 计数分类
func groupAnagrams(strs []string) [][]string {
	listMap := make(map[[26]int][]string)
	for _, str := range strs {
		key := strArray(str)
		if _, ok := listMap[key]; !ok {
			listMap[key] = []string{str}
			continue
		}
		listMap[key] = append(listMap[key], str)
	}
	res := make([][]string, 0, len(listMap))
	for _, list := range listMap {
		res = append(res, list)
	}
	return res
}

func strArray(s string) [26]int {
	res := [26]int
	for _, v := range s {
	  res[v - 'a']++
	}
	return res
}
  
// @lc code=end

