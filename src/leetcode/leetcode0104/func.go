package leetcode104

import . "core/src/datastruct"

// 方法1：广搜，看层数
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	layer := 0
	for len(queue) > 0 {
		layer++
		newQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return layer
}

// 方法2：深搜，记录最大层数
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	layer := 1
	depth := 1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// process
		layer = max(depth, layer)
		depth++
		dfs(root.Left)
		dfs(root.Right)
		// 还原现场
		depth--
	}
	dfs(root)
	return layer
}

// 方法3：递归
// 思想：树的最大深度，等于左子树和右子树的最大深度+1(根节点自身)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
