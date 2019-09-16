# coding=utf8
# 67. 二进制求和
# 给定两个二进制字符串，返回他们的和（用二进制表示）。
# 输入为非空字符串且只包含数字 1 和 0。
# 示例 1:
# 输入: a = "11", b = "1"
# 输出: "100"
# 示例 2:
# 输入: a = "1010", b = "1011"
# 输出: "10101"


class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        a = list(reversed(a))
        b = list(reversed(b))
        lenA = len(a)
        lenB = len(b)
        maxLen = max([lenA, lenB])
        res = []
        plus = 0
        for i in range(maxLen):
            ai = int(a[i]) if i < lenA else 0
            bi = int(b[i]) if i < lenB else 0
            res.append(str((ai + bi + plus) % 2))
            plus = 1 if ai + bi + plus > 1 else 0
        if plus:
            res.append(str(plus))
        return "".join(reversed(res))


s = Solution()
print s.addBinary('1010', '1011')
