<?php

use Phalcon\Di;

use Phalcon\Events\Event;
use Phalcon\Events\Manager as EventsManager;

use Phalcon\Mvc\Model;
use Phalcon\Mvc\Model\Manager as ModelsManager;
use Phalcon\Mvc\Model\Metadata\Memory as MetaData;

use Phalcon\Logger;
use Phalcon\Logger\Adapter\File as FileLogger;
use Phalcon\Logger\Formatter\Line as FormatterLine;

const EOL = PHP_EOL;

function pr($var) { var_export($var); echo EOL; }

$di = new Di();

$di->set('db', function () {
    $config = [
        "host"     => "127.0.0.1",
        "username" => "root",
        "password" => "",
        "dbname"   => "test",
        "options"  => [ PDO::ATTR_CASE => PDO::CASE_LOWER ]
    ];

    $db = new \Phalcon\Db\Adapter\Pdo\Mysql($config);

    return $db;
});

$di->set("modelsManager",  new ModelsManager());
$di->set("modelsMetadata", new MetaData());

class Robots extends Model
{
    public $id;
    public $name;

    public function initialize()
    {
        $this->hasMany("id", "RobotsParts", "robots_id");

       #$this->hasManyToMany(
       #    "id",
       #    "RobotsParts",
       #    "robots_id", "parts_id",
       #    "Parts",
       #    "id"
       #);
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

    public function onConstruct() { }

    public function initialize()
    {
        $this->belongsTo("robots_id", "Robots", "id");
        $this->belongsTo("parts_id",  "Parts",  "id");
    }
}

########################################

$robot = Robots::findFirst("type='type-2'");

$robots = Robots::find();

foreach ($robots as $robot) {
    echo $robot->name, ': ';
    foreach ($robot->robotsParts as $robotPart) {
        echo $robotPart->parts->name, ' ';
    }
    echo EOL;
}

########################################

#$part = Parts::findFirst();
#foreach ($part->robotsParts as $robotPart) {
#    echo $part->name, ', ', $robotPart->parts->name, EOL;
#}
