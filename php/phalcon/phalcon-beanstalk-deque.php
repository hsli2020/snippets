<?php

// Connect to the queue
$queue = new Phalcon\Queue\Beanstalk(
    array(
        'host' => 'localhost',
        'port' => '11300'
    )
);

while (1) {
    while (($job = $queue->peekReady()) !== false) {

        $message = $job->getBody();

        print_r($message);

        $job->delete();
    }
    usleep(200000);
}
