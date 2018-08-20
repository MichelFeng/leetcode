<?php

function verifySequenceOfBST($sequence)
{
    if (empty($sequence)) {
        return false;
    }
    $count = count($sequence);
    $root = $sequence[$count - 1];

    for ($i = 0; $i < $count - 1; $i++) {
        if ($sequence[$i] > $root) {
            break;
        }
    }

    for ($j = $i; $j < $count - 1; $j++) {
        if ($sequence[$j] < $root) {
            return false;
        }
    }
    $isLeftValid = true;
    if ($i > 0) {
        $isLeftValid = verifySequenceOfBST(array_slice($sequence, 0, $i));
    }

    $isRightValid = true;
    if ($j > 0) {
        $isRightValid = verifySequenceOfBST(array_slice($sequence, $i, $count - $i - 1));
    }

    return $isLeftValid && $isRightValid;
}

var_dump(verifySequenceOfBST([5, 7, 6, 9, 11, 10, 8]));