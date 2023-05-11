#
# @lc app=leetcode.cn id=442 lang=python3
#
# [442] 数组中重复的数据
#


# @lc code=start
class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        nums.sort()
        res = []
        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                res.append(nums[i])
        return res


# @lc code=end
