<?php

function pr($var)
{
    ob_start();
    var_dump($var);
    $str = preg_replace("/=>\n(\s+)/", " => ", ob_get_clean());
    echo $str, PHP_EOL;
}

function CsvFileLoad($file)
{
    $fp = fopen($file, 'r');

    $columns = fgetcsv($fp);
    $columns = array_map('trim', $columns);

    while (($values = fgetcsv($fp)) !== false) {
        $values = array_map('trim', $values);
        if (count($columns) != count($values)) {
            pr($values);
            continue;
        }

        $fields = array_combine($columns, $values);
        yield $fields;
    }

    fclose($fp);
}

$rows = CsvFileLoad("e:/data.csv");
foreach ($rows as $row) {
    print_r($row);
}
