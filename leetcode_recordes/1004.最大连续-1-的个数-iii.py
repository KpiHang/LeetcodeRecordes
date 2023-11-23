#
# @lc app=leetcode.cn id=1004 lang=python3
#
# [1004] 最大连续1的个数 III
#

# @lc code=start
class Solution:
    # 寻找最长； self 熟练使用模版
    def longestOnes(self, nums: List[int], k: int) -> int:
        left, right = 0, 0
        _0_nums = 0
        longest = 0
        while right < len(nums):
            if nums[right] == 0:
                _0_nums += 1
            # window 条件：窗口内0的个数<=k;
            while _0_nums > k: # 最长不满足条件，left才移动；
                if nums[left] == 0:
                    _0_nums -= 1
                left += 1
            right += 1
            longest = max(longest, right-left)
        return longest
    # self 熟练使用模版，一次AC；
# @lc code=end