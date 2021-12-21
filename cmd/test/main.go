package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := [][]int{{}, {7}, {5}, {5, 7}, {3}, {3, 7}, {3, 5}, {3, 5, 7}, {0}, {0, 7}, {0, 5}, {0, 5, 7}, {0, 3}, {0, 3, 7}, {0, 3, 5}, {0, 3, 5, 7}, {9}, {9, 7}, {9, 5}, {9, 5, 7}, {9, 3}, {9, 3, 7}, {9, 3, 5}, {9, 3, 5, 7}, {9, 0}, {9, 0, 7}, {9, 0, 5}, {9, 0, 5, 7}, {9, 0, 3}, {9, 0, 3, 5}, {9, 0, 3, 5}, {9, 0, 3, 5, 7}}
	for i := range arr {
		sort.Slice(arr[i], func(m, n int) bool {
			return arr[i][m] < arr[i][n]
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	fmt.Println(arr)
	fmt.Println("==============================================================")
	arr2 := [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}
	for i := range arr2 {
		sort.Slice(arr2[i], func(m, n int) bool {
			return arr2[i][m] < arr2[i][n]
		})
	}
	sort.Slice(arr2, func(i, j int) bool {
		return len(arr2[i]) < len(arr2[j])
	})
	fmt.Println(arr2)

	map1 := convertToMap(arr)
	map2 := convertToMap(arr2)
	fmt.Println(equalMap(map1, map2))
}

func convertToMap(arrs [][]int) map[string]bool {
	ans := make(map[string]bool)
	for _, arr := range arrs {
		key := ""
		for _, v := range arr {
			key += strconv.Itoa(v)
		}
		if _, exist := ans[key]; exist {
			fmt.Printf("key:%s, arr:%v", key, arr)
		}
		ans[key] = true
	}
	return ans
}

func equalMap(m1, m2 map[string]bool) bool {
	if len(m1) != len(m2) {
		fmt.Println("len-m1:", len(m1), "len-m2:", len(m2))
		return false
	}
	for key := range m1 {
		if _, exist := m2[key]; !exist {
			fmt.Println("key:", key)
			return false
		}
	}
	return true
}
