<?php
/**
 * @param array $arr 二维数组
 * @param int $columns 列数
 * @param int $rows 行数
 */
function printMatrixClockwisely(array $arr, int $columns, int $rows)
{
    if ($arr == null || $columns == 0 || $rows == 0) {
        return;
    }
    $start = 0;
    while ($columns > $start * 2 && $rows > $start * 2) {
        printMatrixInCircle($arr, $columns, $rows, $start);
        $start++;
    }
}

/**
 * @param array $arr 二维数组
 * @param int $columns 列数
 * @param int $rows 行数
 * @param int $start 圈号
 */
function printMatrixInCircle(array $arr, int $columns, int $rows, int $start)
{
    // 从0开始，需要额外减1
    $endRow = $rows - $start - 1;
    $endColumn = $columns - $start - 1;

    // 从左至右输出
    for ($i = $start; $i <= $endColumn; $i++) {
        $num = $arr[$start][$i];
        printNum($num);
    }

    // 从上至下输出
    if ($endRow > $start) {
        for ($i = $start + 1; $i <= $endRow; $i++) {
            $num = $arr[$i][$endColumn];
            printNum($num);
        }
    }
    // 从右至左输出
    if ($endRow > $start && $endColumn > $start) {
        for ($i = $endColumn - 1; $i >= $start; $i--) {
            $num = $arr[$endRow][$i];
            printNum($num);
        }
    }

    // 从下至上输出
    if ($endColumn > $start && $endRow - $start > 1) {
        for ($i = $endRow - 1; $i > $start; $i--) {
            $num = $arr[$i][$start];
            printNum($num);
        }
    }
}

function printNum($num)
{
    printf('%3d', $num);
}


$arr = [[1]];
printMatrixClockwisely($arr, 1,1);