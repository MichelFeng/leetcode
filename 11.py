class Solution(object):
    def maxArea(self, height):
        """
        :type height: list[int]
        :rtype: int
        """
        i = 0
        j = len(height) -1

        max_area = 0
        while i < j:
            move_right = height[i] < height[j]
            h = height[i] if move_right else height[j]
            t = (j - i) * h
            if t > max_area:
                max_area = t
            if move_right:
                i += 1
            else:
                j -= 1
        return max_area


if __name__ == '__main__':
    s = Solution()
    height = [1, 8 ,6, 2,5,4,8,3,7]
    print s.maxArea(height)
