<?php

function array_flatten($array) 
{
    $iter = new \RecursiveIteratorIterator(new \RecursiveArrayIterator($array));

    $result = array();
    foreach ($iter as $leafValue) {
        $keys = array();
        foreach (range(0, $iter->getDepth()) as $depth) {
            $keys[] = $iter->getSubIterator($depth)->key();
        }
        $result[join('.', $keys)] = $leafValue;
    }

    return $result;
}

$users_array = include 'users_array.php';
var_export(array_flatten($users_array));
