package main

import (
	"fmt"
	"math/rand"

	quicksort "core/src/sort/quick-sort"
)

// main下面主要用于验证各个包下的demo
func main() {
	leetcodeTest()
	// fmt.Println(int(1e10) <= math.MaxInt32)
}

func leetcodeTest() {
	fmt.Println(quicksort.TestDemo(
		randomNums(1e6),
	))
}

func randomNums(num int) []int {
	nums := make([]int, num)
	for i := range nums {
		nums[i] = rand.Intn(num)
	}
	return nums
}
