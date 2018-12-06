<?php

# cd /e/LHS/Repo/beanstalkd-win/
# ./beanstalkd.exe -l 0.0.0.0 -p 11300 -b e:/ > ./error.txt &

// Connect to the queue
$queue = new Phalcon\Queue\Beanstalk(
    array(
        'host' => 'localhost',
        'port' => '11300'
    )
);

// Insert the job in the queue
#for ($i=1000; $i<1010; $i++) {
    $queue->put(
        array(
            'TestJob' => 1,
        )
    );
#}

// Insert the job in the queue with options
/*
$queue->put(
    array(
        'processVideo' => 4871
    ),
    array(
        'priority' => 250,
        'delay'    => 10,
        'ttr'      => 3600
    )
);
*/
