<?php

class Stack
{
    private $dataArray = [];
    private $minArray = [];

    function min()
    {
        if (count($this->minArray) > 0 && count($this->dataArray) > 0) {
            return $this->minArray[0];
        }

        return null;
    }

    function pop()
    {
        array_shift($this->minArray);

        return array_shift($this->dataArray);
    }

    function push($value)
    {
        array_unshift($this->dataArray, $value);
        if (count($this->minArray) == 0 || $this->minArray[0] > $value) {
            array_unshift($this->minArray, $value);
        } else {
            array_unshift($this->minArray, $this->minArray[0]);
        }
    }
}

$stack = new Stack();
$stack->push(3);
$stack->push(5);
$stack->push(2);
$stack->push(4);
$stack->push(1);
echo $stack->min();
$stack->pop();
echo $stack->min();
$stack->pop();
$stack->pop();
echo $stack->min();
