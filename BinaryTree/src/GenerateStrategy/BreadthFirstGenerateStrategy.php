<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/21
 * Time: 00:20
 */

namespace LeetCode\GenerateStrategy;

use LeetCode\Node;

class BreadthFirstGenerateStrategy extends GenerateStrategy
{

    /**
     * 根据传入的数组生成树
     * @param $arr1
     * @param $arr2
     * @return mixed
     */
    public function generate($arr1, $arr2)
    {
        if (empty($arr1)) {
            return null;
        }

        return $this->doGenerate($arr1, 1);
    }

    /**
     * 递归操作生成树
     * @param $arr
     * @param $index
     * @return Node
     */
    private function doGenerate($arr, $index)
    {
        $node = new Node($arr[$index - 1]);
        $len = count($arr);
        if ($index <= $len) {
            $leftIndex = 2 * $index;
            if ($leftIndex <= $len) {
                $leftNode = $this->doGenerate($arr, $leftIndex);
                $node->setLeft($leftNode);
            }
            $rightIndex = 2 * $index + 1;
            if ($rightIndex <= $len) {
                $rightNode = $this->doGenerate($arr, $rightIndex);
                $node->setRight($rightNode);
            }
        }

        return $node;
    }

}