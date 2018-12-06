<?php
$start = microtime(true);
$arr = [];
for ($i = 0; $i < 1000000; $i++) {
    $time = microtime(true);
    $arr[$i.'_'.$time] = $time;
}
echo (microtime(true) - $start), "\n";
