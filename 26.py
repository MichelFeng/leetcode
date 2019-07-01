#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution(object):
    def removeDuplicates(self, nums):
        """TODO: Docstring for removeDuplicates.
        :returns: TODO

        """
        n = len(nums)
        if n <= 1:
            return n

        remove = 0
        for i in range(1, n):
            if nums[i] == nums[i-1]:
                remove += 1
                continue
            nums[i-remove] = nums[i]
        if remove > 0:
            nums = nums[0:n-remove]
        return len(nums)

if __name__ == "__main__":
    s = Solution()
    nums = [0,0,1,1,1,2,2,3,3,4]
    print(s.removeDuplicates(nums))


