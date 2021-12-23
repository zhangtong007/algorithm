package main

import "fmt"

func main() {
	test(1<<32, -1<<32)
}

func test(a, b int) {
	fmt.Println(a, b)
}
