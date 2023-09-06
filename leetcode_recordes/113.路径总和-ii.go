/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

// Self：有回溯一定有递归函数传参；
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	// ans 是什么，ans 是一个 指针区间 的集合
	if root == nil {
		return ans
	}
	sum_ls := []int{}
	var traversal func(curr *TreeNode, sum_ls []int, count int)
	traversal = func(curr *TreeNode, sum_ls []int, count int) {
		// N
		sum_ls = append(sum_ls, curr.Val)
		count -= curr.Val

		if curr.Left == nil && curr.Right == nil && count == 0 {
			cp := make([]int, len(sum_ls)) // 这里为什么不直接append sum_ls？因为要时刻记住：切片是底层数组的视图
			copy(cp, sum_ls)
			ans = append(ans, cp) // ans是匿名函数外的闭包；
		}
		// L
		if curr.Left != nil {
			traversal(curr.Left, sum_ls, count)
		}
		// R
		if curr.Right != nil {
			traversal(curr.Right, sum_ls, count)
		}
	}

	traversal(root, sum_ls, targetSum)
	return ans
}

// @lc code=end

