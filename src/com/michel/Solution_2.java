package com.michel;

public class Solution_2 {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {

        ListNode head = new ListNode(-1);
        ListNode p = head;
        int left = 0;
        while (true) {
            int l1num = 0;
            if (l1 != null){
                l1num = l1.val;
                l1 = l1.next;
            }
            int l2num = 0;
            if (l2 != null){
                l2num = l2.val;
                l2 = l2.next;
            }
            int sum = l1num + l2num + left;
            left = 0;
            if (sum > 9){
                left = 1;
                sum -= 10;
            }

            ListNode node = new ListNode(sum);
            p.next=node;
            p = p.next;

            if (l1 == null && l2 == null){
                if (left > 0){
                    ListNode node2 = new ListNode(left);
                    p.next=node2;
                    p = p.next;
                }
                return head.next;
            }
        }
    }

    public static void main(String[] args) {
        ListNode l1 = new ListNode(9);

        ListNode l2 = new ListNode(1);
        ListNode l21 = new ListNode(9);
        l2.next = l21;
        ListNode l22 = new ListNode(9);
        l21.next = l22;
        ListNode l23 = new ListNode(9);
        l22.next = l23;
        ListNode l24 = new ListNode(9);
        l23.next = l24;

        Solution_2 solution = new Solution_2();
        solution.print(l1);
        solution.print(l2);
        ListNode result = solution.addTwoNumbers(l1, l2);
        solution.print(result);

    }

    public void print(ListNode node) {
        while (node != null) {
            System.out.print(node.val);
            node = node.next;
        }
        System.out.println();
    }
}