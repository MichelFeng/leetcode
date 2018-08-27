<?php

class TreeNode
{
    private $val;
    private $left;
    private $right;

    public function __set($name, $value)
    {
        $this->$name = $value;
    }

    public function __get($name)
    {
        return $this->$name;
    }
}

function getTree()
{
    $node1 = new TreeNode();
    $node1->val = 10;
    $node2 = new TreeNode();
    $node2->val = 6;
    $node3 = new TreeNode();
    $node3->val = 14;
    $node1->left = $node2;
    $node1->right = $node3;
    $node4 = new TreeNode();
    $node4->val = 4;
    $node5 = new TreeNode();
    $node5->val = 8;
    $node2->left = $node4;
    $node2->right = $node5;
    $node6 = new TreeNode();
    $node6->val = 12;
    $node7 = new TreeNode();
    $node7->val = 16;
    $node3->left = $node6;
    $node3->right = $node7;

    return $node1;
}

function Convert($root)
{
    $lastInList = null;
    ConvertNode($root, $lastInList);

    $head = $lastInList;
    while ($head != null && $head->left != null) {
        $head = $head->left;
    }

    return $head;
}

function ConvertNode($root, &$lastInList)
{
    if ($root == null) {
        return;
    }
    $current = $root;
    if ($current->left != null) {
        ConvertNode($current->left, $lastInList);
    }

    $current->left = $lastInList;
    if ($lastInList != null) {
        $lastInList->right = $current;
    }
    $lastInList = $current;
    if ($current->right != null) {
        ConvertNode($current->right, $lastInList);
    }
}

function printList($head)
{
    $last = $head;
    while ($head != null) {
        echo $head->val.' ';
        $head = $head->right;
        if ($head != null) {
            $last = $head;
        }
    }
    echo PHP_EOL;
    while ($last != null) {
        echo $last->val.' ';
        $last = $last->left;
    }
}

printList(Convert(getTree()));