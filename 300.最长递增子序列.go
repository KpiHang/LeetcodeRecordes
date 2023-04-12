/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start

// 代录：动态规划求解
func lengthOfLIS(nums []int) int {
	// dp数组的定义 dp[i]表示取第i个元素的时候，表示子序列的长度，其中包括 nums[i] 这个元素
	dp := make([]int, len(nums))

	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// 代录-self；动态规划；
func lengthOfLIS_self(nums []int) int {
	// dp[i]：以nums[i]结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	// 初始化
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	// 返回最长递增子序列长度
	res := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

