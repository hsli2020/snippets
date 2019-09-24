<?php

try {
    $mng = new MongoDB\Driver\Manager("mongodb://localhost:27017");

    // dbstats
    $stats = new MongoDB\Driver\Command(["dbstats" => 1]);
    $res = $mng->executeCommand("test", $stats);
    $stats = current($res->toArray());
    print_r($stats);
    echo str_repeat('-', 40), PHP_EOL;

    // list_databases
    $listdatabases = new MongoDB\Driver\Command(["listDatabases" => 1]);
    $res = $mng->executeCommand("admin", $listdatabases);
    $databases = current($res->toArray());
    foreach ($databases->databases as $el) {
        echo $el->name . "\n";
    }
    echo str_repeat('-', 40), PHP_EOL;

    // read_all
    $query = new MongoDB\Driver\Query([]); 
    $rows = $mng->executeQuery("test.trainers", $query);
    foreach ($rows as $row) {
        echo "$row->name : $row->age, $row->city\n";
    }
    echo str_repeat('-', 40), PHP_EOL;

    // filtering
    $filter = [ 'name' => 'Brock' ]; 
    $query = new MongoDB\Driver\Query($filter);     
    $res = $mng->executeQuery("test.trainers", $query);
    $car = current($res->toArray());
    
    if (!empty($car)) {
        echo "$row->name : $row->age, $row->city\n";
    } else {
        echo "No match found\n";
    }
    echo str_repeat('-', 40), PHP_EOL;
    
    // projection
    $filter = [];
    $options = ["projection" => ['_id' => 0]];
    $query = new MongoDB\Driver\Query($filter, $options);
    $rows = $mng->executeQuery("test.trainers", $query);
    foreach ($rows as $row) {
       print_r($row);
    }    
    echo str_repeat('-', 40), PHP_EOL;

    // read_limit
    $query = new MongoDB\Driver\Query([], ['sort' => [ 'name' => 1], 'limit' => 5]);     
    $rows = $mng->executeQuery("test.trainers", $query);
    foreach ($rows as $row) {
        echo "$row->name : $row->age\n";
    }

    // bulkwrite
    $bulk = new MongoDB\Driver\BulkWrite;
    $doc = ['_id' => new MongoDB\BSON\ObjectID, 'name' => 'Toyota', 'price' => 26700];
    $bulk->insert($doc);
    $bulk->update(['name' => 'Audi'], ['$set' => ['price' => 52000]]);
    $bulk->delete(['name' => 'Hummer']);
    $mng->executeBulkWrite('test.trainers', $bulk);
}
catch (MongoDB\Driver\Exception\Exception $e) {
    $filename = basename(__FILE__);
    
    echo "The $filename script has experienced an error.\n"; 
    echo "It failed with the following exception:\n";
    
    echo "Exception: ", $e->getMessage(), "\n";
    echo "In file:   ", $e->getFile(), "\n";
    echo "On line:   ", $e->getLine(), "\n";       
}
