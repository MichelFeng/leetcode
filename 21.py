# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = result = None
        while l1 and l2:
            if l1.val < l2.val:
                node = ListNode(l1.val)
                l1 = l1.next
            else:
                node = ListNode(l2.val)
                l2 = l2.next
            if result:
                result.next = node
                result = result.next
            else:
                result = node
                head = result
        if l1:
            if not result:
                head = result = l1
            else:
                result.next = l1
        if l2:
            if not result:
                head = result = l2
            else:
                result.next = l2
        return head
