package main

import (
	leetcode232 "core/src/leetcode/0232"
)

const (
	Byte = 1
	KB   = Byte << 10
	MB   = KB << 10
	GB   = MB << 10
)

// main下面主要用于验证各个包下的demo
func main() {
	queue := leetcode232.Constructor()
	queue.Push(1)
	queue.Print()
	queue.Push(2)
	queue.Print()
	queue.Peek()
	queue.Print()
	queue.Pop()
	queue.Print()
	queue.Empty()
	queue.Print()
}
