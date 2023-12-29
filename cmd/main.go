package main

import (
	"fmt"
	"math"

	lcr103 "core/src/LCR/LCR103"
)

func main() {
	// leetcodeTest()
	fmt.Println(int(1e10) <= math.MaxInt32)
}

func leetcodeTest() {
	fmt.Println(lcr103.TestDemo(
		[]int{2},
		3,
	))
}
