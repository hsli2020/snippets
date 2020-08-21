<?php

// Since PHP 7.1+

$products = [
    [ 'id' => '#00001', 'category' => 'A' ],
    [ 'id' => '#00002', 'category' => 'B' ],
    [ 'id' => '#00003', 'category' => 'C' ],
];

foreach ($products as [ 'id' => $id, 'category' => $category]) {
    echo $id, ' ', $category, "\n";
}
