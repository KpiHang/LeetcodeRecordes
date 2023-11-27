#
# @lc app=leetcode.cn id=1658 lang=python3
#
# [1658] 将 x 减到 0 的最小操作数
#

# @lc code=start
class Solution:
    # 方法一：逆向思维 先转化问题，再滑动窗口；
    # https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/solutions/2048811/ni-xiang-si-wei-pythonjavacgo-by-endless-b4jt/
    def minOperations(self, nums: List[int], x: int) -> int:
        target = sum(nums) - x
        if target < 0:
            return -1
        longest = 0
        left = win_sum = 0
        for right, num in enumerate(nums):
            win_sum += num
            while left <= right and win_sum > target:
                win_sum -= nums[left]
                left += 1
            if win_sum == target:
                longest = max(longest, right - left + 1)
        return -1 if longest < 0 else len(nums) - longest
# @lc code=end

