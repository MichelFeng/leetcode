<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/20
 * Time: 23:58
 */

namespace LeetCode\TraversalStrategy;


class BreadthFirstTraversalTraversalStrategy extends TraversalStrategy
{

    public function traversal($root)
    {
        $result = array_fill(0, $this->getNodeCount($root) + 1, null);
        $this->doTraversal($root, $result, 1);
        array_shift($result);

        return $result;
    }

    private function doTraversal($root, &$result, $index)
    {
        if (isset($root)) {
            $data = $root->getData();
            $result[$index] = $data;
            $left = $root->getLeft();
            $this->doTraversal($left, $result, $index * 2);
            $right = $root->getRight();
            $this->doTraversal($right, $result, $index * 2 + 1);
        }
    }

    private function getNodeCount($root)
    {
        return pow(2, $this->getDepth($root)) - 1;
    }

    private function getDepth($root, $depth = 0)
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
}