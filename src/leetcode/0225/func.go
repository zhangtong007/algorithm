package leetcode225

import "container/list"

// 模拟一个queue 只允许pushBack, pop, size & isEmpty 方法
type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) pushBack(v int) {
	q.list.PushBack(v)
}

func (q *Queue) pop() int {
	return q.list.Remove(q.list.Front()).(int)
}

func (q *Queue) peek() int {
	return q.list.Front().Value.(int)
}

func (q *Queue) size() int {
	return q.list.Len()
}

func (q *Queue) isEmpty() bool {
	return q.list.Len() == 0
}

type MyStack struct {
	q1, q2 *Queue
}

func Constructor() MyStack {
	return MyStack{
		q1: NewQueue(),
		q2: NewQueue(),
	}
}

// 返回主队列和副队列
func (this *MyStack) mainAndBackup() (*Queue, *Queue) {
	if this.q1.isEmpty() {
		return this.q2, this.q1
	}
	return this.q1, this.q2
}

func (this *MyStack) Push(x int) {
	main, _ := this.mainAndBackup()
	main.pushBack(x)
}

func (this *MyStack) Pop() int {
	main, backup := this.mainAndBackup()
	for main.size() > 1 {
		backup.pushBack(main.pop())
	}
	return main.pop()
}

func (this *MyStack) Top() int {
	main, backup := this.mainAndBackup()
	for main.size() > 1 {
		backup.pushBack(main.pop())
	}
	v := main.pop()
	backup.pushBack(v)
	return v
}

func (this *MyStack) Empty() bool {
	return this.q1.isEmpty() && this.q2.isEmpty()
}
