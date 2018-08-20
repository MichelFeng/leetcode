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
$node5->next = $node3;

function meetingNode($head)
{
    if (empty($head)) {
        return null;
    }

    $slow = $head->next;
    if (empty($slow)) {
        return null;
    }

    $fast = $slow->next;
    while (!empty($fast) && !empty($slow)) {
        if ($fast == $slow) {
            return $fast;
        }
        $slow = $slow->next;
        $fast = $fast->next;
        if (!empty($fast)) {
            $fast = $fast->next;
        }
    }

    return null;
}

function findEntryNode($head)
{
    $meetingNode = meetingNode($head);
    if (empty($meetingNode)) {
        return null;
    }

    // 计算环中节点数
    $nodesInLoop = 1;
    $node = $meetingNode;
    while ($node->next != $meetingNode) {
        $node = $node->next;
        ++$nodesInLoop;
    }

    $node = $head;
    for ($i = 0; $i < $nodesInLoop; $i++) {
        $node = $node->next;
    }

    $node2 = $head;
    while ($node != $node2) {
        $node = $node->next;
        $node2 = $node2->next;
    }

    return $node;
}

var_dump(findEntryNode($node1));