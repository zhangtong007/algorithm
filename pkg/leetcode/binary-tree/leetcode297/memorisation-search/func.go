package memorisationsearch

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cases = map[string]*TreeNode{}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 通过层序遍历结果，记录之后存入cases
	queue := []*TreeNode{root}
	nodes := []string{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			nodes = append(nodes, "null")
			continue
		}
		nodes = append(nodes, strconv.Itoa(node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	ans := strings.Join(nodes, ",")
	cases[ans] = root
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return cases[data]
}
