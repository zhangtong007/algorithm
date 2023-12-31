package leetcode167

import "sort"

func twoSum1(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}

// 双指针模板：一端固定，一端减小
func twoSum2(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

// 假设原数据不是排序数组，可以先对其进行排序
type Node struct {
	Idx int
	Val int
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	newNumbers := make([]*Node, n)
	for i, v := range numbers {
		newNumbers[i] = &Node{
			Idx: i,
			Val: v,
		}
	}
	// 对数组进行排序
	sort.Slice(newNumbers, func(i, j int) bool {
		return newNumbers[i].Val < newNumbers[j].Val
	})
	// 双指针扫描
	j := n - 1
	for i := 0; i < n; i++ {
		sum := newNumbers[i].Val + newNumbers[j].Val
		for i < j && sum > target {
			j--
		}
		if i < j && sum == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
