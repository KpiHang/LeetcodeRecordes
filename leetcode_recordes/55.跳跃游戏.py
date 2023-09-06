#
# @lc app=leetcode.cn id=55 lang=python3
#
# [55] 跳跃游戏
#

# @lc code=start

# 贪心算法，找局部最优解；
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) <= 1: return True
        start = 0
        for end in range(len(nums)):
            if nums[start]-(end-start) <= nums[end]:
                start = end
                if nums[start] == 0 and start != len(nums)-1: 
                    return False
        return True
# @lc code=end
