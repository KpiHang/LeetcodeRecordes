#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    # 滑动窗口，寻找最长；
    def lengthOfLongestSubstring(self, s: str) -> int:
        left, right = 0, 0
        window, longest = {}, 0

        while right < len(s):
            if s[right] not in window:
                window[s[right]] = 1
            else: window[s[right]] += 1

            while window[s[right]] > 1:
                window[s[left]] -= 1
                left += 1
            
            longest = max(longest, right - left + 1)
            right += 1 

        return longest


# @lc code=end

