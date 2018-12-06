<?php

use Phalcon\Logger;
use Phalcon\Logger\Adapter\File as FileAdapter;

$logger = new FileAdapter("test.log");
$logger->setLogLevel(Logger::CRITICAL);

// These are the different log levels available:
#echo 'Logger::EMERGENCY = ', Logger::EMERGENCY, PHP_EOL;
#echo 'Logger::CRITICAL  = ', Logger::CRITICAL, PHP_EOL;
#echo 'Logger::ALERT     = ', Logger::ALERT, PHP_EOL;
#echo 'Logger::ERROR     = ', Logger::ERROR, PHP_EOL;
#echo 'Logger::WARNING   = ', Logger::WARNING, PHP_EOL;
#echo 'Logger::NOTICE    = ', Logger::NOTICE, PHP_EOL;
#echo 'Logger::INFO      = ', Logger::INFO, PHP_EOL;
#echo 'Logger::DEBUG     = ', Logger::DEBUG, PHP_EOL;

// These are the different log levels available:
$logger->emergency("This is an emergency message");
$logger->critical("This is a critical message");
$logger->alert("This is an alert message");
$logger->error("This is an error message");
$logger->warning("This is a warning message");
$logger->notice("This is a notice message");
$logger->info("This is an info message");
$logger->debug("This is a debug message");

// You can also use the log() method with a Logger constant:
#$logger->log("This is another error message", Logger::ERROR);

// If no constant is given, DEBUG is assumed.
#$logger->log("This is a message");
