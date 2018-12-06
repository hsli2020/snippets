<?php

const EOL = PHP_EOL;

function pr($d) { var_export($d); echo EOL; }

function insertInto($table, $columns, $data)
{
    $columnList = '`' . implode('`, `', $columns) . '`';

    $query = "INSERT INTO `$table` ($columnList) VALUES\n";

    $values = array();

    foreach($data as $row) {
        foreach($row as &$val) {
            $val = addslashes($val);
        }
        $values[] = "('" . implode("', '", $row). "')";
    }

    $update = implode(', ', 
        array_map(function($name) {
            return "`$name`=VALUES(`$name`)"; 
        }, $columns)
    );

    return $query . implode(",\n", $values) . "\nON DUPLICATE KEY UPDATE " . $update . ';';
}

$columns = [ 'name', 'url', 'rank' ];

$data = [
    [ 'A1', "B'2", 'C"3' ],
    [ 'A3', 'B4', 'C5' ],
    [ 'A5', 'B6', 'C7' ],
    [ 'A7', 'B8', 'C9' ],
];

echo insertInto('test', $columns, $data);

##-----------------------------------------------

$item = [ 'name' => 1, 'url' => 2, 'rank' => 3 ];

$update = implode(', ', 
    array_map(function($name) { 
        return "$name=VALUES($name)"; 
    }, array_keys($item))
);

#pr($update);
