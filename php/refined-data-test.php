<?php
/*
Question

Write a function in PHP, named printList, that takes 2 parameters (only!): a string,
and a nested array. The function will print the value of each element on a separate
line. The value will be preceded by the string and also indicate the position and
index within the array. ie.

$list = array("stringA", array("a", "b", "c"), "stringB", array("stringC"));
with a call like printList("Foo", $list) would result in:

Foo[0] = stringA
Foo[1][0] = a
Foo[1][1] = b
Foo[1][2] = c
Foo[2] = stringB
Foo[3][0] = stringC
*/

function printList($label, $list)
{
    foreach($list as $index => $value) {
        if (is_array($value)) {
            $newLabel = $label . "[$index]";
            printList($newLabel, $value);
        } else {
            echo $label, "[$index] = $value\n";
        }
    }
}

$list = array("stringA", array("a", "b", "c"), "stringB", array("stringC"));

printList("Foo", $list);
