# coding=utf8
# 160. 相交链表
# 编写一个程序，找到两个单链表相交的起始节点。
# ：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
# 输出：Reference of the node with value = 8
# 输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为[4, 1, 8, 4, 5]，链表 B 为[5, 0, 1, 8, 4, 5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

# 输入：intersectVal = 2, listA = [0, 9, 1, 2, 4], listB = [3, 2, 4], skipA = 3, skipB = 1
# 输出：Reference of the node with value = 2
# 输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为[0, 9, 1, 2, 4]，链表 B 为[3, 2, 4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        双指针法
        pA = headA + headB
        pB = headB + headA
        长度一致，如果遍历过程中遇到相同的节点则终止，此时节点为相交节点
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        if not headA or not headB:
            return None
        pA = headA
        pB = headB
        while pA != pB:
            pA = pA.next if pA else headB
            pB = pB.next if pB else headA

        return pA


s = Solution()
ha = ListNode(0)
n1 = ListNode(9)
n2 = ListNode(1)
n3 = ListNode(2)
n4 = ListNode(4)
ha.next = n1
n1.next = n2
n2.next = n3
n3.next = n4
hb = ListNode(3)
hb.next = n3

print s.getIntersectionNode(ha, hb)
