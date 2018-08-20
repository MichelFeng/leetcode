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


function reverseList($head)
{
    $prev = null;
    $node = $head;
    $reverseHead = null;
    while (!empty($node)) {
        $next = $node->next;
        if (empty($next)) {
            $reverseHead = $node;
        }

        $node->next = $prev;
        $prev = $node;
        $node = $next;
    }

    return $reverseHead;
}

var_dump(reverseList(null));
var_dump(reverseList($node1));
var_dump(reverseList($node5));