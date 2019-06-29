# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        dummy = ListNode(0)
        dummy.next = head
        length = self.getLength(head)
        
        length -= n
        first = dummy
        while length > 0:
                length -= 1
                first = first.next
        
        first.next = first.next.next
        return dummy.next
        
    
    def getLength(self, head):
        length = 0
        while head:
            length += 1
            head = head.next
        return length
