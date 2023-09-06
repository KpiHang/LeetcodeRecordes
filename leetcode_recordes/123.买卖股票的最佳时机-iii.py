#
# @lc app=leetcode.cn id=123 lang=python3
#
# [123] 买卖股票的最佳时机 III
#


# @lc code=start
# 代录；动态规划；股票买卖问题；滚动 数组优化 存储；
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0
        dp = [0] * 5
        dp[1] = -prices[0]
        dp[3] = -prices[0]
        for i in range(1, len(prices)):
            dp[1] = max(dp[1], dp[0] - prices[i])
            dp[2] = max(dp[2], dp[1] + prices[i])
            dp[3] = max(dp[3], dp[2] - prices[i])
            dp[4] = max(dp[4], dp[3] + prices[i])
        return dp[4]

    # self 这种方法可以看到，其实只需要维护dp[i],和dp[i-1]即可，
    # 用滚动数组的话只需要一个dp[i]的状态即可，如上
    def maxProfit_self(self, prices: List[int]) -> int:
        n = len(prices)
        dp = [[0] * 5 for _ in range(n)]
        # 初始化
        dp[0][0] = 0
        dp[0][1] = -prices[0]
        dp[0][2] = 0
        dp[0][3] = -prices[0]
        dp[0][4] = 0

        for i in range(1, n):
            dp[i][1] = max(dp[i - 1][0] - prices[i], dp[i - 1][1])  # 可以滚动数组
            dp[i][2] = max(dp[i - 1][1] + prices[i], dp[i - 1][2])  # 可以滚动数组
            dp[i][3] = max(dp[i - 1][2] - prices[i], dp[i - 1][3])  # 可以滚动数组
            dp[i][4] = max(dp[i - 1][3] + prices[i], dp[i - 1][4])  # 可以滚动数组
        return max(dp[n - 1])


# @lc code=end
