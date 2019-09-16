class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        if x == 0:
            return 0
        i = 1
        while(i <= x):
            if i * i > x:
                return i-1
            i += 1


s = Solution()
print s.mySqrt(2)
print s.mySqrt(4)
print s.mySqrt(8)
