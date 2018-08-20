<?php

class ListNode
{
    private $val;
    private $next;

    public function __set($name, $value)
    {
        $this->$name = $value;
    }

    public function __get($name)
    {
        return $this->$name;
    }
}

function getList1()
{

    $node1 = new ListNode();
    $node1->val = 1;
    $node2 = new ListNode();
    $node2->val = 3;
    $node1->next = $node2;
    $node3 = new ListNode();
    $node3->val = 5;
    $node2->next = $node3;
    $node4 = new ListNode();
    $node4->val = 7;
    $node3->next = $node4;

    return $node1;
}

function getList2()
{
    $node1 = new ListNode();
    $node1->val = 2;
    $node2 = new ListNode();
    $node2->val = 4;
    $node1->next = $node2;
    $node3 = new ListNode();
    $node3->val = 6;
    $node2->next = $node3;
    $node4 = new ListNode();
    $node4->val = 8;
    $node3->next = $node4;

    return $node1;
}

function mergeList($head1, $head2)
{
    if ($head1 == null) {
        return $head2;
    }

    if ($head2 == null) {
        return $head1;
    }

    $mergeHead = null;
    if ($head1->val < $head2->val) {
        $mergeHead = $head1;
        $mergeHead->next = mergeList($head1->next, $head2);
    } else {
        $mergeHead = $head2;
        $mergeHead->next = mergeList($head1, $head2->next);
    }

    return $mergeHead;
}

var_dump(mergeList(getList1(), getList2()));