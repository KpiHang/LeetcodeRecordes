/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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

// 参考答案：迭代法前序遍历二叉树；
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New() // 用list内置库创建的一个带头结点的循环双链表； https://juejin.cn/post/7042729165400834056#heading-2
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode) //这里是一个类型断言；remove返回删除的元素；

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

// SELF: 迭代法前序遍历二叉树；
func preorderTraversa_3(root *TreeNode) (res []int) {
	stack := []*TreeNode{} // 创建栈
	if root == nil {
		return res
	}
	stack = append(stack, root) // 根节点入栈；

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈
		res = append(res, curr.Val)  // 访问该元素

		// 先右后左，出栈时遍历顺序才是先左后右；
		if curr.Right != nil { // 右子树不空，右子树入栈
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil { // 左子树不空，左子树入栈
			stack = append(stack, curr.Left)
		}
	}
	return res
}

// 分治法
func preorderTraversal_2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 1. 拆分成子问题
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	// 2. 子问题结果合并
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)

	return res

}

// 递归法
func preorderTraversal_1(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// @lc code=end

