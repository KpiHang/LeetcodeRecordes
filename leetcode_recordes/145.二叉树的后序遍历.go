/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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

// 简单的迭代法，利用先序中左右，调整顺序中右左，然后reverse；
func postorderTraversal(root *TreeNode) (res []int) {
	res = preorderTraversa(root)
	reverse(res)
	return
}
func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}

// 迭代法，前序遍历二叉树；但更改顺序为中右左；
func preorderTraversa(root *TreeNode) (res []int) {
	stack := []*TreeNode{} // 创建栈
	if root == nil {
		return res
	}
	stack = append(stack, root) // 根节点入栈；

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈
		res = append(res, curr.Val)  // 访问该元素

		if curr.Left != nil { // 左子树不空，左子树入栈
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil { // 右子树不空，右子树入栈
			stack = append(stack, curr.Right)
		}
	}
	return res
}

// 上法结束；；；

// 参考答案：后序遍历二叉树；
func postorderTraversal_3(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev { // 右子树已访问的情况，访问根结点；
			res = append(res, root.Val)
			prev = root
			root = nil
		} else { // 右子树还未访问，根重新入栈，先右，再根；
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

// 分治法
func postorderTraversal_2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	// 把问题拆分为左右两个子问题；
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	// 子问题结果合并；   左        右         中
	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}

// 递归法
func postorderTraversal_1(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 后序遍历时的位置调整
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

// @lc code=end

