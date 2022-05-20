package leetcode98

import . "core/src/datastruct"

// 方法1：暴力破解
// 二叉搜索树中序遍历一定升序排列，先得到中序遍历数组，而后验证数组有序性
func isValidBST1(root *TreeNode) bool {
	vals := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		vals = append(vals, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	// 验证有序性
	for i := 1; i < len(vals); i++ {
		if vals[i-1] >= vals[i] {
			return false
		}
	}
	return true
}

// 方法2：递归
func isValidBST(root *TreeNode) bool {
	// 初始值足够大，保证数据范围能够包含所有元素
	return recur(root, 1<<32, -1<<32)
}

// 遍历节点时缩小范围，原则：左边的都比根节点小，右边的都比根节点大，作比较区间收缩
func recur(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		// 如果区间不符合，则不满足
		return false
	}
	// 作区间搜索，左子树小于根，右子树大于根, 如果有不满足则返回，作与
	return recur(root.Left, root.Val, min) && recur(root.Right, max, root.Val)
}
