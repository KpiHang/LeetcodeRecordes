#
# @lc app=leetcode.cn id=53 lang=python3
#
# [53] 最大子数组和
#

# @lc code=start

INT_MIN_VALUE = -sys.maxsize - 1

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        tmp = 0
        res = INT_MIN_VALUE
        for i in range(len(nums)):
            tmp += nums[i]
            res = max(tmp, res)
            if tmp < 0: tmp = 0
        return res
# @lc code=end

