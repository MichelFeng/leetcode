# coding=utf8
# 74. 搜索二维矩阵
# 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
# 每行中的整数从左到右按升序排列。
# 每行的第一个整数大于前一行的最后一个整数。
# 示例 1:
# 输入:
# matrix = [
#     [1,   3,  5,  7],
#     [10, 11, 16, 20],
#     [23, 30, 34, 50]
# ]
# target = 3
# 输出: true
# 示例 2:
# 输入:
# matrix = [
#     [1,   3,  5,  7],
#     [10, 11, 16, 20],
#     [23, 30, 34, 50]
# ]
# target = 13
# 输出: false


class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix:
            return False
        rows = len(matrix) - 1
        if not matrix[0]:
            return False
        cols = len(matrix[0])
        for i in range(cols):
            while rows >= 0:
                if matrix[rows][i] == target:
                    return True
                elif matrix[rows][i] > target:
                    rows -= 1
                else:
                    break
        return False


s = Solution()
matrix = [
    [1,   3,  5,  7],
    [10, 11, 16, 20],
    [23, 30, 34, 50]
]
print s.searchMatrix(matrix, 13)
print s.searchMatrix(matrix, 16)
