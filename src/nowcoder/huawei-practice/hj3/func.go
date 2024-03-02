package main

// hw机试题：明明的随机数
// 思路：基数排序思想
import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)
	n, _ := strconv.Atoi(input)
	bases := make([]int, 501)
	var number int
	for i := 0; i < n; i++ {
		fmt.Scanln(&input)
		number, _ = strconv.Atoi(input)
		bases[number] = 1
	}
	// 打印结果
	for i, v := range bases {
		if v == 1 {
			fmt.Println(i)
		}
	}
}
