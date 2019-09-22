# coding=utf8
# 反转int类型的值x，不要借用String，只用int 即可。


class Solution(object):
    def reverse(self, val):
        if val == 0:
            return 0
        flag = val < 0
        if flag:
            val *= -1
        res = 0
        while val:
            res = res * 10 + val % 10
            val /= 10
        return (res * -1) if flag else res


s = Solution()
print s.reverse(0)
print s.reverse(1)
print s.reverse(-1)
