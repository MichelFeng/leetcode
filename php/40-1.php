<?php
/**
 * 使用Partition求最下的k个数
 */


/**
 * @param array $numbers
 * @param int $k
 */
function getLeastNumbers(array $numbers, int $k)
{
    if (empty($numbers) || $k > count($numbers) || $k <= 0) {
        return;
    }
    $start = 0;
    $len = count($numbers);
    $end = $len - 1;
    $index = partition($numbers, $k, $start, $end);
    while ($index != $k - 1) {
        if ($index < $k - 1) {
            $start = $index + 1;
            $index = partition($numbers, $k, $start, $end);
        } else {
            $end = $index - 1;
            $index = partition($numbers, $k, $start, $end);
        }
    }

    return array_slice($numbers, 0, $k);
}

function partition(&$numbers, $k, $start, $end)
{
    $num = $numbers[$k];
    $i = $start;
    $j = $end;
    while ($i < $j) {
        if ($numbers[$i] < $num) {
            $i++;
        }
        if ($numbers[$j] > $num) {
            $j--;
        }
        $tmp = $numbers[$i];
        $numbers[$i] = $numbers[$j];
        $numbers[$j] = $tmp;
    }

    return $i;
}

$numbers = [4, 5, 1, 6, 2, 7, 3, 8];
$res = getLeastNumbers($numbers, 6);
var_dump($res);

