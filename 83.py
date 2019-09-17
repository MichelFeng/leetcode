# coding=utf8
# 83. 删除排序链表中的重复元素
# 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
# 示例 1:
# 输入: 1 -> 1 -> 2
# 输出: 1 -> 2
# 示例 2:
# 输入: 1 -> 1 -> 2 -> 3 -> 3
# 输出: 1 -> 2 -> 3
# Definition for singly-linked list.


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return head
        c = head
        n = head.next
        while n:
            if c.val == n.val:
                c.next = n.next
            else:
                c = c.next
            n = n.next
        return head

s = Solution()
h1 = ListNode(1)
n2 = ListNode(1)
n3 = ListNode(2)
n4 = ListNode(3)
n5 = ListNode(3)
h1.next = n2
n2.next = n3
n3.next = n4
n4.next = n5
r = s.deleteDuplicates(h1)
while r:
    print r.val
    r = r.next
