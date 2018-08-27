<?php
function permutation($str)
{
    if (empty($str)) {
        return;
    }

    doPermutation($str, 0);

}

function doPermutation($str, $pos)
{
    if ($pos == mb_strlen($str) - 1) {
        echo $str.PHP_EOL;
    } else {
        for ($i = $pos; $i < mb_strlen($str); $i++) {
            $tmp = $str[$pos];
            $str[$pos] = $str[$i];
            $str[$i] = $tmp;

            doPermutation($str, $pos + 1);

            $tmp = $str[$pos];
            $str[$pos] = $str[$i];
            $str[$i] = $tmp;
        }
    }
}

permutation('abc');