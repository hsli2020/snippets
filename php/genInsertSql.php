<?php

const EOL = "\n";  // unix, PHP_EOL is os-specific

// Create a connection with PDO options
$db = new \Phalcon\Db\Adapter\Pdo\Mysql(
    array(
        "host"     => "localhost",
        "username" => "root",
        "password" => "",
        "dbname"   => "bte",
        "options"  => array(
#           PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES \'UTF8\'",
#           PDO::ATTR_CASE               => PDO::CASE_LOWER
        )
    )
);

function genInsertSql($table, $columns, $data)
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
