# -*- coding: utf-8 -*-


class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        digits = digits[-1::-1]
        res = []
        plus = 1
        for i, digit in enumerate(digits):
            if digit + plus > 9:
                res.append(digit - 9)
                plus = 1
            else:
                res.append(digit + plus)
                plus = 0
        if plus:
            res.append(plus)
        return res[-1::-1]


s = Solution()
print s.plusOne([0])
print s.plusOne([1, 2, 3])
print s.plusOne([4, 3, 2, 1])
print s.plusOne([1, 2, 9])
print s.plusOne([9])
