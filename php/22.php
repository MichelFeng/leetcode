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

$node1 = new ListNode();
$node1->val = 1;
$node2 = new ListNode();
$node2->val = 2;
$node1->next = $node2;
$node3 = new ListNode();
$node3->val = 3;
$node2->next = $node3;
$node4 = new ListNode();
$node4->val = 4;
$node3->next = $node4;
$node5 = new ListNode();
$node5->val = 5;
$node4->next = $node5;

function findKthNode(ListNode $head = null, $k)
{
    if (empty($head) || $k <= 0) {
        return '输入参数错误！';
    }
    $first = $head;

    for ($i = 0; $i < $k - 1; $i++) {
        if ($first->next) {
            $first = $first->next;
        } else {
            return "输入参数k的值超过链表长度";
        }
    }
    $second = $head;
    while ($first->next) {
        $second = $second->next;
        $first = $first->next;
    }

    return $second->val;
}

echo findKthNode($node1, 1);