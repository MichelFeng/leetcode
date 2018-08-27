<?php

class ListNode
{
    private $val;
    private $next;
    private $sibling;

    public function __get($name)
    {
        return $this->$name;
    }

    public function __set($name, $value)
    {
        $this->$name = $value;
    }
}

function getList()
{
    $node1 = new ListNode();
    $node1->val = 'A';

    $node2 = new ListNode();
    $node2->val = 'B';

    $node3 = new ListNode();
    $node3->val = 'C';

    $node4 = new ListNode();
    $node4->val = 'D';

    $node5 = new ListNode();
    $node5->val = 'E';

    $node1->next = $node2;
    $node1->sibling = $node3;

    $node2->next = $node3;
    $node2->sibling = $node5;

    $node3->next = $node4;
    $node3->sibling = null;

    $node4->next = $node5;
    $node4->sibling = $node2;

    $node5->next = null;
    $node5->sibling = null;

    return $node1;
}

function cloneList($head)
{
    cloneNext($head);
    cloneSibling($head);
    $cloneList = splitList($head);
//    printList($head);
//    printList($cloneList);
}

function cloneNext(&$head)
{
    $node = $head;
    while ($node != null) {
        $cloneNode = new ListNode();
        $cloneNode->val = $node->val;
        $cloneNode->next = $node->next;
        $cloneNode->sibling = null;

        $node->next = $cloneNode;
        $node = $cloneNode->next;
    }
}

function cloneSibling(&$head)
{
    $node = $head;
    while ($node != null) {
        $clone = $node->next;
        if ($node->sibling != null) {
            $clone->sibling = $node->sibling->next;
        }
        $node = $clone->next;
    }
}

function splitList(&$head)
{
    $node = $head;
    $cloneHead = null;
    $cloneNode = null;
    if ($node != null) {
        $cloneHead = $cloneNode = $node->next;
        $node->next = $cloneNode->next;
        $node = $node->next;
    }

    while ($node != null) {
        $cloneNode->next = $node->next;
        $cloneNode = $cloneNode->next;
        $node->next = $cloneNode->next;
        $node = $node->next;
    }

    return $cloneHead;
}

function printList($head)
{
    $node = $head;
    while ($node != null) {
        echo $node->val.' ';
        if ($node->sibling != null) {
            echo $node->sibling->val.' ';
        }
        $node = $node->next;
    }
}

$head = getList();
cloneList($head);
