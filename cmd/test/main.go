package main

import "fmt"

func main() {
	str := "aaabbb"
	a := str[2:2]
	fmt.Printf("t:%T, v:%v\n", a, a)
}
