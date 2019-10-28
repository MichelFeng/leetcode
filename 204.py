# coding=utf8
# 204. 计数质数
# 统计所有小于非负整数 n 的质数的数量。
# 统计所有小于非负整数 n 的质数的数量。
# 示例:
# 输入: 10
# 输出: 4
# 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。


class Solution(object):
    def countPrimes(self, n):
        """
        超时了
        :type n: int
        :rtype: int
        """

        dp = [0, 0, ]
        for i in range(2, n + 1):
            dp.append(dp[i-1] + self.is_prime(i))
        return dp[n-1]

    def is_prime(self, n):
        for i in range(2, n//2 + 1):
            if n % i == 0:
                return 0
        return 1

    def countPrimes2(self, n):
        """
        厄拉多塞筛法
        """
        res = 0
        flag = [True for _ in range(n)]
        for i in range(2, n):
            if flag[i]:
                res += 1
                for j in range(i*2, n, i):
                    flag[j] = False
        return res


s = Solution()
print(s.countPrimes(10))
print(s.countPrimes(15))
print(s.countPrimes2(2))
print(s.countPrimes2(10))
print(s.countPrimes2(15))
