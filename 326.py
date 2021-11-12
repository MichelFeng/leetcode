class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        if n == 1:
            return True
        if n < 3:
            return False
        return n % 3 == 0 and self.isPowerOfThree(n/3)


if __name__ == "__main__":
    print(divmod(16, 3))
    print(16/3)
    s = Solution()
    print(s.isPowerOfThree(45))
    print(s.isPowerOfThree(0))
    print(s.isPowerOfThree(1))
    print(s.isPowerOfThree(3))
    print(s.isPowerOfThree(9))
    print(s.isPowerOfThree(16))
