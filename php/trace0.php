<?php

// trace(__FILE__);
// trace(__METHOD__);
// trace($var, 'return value is');
// trace($var, __METHOD__);
// trace($var, __METHOD__, __LINE__);

function trace($var, $func='', $ln=0)
{
    static $_first = true;

    $str = '';

    if ($_first) {
        $_first = false;
        $str .= str_repeat('-', 80)."\n".date('Y-m-d H:i:s')."\n\n";
    }

    if (!empty($func)) {
        $str .= $func;
        if (!empty($ln)) $str .= " # ($ln)";
        $str .= "\n";
    }

    if (is_array($var) || is_object($var))
        $str .= print_r($var, true);
    else
        $str .= $var;

    $str = rtrim($str)."\n\n";
    file_put_contents('/vagrant/trace.log', $str, FILE_APPEND);
    #file_put_contents('/vagrant/trace.log', $str); // overwrite
}

function dpr($var, $label='')
{
    echo "<pre>\n";
    if ($label) echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

