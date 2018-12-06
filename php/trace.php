<?php

if (!function_exists('fpr')) {

function pr($var) // Same as dpr()
{
    dpr($var);
}

function dpr($var) // print and log
{
    print_r($var);
    echo PHP_EOL;
    fpr($var);
}

function fpr()
{
    static $first = true;

    $filename = dirname(__DIR__) . '/app/logs/zzz.log';

    if ($first) {
        $first = false;
        $str = sprintf("%'-30s %s %'-30s\n", '-', date('Y-m-d H:i:s'), '-');
        if (php_sapi_name() != 'cli') {
            $str .= "\tHTTP_HOST    = ".$_SERVER['HTTP_HOST']."\n";
            $str .= "\tREQUEST_URI  = ".$_SERVER['REQUEST_URI']."\n";
            $str .= "\tQUERY_STRING = ".$_SERVER['QUERY_STRING']."\n";
#           $str .= "\tSERVER_NAME  = ".$_SERVER['SERVER_NAME']."\n";
#           $str .= "\tUSER_AGENT   = ".$_SERVER['HTTP_USER_AGENT']."\n";
#           if (isset($_SERVER['HTTP_REFERER']))
#               $str .= "\tHTTP_REFERER = ".$_SERVER['HTTP_REFERER']."\n";
            $str .= "\n";
        }
        error_log($str, 3, $filename);
    }

    $args = func_get_args();
    foreach ($args as $var) {
        $str = var_export($var, true);
        $str = preg_replace("/=> \n(\s+)/", "=> ", $str);
        error_log(trim($str, "'")."\n", 3, $filename);
    }
    error_log("\n", 3, $filename);
}

function ftr($msg)
{
    fpr($msg);

    $files = "{>>>\n";
    $trace = debug_backtrace();
    foreach ($trace as $entry) {
        if (isset($entry['file'])) {
            $files .= $entry['file'] .':'. $entry['line'] . "\n";
        }
    }
    $files .= "<<<}";
    fpr($files);
}

function &timer_fetch()
{
    static $timers = [ '__names__' => [] ];
    return $timers;
}

function timer_start($name)
{
    $timers = &timer_fetch();
    array_push($timers['__names__'], $name);
    $timers[$name] = microtime(true);
}

function timer_end($name = null)
{
    $timers = &timer_fetch();
    if (!$name) {
        $name = array_pop($timers['__names__']);
    } else {
        $key = array_search($name, $timers['__names__']);
        unset($timers['__names__'][$key]);
    }
    $start = $timers[$name];
    $end = microtime(true);
    $timers[$name] = number_format($end - $start, 3);
}

}
