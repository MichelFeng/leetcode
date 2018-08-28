<?php

function buildHeap(&$numbers)
{
    $len = count($numbers);
    $half = $len >> 1;
    for ($i = $half - 1; $i >= 0; $i--) {

        if (($i << 1) + 1 < $len) {
            $min = ($i << 1) + 1;
            //如果有右子结点,比较左右结点的大小,如果右子结点更小,将其结点的下标记录进最小值$min
            if (($i << 1) + 2 < $len) {
                if ($numbers[($i << 1) + 2] < $numbers[$min]) {
                    $min = ($i << 1) + 2;
                }
            }
            echo $min;
            //将子结点中较小的和父结点比较,若子结点较小,与父结点交换位置,同时更新较小
            if ($numbers[$min] < $numbers[$i]) {
                $tmp = $numbers[$min];
                $numbers[$min] = $numbers[$i];
                $numbers[$i] = $tmp;

                buildHeap($numbers);
            }
        }
    }
}

$numbers = [6, 5, 2, 7, 8, 3, 1];
buildHeap($numbers);
var_dump($numbers);
