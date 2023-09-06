#
# @lc app=leetcode.cn id=188 lang=python3
#
# [188] 买卖股票的最佳时机 IV
#
# @lc code=start


class Solution:
    # 代录
    def maxProfit(self, k: int, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0
        dp = [0] * (2 * k + 1)
        for i in range(1, 2 * k, 2):  # 改进；
            dp[i] = -prices[0]
        for i in range(1, len(prices)):
            for j in range(1, 2 * k + 1):
                if j % 2:
                    dp[j] = max(dp[j], dp[j - 1] - prices[i])
                else:
                    dp[j] = max(dp[j], dp[j - 1] + prices[i])
        return dp[2 * k]

    # self; 由123题代录，滚动数组的基础；
    def maxProfit_self(self, k: int, prices: List[int]) -> int:
        n = len(prices)
        k = 2 * k + 1
        dp = [0] * k
        # 初始化；
        for i in range(k):
            if i % 2 != 0:
                dp[i] = -prices[0]

        for i in range(1, n):
            for j in range(1, k):
                if j % 2 == 0:
                    dp[j] = max(dp[j], dp[j - 1] + prices[i])
                else:
                    dp[j] = max(dp[j], dp[j - 1] - prices[i])
        return max(dp)


# @lc code=end
