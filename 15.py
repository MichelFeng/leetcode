# coding=utf-8

class Solution(object):
    def threeSum(self, nums):
        n = len(nums)
        if n < 3:
            return []
        nums.sort()

        res = []
        for i in range(n-1):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            if sum(nums[i: i+3]) > 0:
                break
            l = i + 1
            r = n - 1
            while l < r:
                t = nums[i] + nums[l] + nums[r]
                if t == 0:
                    res.append([nums[i], nums[l], nums[r]])
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1
                    l += 1
                    r -= 1
                elif t < 0:
                    l += 1
                else:
                    r -= 1

        return res


if __name__ == '__main__':
    s = Solution()
    nums = [-1, 0, 1, 2, -1, -4]
    print s.threeSum(nums)
