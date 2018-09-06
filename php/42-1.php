<?php
/**
 * 顺序累加的方式求最大和
 */

/**
 * @param $numbers
 */
function getGreatestSumOfSubArray($numbers)
{
    if (empty($numbers)) {
        return;
    }

    $sum = 0;
    $max = PHP_INT_MIN;

    $len = count($numbers);
    for ($i = 0; $i < $len; $i++) {
        if ($sum <= 0) {
            $sum = $numbers[$i];
        } else {
            $sum += $numbers[$i];
        }

        if ($sum > $max) {
            $max = $sum;
        }
    }

    return $max;
}

echo getGreatestSumOfSubArray([1, -2, 3, 10, -4, 7, 2, -5]);