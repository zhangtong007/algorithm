package leetcode0152

import "testing"

func Test_maxProduct(t *testing.T) {
	tasks := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{2, 3, -2, 4},
			expect: 6,
		},
		{
			nums:   []int{-2, 0, -1},
			expect: 0,
		},
	}
	for _, task := range tasks {
		actual := maxProduct(task.nums)
		if task.expect != actual {
			t.Errorf("expect: %d, but: %d\n", task.expect, actual)
		}
	}
}
