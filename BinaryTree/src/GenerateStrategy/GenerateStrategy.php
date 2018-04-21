<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/21
 * Time: 00:18
 */

namespace LeetCode\GenerateStrategy;

abstract class GenerateStrategy
{
    /**
     * 根据传入的数组生成树
     * @param $arr1 array 如果是根据广度优先遍历数组生成树，则$arr1是广度优先遍历数组，否则是中序遍历数组
     * @param $arr2 array 如果是根据广度优先遍历数组生成树，则$arr1是null，否则是前序或后序遍历数组
     * @return mixed
     */
    abstract public function generate($arr1, $arr2);

    protected function searchKey($arr, $key, $start, $end)
    {
        for ($i = $start; $i <= $end; $i++) {
            if ($key == $arr[$i]) {
                return $i;
            }
        }

        return false;
    }
}