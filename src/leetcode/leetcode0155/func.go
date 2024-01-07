package leetcode155

type MinStack struct {
	stack []int
	top   int
	// 记录前缀最小值
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0, 16),
		top:      -1,
		minStack: make([]int, 0, 16),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 如果栈为空，直接入栈最小值
	if this.top == -1 {
		this.minStack = append(this.minStack, val)
	} else {
		// 如果非空，与前缀最小值比较决定入栈谁
		this.minStack = append(this.minStack, min(this.minStack[this.top], val))
	}
	this.top++
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.minStack = this.minStack[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.stack[this.top]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.top]
}
