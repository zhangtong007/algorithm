package leetcode1261

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type FindElements struct {
	vals map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	vals := make(map[int]bool)
	vals[0] = true
	root.Val = 0
	dfs(vals, root)
	return FindElements{
		vals: vals,
	}
}

func dfs(vals map[int]bool, node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		node.Left.Val = node.Val*2 + 1
		vals[node.Left.Val] = true
		dfs(vals, node.Left)
	}
	if node.Right != nil {
		node.Right.Val = node.Val*2 + 2
		vals[node.Right.Val] = true
		dfs(vals, node.Right)
	}
}

func (this *FindElements) Find(target int) bool {
	return this.vals[target]
}
