<?php

function reOrder(array $arr, Closure $compare)
{
    if (empty($arr)) {
        return;
    }
    $start = 0;
    $end = count($arr) - 1;

    while ($start < $end) {
        while ($start < $end && $compare($arr[$start]) == 1) {
            $start++;
        }

        while ($start < $end && $compare($arr[$end]) == 0) {
            $end--;
        }

        if ($start < $end) {
            $tmp = $arr[$start];
            $arr[$start] = $arr[$end];
            $arr[$end] = $tmp;
        }
    }

    return $arr;
}


var_dump(reOrder([1,2,3,4,5,6,7], function($param){
    return $param & 1;
}));