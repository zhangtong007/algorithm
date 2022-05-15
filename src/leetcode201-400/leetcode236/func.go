package leetcode236

// 二叉树两个节点的公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Ans struct {
	node *TreeNode
}

func (ans *Ans) IsEmpty() bool {
	return ans.node == nil
}

func (ans *Ans) Result() *TreeNode {
	return ans.node
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := new(Ans)
	dfs(root, p, q, ans)
	return ans.Result()
}

// dfs 返回值 表示是否找到p, q
func dfs(root, p, q *TreeNode, ans *Ans) (bool, bool) {
	if root == nil {
		return false, false
	}
	left_p, left_q := dfs(root.Left, p, q, ans)
	right_p, right_q := dfs(root.Right, p, q, ans)
	pFound, qFound := left_p || right_p || root.Val == p.Val, left_q || right_q || root.Val == q.Val
	if pFound && qFound && ans.IsEmpty() {
		ans.node = root
	}
	return pFound, qFound
}
