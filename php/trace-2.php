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

function dump($var, $func='', $ln=0)
{
    static $_buffer = '';
    static $_first = true;

    if ($_first) {
        $_first = false;
        register_shutdown_function('dump', '#END#');
    }

    if ($var !== '#END#') {
        if (!empty($func)) {
            $_buffer .= $func;
            if (!empty($ln)) $_buffer .= " # ($ln)";
            $_buffer .= "\n";
        }

        if (is_array($var) || is_object($var))
            $_buffer .= print_r($var, true);
        else
            $_buffer .= $var;

        $_buffer = rtrim($_buffer)."\n\n";
    } 
    else {
        if (empty($_buffer)) return;

        $body = "TIMESTAMP: ".date('Y-m-d H:i:s')."\n".$_buffer."\n";
        file_put_contents('/vagrant/trace.log', $body);
    }
}

function dpr($var, $label='')
{
    echo "<pre>\n";
    if ($label) echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

