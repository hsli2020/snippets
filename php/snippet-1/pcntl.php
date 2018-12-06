<?php

$pid = pcntl_fork(); // This function doesn't exist in windows

if ($pid == -1) {
    die('could not fork');
} else if ($pid) {
    // we are the parent
    pcntl_wait($status); //Protect against Zombie children
} else {
    // we are the child
    echo "output from child process\n";
}
