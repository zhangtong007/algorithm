package leetcode1457

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// ==========================================================
func pseudoPalindromicPaths(root *TreeNode) int {
	ans := 0
	// dfs 搜索路径，而后进行伪回文判断 统计每一个value出现次数，最后判断奇数次数<=1
	var dfs func(root *TreeNode, counts []int)
	dfs = func(root *TreeNode, counts []int) {
		// 当前递归结束时，恢复状态
		defer func() { counts[root.Val]-- }()
		if root.Left == nil && root.Right == nil {
			counts[root.Val]++
			if isPseudoPalindromicPath(counts) {
				ans++
			}
			return
		}
		counts[root.Val]++
		if root.Left != nil {
			dfs(root.Left, counts)
		}
		if root.Right != nil {
			dfs(root.Right, counts)
		}
	}
	dfs(root, make([]int, 10))
	return ans
}

func isPseudoPalindromicPath(counts []int) bool {
	odd := 0
	for _, v := range counts {
		if v&1 == 1 {
			odd++
		}
	}
	return odd <= 1
}
