<?php

namespace LeetCode;


use LeetCode\GenerateStrategy\GenerateStrategy;
use LeetCode\TraversalStrategy\MiddleOrderTraversalTraversalStrategy;
use LeetCode\TraversalStrategy\TraversalStrategy;

class Tree
{
    private $root;

    public function __construct()
    {
        $this->root = null;
    }

    /**
     *  根据具体策略生成对应的树
     * @param $strategy
     * @param $arr1
     * @param $arr2
     * @return mixed|null
     */
    public function generate($strategy, $arr1, $arr2)
    {
        $result = null;
        if ($strategy instanceof GenerateStrategy) {
            $result = $strategy->generate($arr1, $arr2);
            $this->root = $result;
        }

        return $result;
    }


    /**
     * 根据具体策略遍历树
     * @param $strategy
     * @return null
     */
    public function traversal($strategy)
    {
        $result = null;
        if ($strategy instanceof TraversalStrategy) {
            $result = $strategy->traversal($this->root);
        }

        return $result;
    }

    /**
     * 打印输出
     * @param $result
     */
    public function print($result)
    {
        if (isset($result)) {
            echo PHP_EOL.json_encode($result).PHP_EOL;
        } else {
            echo '空树';
        }
    }

    /**
     * 获取树的深度
     * @param $root
     * @param int $depth
     * @return int
     */
    public function getDepth($root, $depth = 0)
    {
        if (isset($root)) {
            $left = $root->getLeft();
            $depthLeft = $this->getDepth($left, $depth + 1);
            $right = $root->getRight();
            $depthRight = $this->getDepth($right, $depth + 1);
            if ($depthLeft >= $depthRight) {
                $depth = $depthLeft;
            } else {
                $depth = $depthRight;
            }
        }

        return $depth;
    }

    /**
     * 判断树是不是对称的
     * @return bool
     */
    public function isSymmetric()
    {
        $isSymmetric = true;
        if (isset($this->root)) {
            $left = $this->root->getLeft();
            $right = $this->root->getRight();
            if ((isset($left) && !isset($right)) || (!isset($left) && isset($right))) {
                $isSymmetric = false;
            } else {
                // 中序遍历后，判断结果是不是回文数
                $result = $this->traversal(new MiddleOrderTraversalTraversalStrategy());
                $data = [];
                foreach ($result as $item) {
                    if (isset($item)) {
                        $data[] = $item;
                    }
                }
                $str1 = implode('', $data);
                $str2 = implode('', array_reverse($data));
                if ($str1 != $str2) {
                    $isSymmetric = false;
                }
            }
        }

        return $isSymmetric;
    }

    public function rightView($results)
    {
        $result = [];
        foreach ($results as $level) {
            $count = count($level);
            $result[] = $level[$count - 1];
        }

        return $result;
    }

    public function leftView($results)
    {
        $result = [];
        foreach ($results as $level) {
            $result[] = $level[0];
        }

        return $result;
    }
}