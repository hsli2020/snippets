<?php

class JobQueue
{
    protected $redis = null;

    const JOBQUEUE = "JOB:QUEUE";
    const JOBDELAY = "JOB:DELAY";

    public function __construct()
    {
        $this->redis = new \Redis();
        $this->redis->connect('127.0.0.1');
    }

    public function push($name, $args = '')
    {
        $job['time'] = 0; // no delay
        $job['name'] = $name;
        $job['args'] = $args;

        $info = json_encode($job);

        $this->redis->rpush(self::JOBQUEUE, $info);
    }

    public function pop($timeout = 5)
    {
        $this->checkDelayJob();

        $info = $this->redis->BLPop(self::JOBQUEUE, $timeout);

        /**
         * [
         *   [0] => JOB:QUEUE
         *   [1] => {"time":0,"name":"Test","args":"Args"}
         * ]
         */

        if ($info) {
            $job = json_decode($info[1], 1);
            return $job;
        }

        return false;
    }

    /**
     * $time: seconds or timestamp
     */
    public function delay($time, $name, $args = '')
    {
        $now = time();

        if ($time < $now) {
            $time += $now; // seconds
        }

        $job['time'] = $time;
        $job['name'] = $name;
        $job['args'] = $args;

        $info = json_encode($job);

        $this->redis->zAdd(self::JOBDELAY, $time, $info);
    }

    public function checkDelayJob()
    {
        $info = $this->redis->zRange(self::JOBDELAY, 0, 0);

        /**
         * [
         *   [0] => {"time":1592944718,"name":"Delay","args":"Args"}
         * ]
         */
        if (!$info) {
            return false;
        }

        $job = json_decode($info[0], 1);
        if (time() >= $job['time']) {
            //print_r($job);
            $this->push($job['name'], $job['args']);
            $this->redis->zRemRangeByRank(self::JOBDELAY, 0, 0);
            return true;
        }

        return false;
    }
}

//$jobque = new JobQueue();

//$jobque->push("Test", "Args");
//print_r($jobque->pop());

//$jobque->delay(10, "Delay", "Args");
//if ($jobque->checkDelayJob()) {
//    print_r($jobque->pop());
//}
