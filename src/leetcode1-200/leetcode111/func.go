package leetcode111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：dfs
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minLayer := 100000
	depth := 1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			minLayer = min(minLayer, depth)
		}
		depth++
		dfs(root.Left)
		dfs(root.Right)
		// 还原现场
		depth--
	}
	dfs(root)
	return minLayer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法2：广搜：层序遍历  按层遍历，找到第一个左右子树均空的节点，返回当前深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	depth := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := []*TreeNode{}
		depth++
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return depth
}
