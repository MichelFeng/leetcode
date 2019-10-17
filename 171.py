# coding=utf8
# 给定一个Excel表格中的列名称，返回其相应的列序号。
# 例如，
# A -> 1
# B -> 2
# C -> 3
# ...
# Z -> 26
# AA -> 27
# AB -> 28
# ...
# 示例 1:
# 输入: "A"
# 输出: 1
# 示例 2:
# 输入: "AB"
# 输出: 28
# 示例 3:
# 输入: "ZY"
# 输出: 701


class Solution(object):
    def titleToNumber(self, s):
        """
        26进制
        :type s: str
        :rtype: int
        """

        l = len(s)
        r = 0
        for i in range(0, l):
            r += (ord(s[i])-64) * 26**(l-i-1)
        return r


s = Solution()
print(s.titleToNumber('A'))
print(s.titleToNumber('AB'))
print(s.titleToNumber('ZY'))
print(s.titleToNumber('ABC'))

# 1 2 3
# 1*10**(3-1)  2*10**(2-1)  3*10**(1-1)

# ABC
# 1*26**(3-1)+2*26**(2-1)+3*26**(1-1) = 26**2 + 26*2+3 = 731
