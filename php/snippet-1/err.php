<?php

function logError($msg)   { trigger_error($msg, E_USER_ERROR); }
function logWarning($msg) { trigger_error($msg, E_USER_WARNING); }
function logNotice($msg)  { trigger_error($msg, E_USER_NOTICE); }

function errorHandler($num, $str, $file, $line, $context = null)
{
    $types = [
        E_USER_ERROR => "Error: ",
        E_USER_WARNING => "Warning: ",
        E_USER_NOTICE => "Notice: ",

        E_ERROR => "ERROR: ",
        E_WARNING => "WARNING: ",
        E_NOTICE => "NOTICE: ",
    ];

    $type = "ERROR: ";

    if (isset($types[$num])) {
        $type = $types[$num];
    }

    exceptionHandler(new ErrorException($type.$str, 0, $num, $file, $line));
}

function exceptionHandler(Exception $e)
{
    $file = "errors.log";

    $message  = $e->getMessage() . PHP_EOL;
    $message .= "\t";
    $message .= str_replace($e->getFile(), '\\', '/').':'.$e->getLine().PHP_EOL;

    file_put_contents($file, $message, FILE_APPEND);
}

function checkForFatal()
{
    $error = error_get_last();
    if ($error["type"] == E_ERROR)
        errorHandler($error["type"], $error["message"], $error["file"], $error["line"]);
}

register_shutdown_function("checkForFatal");
set_error_handler("errorHandler");
set_exception_handler("exceptionHandler");
ini_set("display_errors", "off");
error_reporting(E_ALL);

file_get_contents("c:\aa.txt");
logError("file is missing");
$a = [];
$b = $a[2];
