#!/usr/bin/env python
# -*- coding: utf-8 -*-
class Solution(object):
    """description"""
    def removeElement(self, nums, val):
        """TODO: Docstring for removeElement.
        :returns: TODO

        """
        if not nums:
            return 0
        n = len(nums)
        l = 0
        while l < n:
            if nums[l] != val:
                l += 1
            else:
                nums[l] = nums[n-1]
                n -= 1
        return n
    

if __name__ == "__main__":
    s = Solution()
    nums = [0,1,2,2,3,0,4,2]
    val = 2
    print(s.removeElement(nums, val))
