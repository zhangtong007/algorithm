package leetcode232

import "fmt"

func (this *MyQueue) Print() {
	fmt.Println(*(this.s1), *(this.s2))
}

// 实现标准栈操作
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

// =======================================================

type MyQueue struct {
	s1, s2 *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: &Stack{},
		s2: &Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	opt, _ := this.check()
	opt.Push(x)
}

func (this *MyQueue) Pop() int {
	opt, backup := this.check()
	for !opt.isEmpty() {
		backup.Push(opt.Pop())
	}
	v := backup.Pop()
	// 不能改变原有顺序，将顺序倒置回去
	for !backup.isEmpty() {
		opt.Push(backup.Pop())
	}
	return v
}

func (this *MyQueue) Peek() int {
	opt, backup := this.check()
	for !opt.isEmpty() {
		backup.Push(opt.Pop())
	}
	v := backup.Peek()
	// 不能改变原有顺序，将顺序倒置回去
	for !backup.isEmpty() {
		opt.Push(backup.Pop())
	}
	return v
}

func (this *MyQueue) Empty() bool {
	return this.s1.isEmpty() && this.s2.isEmpty()
}

func (this *MyQueue) check() (*Stack, *Stack) {
	if this.s1.isEmpty() {
		return this.s2, this.s1
	}
	return this.s1, this.s2
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
