package main

import (
	"fmt"
)

const (
	Byte = 1
	KB   = Byte << 10
	MB   = KB << 10
	GB   = MB << 10
)

// main下面主要用于验证各个包下的demo
func main() {
	fmt.Println(toBit("a"))
	fmt.Println(toBit("abc"))
	fmt.Println(toBit("aaa"))
}

func toBit(word string) int {
	ans := 0
	for i := range word {
		ans |= 1 << (word[i] - 'a')
	}
	return ans
}
