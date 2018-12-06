<?php

// http://stackoverflow.com/questions/1397036/how-to-convert-array-to-simplexml
// http://www.lalit.org/lab/convert-php-array-to-xml-with-attributes/
// http://snipplr.com/view/3491/convert-php-array-to-xml-or-simple-xml-object-if-you-wish/
// http://www.redips.net/php/convert-array-to-xml/

$array = include "users_array.php";

//echo arrayToXml("response", $array);
//echo array2Xml("response", $array);

function arrayToXml($nodeName, $input, $indentLevel=0)
{
    if (is_numeric($nodeName)) {
        throw new Exception("cannot parse into xml. remainder :". print_r($input, true));
    }

    $indent = str_repeat('  ', $indentLevel);

    if (!(is_array($input) || is_object($input))) {
        $input = htmlspecialchars($input);
        return "$indent<$nodeName>$input</$nodeName>\n";
    }
    else {
        $newNode="$indent<$nodeName>\n";
        foreach ($input as $key => $value) {
            if (is_numeric($key)) {
                $key = substr($nodeName, 0, -1);
            }
            $newNode .= arrayToXml($key, $value, $indentLevel + 1);
        }
        $newNode .= "$indent</$nodeName>\n";
        return $newNode;
    }
}

function array2Xml($nodeName, $input, $indentLevel=0)
{
    if (is_numeric($nodeName)) {
        throw new Exception("cannot parse into xml. remainder :". print_r($input, true));
    }

    $indent = str_repeat('  ', $indentLevel);

    if (!(is_array($input) || is_object($input))) {
        $input = htmlspecialchars($input);
        return "$indent<$nodeName>$input</$nodeName>\n";
    }
    else if (isNumericIndex($input)) {
        $newNode = '';
        foreach ($input as $key => $value) {
            $newNode .= array2Xml($nodeName, $value, $indentLevel + 1);
        }
        return $newNode;
    }
    else {
        $newNode="$indent<$nodeName>\n";
        foreach ($input as $key => $value) {
            if (is_numeric($key)) {
                $key = substr($nodeName, 0, -1);
            }
            $newNode .= array2Xml($key, $value, $indentLevel + 1);
        }
        $newNode .= "$indent</$nodeName>\n";
        return $newNode;
    }
}

// this is wrong: var_dump(isNumericIndex(['a' => 1]));
function isNumericIndex(array $array)
{
    $keys = array_keys($array);
    return array_keys($keys) == $keys;
}
