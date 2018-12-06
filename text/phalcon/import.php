<?php

const EOL = PHP_EOL;

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


$columns = [
    'time',
    'error',
    'low_alarm',
    'high_alarm',
    'dcvolts',
    'kw',
    'kwh'
];

foreach (glob("csv/*.csv") as $filename) {
    $devid = substr($filename, 0, 6);
    echo $devid, PHP_EOL;
    if (($handle = fopen($filename, "r")) !== FALSE) {
        fgetcsv($handle); // skip first line
        while (($fields = fgetcsv($handle, 1000, ",")) !== FALSE) {
            $data = array_combine($columns, $fields);
            $data['devid'] = $devid;
            $db->insertAsDict('solar', $data);
        }
        fclose($handle);
    }
}
