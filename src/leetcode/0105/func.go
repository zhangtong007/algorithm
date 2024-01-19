package leetcode105

import . "core/src/datastruct"

/**
基于二叉树各个值不同的场景：
根据先序或者后序， 加上中序 序列化二叉树问题
递归，分解为子问题
先序：  root  left   right
中序:	left  root   right
后序：  left  right  left

根据先序和后序可以找到root  (分别找首, 尾)
找到root后可以根据中序，确认 left和right的长度 len
划分子问题：
先序:
	left -> 从root后面 寻len个长度元素
    right -> 取完left之后的所有元素

后序:
	left -> 从0开始，寻len个长度元素
	right -> 寻完left之后的所有元素，除去末尾元素  (末尾元素是根节点，已经被统计)
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	mid := 0
	for inorder[mid] != preorder[0] {
		mid++
	}
	root.Left = buildTree(preorder[1:mid+1], inorder[0:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
