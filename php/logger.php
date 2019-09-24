<?php

class Logger
{
    protected static $filename = "app.log";
    protected static $echo = true;
    protected static $disabled = false;

    public static function setFilename($filename)
    {
        self::$filename = $filename;
    }

    public static function setEcho($echo)
    {
        self::$echo = $echo;
    }

    public static function disable()
    {
        self::$disabled = true;
    }

    public static function error($message)
    {
        self::log('ERROR', $message);
    }

    public static function info($message)
    {
        self::log('INFO', $message);
    }

    public static function debug($message)
    {
        self::log('DEBUG', $message);
    }

    public static function log($level, $message)
    {
        if (self::$disabled) {
            return;
        }

        if (self::$echo) {
            echo $message, "\n";
        }

        $filename = self::$filename;
        $msg = date('Y-m-d H:i:s'). " [$level] $message\n";
        error_log($msg, 3, $filename);
    }
}

/*
Logger::setFilename("new.log");
Logger::setEcho(false);
Logger::disable();

Logger::info("Hello");
Logger::info("Bye");

Logger::error("Hello");
Logger::error("Bye");

Logger::debug("Hello");
Logger::debug("Bye");
//*/
