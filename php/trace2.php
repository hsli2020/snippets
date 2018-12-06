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
       #$str = str_repeat('-', 30).' '.date('Y-m-d H:i:s').' '.str_repeat('-', 30)."\n";
        $str = sprintf("%'-30s %s %'-30s\n", '-', date('Y-m-d H:i:s'), '-');
        $str .= "\tHTTP_HOST    = ".$_SERVER['HTTP_HOST']."\n";
        $str .= "\tREQUEST_URI  = ".$_SERVER['REQUEST_URI']."\n";
        $str .= "\tQUERY_STRING = ".$_SERVER['QUERY_STRING']."\n";
        if (isset($_SERVER['HTTP_REFERER']))
            $str .= "\tHTTP_REFERER = ".$_SERVER['HTTP_REFERER']."\n";
        $str .= "\n";
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
    fpr(' ');
}

