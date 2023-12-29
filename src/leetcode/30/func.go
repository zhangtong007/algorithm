package leetcode30

/*
可优化点：
当前设计是直接以matchLength暴力遍历每一个长度匹配的子串去判断，转化为map的时候不能复用
可考虑分组优化
*/
func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)
	// 每一个单词长度一致，总长度可以直接相乘
	wordLen := len(words[0])
	matchLength := wordLen * len(words)
	for _, word := range words {
		wordsMap[word]++
	}
	ans := []int{}
	// 每matchLenth一组去进行匹配
	for i := 0; i <= len(s)-matchLength; i++ {
		if match(s[i:i+matchLength], wordsMap, wordLen) {
			ans = append(ans, i)
		}
	}
	return ans
}

func match(str string, wordsMap map[string]int, wordLen int) bool {
	// 将str转化为一个待匹配map，与wordMap做比较即可
	// 由于单词长度一致，所以可以直接将str进行拆分
	strMap := make(map[string]int)
	for i := 0; i < len(str); i += wordLen {
		strMap[str[i:i+wordLen]]++
	}
	return mapEqual(strMap, wordsMap)
}

func mapEqual(strMap, wordsMap map[string]int) bool {
	// 如果key数量不等，则不相等
	if len(strMap) != len(wordsMap) {
		return false
	}
	// 遍历其中一个map，查询在另一个map中是否存在key，且value一致
	for k, v := range strMap {
		// 如果key在另一个map不存在，或者值不相等，则false
		if val, exist := wordsMap[k]; !exist || val != v {
			return false
		}
	}
	return true
}
