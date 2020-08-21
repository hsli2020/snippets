<?php

function dd()
{
   array_map(function($x) { var_dump($x); }, func_get_args()); die;
}

function dd()
{
    array_map(function ($v) {
        $str = json_encode($v, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
        echo $str, PHP_EOL;
    }, func_get_args());

    die(0);
}

function dd(...$args)
{
	die(var_dump(...$args));
}

function fpr()
{
    $list = array_map(function($v) {
        $time = date("Y-m-d H:i:s");
        $json = json_encode($v, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
        return $time." ".$json;
    }, func_get_args());

    $list[] = PHP_EOL;

    $msg = implode(PHP_EOL, $list);
    error_log($msg, 3, "notice.log");
}

//$str = "Something";
//$arr = [ "One" => 1, "Two" => 2, "Month" => [ "Jan", "Feb", "Mar" ], 11 => 22 ];

//fpr($arr, $arr, $str, $str, $arr);
//dd($arr, $str);
