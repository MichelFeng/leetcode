## 3. 无重复字符的最长子串
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。


#### 解题思路
从第一个字符开始向后遍历，若当前字符不出现在已遍历的字符串中，则继续向后移动，如果出现在已遍历的字符串，则计算字符串长度与最大长度比较，若大于则替换，
否则将起始标志位设置在该字符在前面字符串中的索引+1的位置，然后继续向后遍历字符串
