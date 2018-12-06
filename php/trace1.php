<?php

function pr($label, $var='')
{
    echo "<pre>\n";
    echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

function dpr()
{
    echo "<pre>\n";
    $args = func_get_args();
    foreach ($args as $var) {
        print_r($var);
    }
    echo "</pre>\n";
}

function fpr()
{
    static $first = true;

    $filename = '/tmp/ztrace.log';

    if ($first) {
        $first = false;
        $str = str_repeat('-', 80)."\n";
        error_log($str, 3, $filename);
    }

    $args = func_get_args();
    foreach ($args as $var) {
        if (is_null($var))
            $str = "NULL";
        else if ($var === FALSE)
            $str = "FALSE";
        else if ($var === '')
            $str = "''";
        else
            $str = print_r($var, true);

        error_log($str."\n", 3, $filename);
    }
}

function ftr($msg)
{
    fpr($msg);

    $trace = debug_backtrace();
    foreach ($trace as $entry) {
        if (isset($entry['file'])) {
            fpr($entry['file'] .':'. $entry['line']);
        }
    }
    fpr('');
}
