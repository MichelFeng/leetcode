<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/21
 * Time: 15:38
 */

namespace LeetCode\TraversalStrategy;


class ZigzagTraversalStrategy extends TraversalStrategy
{

    public function traversal($root)
    {
        $traversalResult = array_fill(0, $this->getNodeCount($root) + 1, null);
        $this->doTraversal($root, $traversalResult, 1);
        array_shift($traversalResult);

        $levelResults = [];
        $depth = $this->getDepth($root);
        for ($idx = 0; $idx < $depth; $idx++) {
            $startIdx = pow(2, $idx) - 1;
            $endIdx = pow(2, $idx + 1) - 1;
            $levelResult = [];
            for ($i = $startIdx; $i < $endIdx; $i++) {
                if (isset($traversalResult[$i])) {
                    $levelResult[] = $traversalResult[$i];
                }
            }
            $levelResults[] = $levelResult;
        }
        $result = [];
        foreach ($levelResults as $key => $level) {
            if ($key % 2 == 1) {
                $result[] = array_reverse($level);
            } else {
                $result[] = $level;
            }
        }

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