package leetcode49

import (
	"sort"
)

// 字母只使用一次，可以计算二进制
func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	// 用来存储相同hash的数组在ans内的索引，如果不存在，ans里面append
	idxMap := make(map[string]int)
	for _, str := range strs {
		key := calHash(str)
		idx, exist := idxMap[key]
		if !exist {
			ans = append(ans, []string{str})
			idxMap[key] = len(ans) - 1
			continue
		}
		ans[idx] = append(ans[idx], str)
	}
	return ans
}

// 返回字符串的按字母升序排列
// abc -> hash = "abc"
// cab -> hash = "abc"
func calHash1(val string) string {
	bytes := []byte(val)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

// 统计字符出现个数作为hash
func calHash(val string) string {
	bytes := make([]byte, 26)
	for i := range val {
		bytes[val[i]-'a']++
	}
	return string(bytes)
}
