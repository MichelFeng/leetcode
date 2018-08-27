<?php
function moreThanHalfNumber($numbers)
{
    if (empty($numbers)) {
        return -1;
    }

    $len = count($numbers);
    $res = $numbers[0];
    $times = 1;
    for ($i = 1; $i < $len; $i++) {
        if ($times == 0) {
            $res = $numbers[$i];
            $times = 1;
        } elseif ($res == $numbers[$i]) {
            $times++;
        } else {
            $times--;
        }
    }

    $times = 0;
    for ($i = 0; $i < $len; $i++) {
        if ($numbers[$i] == $res) {
            $times++;
        }
    }

    if ($times << 1 <= $len) {
        return -1;
    }

    return $res;
}
echo moreThanHalfNumber([1,2,2,2,3]); // 2
echo moreThanHalfNumber([1,2,2,3,3,4]); // -1