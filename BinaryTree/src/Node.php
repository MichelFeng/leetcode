<?php

namespace LeetCode;

/**
 * Class Node
 * 树的结点类
 * @package LeetCode
 */
class Node
{
    private $left; // 左孩子
    private $right; // 右孩子
    private $data; // 数据域

    public function __construct($data, $left = null, $right = null)
    {
        $this->left = $left;
        $this->right = $right;
        $this->data = $data;
    }

    public function getLeft()
    {
        return $this->left;
    }

    public function setLeft($left)
    {
        $this->left = $left;
    }

    public function getRight()
    {
        return $this->right;
    }

    public function setRight($right)
    {
        $this->right = $right;
    }

    public function getData()
    {
        return $this->data;
    }
}