<?php

$order['qty'] = 1;

$note='  DH-123   3  ; SYN-456   2';
$note='DH-123';

/*
$parts = explode(';', trim($note));
if (count($parts) > 1) {
    foreach ($parts as $part) {
        $items = explode(' ', trim($part));
        if (count($items) > 1) {
            $sku = trim($items[0]);
            $qty = trim($items[1]);
        } else {
            $sku = trim($items[0]);
            $qty = $order['qty'];
        }
        echo $sku, ' ', $qty, PHP_EOL;
    }
}
*/

$sep = ' ;,';
$tok = strtok($note, $sep);

while ($tok !== false) {
    $sku = '';
    $qty = $order['qty'];

    // first token
    if (is_numeric($tok)) {
        $qty = $tok;
    } else if ($tok) {
        $sku = $tok;
    }

    // second token
    $tok = strtok($sep);
    if (is_numeric($tok)) {
        $qty = $tok;
    } else if ($tok) {
        $sku = $tok;
    }
    echo $sku, ' ', $qty, PHP_EOL;

    // next pair of tokens
    $tok = strtok($sep);
}
