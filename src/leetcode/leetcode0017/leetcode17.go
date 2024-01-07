package leetcode17

var (
	mapping = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ret := []string{}
	dfs(&ret, make([]byte, len(digits)), digits, 0)
	return ret
}

func dfs(ret *[]string, strs []byte, digits string, pos int) {
	if pos == len(digits) {
		*ret = append(*ret, string(strs))
		return
	}

	for i := range mapping[digits[pos]] {
		strs[pos] = mapping[digits[pos]][i]
		dfs(ret, strs, digits, pos+1)
	}
}
