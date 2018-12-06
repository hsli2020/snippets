<?php

use Phalcon\Di;

use Phalcon\Mvc\Model;
use Phalcon\Mvc\Model\Manager as ModelsManager;
use Phalcon\Mvc\Model\Metadata\Memory as MetaData;

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

class Artists extends Model
{
    public $name;
    public $country;

    public function initialize()
    {
        $this->hasMany("id", "Albums", "artist_id");
        $this->hasMany("id", "Songs",  "artist_id");
    }
}

class Albums extends Model
{
    public $name;

    public function initialize()
    {
        $this->belongsTo('artist_id', 'Artists', 'id', ['alias' => 'artist']);
        $this->hasMany('id', 'Songs', 'album_id');
    }
}

class Songs extends Model
{
    public $name;
    public $duration;

    public function initialize()
    {
        $this->belongsTo('artist_id', 'Artists', 'id', ['alias' => 'artist']);
        $this->belongsTo('album_id',  'Albums',  'id', ['alias' => 'album']);
    }
}

########################################

#$artists = Artists::find();
#
#foreach ($artists as $artist) {
#
#    $albums = $artist->albums;
#
#    foreach ($albums as $album) {
#        foreach ($album->songs as $song) {
#            echo $artist->name, ' ';
#            echo $artist->country, "\t";
#
#            echo $album->name, ': ';
#
#            echo $song->name,   ' ';
#            echo $song->duration,   ' ';
#            echo EOL;
#        }
#    }
#}
#
#echo EOL;

$artist = Artists::findFirst("name = 'artist-1'");

// Create an album
$album = new Albums();

$album->name   = "The One";
$album->artist = $artist;

$songs = [];

// Create a first song
$songs[0]           = new Songs();
$songs[0]->name     = "Star Guitar";
$songs[0]->duration = "5:54";
$songs[0]->artist   = $artist;

// Create a second song
$songs[1]           = new Songs();
$songs[1]->name     = "Last Days";
$songs[1]->duration = "4:29";
$songs[1]->artist   = $artist;

// Assign the songs array
$album->songs = $songs;

// Save the album + its songs
$album->save();
