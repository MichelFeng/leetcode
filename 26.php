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

function getTree1()
{
    $node1 = new TreeNode();
    $node1->val = 6;
    $node2 = new TreeNode();
    $node2->val = 7;
    $node3 = new TreeNode();
    $node3->val = 8;
    $node1->left = $node2;
    $node1->right = $node3;
    $node4 = new TreeNode();
    $node4->val = 2;
    $node5 = new TreeNode();
    $node5->val = 9;
    $node3->left = $node4;
    $node3->right = $node5;
    $node6 = new TreeNode();
    $node6->val = 4;
    $node4->left = $node6;
    $node7 = new TreeNode();
    $node7->val = 5;
    $node6->left = $node7;

    return $node1;
}

function getTree2()
{
    $node1 = new TreeNode();
    $node1->val = 8;
    $node2 = new TreeNode();
    $node2->val = 2;
    $node3 = new TreeNode();
    $node3->val = 9;
    $node1->left = $node2;
    $node1->right = $node3;

    return $node1;
}

function hasSubTree($tree1, $tree2)
{
    $result = false;
    if ($tree1 != null && $tree2 != null) {
        if (isEqual($tree1->val, $tree2->val)) {
            $result = doesTree1HaveTree2($tree1, $tree2);
        }
        if (!$result) {
            $result = hasSubTree($tree1->left, $tree2);
        }
        if (!$result) {
            $result = hasSubTree($tree1->right, $tree2);
        }
    }

    return $result;
}

function isEqual($val1, $val2)
{
    return ($val1 - $val2 > -10e-7) && ($val1 - $val2 < 10e-7);
}

function doesTree1HaveTree2($tree1, $tree2)
{
    if ($tree2 == null) {
        return true;
    }
    if ($tree1 == null) {
        return false;
    }
    if (!isEqual($tree1->val, $tree2->val)) {
        return false;
    }

    return doesTree1HaveTree2($tree1->left, $tree2->left) && doesTree1HaveTree2($tree1->right, $tree2->right);
}

$tree1 = getTree1();
$tree2 = getTree2();
var_dump(hasSubTree($tree1, $tree2));