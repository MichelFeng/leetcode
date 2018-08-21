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
    $node2->val = 5;
    $node3 = new TreeNode();
    $node3->val = 12;
    $node1->left = $node2;
    $node1->right = $node3;
    $node4 = new TreeNode();
    $node4->val = 4;
    $node5 = new TreeNode();
    $node5->val = 7;
    $node2->left = $node4;
    $node2->right = $node5;

    return $node1;
}
