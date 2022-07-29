<?php

namespace App\Support;

class Session
{
    public function set($key, $val) { echo "Session::set($key, $val)\n"; }
    public function get($key)       { echo "Session::get($key)\n"; }
}

class Cookie
{
    public function set($key, $val) { echo "Cookie::set($key, $val)\n"; }
    public function get($key)       { echo "Cookie::get($key)\n"; }
}

namespace App\Facade;

abstract class Facade
{
    /**
     * Get the registered name of the component.
     *
     * @return string
     * @throws \RuntimeException
     */
    public static function getFacadeAccessor()
    {
        throw new RuntimeException('Facade does not implement getFacadeAccessor method.');
    }

   #abstract public static function getFacadeAccessor();

    public static function getService()
    {
        return \Phalcon\Di::getDefault()->get(static::getFacadeAccessor());
    }

    public static function __callStatic($method, $args)
    {
        $instance = static::getService();
        if (! $instance) {
            throw new RuntimeException('A facade root has not been set.');
        }
        return $instance->$method(...$args);
       #return call_user_func_array([$instance, $method], $args);
    }
}

// SessionFacade
class Session extends Facade
{
    public static function getFacadeAccessor()
    {
        return 'session';
    }
}

// CookieFacade
class Cookie extends Facade
{
    public static function getFacadeAccessor()
    {
        return 'cookie';
    }
}

namespace App\Demo;

use \App\Facade\Session;
use \App\Facade\Cookie;

$di = new \Phalcon\DI\FactoryDefault();
#$app = new \Phalcon\Mvc\Application($di);

$di->set('session', function() { return new \App\Support\Session(); });
$di->set('cookie',  function() { return new \App\Support\Cookie(); });

Session::set('user', 'Joe');
Session::get('user');

Cookie::set('token', '1234asdf5678zxcv');
Cookie::get('token');
