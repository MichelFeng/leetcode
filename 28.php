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
    $node1->val = 7;
    $node2 = new TreeNode();
    $node2->val = 7;
    $node3 = new TreeNode();
    $node3->val = 7;
    $node1->left = $node2;
    $node1->right = $node3;
    $node4 = new TreeNode();
    $node4->val = 7;
    $node5 = new TreeNode();
    $node5->val = 7;
    $node2->left = $node4;
    $node2->right = $node5;
    $node6 = new TreeNode();
    $node6->val = 7;
//    $node7 = new TreeNode();
//    $node7->val = 7;
    $node3->left = $node6;
//    $node3->right = $node7;

    return $node1;
}

function isSymmetrical($tree)
{
    return checkTreeIsSymmetrical($tree, $tree);
}

function checkTreeIsSymmetrical($tree1, $tree2)
{
    if ($tree1 == null && $tree2 == null) {
        return true;
    }
    if ($tree1 == null || $tree2 == null) {
        return false;
    }
    if ($tree1->val != $tree2->val){
        return false;
    }
    return checkTreeIsSymmetrical($tree1->left, $tree2->right)
        && checkTreeIsSymmetrical($tree1->right, $tree2->left);

}

$tree = getTree();

var_dump(isSymmetrical($tree));