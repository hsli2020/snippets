<?php

$di = new Phalcon\DI\FactoryDefault();

$di->set('logger', function ($filename = null, $format = null) {
    // format
    $lineFmt = '%date% [%type%] %message%';
    $dateFmt = 'Y-m-d H:i:s';

    $formatter = new Phalcon\Logger\Formatter\Line($lineFmt, $dateFmt);

    // logger
    $filename = "app.log";

    $logger = new Phalcon\Logger\Adapter\File($filename);
    $logger->setFormatter($formatter);

    return $logger;
});

class Foo extends Phalcon\Di\Injectable
{
    public function log($msg)
    {
        echo $msg, "\n";
        $this->logger->info($msg);
    }

    public function error($msg)
    {
        echo 'ERROR: ', $msg, "\n";
        $this->logger->error($msg);
    }
}

$foo = new Foo(); // DI injected !!!

$foo->log("Why does this work?");
$foo->error("I cann't figure out");
