<?php

$subject = 'Order cancellation request for Order ID: 702-4656356-9001020';

$kwd = 'Order cancellation request';

if (strstr($subject, $kwd)) {
    // This Works
    $pat = '/Order ID: ([0-9\-]*)/';
    if (preg_match($pat, $subject, $match)) {
        print_r($match);
        # Array
        # (
        #     [0] => Order ID: 702-4656356-9001020
        #     [1] => 702-4656356-9001020
        # )
    }

    // This Works
    $pat = '/: ([0-9\-]*)/';
    if (preg_match($pat, $subject, $match)) {
        print_r($match);
        # Array
        # (
        #     [0] => : 702-4656356-9001020
        #     [1] => 702-4656356-9001020
        # )
    }

    // This DOESN'T Works
    $pat = '/([0-9\-]*)/';

    // This Works, but return different
    $pat = '/\d{3}-\d{7}-\d{7}/';
    if (preg_match($pat, $subject, $match)) {
        print_r($match);
        # Array
        # (
        #     [0] => 702-4656356-9001020
        # )
    }
}
