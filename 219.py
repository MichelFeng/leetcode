# coding=utf8
# 219. 存在重复元素 II
# 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums[i] = nums[j]，并且 i 和 j 的差的绝对值最大为 k。
# 示例 1:
# 输入: nums = [1, 2, 3, 1], k = 3
# 输出: true
# 示例 2:
# 输入: nums = [1, 0, 1, 1], k = 1
# 输出: true
# 示例 3:
# 输入: nums = [1, 2, 3, 1, 2, 3], k = 2
# 输出: false


class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        d = dict()
        for i, n in enumerate(nums):
            if n in d:
                if k >= abs(i - d[n]):
                    return True
                d[n] = i
            else:
                d[n] = i
        return False


s = Solution()

nums = [1, 0, 1, 1]
k = 1
print(s.containsNearbyDuplicate(nums, k))

nums = [1, 2, 3, 1]
k = 3
print(s.containsNearbyDuplicate(nums, k))

nums = [1, 2, 3, 1, 2, 3]
k = 2
print(s.containsNearbyDuplicate(nums, k))
