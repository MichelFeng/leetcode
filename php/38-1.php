<?php

function combination($str)
{
    if (empty($str)) {
        return;
    }

    doCombination($str, mb_strlen($str), []);
}

function doCombination($str, $len, $arr)
{
    if (empty($str) || $len == 0) {
        if (!empty($arr)) {
            echo implode('', $arr).PHP_EOL;
        }

        return;
    }

    array_push($arr, $str[0]);
    doCombination(mb_substr($str, 1), $len - 1, $arr);

    array_pop($arr);
    doCombination(mb_substr($str, 1), $len, $arr);
}

combination('abcd');