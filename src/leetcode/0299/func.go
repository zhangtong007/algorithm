package leetcode299

import "fmt"

// 翻译：
// bulls：位置和值均正确
// cows: 位置不正确但值正确
func getHint(secret string, guess string) string {
	// 直接遍历统计
	bulls := 0
	// 分别记录相同位置，值不相同的其他值出现的次数
	secrets, guesses := make([]int, 10), make([]int, 10)
	for i := range secret {
		if secret[i] == byte(guess[i]) {
			bulls++
			continue
		}
		secrets[secret[i]-'0']++
		guesses[guess[i]-'0']++
	}
	cows := 0
	for i := range secrets {
		// 相同字母出现次数，少的即是两个公共的
		cows += min(secrets[i], guesses[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
