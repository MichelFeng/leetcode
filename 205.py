# coding=utf8
# 205. 同构字符串
# 给定两个字符串 s 和 t，判断它们是否是同构的。
# 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
# 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
# 示例 1:
# 输入: s = "egg", t = "add"
# 输出: true
# 示例 2:
# 输入: s = "foo", t = "bar"
# 输出: false
# 示例 3:
# 输入: s = "paper", t = "title"
# 输出: true
# 说明:
# 你可以假设 s 和 t 具有相同的长度。


class Solution(object):
    def isIsomorphic(self, s, t):
        """
        思路：字符在s和t中首次出现的位置是否相同
        :type s: str
        :type t: str
        :rtype: bool
        """
        for i, c in enumerate(s):
            if s.find(c) != t.find(t[i]):
                return False
        return True


so = Solution()
s = "egg"
t = "add"
print(so.isIsomorphic(s, t))
s = "foo"
t = "bar"
print(so.isIsomorphic(s, t))
s = "paper"
t = "title"
print(so.isIsomorphic(s, t))
