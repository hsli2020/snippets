<?php

use Phalcon\Cache\Backend\Redis;
use Phalcon\Cache\Frontend\Data as FrontData;

function test_Redis_in_Phalcon() // slow
{
    // Cache data for 2 days
    $frontCache = new FrontData([
        'lifetime' => 172800
    ]);

    // Create the Cache setting redis connection options
    $cache = new Redis($frontCache, [
        'host' => 'localhost',
        'port' => 6379,
    #   'auth' => 'foobared',
        'persistent' => false,
        'index' => 0,
    ]);

    // Cache arbitrary data
    $cache->save('my-data', [1, 2, 3, 45]);

    // Get data
    $data = $cache->get('my-data');
    print_r($data);

    #$data = $cache->get('foo');
    #var_dump($data);
}

function test_Redis_in_PHP() // faster
{
    $redis = new \Redis();
    $redis->connect('127.0.0.1');

    $data = [1,2,3,4];
    $redis->set('foo', json_encode($data));

    $value = $redis->get('foo');
    print_r(json_decode($value));

    #$value = $redis->get('my-data');
    #print_r($value);
}

//test_Redis_in_PHP();
//test_Redis_in_Phalcon();
