package leetcode106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	mid := 0
	for inorder[mid] != root.Val {
		mid++
	}
	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return root
}
