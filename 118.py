# coding=utf8
# 118. 杨辉三角
# 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
# 在杨辉三角中，每个数是它左上方和右上方的数的和。
# 示例:
# 输入: 5
# 输出:
# [
#     [1],
#     [1, 1],
#     [1, 2, 1],
#     [1, 3, 3, 1],
#     [1, 4, 6, 4, 1]
# ]


class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        res = []
        i = 1
        while i <= numRows:
            j = 1
            row = []
            while j <= i:
                if j == 1 or j == i:
                    row.append(1)
                else:
                    row.append(res[i-2][j-1] + res[i-2][j-2])
                j += 1
            res.append(row)
            i += 1
        return res


s = Solution()
print s.generate(0)
print s.generate(1)
print s.generate(2)
print s.generate(3)
print s.generate(5)
