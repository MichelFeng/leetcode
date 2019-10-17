# coding=utf8
# 300. 最长上升子序列
# 给定一个无序的整数数组，找到其中最长上升子序列的长度。
# 示例:
# 输入: [10, 9, 2, 5, 3, 7, 101, 18]
# 输出: 4
# 解释: 最长的上升子序列是 [2, 3, 7, 101]，它的长度是 4。
# 说明:
# 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
# 你算法的时间复杂度应该为 O(n2) 。
# 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?


class Solution(object):
    def lengthOfLIS(self, nums):
        """
        时间复杂度 n^2
        :type nums: List[int]
        :rtype: int
        """
        length = len(nums)
        if not length:
            return 0
        dp = [0] * length
        dp[0] = 1

        for i in range(1, length):
            dp[i] = 1
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max([dp[j]+1, dp[i]])

        return max(dp)

    def lengthOfLIS2(self, nums):
        """
        时间复杂度 nlogn
        :type nums: List[int]
        :rtype: int
        """
        length = len(nums)
        if not length:
            return 0
        dp = [nums[0]] * length
        l = 1
        for i in range(1, length):
            if nums[i] > dp[l-1]:
                dp[l] = nums[i]
                l += 1
            else:
                pos = self.biSearch(dp, l, nums[i])
                dp[pos] = nums[i]
        return l

    def biSearch(self, arr, length, num):
        l = 0
        r = length-1
        while l <= r:
            mid = l + (r-l)//2
            if arr[mid] > num:
                r = mid - 1
            elif arr[mid] < num:
                l = mid + 1
            else:
                return mid
        return l


s = Solution()
print(s.lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]))
print(s.lengthOfLIS2([10, 9, 2, 5, 3, 7, 101, 18]))
