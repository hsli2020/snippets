<?php # https://www.youtube.com/watch?v=-RmHhQy2m2M

// bulkInsert: 通过使用 yield 和 closure, 极大地改进可读性和执行效率

function chunkFile(string $filename, callable $generator, int $chunkSize)
{
   #$count = 0;
    $data = [];

    $file = fopen($filename, 'r');

    while (($row = fgetcsv($file, null, ',') !== false) {
        $data[] = $generator($row);
       #$count++;

       #if ($count % $chunkSize == 0) {
        if (count($data) == $chunkSize) {
            yield $data;
            $data = [];
        }
    }

    if (count($data) > 0) {
        yield $data;
    }

    fclose($file);
}

// usage

$filename = 'large-set.csv';

$generateRow = function($row) {
    return [
        'log_date'           => date('Y-m-d', strtotime($row[0])),
        'log_time'           => $row[1],
        'gauge_one_pressure' => $row[2],
        'gauge_one_temp'     => $row[3],
        'gauge_two_pressure' => $row[4],
        'gauge_two_temp'     => $row[5],
    ];
};

DB::statement('SET FOREIGN_KEY_CHECKS=0');
DB::statement('ALTER TALBE gauge_reading DISABLE KEYS');

foreach (chunkFile($filename, $generateRow, 1000) as $chunk) {
    //bulkInsert($table, $chunk);
    GaugeReading::insert($chunk); // Laravel
}

DB::statement('ALTER TALBE gauge_reading ENABLE KEYS');
DB::statement('SET FOREIGN_KEY_CHECKS=1');

// ---------------------------------------------------------
// Fastest way: use 'LOAD DATA INFILE' (mysql only)
// my.ini: local_infile = 1
// PDO::MYSQL_ATTR_LOCAL_INFILE => true

$filepath = DB::getPdo()->quote($filename);

DB::statement("
    LOAD DATA LOCAL INFILE {$filepath}
    INTO TABLE gauge_reading 
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\\n'
    (@date_var, log_time, gauge_one_pressure, gauge_one_temp, gauge_two_pressure, gauge_two_temp)
    SET log_date = STR_TO_DATE(@date_var, '%m/%d/%Y')
");
