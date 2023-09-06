#
# @lc app=leetcode.cn id=647 lang=python3
#
# [647] 回文子串
#


# @lc code=start
class Solution:
    # 代录 DP 从下到上 从左到右
    def countSubstrings(self, s: str) -> int:
        dp = [[False] * len(s) for _ in range(len(s))]
        res = 0
        for i in range(len(s) - 1, -1, -1):
            for j in range(i, len(s)):
                if s[i] == s[j]:
                    if j - i <= 1:
                        dp[i][j] = True
                        res += 1
                    else:
                        if dp[i + 1][j - 1]:
                            dp[i][j] = True
                            res += 1
        return res


# @lc code=end
