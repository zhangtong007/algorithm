package main

// hw机试：简单错误记录
// hash表+链表, 实现顺序淘汰
import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Key struct {
	path   string
	number string
}

type Node struct {
	key   Key
	count int
}

type Recorder struct {
	kvs      map[Key]*list.Element
	list     *list.List
	first    *list.Element
	capicity int
}

func NewRecorder(capicity int) *Recorder {
	return &Recorder{
		kvs:      make(map[Key]*list.Element),
		list:     list.New(),
		first:    nil,
		capicity: capicity,
	}
}

func (r *Recorder) update(node Node) {
	// 查找kv是否存在
	e, exist := r.kvs[node.key]
	// 存在，那么更新即可
	if exist {
		cur := e.Value.(Node)
		cur.count += node.count
		e.Value = cur
		return
	}
	/*
		此处题意理解错误，记录器里不对数据作淘汰，只是最后只返回最后出现的8条数据
		// 不存在，如果达到最大长度，那么先删除最早
		if r.list.Len() == r.capicity {
			delete(r.kvs, r.list.Remove(r.list.Front()).(Node).key)
		}
	*/
	// 插入
	r.kvs[node.key] = r.list.PushBack(node)
	// 更新记录first，用于作结果输出打印
	// 如果第一次插入，那么first记录为第一个元素
	if r.list.Len() == 1 {
		r.first = r.list.Front()
		return
	}
	// 如果长度已经超过capicity，其实就是8个。那么first后移，保证first到末尾最多8个
	if r.list.Len() > r.capicity {
		r.first = r.first.Next()
	}
}

func (r *Recorder) print() {
	e := r.first
	var node Node
	for e != nil {
		node = e.Value.(Node)
		fmt.Printf("%s %s %d\n", node.key.path, node.key.number, node.count)
		e = e.Next()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 最大长度8
	recorder := NewRecorder(8)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		// 处理得到path和number
		temp := strings.Split(input, " ")
		path, number := temp[0], temp[1]
		// 再对path处理
		temp = strings.Split(path, "\\")
		path = temp[len(temp)-1]
		if len(path) > 16 {
			path = path[len(path)-16:]
		}
		recorder.update(Node{
			key: Key{
				path:   path,
				number: number,
			},
			count: 1,
		})
	}
	// 打印
	recorder.print()
}
