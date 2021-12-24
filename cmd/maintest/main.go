package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(50.0, -2147483648))
}

func test(a, b int) {
	fmt.Println(-1 << 1)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	/*
		go 语言 int 可以自动进行int32和int64转换，n=-2147483648 -n = 2147483648
		如果是 int32 或者 是其他语言的int，-n 溢出范围，会导致无法处理
		if n == -1<<31 {
			fmt.Println(-(n + 1))
			return 1.0 / myPow(x, -(n+1))
		}
	*/
	if n < 0 {
		fmt.Println(-n)
		return 1.0 / myPow(x, -n)
	}
	temp := myPow(x, n/2)
	ans := temp * temp
	if n&1 == 1 {
		ans *= x
	}
	return ans
}
