package preorder

import (
	"strconv"
	"strings"

	. "core/src/datastruct"
)

// 先序遍历
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	nodes := []string{}
	serializeDfs(&nodes, root)
	return strings.Join(nodes, ",")
}

func serializeDfs(nodes *[]string, root *TreeNode) {
	if root == nil {
		*nodes = append(*nodes, "null")
		return
	}
	*nodes = append(*nodes, strconv.Itoa(root.Val))
	serializeDfs(nodes, root.Left)
	serializeDfs(nodes, root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	cur := -1
	root := deserializeDfs(nodes, &cur)
	return root
}

func deserializeDfs(nodes []string, cur *int) *TreeNode {
	*cur++
	if *cur >= len(nodes) {
		return nil
	}
	if nodes[*cur] == "null" {
		return nil
	}
	root := new(TreeNode)
	val, _ := strconv.Atoi(nodes[*cur])
	root.Val = val
	root.Left = deserializeDfs(nodes, cur)
	root.Right = deserializeDfs(nodes, cur)
	return root
}
