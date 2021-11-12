# coding=utf8
# 392. 判断子序列
# 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
# 你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~ = 500, 000），而 s 是个短字符串（长度 <= 100）。
# 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
# 示例 1:
# s = "abc", t = "ahbgdc"
# 返回 true.
# 示例 2:
# s = "axc", t = "ahbgdc"
# 返回 false.
# 后续挑战:
# 如果有大量输入的 S，称作S1, S2, ..., Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？


class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        len_s = len(s)
        len_t = len(t)
        dp = []
        for i in range(len_s+1):
            tmp = []
            for j in range(len_t+1):
                if i == 0:
                    tmp.append(1)
                else:
                    tmp.append(0)
            dp.append(tmp)

        for i in range(1, len_s+1):
            for j in range(1, len_t+1):
                if t[j-1] == s[i-1]:
                    dp[i][j] = dp[i-1][j]
                else:
                    dp[i][j] = dp[i][j-1]

        return dp[len_s][len_t]


s = Solution()
print s.isSubsequence('abc', 'ahbgcd')
print s.isSubsequence('axc', 'ahbgdc')
print s.isSubsequence('leeeeetcode', 'leyyyyeetcode')
