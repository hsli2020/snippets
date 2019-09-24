<?php

$redis = new Redis();
$redis->connect('127.0.0.1');

$quename = "delay:job";

//delayJob(15*60, "Test.php", "arg1", "arg2");
function delayJob()
{
    global $redis;

    $args = func_get_args();

    $now = time();
    if ($args[0] < $now) {
        $args[0] += $now;
    }

    if (substr($args[1], -4) != '.php') {
        $args[1] .= '.php';
    }
    $msg = json_encode($args);
    //print_r($msg);
    $redis->zAdd("delay:job", $args[0], $msg);
}

$now = time();
/*
for ($i=0; $i<10; $i++) {
    $time = $now + $i*10;
    $msg = json_encode([$time, "Test.php", "arg1", "arg2"]);
    $redis->zadd($quename, $time, $msg);
}
//*/

echo $now, PHP_EOL;

$msg = $redis->zRange($quename, 0, 0, "withscores");
if ($msg) {
    print_r($msg);

    $job = key($msg);
    $time = current($msg);

    if ($time < $now) {
        $job = json_decode($job, true);
        array_shift($job);
        $job = json_encode($job);
        echo $job;
        $redis->zDeleteRangeByRank($quename, 0, 0);
    }
}
//$redis->zDeleteRangeByRank($quename, 0, 0);

//$d = $redis->zRange($quename, 0, 0, "withscores");
//print_r($d);
//$redis->zDeleteRangeByRank($quename, 0, 0);

//$redis->zDeleteRangeByRank($quename, 0, 10);
//print_r($d);

/*
$redis->zadd($quename, 6666, "Alan Kay");
$redis->zadd($quename, 8888, "Sophie Wilson");
$redis->zadd($quename, 7777, "Richard Stallman");
$redis->zadd($quename, 9999, "Anita Borg");
$redis->zadd($quename, 1111, "Yukihiro Matsumoto");
$redis->zadd($quename, 5555, "Hedy Lamarr");
$redis->zadd($quename, 2222, "Claude Shannon");
$redis->zadd($quename, 4444, "Linus Torvalds");
$redis->zadd($quename, 3333, "Alan Turing");
//*/
