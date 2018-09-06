<?php
/**
 * 数字序列中某一位的数字
 * User: michelfeng
 * Date: 2018/9/7
 * Time: 07:09
 */

function digitAtIndex($index)
{
    if (!(isset($index) && is_numeric($index) && (int)$index >= 0)) {
        return -1;
    }

    $digits = 1;
    while (true) {
        $numbers = countOfIntegers($digits);
        if ($index < $numbers * $digits) {
            return recursiveDigitAtIndex($index, $digits);
        }
        $index -= $digits * $numbers;
        $digits++;
    }

    return -1;
}

function countOfIntegers($digits)
{
    if ($digits == 1) {
        return 10;
    }

    $count = pow(10, $digits - 1);

    return 9 * $count;
}

function recursiveDigitAtIndex($index, $digits)
{
    $number = beginNumber($digits) + (int)($index / $digits);
    $indexFromRight = $digits - $index % $digits;
    for ($i = 1; $i < $indexFromRight; $i++) {
        $number /= 10;
    }

    return $number % 10;
}

function beginNumber($digits)
{
    if ($digits == 1) {
        return 0;
    }

    return (int)pow(10, $digits - 1);
}

echo digitAtIndex(1001);