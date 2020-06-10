<?php

include 'JobQueue.php';

$jobQueue = new JobQueue();

echo "Job worker is running\n";

$pause = 0;
$start = time();

while (1) {
    $job = $jobQueue->pop();

    if ($job) {
        $name = $job['name'];
        $args = $job['args'];
        $file = "$name.php";

        prlog("Run Job: $name $args");

        if (file_exists($file)) {
           #exec('psexec -d c:/xampp/php64/php.exe ' . $file);
            exec("c:/xampp/php64/php.exe $file $args");
        } else {
            prlog("Error: $file not found");
        }

        prlog("Job End: $name\n");
    }

    sleep(1);

    if (time() - $start > 290) { // 5 minutes
        break;
    }
}

function prlog($message)
{
    $filename = 'app/logs/job-worker.log';
    $filename = 'job-worker.log';

    echo "$message\n";

    if (file_exists($filename)) {
        // if date changed, add a blank line to seperate
        if (date('d') != date('d', filemtime($filename))) {
            error_log("\n", 3, $filename);
        }

        // if file is too big, purge it
        if (filesize($filename) > 128*1024) {
            unlink($filename);
        }
    }

    error_log(date('Y-m-d H:i:s').' '.$message. "\n", 3, $filename);
}
