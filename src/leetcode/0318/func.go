package leetcode318

func maxProduct(words []string) int {
	// 求解words对应的二进制 a -> 1, b -> 10
	bits := make([]int, len(words))
	for i, word := range words {
		bits[i] = toBit(word)
	}
	ans := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			// 如果i, j 不同，则计算
			if bits[i]&bits[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toBit(word string) int {
	ans := 0
	for i := range word {
		ans |= 1 << (word[i] - 'a')
	}
	return ans
}
