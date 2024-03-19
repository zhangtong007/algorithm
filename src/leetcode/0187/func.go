package leetcode187

// 哈希表
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}
	exist := map[string]bool{s[:10]: true}
	ans := map[string]struct{}{}
	for i := 11; i <= len(s); i++ {
		if exist[s[i-10:i]] {
			ans[s[i-10:i]] = struct{}{}
			continue
		}
		exist[s[i-10:i]] = true
	}
	ret := make([]string, 0, len(ans))
	for k := range ans {
		ret = append(ret, k)
	}
	return ret
}
