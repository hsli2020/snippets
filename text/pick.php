<?php

function pick($arr, $keys)
{
    $result = array_intersect_key($arr, $keys);
    foreach ($keys as $old => $new) {
        $result[$new] = $result[$old];
        unset($result[$old]);
    }
    return $result;
}

$arr = [
    'u' => 'user',
    'p' => 'pass',
    'e' => 'email',
    'x' => '##',
    'y' => '##',
];

$keys = [
    'u' => 'username',
    'p' => 'password',
    'e' => 'email',
];

print_r(pick($arr, $keys));
