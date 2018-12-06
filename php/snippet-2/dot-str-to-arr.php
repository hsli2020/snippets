<?php

$str = 'one.two.three.four.five.six';
$value = 'is val';

$arr = explode('.', $str);

array_push($arr, $value);

while(count($arr) > 1){
    $val = array_pop($arr);
    $key = array_pop($arr);
    array_push($arr, [ $key => $val ]);
}
$res = array_pop($arr);
print_r($res);
