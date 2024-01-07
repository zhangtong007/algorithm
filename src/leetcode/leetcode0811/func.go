package leetcode811

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	ansMap := make(map[string]int)
	for _, str := range cpdomains {
		strs := strings.Split(str, " ")
		num, _ := strconv.Atoi(strs[0])
		strs[1] = "." + strs[1]
		for i := 0; i < len(strs[1]); i++ {
			if strs[1][i] == '.' {
				ansMap[strs[1][i+1:]] += num
			}
		}
	}
	ans := make([]string, 0, len(ansMap))
	for key, v := range ansMap {
		val := strconv.Itoa(v) + " " + key
		ans = append(ans, val)
	}
	return ans
}
