<?php

// Max 100 IDs per second
function generateOrderNumber()
{
    $year = date('Y') - 2020; // 0->2020
    $ytod = date('z');
    $hms  = date('siH');
    $rand = rand(1, 100);

    if ($year >= 10) {
        $year = chr($year - 10 + ord('A'));
    }

    $orderId = sprintf("%s%03d%s%02d", $year, $ytod, $hms, $rand);

    return $orderId;
}

$orders = [];

for ($i=0; $i<20; $i++) {
    do {
        $id = generateOrderNumber();
    } while (isset($orders[$id]));

    $orders[$id] = $id;
    echo $id, PHP_EOL;
}
/*
400340091980
400340091989
400340091904
400340091975
400340091973
*/
