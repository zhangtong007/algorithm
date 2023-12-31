package leetcode255

import "math"

func verifyPreorder(preorder []int) bool {
	min := math.MinInt32
	stack := []int{}
	for i := 0; i < len(preorder); i++ {
		if preorder[i] < min {
			return false
		}
		for len(stack) != 0 && preorder[i] > stack[len(stack)-1] {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, preorder[i])
	}
	return true
}
