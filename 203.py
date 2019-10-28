# coding=utf8
# 203. 移除链表元素
# 删除链表中等于给定值 val 的所有节点。
# 示例:
# 输入: 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6, val = 6
# 输出: 1 -> 2 -> 3 -> 4 -> 5


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        if not head:
            return None

        while head and head.val == val:
            head = head.next
        if not head:
            return None
        p = head
        q = head.next

        while q:
            if q.val == val:
                p.next = q.next
            else:
                p = q
            q = q.next
        return head

    def p(self, head):
        while head:
            print("%d \t" % head.val)
            head = head.next


s = Solution()
h = ListNode(1)
n2 = ListNode(2)
n3 = ListNode(2)
n4 = ListNode(1)
n5 = ListNode(5)
n6 = ListNode(6)
h.next = n2
n2.next = n3
n3.next = n4
n4.next = n5
n5.next = n6
s.p(s.removeElements(h, 2))
s.p(s.removeElements(h, 1))
