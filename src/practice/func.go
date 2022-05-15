package practice

import "container/list"

type Node struct {
	Key int
	Val int
}

type LRUCache struct {
	list       *list.List
	elementMap map[int]*list.Element
	cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:       list.New(),
		elementMap: make(map[int]*list.Element),
		cap:        capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	element, exist := this.elementMap[key]
	if !exist {
		return -1
	}
	// 删除，且putFront
	node := element.Value.(*Node)
	delete(this.elementMap, node.Key)
	this.list.Remove(element)
	this.elementMap[node.Key] = this.list.PushFront(node)
	return element.Value.(*Node).Val
}

func (this *LRUCache) Put(key int, value int) {
	element, exist := this.elementMap[key]
	if exist {
		node := element.Value.(*Node)
		node.Val = value
		delete(this.elementMap, key)
		this.list.Remove(element)
		this.elementMap[key] = this.list.PushFront(node)
		return
	}
	if this.list.Len() == this.cap {
		node := this.list.Remove(this.list.Back()).(*Node)
		delete(this.elementMap, node.Key)
	}
	node := &Node{
		Key: key,
		Val: value,
	}
	this.elementMap[key] = this.list.PushFront(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
