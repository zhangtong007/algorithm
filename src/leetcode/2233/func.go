package leetcode2233

import (
	"container/heap"
)

func Execute() {
	maximumProduct([]int{3, 1, 6, 4, 3, 2}, 4)
}

// 贪心：每次对最小值+1
/*
证明：
	假设 非负整数a， b  a<b
	(a+1)*b = ab + b
	a*(b+1) = ab + a
	(a+1)*b > a*(b+1)
*/
// 贪心+小根堆

const mod = 1e9 + 7

type HeapImp []int

func (h HeapImp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HeapImp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HeapImp) Len() int {
	return len(h)
}

func (h *HeapImp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapImp) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func maximumProduct(nums []int, k int) int {
	h := HeapImp(nums)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		h[0]++
		heap.Fix(&h, 0)
	}
	ans := 1
	for _, v := range h {
		ans = ans * v % mod
	}
	return ans
}
