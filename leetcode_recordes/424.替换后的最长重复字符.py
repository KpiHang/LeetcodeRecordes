#
# @lc app=leetcode.cn id=424 lang=python3
#
# [424] 替换后的最长重复字符
#

# @lc code=start
class Solution:
    # 滑动窗口，一次直接AC
    def characterReplacement(self, s: str, k: int) -> int:
        left, right = 0, 0
        win = {}
        ans = 0
        while right < len(s):
            if s[right] not in win:
                win[s[right]] = 1
            else:
                win[s[right]] += 1
            
            while (right-left+1)-max(win.values()) > k:
                win[s[left]] -= 1
                if win[s[left]] == 0:
                    del win[s[left]]
                left += 1

            ans = max(ans, right-left+1)
            right += 1
        return ans
# @lc code=end

