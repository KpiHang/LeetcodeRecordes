/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
// SELF：两个队列层序遍历
func levelOrder(root *TreeNode) (res [][]int) {
	currQueue := []*TreeNode{} // 当前层队列；
	nextQueue := []*TreeNode{} // 暂时存放下一层队列；
	if root == nil {
		return
	}
	currQueue = append(currQueue, root)

	for len(currQueue) > 0 || len(nextQueue) > 0 {
		curr_res := []int{} // 当前层的层序结果
		for len(currQueue) > 0 {
			curr := currQueue[0]
			currQueue = currQueue[1:] // 出队
			curr_res = append(curr_res, curr.Val)
			if curr.Left != nil {
				nextQueue = append(nextQueue, curr.Left)
			}
			if curr.Right != nil {
				nextQueue = append(nextQueue, curr.Right)
			}
		}
		res = append(res, curr_res)
		currQueue, nextQueue = nextQueue, currQueue
	}
	return
}

// @lc code=end
