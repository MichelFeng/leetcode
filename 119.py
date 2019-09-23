# coding=utf8
# 119. 杨辉三角 II
# 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
# 在杨辉三角中，每个数是它左上方和右上方的数的和。
# 示例:
# 输入: 3
# 输出: [1, 3, 3, 1]


class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        res = []
        i = 1
        rowIndex += 1
        while i <= rowIndex:
            j = 1
            row = []
            while j <= i:
                if j == 1 or j == i:
                    row.append(1)
                else:
                    row.append(res[i-2][j-1] + res[i-2][j-2])
                j += 1
            i += 1
            res.append(row)
        return res[rowIndex-1] if rowIndex > 0 else res


s = Solution()
print s.getRow(0)
print s.getRow(1)
print s.getRow(2)
print s.getRow(3)
print s.getRow(5)
