/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
// 代录：贪心算法；
func candy(ratings []int) int {
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	need := make([]int, len(ratings))
	// 初始化(每个人至少一个糖果)
	for i := range need {
		need[i] = 1
	}

	// 1. 从左到右，当右边大于左边就加1；
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			need[i] = need[i-1] + 1
		}
	} // 保证右边大于左边，分发糖果的情况；

	// 2. 再从右到左，当左边大于右边的就右边加一，但同时要满足右大于左；
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}

	sum := 0
	for i := range need {
		sum += need[i]
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end

