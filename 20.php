<?php
function isNumeric($str)
{
    if (!isset($str)) {
        return false;
    }
    $idx = 0;
    $numeric = scanInteger($str, $idx);

    if ($str[$idx] == '.') {
        $idx++;

        /**
         * 小数可以没有整数部分，如.123 等于 0.123
         * 小数点后面可以没有数字，如123. 等于 123.0
         * 小数点前后都有数字
         */
        $numeric = scanUnsignedInteger($str, $idx) || $numeric;
    }
    if (isset($str[$idx]) && ($str[$idx] == 'e' || $str[$idx] == 'E')) {
        $idx++;
        /**
         * e|E 后面必须有数字，且必须是整数
         */
        $numeric = $numeric && scanInteger($str, $idx);
    }

    return $numeric && !isset($str[$idx]);
}

function scanInteger($str, &$idx)
{
    if ($str[$idx] == '+' || $str[$idx] == '-') {
        $idx++;
        return scanUnsignedInteger($str, $idx);
    }

    return scanUnsignedInteger($str, $idx);
}

function scanUnsignedInteger($str, &$idx)
{
    $hasNumeric = false;
    while (isset($str[$idx]) && $str[$idx] >= '0' && $str[$idx] <= '9') {
        $idx++;
        $hasNumeric = true;
    }

    return $hasNumeric;
}

var_dump(isNumeric('123.'));
var_dump(isNumeric('.123'));
var_dump(isNumeric('-0.123e-32'));
var_dump(isNumeric('-0.123e-1.32'));
var_dump(isNumeric('-0.1a23e-32'));

