<?php

class TreeNode
{
    public $val;
    public $left;
    public $right;

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

function serializeTree($root)
{
    if ($root == null) {
        echo '$,';

        return;
    }
    echo $root->val.',';
    serializeTree($root->left);
    serializeTree($root->right);
}

serializeTree(getTree());

function deserializeTree(&$root, &$serializeStr)
{
    if (($number = readChar($serializeStr))) {
        $node = new TreeNode();
        $node->val = $number;

        $root = $node;
        deserializeTree($left, $serializeStr);
        $root->left = $left;
        deserializeTree($right, $serializeStr);
        $root->right = $right;
    }
}

function readChar(&$serializeStr)
{
    if (empty($serializeStr)) {
        return false;
    }
    $tmp = explode(',', $serializeStr);
    $number = array_shift($tmp);
    if (!is_numeric($number)) {
        $number = false;
    }

    $serializeStr = implode(',', $tmp);

    return $number;
}
$root = null;
$str = '10,6,4,$,$,8,$,$,14,12,$,$,16,$,$,';
deserializeTree($root, $str);

var_dump($root);
