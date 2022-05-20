package leetcode433

import "fmt"

var (
	genes = []byte{'A', 'C', 'G', 'T'}
)

func minMutation(start string, end string, bank []string) int {
	depth := map[string]int{}
	bankMap := map[string]bool{}
	for _, v := range bank {
		bankMap[v] = true
	}
	if !bankMap[end] {
		return -1
	}
	queue := []string{start}
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		for i := 0; i < 8; i++ {
			for j := 0; j < 4; j++ {
				if word[i] == genes[j] {
					continue
				}
				wordBytes := []byte(word)
				wordBytes[i] = genes[j]
				newWord := string(wordBytes)
				fmt.Println(newWord)
				if !bankMap[newWord] {
					continue
				}
				if _, exist := depth[newWord]; exist {
					continue
				}
				depth[newWord] = depth[word] + 1
				if newWord == end {
					return depth[newWord]
				}
				queue = append(queue, newWord)
			}
		}
	}
	return -1
}
