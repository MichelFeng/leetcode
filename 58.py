# 58. 最后一个单词的长度
class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = s.split()
        if len(res):
            return len(res[-1])
        return 0


s = Solution()
print s.lengthOfLastWord('Hello World  ')
