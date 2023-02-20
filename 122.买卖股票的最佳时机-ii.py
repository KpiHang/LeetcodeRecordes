#
# @lc app=leetcode.cn id=122 lang=python3
#
# [122] 买卖股票的最佳时机 II
#

# @lc code=start

# 贪心算法，从局部最优-->全局最优；
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans_profit = 0
        for i in range(1, len(prices)):
            ans_profit = ans_profit + max(0, prices[i]-prices[i-1])
        return ans_profit

# @lc code=end

