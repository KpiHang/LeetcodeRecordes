/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代法中序遍历二叉树；
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{} // 创建栈
	curr := root
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, curr.Val) // 访问该元素； 中；
			curr = curr.Right           // 该访问右子树了；右；
		}
	}
	return res
}

// 分治法
func inorderTraversal_2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	// 拆分成左右子树两个子问题；
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	// 子问题结果合并；
	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)

	return res
}

// 递归法
func inorderTraversal_1(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历时的位置调整；
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// @lc code=end

