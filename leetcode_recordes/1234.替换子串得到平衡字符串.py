#
# @lc app=leetcode.cn id=1234 lang=python3
#
# [1234] 替换子串得到平衡字符串
#

# @lc code=start
from collections import Counter


class Solution:
    # 寻找最短进阶；比直接使用模版更复杂一些的问题；score: 1878
    # 笔记 + 视频：https://www.bilibili.com/video/BV1TM411e7W7/
    def balancedString(self, s: str) -> int:
        n = len(s)
        letter_to_idx = {
            "Q": 0,
            "W": 1,
            "E": 2,
            "R": 3,
        }
        idx_to_letter = {}
        for k, v in letter_to_idx.items():
            idx_to_letter[v] = k

        total = Counter(s)  # eg, Counter({'Q': 5, 'W': 2, 'E': 1, 'R': 3})
        outside = [0] * 4  # 0，1,2,3 对应代表 QWER
        for k, v in idx_to_letter.items():
            outside[k] = total[v]

        if max(outside) == n // 4:  # 初始已经是不用替换的。
            return 0

        left, right = 0, 0
        shortest = float("inf")
        while left < n and right < n:
            while right < n and max(outside) > n // 4:
                outside[letter_to_idx[s[right]]] -= 1
                right += 1
            # NOTE left < right
            while left < right and max(outside) <= n // 4:  # 虽然这里写的<= 但max不可能< 循环条件只用到了==
                # 寻找最短，满足条件，尽可能短
                outside[letter_to_idx[s[left]]] += 1
                left += 1
            shortest = min(shortest, right - left + 1)  # left+1导致跳出循环

        return shortest


# @lc code=end
