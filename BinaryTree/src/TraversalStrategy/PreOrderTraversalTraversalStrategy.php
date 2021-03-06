<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/20
 * Time: 23:32
 */

namespace LeetCode\TraversalStrategy;


use LeetCode\Node;

class PreOrderTraversalTraversalStrategy extends TraversalStrategy
{

    public function traversal($root)
    {
        if (isset($root)) {
            $this->result = [];
            $this->doTraversal($root);

            return $this->result;
        }

        return null;
    }

    private function doTraversal(Node $root)
    {
        if (!empty($root->getData())) {
            $this->result[] = $root->getData();
        }
        $left = $root->getLeft();
        if (isset($left)) {
            $this->doTraversal($left);
        }
        $right = $root->getRight();
        if (isset($right)) {
            $this->doTraversal($right);
        }
    }
}