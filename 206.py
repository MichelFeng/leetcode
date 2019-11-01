# coding=utf8
# 206. 反转链表
# 示例:
# 输入: 1 -> 2 -> 3 -> 4 -> 5 -> NULL
# 输出: 5 -> 4 -> 3 -> 2 -> 1 -> NULL
# 进阶:
# 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return head

        p = None
        q = head
        while q:
            t = q.next
            q.next = p
            p = q
            q = t
        return p

    def pf(self, head):
        while head:
            print(head.val)
            head = head.next


head = ListNode(1)
n1 = ListNode(2)
n2 = ListNode(3)
n3 = ListNode(4)
n4 = ListNode(5)
head.next = n1
n1.next = n2
n2.next = n3
n3.next = n4

s = Solution()
h = s.reverseList(head)
s.pf(h)
