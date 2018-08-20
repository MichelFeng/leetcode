<?php
function isPopOrder($in, $out)
{
    if (empty($in) || empty($out) || count($in) != count($out)) {
        return false;
    }

    $arr = [];
    $inIdx = 0;
    $outIdx = 0;
    $inCount = count($in);
    $outCount = count($out);
    while ($outIdx < $outCount) {
        while (empty($arr) || $arr[0] != $out[$outIdx]) {
            if ($inIdx == $inCount) {
                break;
            }
            array_unshift($arr, $in[$inIdx]);
            $inIdx++;
        }
        if ($arr[0] != $out[$outIdx]) {
            break;
        }
        array_shift($arr);
        $outIdx++;
    }
    if (empty($arr) && $outIdx == $outCount) {
        return true;
    }

    return false;
}

var_dump(isPopOrder([1,2,3,4,5], [5,4,3,2,1]));
var_dump(isPopOrder([1,2,3,4,5], [4,5,3,2,1]));
var_dump(isPopOrder([1,2,3,4,5], [5,4,3,1,2]));