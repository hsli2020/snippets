<?php

class Timer
{
	protected static $timers = [];

    public static function start($name)
    {
        self::$timers[$name] = microtime(true);
    }

    public static function stop($name)
    {
        $start = self::$timers[$name];
        $end = microtime(true);
        self::$timers[$name] = number_format($end - $start, 2);
    }

    public static function get($name = '')
    {
        if ($name) {
            return self::$timers[$name] ?? '';
        }
        return self::$timers;
    }
}

/*
Timer::start('t1');
Timer::start('t2');

usleep(100000);
Timer::stop('t2');

usleep(100000);
Timer::stop('t1');

echo Timer::get('t1'), "\n";
echo Timer::get('t2'), "\n";
//*/
