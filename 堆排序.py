# coding=utf8


class HeapSort(object):

    def create(self, nums):
        if not nums:
            return None

        l = int(len(nums)/2)
        for i in range(l, -1, -1):
            self.heapify(nums, i)
        return nums

    def heapify(self, nums, i):
        left = 2 * i + 1
        right = 2 * i + 2
        largest = i
        l = len(nums)
        if left < l and nums[left] > nums[largest]:
            largest = left
        if right < l and nums[right] > nums[largest]:
            largest = right
        if largest != i:
            self.swap(nums, i, largest)
            self.heapify(nums, largest)

    def swap(self, nums, i, j):
        nums[i], nums[j] = nums[j], nums[i]

    def sort(self, nums):
        nums = self.create(nums)
        res = []
        while nums:
            res.append(nums.pop(0))
            self.create(nums)

        return res

    def insert(self, nums, val):
        nums = self.create(nums)
        nums.append(val)
        self.heapify(nums, )


s = HeapSort()
nums = [1, 5, 2, 6, 3, 4, 7]
print s.create(nums)
print s.sort(nums)
