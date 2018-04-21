<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/20
 * Time: 23:31
 */

namespace LeetCode\TraversalStrategy;

abstract class TraversalStrategy
{
    protected $result;

    abstract public function traversal($root);
}