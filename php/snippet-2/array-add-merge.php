<?php

$default = [ 'a' => 11, 'b' => 22, 'c' => 33 ];

$options = [ 'a' => 111, 'd' => 44, 'c' => 333 ];
$options = [ ];

print_r($default + $options);
print_r(array_merge($default, $options));
