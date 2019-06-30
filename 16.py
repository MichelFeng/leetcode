#!/usr/bin/env python
# -*- coding: utf-8 -*-
class Solution(object):
    def threeSumClosest(self, nums, target):
        n = len(nums)
        if n < 3:
            return []
        nums.sort()
        closest = 2 ** 32 - 1
        res = 0
        for i in range(n-2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            l = i + 1
            r = n - 1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                t = target - s

                if t > 0:
                    l += 1
                else:
                    r -= 1
                if abs(t) < closest:
                    closest = abs(t) 
                    res = s
        return res


if __name__ == "__main__":
    nums = [-1, 2, 1, -4]
    target = 1
    print Solution().threeSumClosest(nums, target)

    nums = [0,0,0]
    target = 1
    print Solution().threeSumClosest(nums, target)
