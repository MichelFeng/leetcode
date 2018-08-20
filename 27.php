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
    $node1->val = 8;
    $node2 = new TreeNode();
    $node2->val = 6;
    $node3 = new TreeNode();
    $node3->val = 10;
    $node1->left = $node2;
    $node1->right = $node3;
    $node4 = new TreeNode();
    $node4->val = 5;
    $node5 = new TreeNode();
    $node5->val = 7;
    $node2->left = $node4;
    $node2->right = $node5;
    $node6 = new TreeNode();
    $node6->val = 9;
    $node7 = new TreeNode();
    $node7->val = 11;
    $node3->left = $node6;
    $node3->right = $node7;

    return $node1;
}

function mirror(&$tree)
{
    if ($tree != null) {
        $tmp = $tree->left;
        $tree->left = $tree->right;
        $tree->right = $tmp;
        if ($tree->left!=null) {
            mirror($tree->left);
        }
        if ($tree->right!=null){
            mirror($tree->right);
        }
    }
}

$tree = getTree();
mirror($tree);
var_dump($tree);