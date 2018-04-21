<?php
/**
 * Created by PhpStorm.
 * User: michelfeng
 * Date: 2018/4/20
 * Time: 23:42
 */
include "vendor/autoload.php";

use LeetCode\GenerateStrategy\BreadthFirstGenerateStrategy;
use LeetCode\GenerateStrategy\PostAndMiddleOrderGenerateStrategy;
use LeetCode\GenerateStrategy\PreAndMiddleOrderGenerateStrategy;
use LeetCode\TraversalStrategy\BreadthFirstTraversalTraversalStrategy;
use LeetCode\TraversalStrategy\LevelTraversalStrategy;
use LeetCode\TraversalStrategy\MiddleOrderTraversalTraversalStrategy;
use LeetCode\TraversalStrategy\PostOrderTraversalTraversalStrategy;
use LeetCode\TraversalStrategy\ZigzagTraversalStrategy;
use LeetCode\Tree;
use LeetCode\TraversalStrategy\PreOrderTraversalTraversalStrategy;
use Monolog\Logger;
use Monolog\Handler\StreamHandler;

// create a log channel
$logger = new Logger('Binary Tree');
$logger->pushHandler(new StreamHandler(__DIR__.'/logs/result.log', Logger::DEBUG));

function question1(Logger $logger)
{
    $arr = [1, null, 2, null, null, 3, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new PreOrderTraversalTraversalStrategy());
    $logger->debug('前序遍历', $result);
    $tree->print($result);
}

function question2(Logger $logger)
{
    $arr = [1, null, 2, null, null, 3, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new MiddleOrderTraversalTraversalStrategy());
    $logger->debug('中序遍历', $result);
    $tree->print($result);
}

function question3(Logger $logger)
{
    $arr = [1, null, 2, null, null, 3, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new PostOrderTraversalTraversalStrategy());
    $logger->debug('后序遍历', $result);
    $tree->print($result);
}

function question4(Logger $logger)
{
    $arr = [1, null, 2, null, null, 3, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new BreadthFirstTraversalTraversalStrategy());
    $logger->debug('广度优先遍历', $result);
    $tree->print($result);
}

function question5(Logger $logger)
{
    $preOrder = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    $middleOrder = [3, 2, 5, 4, 1, 7, 9, 8, 10, 6];

    $tree = new Tree();
    $tree->generate(new PreAndMiddleOrderGenerateStrategy(), $middleOrder, $preOrder);
    $result = $tree->traversal(new PostOrderTraversalTraversalStrategy());
    $logger->debug('根据前序和中序遍历生成树', $result);
    $tree->print($result);
}

function question6(Logger $logger)
{
    $postOrder = [3, 5, 4, 2, 9, 10, 8, 7, 6, 1];
    $middleOrder = [3, 2, 5, 4, 1, 7, 9, 8, 10, 6];

    $tree = new Tree();
    $tree->generate(new PostAndMiddleOrderGenerateStrategy(), $middleOrder, $postOrder);
    $result = $tree->traversal(new PreOrderTraversalTraversalStrategy());
    $logger->debug('根据后序和中序遍历生成树', $result);
    $tree->print($result);
}

function question7(Logger $logger)
{
    $arr = [1, 2, 2, 3, 4, 4, 3];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = ($tree->isSymmetric() ? '是' : '不是').'对称的';
    $logger->debug('根据后序和中序遍历生成树', [$result]);
    echo $result.PHP_EOL;
}

function question8(Logger $logger)
{
    $arr = [1, 2, 2, null, 3, null, 3];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = ($tree->isSymmetric() ? '是' : '不是').'对称的';
    $logger->debug('根据后序和中序遍历生成树', [$result]);
    echo $result.PHP_EOL;
}

function question9(Logger $logger)
{
    $arr = null;
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = ($tree->isSymmetric() ? '是' : '不是').'对称的';
    $logger->debug('根据后序和中序遍历生成树', [$result]);
    echo $result.PHP_EOL;
}

function question10(Logger $logger)
{
    $arr = [3, 9, 20, null, null, 15, 7];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new ZigzagTraversalStrategy());
    $logger->debug('锯齿形层次遍历', $result);
    $tree->print($result);
}

function question11(Logger $logger)
{
    $arr = [1, 2, 3, 4, null, 5, null, null, 6, null, null, null, 7, null, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $result = $tree->traversal(new ZigzagTraversalStrategy());
    $logger->debug('锯齿形层次遍历', $result);
    $tree->print($result);
}

function question12(Logger $logger)
{
    $arr = [1, 2, 3, null, 5, null, 4];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $traversalResult = $tree->traversal(new LevelTraversalStrategy());
    $leftViewResult = $tree->leftView($traversalResult);
    $rightViewResult = $tree->rightView($traversalResult);
    $logger->debug('左视图：', $leftViewResult);
    $logger->debug('右视图：', $rightViewResult);
    $tree->print($leftViewResult);
    $tree->print($rightViewResult);
}

function question13(Logger $logger)
{
    $arr = [1, 2, 3, 4, null, 5, null, null, 6, null, null, null, 7, null, null];
    $tree = new Tree();
    $tree->generate(new BreadthFirstGenerateStrategy(), $arr, null);
    $traversalResult = $tree->traversal(new LevelTraversalStrategy());
    $leftViewResult = $tree->leftView($traversalResult);
    $rightViewResult = $tree->rightView($traversalResult);
    $logger->debug('左视图：', $leftViewResult);
    $logger->debug('右视图：', $rightViewResult);
    $tree->print($leftViewResult);
    $tree->print($rightViewResult);
}

question1($logger);
question2($logger);
question3($logger);
question4($logger);
question5($logger);
question6($logger);
question7($logger);
question8($logger);
question9($logger);
question10($logger);
question11($logger);
question12($logger);
question13($logger);

