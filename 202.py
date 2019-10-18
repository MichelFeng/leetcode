# coding=utf8
# 202. 快乐数
# 编写一个算法来判断一个数是不是“快乐数”。
# 一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
# 示例:
# 输入: 19
# 输出: true
# 解释:
# 1^2 + 9^2 = 82
# 8^2 + 2^2 = 68
# 6^2 + 8^2 = 100
# 1^2 + 0^2 + 0^2 = 1


class Solution(object):
    def isHappy(self, n):
        """
        使用快慢指针思想，判断是否成环，同 141
        :type n: int
        :rtype: bool
        """
        slow = self.calc(n)
        fast = self.calc(self.calc(n))
        while True:
            if fast == 1:
                return True
            if slow == fast:
                return False
            slow = self.calc(slow)
            fast = self.calc(self.calc(fast))

    def calc(self, n):
        s = 0
        while True:
            s = s + (n % 10)**2
            if not n:
                break
            n = n // 10
        return s


s = Solution()
print s.isHappy(19)
print s.isHappy(2)
