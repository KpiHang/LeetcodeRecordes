/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
// 贪心算法；代码随想录“覆盖范围”
// https://www.programmercarl.com/0045.%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8FII.html

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	curDistance := 0  // 当前最远覆盖距离下标；
	nextDistance := 0 // 下一跳最远距离的下标；
	ans := 0
	for i := range nums {
		nextDistance = max(nums[i]+i, nextDistance) // 更新下一步覆盖的最远距离；
		if i == curDistance {                       // 遇到当前跳的最远距离；
			if curDistance < len(nums)-1 { // 当前跳还不能到达终点时；
				ans++ // 跳；
				curDistance = nextDistance
				if nextDistance >= len(nums)-1 { // 下一条的覆盖范围已经可以到达终点；
					break
				}
			} else {
				break
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end

