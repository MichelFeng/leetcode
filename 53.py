# coding=utf8
# 53. 最大子序和
# 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
# 示例:
# 输入: [-2, 1, -3, 4, -1, 2, 1, -5, 4],
# 输出: 6
# 解释: 连续子数组 [4, -1, 2, 1] 的和最大，为 6。


class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = nums[0]
        s = 0
        for n in nums:
            if s > 0:
                s += n
            else:
                s = n
            res = max([res, s])
        return res


s = Solution()
print s.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])
print s.maxSubArray([1, 2])
