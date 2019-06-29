class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return reduce(compute, nums)
    
def compute(a, b):
    return a ^ b
