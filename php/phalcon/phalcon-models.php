<?php

use Phalcon\Mvc\Model;

$robot = Robots::findFirst(
    array(
        "type = 'virtual'",
        "order" => "name DESC",
        "limit" => 30
    )
);

// Return every robot as an array
$robots->setHydrateMode(Resultset::HYDRATE_ARRAYS);

// Return every robot as a stdClass
$robots->setHydrateMode(Resultset::HYDRATE_OBJECTS);

// Return every robot as a Robots instance
$robots->setHydrateMode(Resultset::HYDRATE_RECORDS);

$robots = Robots::find(
    array(
        "conditions" => "name LIKE 'steve%'",
        "columns" => "id, name",
        "bind" => array("status" => "A", "type" => "some-time"),
        "bindTypes" => array(Column::BIND_PARAM_STR, Column::BIND_PARAM_INT),
        "order" => "name DESC, status",
        "limit" => 10,
        "offset" => 5,
        "group" => "name, status",
        "for_update" => true,
        "shared_lock" => true,
        "cache" => array("lifetime" => 3600, "key" => "my-find-key"),
        "hydration" => Resultset::HYDRATE_OBJECTS,
    )
);

$robots = Robots::query()
    ->where("type = :type:")
    ->andWhere("year < 2000")
    ->bind(array("type" => "mechanical"))
    ->order("name")
    ->execute();

findFirstBy<property-name>()

Operation 	        Name 	                    Can stop operation?
Inserting/Updating 	beforeValidation 	        YES
Inserting 	        beforeValidationOnCreate 	YES
Updating 	        beforeValidationOnUpdate 	YES
Inserting/Updating 	onValidationFails 	        YES (already stopped)
Inserting 	        afterValidationOnCreate 	YES
Updating 	        afterValidationOnUpdate 	YES
Inserting/Updating 	afterValidation 	        YES
Inserting/Updating 	beforeSave 	                YES
Updating 	        beforeUpdate 	            YES
Inserting 	        beforeCreate 	            YES
Updating 	        afterUpdate 	            NO
Inserting 	        afterCreate 	            NO
Inserting/Updating 	afterSave 	                NO

trait MyTimestampable
{
    public function beforeCreate()
    {
        $this->created_at = date('r');
    }

    public function beforeUpdate()
    {
        $this->updated_at = date('r');
    }
}

class Robots extends Model
{
    public $id;
    public $name;
    public $status;

    public function beforeSave()
    {
        // Convert the array into a string
        $this->status = join(',', $this->status);
    }

    public function afterFetch()
    {
        // Convert the string to an array
        $this->status = explode(',', $this->status);
    }

    public function afterSave()
    {
        // Convert the string to an array
        $this->status = explode(',', $this->status);
    }

    public function getSource() { return "the_robots"; }  // table name
    public function initialize() { $this->setSource("the_robots"); }

    public function onConstruct() { // ...  }

    public function columnMap()
    {
        // Keys are the real names in the table and
        // the values their names in the application
        return array(
            'id'       => 'code',
            'the_name' => 'theName',
            'the_type' => 'theType',
            'the_year' => 'theYear'
        );
    }
}

hasMany 	    Defines a 1-n relationship
hasOne 	        Defines a 1-1 relationship
belongsTo 	    Defines a n-1 relationship
hasManyToMany 	Defines a n-n relationship

CREATE TABLE `robots` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(70) NOT NULL,
    `type` varchar(32) NOT NULL,
    `year` int(11) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `robots_parts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `robots_id` int(10) NOT NULL,
    `parts_id` int(10) NOT NULL,
    `created_at` DATE NOT NULL,
    PRIMARY KEY (`id`),
    KEY `robots_id` (`robots_id`),
    KEY `parts_id` (`parts_id`)
);

CREATE TABLE `parts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(70) NOT NULL,
    PRIMARY KEY (`id`)
);

- The model “Robots” has many “RobotsParts”.
- The model “Parts” has many “RobotsParts”.
- The model “RobotsParts” belongs to both “Robots” and “Parts” models as a many-to-one relation.
- The model “Robots” has a relation many-to-many to “Parts” through “RobotsParts”.

class Robots extends Model
{
    public $id;
    public $name;

    public function initialize()
    {
        $this->hasMany("id", "RobotsParts", "robots_id");

        $this->hasManyToMany(
            "id",
            "RobotsParts",
            "robots_id", "parts_id",
            "Parts",
            "id"
        );
    }
}

class Parts extends Model
{
    public $id;
    public $name;

    public function initialize()
    {
        $this->hasMany("id", "RobotsParts", "parts_id");
    }
}

class RobotsParts extends Model
{
    public $id;
    public $robots_id;
    public $parts_id;

    public function initialize()
    {
        $this->belongsTo("robots_id", "Robots", "id");
        $this->belongsTo("parts_id", "Parts", "id");
    }
}

Setting multiple databases
==========================

In Phalcon, all models can belong to the same database connection or have an individual one.
Actually, when Phalcon\Mvc\Model needs to connect to the database it requests the “db” service
in the application’s services container. You can overwrite this service setting it in the
initialize method:

use Phalcon\Db\Adapter\Pdo\Mysql as MysqlPdo;
use Phalcon\Db\Adapter\Pdo\PostgreSQL as PostgreSQLPdo;

// This service returns a MySQL database
$di->set('dbMysql', function () {
    return new MysqlPdo(
        array(
            "host"     => "localhost",
            "username" => "root",
            "password" => "secret",
            "dbname"   => "invo"
        )
    );
});

// This service returns a PostgreSQL database
$di->set('dbPostgres', function () {
    return new PostgreSQLPdo(
        array(
            "host"     => "localhost",
            "username" => "postgres",
            "password" => "",
            "dbname"   => "invo"
        )
    );
});

class Robots extends Model
{
    public function initialize()
    {
        $this->setConnectionService('dbPostgres');
    }
}

But Phalcon offers you more flexibility, you can define the connection that must be used
to ‘read’ and for ‘write’. This is specially useful to balance the load to your databases
implementing a master-slave architecture:

use Phalcon\Mvc\Model;

class Robots extends Model
{
    public function initialize()
    {
        $this->setReadConnectionService('dbSlave');
        $this->setWriteConnectionService('dbMaster');
    }
}

The ORM also provides Horizontal Sharding facilities, by allowing you to implement a ‘shard’
selection according to the current query conditions:

use Phalcon\Mvc\Model;

class Robots extends Model
{
    /**
     * Dynamically selects a shard
     *
     * @param array $intermediate
     * @param array $bindParams
     * @param array $bindTypes
     */
    public function selectReadConnection($intermediate, $bindParams, $bindTypes)
    {
        // Check if there is a 'where' clause in the select
        if (isset($intermediate['where'])) {

            $conditions = $intermediate['where'];

            // Choose the possible shard according to the conditions
            if ($conditions['left']['name'] == 'id') {
                $id = $conditions['right']['value'];

                if ($id > 0 && $id < 10000) {
                    return $this->getDI()->get('dbShard1');
                }

                if ($id > 10000) {
                    return $this->getDI()->get('dbShard2');
                }
            }
        }

        // Use a default shard
        return $this->getDI()->get('dbShard0');
    }
}

The method ‘selectReadConnection’ is called to choose the right connection, this method
intercepts any new query executed:

$robot = Robots::findFirst('id = 101');

Logging Low-Level SQL Statements
================================

When using high-level abstraction components such as Phalcon\Mvc\Model to access a database,
it is difficult to understand which statements are finally sent to the database system.
Phalcon\Mvc\Model is supported internally by Phalcon\Db. Phalcon\Logger interacts with
Phalcon\Db, providing logging capabilities on the database abstraction layer, thus allowing
us to log SQL statements as they happen.

use Phalcon\Logger;
use Phalcon\Events\Manager;
use Phalcon\Logger\Adapter\File as FileLogger;
use Phalcon\Db\Adapter\Pdo\Mysql as Connection;

$di->set('db', function () {

    $eventsManager = new EventsManager();

    $logger = new FileLogger("app/logs/debug.log");

    // Listen all the database events
    $eventsManager->attach('db', function ($event, $connection) use ($logger) {
        if ($event->getType() == 'beforeQuery') {
            $logger->log($connection->getSQLStatement(), Logger::INFO);
        }
    });

    $connection = new Connection(
        array(
            "host"     => "localhost",
            "username" => "root",
            "password" => "secret",
            "dbname"   => "invo"
        )
    );

    // Assign the eventsManager to the db adapter instance
    $connection->setEventsManager($eventsManager);

    return $connection;
});

Stand-Alone component
=====================

Using Phalcon\Mvc\Model in a stand-alone mode can be demonstrated below:

use Phalcon\Di;
use Phalcon\Mvc\Model;
use Phalcon\Mvc\Model\Manager as ModelsManager;
use Phalcon\Db\Adapter\Pdo\Sqlite as Connection;
use Phalcon\Mvc\Model\Metadata\Memory as MetaData;

$di = new Di();

// Setup a connection
$di->set('db', new Connection(array("dbname" => "sample.db")));

// Set a models manager
$di->set('modelsManager', new ModelsManager());

// Use the memory meta-data adapter or other
$di->set('modelsMetadata', new MetaData());

// Create a model
class Robots extends Model
{
}

// Use the model
echo Robots::count();

