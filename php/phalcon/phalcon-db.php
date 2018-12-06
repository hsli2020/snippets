Database Abstraction Layer
==========================

Connecting to Databases
-----------------------

<?php  // Mysql

    // Required
    $config = array(
        "host"     => "127.0.0.1",
        "username" => "mike",
        "password" => "sigma",
        "dbname"   => "test_db",
        "options"  => array(
            PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES \'UTF8\'",
            PDO::ATTR_CASE               => PDO::CASE_LOWER
        )
    );

    // Optional
    $config["persistent"] = false;

    // Create a connection
    $connection = new \Phalcon\Db\Adapter\Pdo\Mysql($config);

<?php  // Postgresql

    // Required
    $config = array(
        "host"     => "localhost",
        "username" => "postgres",
        "password" => "secret1",
        "dbname"   => "template"
    );

    // Optional
    $config["schema"] = "public";

    // Create a connection
    $connection = new \Phalcon\Db\Adapter\Pdo\Postgresql($config);

<?php  // Sqlite

    // Required
    $config = array(
        "dbname" => "/path/to/database.db"
    );

    // Create a connection
    $connection = new \Phalcon\Db\Adapter\Pdo\Sqlite($config);

Finding Rows
------------

<?php

    $sql = "SELECT id, name FROM robots ORDER BY name";

    // Send a SQL statement to the database system
    $result = $connection->query($sql);

    // Print each robot name
    while ($robot = $result->fetch()) {
       echo $robot["name"];
    }

    // Get all rows in an array
    $robots = $connection->fetchAll($sql);
    foreach ($robots as $robot) {
       echo $robot["name"];
    }

    // Get only the first row
    $robot = $connection->fetchOne($sql);

Phalcon\Db\Result::setFetchMode()
     Phalcon\Db::FETCH_NUM    | Return an array with numeric indexes
     Phalcon\Db::FETCH_ASSOC  | Return an array with associative indexes
     Phalcon\Db::FETCH_BOTH   | Return an array with both associative and numeric indexes
     Phalcon\Db::FETCH_OBJ    | Return an object instead of an array

<?php

    $sql = "SELECT id, name FROM robots ORDER BY name";
    $result = $connection->query($sql);

    $result->setFetchMode(Phalcon\Db::FETCH_NUM);
    while ($robot = $result->fetch()) {
       echo $robot[0];
    }

Phalcon\Db::query() returns an instance of Phalcon\\Db\\Result\\Pdo.

<?php

    $sql = "SELECT id, name FROM robots";
    $result = $connection->query($sql);

    // Traverse the resultset
    while ($robot = $result->fetch()) {
       echo $robot["name"];
    }

    // Seek to the third row
    $result->seek(2);
    $robot = $result->fetch();

    // Count the resultset
    echo $result->numRows();

Binding Parameters
------------------
Both string and positional placeholders are supported.

<?php

    // Binding with numeric placeholders
    $sql    = "SELECT * FROM robots WHERE name = ? ORDER BY name";
    $result = $connection->query($sql, array("Wall-E"));

    // Binding with named placeholders
    $sql     = "INSERT INTO `robots`(name, year) VALUES (:name, :year)";
    $success = $connection->query($sql, array("name" => "Astro Boy", "year" => 1952));

<?php

    // Binding with PDO placeholders
    $sql    = "SELECT * FROM robots WHERE name = ? ORDER BY name";
    $result = $connection->query($sql, array(1 => "Wall-E"));

Inserting/Updating/Deleting Rows
--------------------------------

<?php

    // Inserting data with a raw SQL statement
    $sql     = "INSERT INTO `robots`(`name`, `year`) VALUES ('Astro Boy', 1952)";
    $success = $connection->execute($sql);

    // With placeholders
    $sql     = "INSERT INTO `robots`(`name`, `year`) VALUES (?, ?)";
    $success = $connection->execute($sql, array('Astro Boy', 1952));

    // Generating dynamically the necessary SQL
    $success = $connection->insert("robots",
       array("Astro Boy", 1952),
       array("name", "year")
    );

    // Generating dynamically the necessary SQL (another syntax)
    $success = $connection->insertAsDict("robots",
       array(
          "name" => "Astro Boy",
          "year" => 1952
       )
    );

    // Updating data with a raw SQL statement
    $sql     = "UPDATE `robots` SET `name` = 'Astro boy' WHERE `id` = 101";
    $success = $connection->execute($sql);

    // With placeholders
    $sql     = "UPDATE `robots` SET `name` = ? WHERE `id` = ?";
    $success = $connection->execute($sql, array('Astro Boy', 101));

    // Generating dynamically the necessary SQL
    $success = $connection->update("robots",
       array("name"),
       array("New Astro Boy"),
       "id = 101" // Warning! In this case values are not escaped
    );

    // Generating dynamically the necessary SQL (another syntax)
    $success = $connection->updateAsDict("robots",
       array(
          "name" => "New Astro Boy"
       ),
       "id = 101" // Warning! In this case values are not escaped
    );

    // With escaping conditions
    $success = $connection->update("robots",
       array("name"),
       array("New Astro Boy"),
       array(
          'conditions' => 'id = ?',
          'bind' => array(101),
          'bindTypes' => array(PDO::PARAM_INT) // Optional parameter
       )
    );
    $success = $connection->updateAsDict("robots",
       array(
          "name" => "New Astro Boy"
       ),
       array(
          'conditions' => 'id = ?',
          'bind' => array(101),
          'bindTypes' => array(PDO::PARAM_INT) // Optional parameter
       )
    );

    // Deleting data with a raw SQL statement
    $sql     = "DELETE `robots` WHERE `id` = 101";
    $success = $connection->execute($sql);

    // With placeholders
    $sql     = "DELETE `robots` WHERE `id` = ?";
    $success = $connection->execute($sql, array(101));

    // Generating dynamically the necessary SQL
    $success = $connection->delete("robots", "id = ?", array(101));

Transactions
------------

<?php

    try {

        // Start a transaction
        $connection->begin();

        // Execute some SQL statements
        $connection->execute("DELETE `robots` WHERE `id` = 101");
        $connection->execute("DELETE `robots` WHERE `id` = 102");
        $connection->execute("DELETE `robots` WHERE `id` = 103");

        // Commit if everything goes well
        $connection->commit();

    } catch (Exception $e) {
        // An exception has occurred rollback the transaction
        $connection->rollback();
    }

