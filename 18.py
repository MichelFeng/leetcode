# coding=utf-8
class Solution(object):
    """
    暴力法
    """
    def fourSum(self, nums, target):
        length = len(nums)
        if length < 4:
            return []
        if length == 4:
            s = sum(nums)
            if s == target:
                return [nums]
            return []
        res = {}
        nums.sort()
        for i in range(length - 3):
            for j in range(1, length-2):
                if j <= i:
                    continue
                for k in range(2, length-1):
                    if k <= i or k <= j:
                        continue
                    for l in range(3, length):
                        if l <= i or l <= j or l <= k:
                            continue
                        if nums[i] + nums[j] + nums[k] + nums[l] == target:
                            t = [nums[i], nums[j], nums[k], nums[l]]
                            key = "".join(map(str, t))
                            res[key] = t

        return res.values()


class Solution2(object):
    def fourSum(self, nums, target):
        length = len(nums)
        if length < 4:
            return []
        if length == 4:
            s = sum(nums)
            if s == target:
                return [nums]
            return []
        nums.sort()
        res = []
        for i in range(length-3):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            if sum(nums[i:i+4]) > target:
                break
            for j in range(i+1, length-2):
                if j-i > 1 and nums[j] == nums[j-1]:
                    continue
                if nums[i] + sum(nums[j:j+3]) > target:
                    break
                for k in range(j+1, length-1):
                    if k-j > 1 and nums[k] == nums[k-1]:
                        continue
                    if nums[i] + nums[j] + nums[k] + nums[k+1] > target:
                        break
                    for l in range(k+1, length):
                        if l-k > 1 and nums[l] == nums[l-1]:
                            continue
                        if nums[i] + nums[j] + nums[k] + nums[l] > target:
                            break
                        if nums[i] + nums[j] + nums[k] + nums[l] == target:
                            res.append([nums[i], nums[j], nums[k], nums[l]])
        return res

class Solution3(object):
    def fourSum(self, nums, target):
        length = len(nums)
        if length < 4:
            return []
        if length == 4:
            s = sum(nums)
            if s == target:
                return [nums]
            return []
        nums.sort()
        res = []
        for i in range(length-3):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            if sum(nums[i:i+4]) > target:
                break
            for j in range(i+1, length-2):
                if j-i > 1 and nums[j] == nums[j-1]:
                    continue
                if nums[i] + sum(nums[j:j+3]) > target:
                    break
                left = j + 1
                right = length - 1
                while left < right:
                    t = nums[i] + nums[j] + nums[left] + nums[right]
                    if t == target:
                        res.append([nums[i], nums[j], nums[left], nums[right]])

                        while left < right and nums[left] == nums[left+1]:
                            left += 1
                        while left < right and  nums[right] == nums[right-1]:
                            right -= 1
                        left += 1
                        right -= 1
                    elif t > target:
                        right -= 1
                    else:
                        left += 1
        return res


if __name__ == '__main__':
    s = Solution3()
    nums =  [-498,-492,-473,-455,-441,-412,-390,-378,-365,-359,-358,-326,-311,-305,-277,-265,-264,-256,-254,-240,-237,-234,-222,-211,-203,-201,-187,-172,-164,-134,-131,-91,-84,-55,-54,-52,-50,-27,-23,-4,0,4,20,39,45,53,53,55,60,82,88,89,89,98,101,111,134,136,209,214,220,221,224,254,281,288,289,301,304,308,318,321,342,348,354,360,383,388,410,423,442,455,457,471,488,488]
    target = -2808
    print s.fourSum(nums, target)
    
    nums = [1, 0, -1, 0, -2, 2]
    target = 0
    print s.fourSum(nums, target)
