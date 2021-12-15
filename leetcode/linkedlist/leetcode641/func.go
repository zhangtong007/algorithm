package leetcode641

// 设计双端队列，通过链表方式设计
type Node struct {
	pre  *Node
	next *Node
	val  int
}

func newNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type MyCircularDeque struct {
	front *Node
	rear  *Node
	cap   int
	size  int
}

func Constructor(k int) MyCircularDeque {
	queue := MyCircularDeque{
		cap:   k,
		rear:  newNode(-1),
		front: newNode(-1),
	}
	queue.rear.pre = queue.front
	queue.front.next = queue.rear
	return queue
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := newNode(value)
	next := this.front.next
	this.front.next = node
	node.next = next
	next.pre = node
	node.pre = this.front
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := newNode(value)
	pre := this.rear.pre
	node.next = this.rear
	node.pre = pre
	pre.next = node
	this.rear.pre = node
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	node := this.front.next
	this.front.next = node.next
	node.next.pre = this.front
	node.next = nil
	node.pre = nil
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	node := this.rear.pre
	node.pre.next = this.rear
	this.rear.pre = node.pre
	node.next = nil
	node.pre = nil
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.front.next.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.rear.pre.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap
}

/*
// 双端队列，通过数组方式设计
type MyCircularDeque struct {
	size int
	cap  int
	vals []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		size: 0,
		cap:  k,
		vals: []int{},
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.vals = append([]int{value}, this.vals...)
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.vals = append(this.vals, value)
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.vals = this.vals[1:]
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.vals = this.vals[:this.size-1]
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.vals[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.vals[this.size-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cap == this.size
}
*/
