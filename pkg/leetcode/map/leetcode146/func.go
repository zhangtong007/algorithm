package leetcode146

import "fmt"

type LRUCache struct {
	list  *List
	cache map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:  NewList(capacity),
		cache: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.cache[key]
	if !exist {
		return -1
	}
	this.list.removeNode(node)
	this.list.putFront(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.cache[key]
	if exist {
		this.list.removeNode(node)
		node.val = value
		this.list.putFront(node)
		return
	}
	if this.list.isFull() {
		delete(this.cache, this.list.Tail.pre.key)
		this.list.removeNode(this.list.Tail.pre)
	}
	node = newNode(key, value)
	this.cache[key] = node
	this.list.putFront(node)
}

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

func newNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

type List struct {
	Head *Node
	Tail *Node
	Size int
	Cap  int
}

func NewList(cap int) *List {
	list := &List{
		Head: newNode(-1, -1),
		Tail: newNode(-1, -1),
		Size: 0,
		Cap:  cap,
	}
	list.Head.next = list.Tail
	list.Tail.pre = list.Head
	return list
}

func (this *List) putFront(node *Node) {
	if this.isFull() {
		return
	}
	node.pre = this.Head
	node.next = this.Head.next
	this.Head.next.pre = node
	this.Head.next = node
	this.Size++
}

func (this *List) removeNode(node *Node) {
	if this.isEmpty() {
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	this.Size--
}

func (this *List) isEmpty() bool {
	return this.Size == 0
}

func (this *List) isFull() bool {
	return this.Size == this.Cap
}

// 调试，用于打印当前缓存内容
// 如果程序中做了print，submit代码时记得删除
// print是O(N)复杂度，必然超时
func (this *LRUCache) print() {
	// 打印cashe
	fmt.Printf("cache:{")
	for key, node := range this.cache {
		fmt.Printf("[key:%d,val:%d]", key, node.val)
	}
	// 打印链表
	fmt.Printf("]}, list:")
	node := this.list.Head
	for i := 0; i < this.list.Size; i++ {
		node = node.next
		fmt.Printf("%d->", node.val)
	}
	fmt.Println("nil")
}
