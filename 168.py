# coding=utf8
# 168. Excel表列名称
# 给定一个正整数，返回它在 Excel 表中相对应的列名称。
# 例如，
# 1 -> A
# 2 -> B
# 3 -> C
# ...
# 26 -> Z
# 27 -> AA
# 28 -> AB
# ...
# 示例 1:
# 输入: 1
# 输出: "A"
# 示例 2:
# 输入: 28
# 输出: "AB"
# 示例 3:
# 输入: 701
# 输出: "ZY"


class Solution(object):
    def convertToTitle(self, n):
        """
        与 171 题对应，26进制
        :type n: int
        :rtype: str
        """
        s = ""
        while n:
            c = n % 26
            if c == 0:
                c = 26
                n -= 26
            s = chr(c+64) + s
            n = int(n / 26)
        return s


s = Solution()
print s.convertToTitle(1)
print s.convertToTitle(28)
print s.convertToTitle(701)
