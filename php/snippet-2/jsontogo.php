<?php

$text=<<<JSON
{
    "status": "OK",
    "data": "done",

    "user": {
        "firstname": "john",
        "lastname":  "smith"
    }
}
JSON;

$json = json_decode($text);
walkjson('Response', $json, 0);

function walkjson($name, $json, $indent)
{
    $varname = makeVarname($name);
    codeln($indent, "type $varname struct {");

    if (is_object($json)) {
        $indent++;
    }

    foreach ($json as $name => $value) {
        $varname = makeVarname($name);

        if (is_object($value)) {
            walkjson($name, $value, $indent);
        } else {
            codeln($indent, "$varname string ". '`json:"'. $name. '"`');
        }
    }

    if (is_object($json)) {
        $indent--;
    }
    codeln($indent, "}");
}

function getVartype($value)
{
    return "string";
}

function makeVarname($name)
{
    $arr = array_map('ucfirst', explode('-', $name));
    return implode('', $arr);
}

function codeln($indent, $code)
{
    $spaces = str_repeat(' ', $indent*4);
    echo $spaces, $code, PHP_EOL;
}
