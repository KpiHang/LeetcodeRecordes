#
# @lc app=leetcode.cn id=309 lang=python3
#
# [309] 最佳买卖股票时机含冷冻期
#


# @lc code=start


# 未完成；
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        k = n // 2  # 最多交易次数

        dp = [0] * (1 + 2 * k + 1)  # 最后位标记是否是今天卖出的；初始为0
        # 初始化
        for i in range(1 + 2 * k, 2):
            dp[i] = -prices[0]

        for i in range(1, n):
            for j in range(1, 1 + 2 * k):
                if dp[-1] == 1:
                    dp[-1] = 0  # 度过冷冻期；把买入的跳过，但卖出的还可以；
                    dp[j] = max(dp[j], dp[j - 1] + prices[i])
                    continue
                if j % 2:
                    dp[j] = max(dp[j], dp[j - 1] - prices[i])
                else:
                    dp[j] = max(dp[j], dp[j - 1] + prices[i])
                    if dp[j] < dp[j - 1] + prices[i]:
                        dp[-1] = 1  # 标记卖出；

        return dp[-2]


# @lc code=end
