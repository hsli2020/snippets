<?php

include 'JobQueue.php';

$jobque = new JobQueue();
$jobque->push("Test", "Args");
