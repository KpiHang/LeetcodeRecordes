#
# @lc app=leetcode.cn id=395 lang=python3
#
# [395] 至少有 K 个重复字符的最长子串
#

# @lc code=start
class Solution:
    # 此题无法直接用滑动窗口求解，理由如下：
    # 变化因素太多了，字符种类会变化，字符个数也会变化；
    # 本题字符就是26种字母，所以可能出现的情况就是1-26种字符，
    # 固定种类数，枚举26种，即：至少有 K 个重复字符的最长子串（子串中的字符种类为n种，n: 1-26）

    # https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solutions/2549261/hua-dong-chuang-kou-jie-gou-qing-xi-by-k-cwu7/
    def longestSubstring(self, s: str, k: int) -> int:
        ans = 0
        for i in range(26):
            ans = max(ans, self.nLongestSubstring(s, k, i + 1))
        return ans

    def nLongestSubstring(self, s: str, k: int, n: int) -> int:
        """
        至少有 k 个重复字符的最长子串（子串中的字母类别只能有 n 种）
        """
        win = {}
        ans = 0
        left, right = 0, 0
        while right < len(s):
            if s[right] not in win:
                win[s[right]] = 1
            else: win[s[right]] += 1

            while len(win) > n:  # r右移导致的不满足，且不满足条件
                win[s[left]] -= 1
                if win[s[left]] == 0:
                    del win[s[left]]  # 因为种类数是len(dict)，虽然value减为0，但还是占用一个长度；
                left += 1
            
            if len(win) == n and min(win.values()) >= k:
                ans = max(ans, right - left + 1)
                
            right += 1
        
        return ans
# @lc code=end

