# 38. 报数
class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        s = '1'
        if n == 1:
            return s
        for i in range(n - 1):
            s = self.read(s)
        return s

    def read(self, s):
        count = 1
        tmp = s[0]
        res = ''
        for _, c in enumerate(s[1:]):
            if c == tmp:
                count += 1
            else:
                res += str(count) + tmp
                tmp = c
                count = 1
        res += str(count) + tmp
        return res


s = Solution()
print s.countAndSay(1)
print s.countAndSay(4)
print s.countAndSay(5)
