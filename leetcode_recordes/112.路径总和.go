/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

// 法二，有回溯一定有函数传参；
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// N
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	// L
	left_bool := hasPathSum(root.Left, targetSum)
	// R
	right_bool := hasPathSum(root.Right, targetSum)

	return left_bool || right_bool
}

// Self: 需要回溯，traversal函数传参；
func hasPathSum_1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var (
		sum     int
		results map[int]struct{} //这里定义map的原因是，go语言中没有 if xx in list 这种东西，只能循环遍历， 所以用map val用空结构体不占用额外内存空间 较为方便；
	) // results 包含了每条从跟到叶子节点的路径和，注意，这种实现方式，是集合，注意路径和相同的情况；
	results = map[int]struct{}{} // 这也是Go语言中集合的实现方式；

	var traversal func(curr *TreeNode, sum int, results map[int]struct{})
	traversal = func(curr *TreeNode, sum int, results map[int]struct{}) {
		// 前
		sum += curr.Val
		if curr.Left == nil && curr.Right == nil {
			results[sum] = struct{}{}
		}
		// 左
		if curr.Left != nil {
			traversal(curr.Left, sum, results)
		}
		// 右
		if curr.Right != nil {
			traversal(curr.Right, sum, results)
		}
	}
	traversal(root, sum, results)
	_, ok := results[targetSum]
	return ok
}

// @lc code=end

