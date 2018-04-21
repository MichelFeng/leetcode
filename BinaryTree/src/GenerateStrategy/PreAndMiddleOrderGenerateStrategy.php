<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/21
 * Time: 00:22
 */

namespace LeetCode\GenerateStrategy;

use LeetCode\Node;

class PreAndMiddleOrderGenerateStrategy extends GenerateStrategy
{
    private $middleArr;
    private $pretArr;
    private $index;

    /**
     * 根据传入的数组生成树
     * @param $arr1 array 如果是根据广度优先遍历数组生成树，则$arr1是广度优先遍历数组，否则是中序遍历数组
     * @param $arr2 array 如果是根据广度优先遍历数组生成树，则$arr1是null，否则是前序或后序遍历数组
     * @return mixed
     */
    public function generate($arr1, $arr2)
    {
        $this->middleArr = $arr1;
        $this->pretArr = $arr2;
        $this->index = 0;

        return $this->doGenerateByPreAndMiddle(0, count($this->pretArr) - 1);
    }

    private function doGenerateByPreAndMiddle($start, $end)
    {
        if ($this->index < count($this->pretArr)) {
            $data = $this->pretArr[$this->index];
            $node = new Node($data);
            $pos = $this->searchKey($this->middleArr, $data, $start, $end);
            if ($pos !== false) {
                $this->index += 1;
                $left = $this->doGenerateByPreAndMiddle($start, $pos - 1);
                if (isset($left)) {
                    $node->setLeft($left);
                }
                $right = $this->doGenerateByPreAndMiddle($pos + 1, $end);
                if (isset($right)) {
                    $node->setRight($right);
                }

                return $node;
            }
        }
    }
}