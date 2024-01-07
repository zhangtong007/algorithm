package main

import (
	"fmt"

	leetcode322 "algorithm/src/leetcode/0322"
)

const (
	Byte = 1
	KB   = Byte << 10
	MB   = KB << 10
	GB   = MB << 10
)

// main下面主要用于验证各个包下的demo
func main() {
	fmt.Println(
		leetcode322.Test(
			[]int{}, 1,
		),
	)
}
