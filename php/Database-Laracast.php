<?php

# https://www.youtube.com/watch?v=PDtBKgOJhGY

class Database
{
    protected $connection;

    public function __construct($config, $username = 'root', $password = '')
    {
         $dsn = 'mysql:'. http_build_query($config, '', ';');

         $this->connection = new PDO($dsn, $username, $password, [
             PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC
         ]);
    }

    public function query($query, $params = [])
    {
        $statement = $this->connection->prepare($query);

        $statement->execute($params);

        return $statement;
    }
}

# $config = require('config.php');

$config = [
    'database' => [
        'host' => 'localhost',
        'port' => 3306,
        'dbname' => 'test',
        'charset' => 'utf8mb4',
    ],

    'something' => [ ]
];

$db = new Database($config['database']);

/*
$post = $db->query("SELECT * FROM posts WHERE id=1")->fetch();
print_r($post);

$posts = $db->query("SELECT * FROM posts")->fetchAll();
print_r($posts);

# SQL injection
$id = $_GET['id'];
$post = $db->query("SELECT * FROM posts WHERE id=$id")->fetch();
print_r($post);

# Anti SQL injection
$id = $_GET['id'];
$post = $db->query("SELECT * FROM posts WHERE id=?", [$id])->fetch();
print_r($post);
//*/

$rows = $db->query("SELECT * FROM date_excluded")->fetchAll();
print_r($rows);
