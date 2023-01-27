/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// Self: 二叉树双队列层序遍历模板（入队是稍加变动）+数值处理
func levelOrder(root *Node) (res [][]int) {
	// res := [][]int{} //  []类型名{}，类型名是[]int
	// 返回值也可以像形式参数一样被命名。在这种情况下，**每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为该类型的零值**（自动内存分配并赋零初值）
	if root == nil {
		return res
	}
	Q := []*Node{root}
	for len(Q) > 0 { // 层遍历
		P := []*Node{}
		res_level := []int{}
		for i := 0; i < len(Q); i++ { // 每层的元素遍历
			curr_node := Q[i]
			res_level = append(res_level, curr_node.Val)
			for _, kid := range curr_node.Children {
				if kid != nil {
					P = append(P, kid)
				}
			}
		}
		res = append(res, res_level)
		Q = P
	}
	return res
}

// @lc code=end

