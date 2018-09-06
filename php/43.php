<?php
function numberOf1Between1AndN($n)
{
    if ($n <= 0) {
        return 0;
    }

    return numberOf1($n);
}

function numberOf1($n)
{
    if (empty($n) || !is_numeric($n)) {
        return 0;
    }

    $n = (string)$n;

    $first = $n[0];
    $len = strlen($n);
    if ($len == 1 && $first == 0) {
        return 0;
    }
    if ($len == 1 && $first > 0) {
        return 1;
    }

    $numFirstDigit = 0;
    if ($first > 1) {
        $numFirstDigit = pow(10, $len - 1);
    } elseif ($first == 1) {
        $numFirstDigit = intval(substr($n, 1)) + 1;
    }
    $numOtherDigits = $first * ($len - 1) * pow(10, $len - 2);
    $numRecursive = numberOf1(substr($n, 1));

    return $numFirstDigit + $numOtherDigits + $numRecursive;
}

echo numberOf1Between1AndN(999);