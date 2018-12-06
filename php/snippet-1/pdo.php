<?php

$host = '127.0.0.1';
$db   = 'bte';
$user = 'root';
$pass = '';
$charset = 'utf8';

// mysql:host=localhost;dbname=test;port=3306;charset=utf8
$dsn = "mysql:host=$host;dbname=$db;charset=$charset";

$opt = [
    PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
    PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
    PDO::ATTR_EMULATE_PREPARES   => false,
];

$db = new PDO($dsn, $user, $pass, $opt);
$db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

try {
    # MS SQL Server and Sybase with PDO_DBLIB
    $DBH = new PDO("mssql:host=$host;dbname=$dbname, $user, $pass");
    $DBH = new PDO("sybase:host=$host;dbname=$dbname, $user, $pass");

    # MySQL with PDO_MYSQL
    $DBH = new PDO("mysql:host=$host;dbname=$dbname", $user, $pass);

    # SQLite Database
    $DBH = new PDO("sqlite:my/database/path/database.db");
}
catch(PDOException $e) {
    echo $e->getMessage();
}

#try {
#    $db->prepare("INSERT INTO users VALUES (NULL,?,?,?,?)")->execute($data);
#} catch (PDOException $e) {
#    if ($e->getCode() == 1062) {
#        // Take some action if there is a key constraint violation, i.e. duplicate name
#    } else {
#        throw $e;
#    }
#}

$stmt = $db->query('SELECT * FROM users');
while ($row = $stmt->fetch()) {
    echo $row['username'], " ", $row['password'], "\n";
}

$stmt = $db->query('SELECT username FROM users');
foreach ($stmt as $row) {
    echo $row['username'] . "\n";
}

#$stmt = $db->prepare('SELECT * FROM users WHERE email = ? AND status=?');
#$stmt->execute([$email, $status]);
#$user = $stmt->fetch();
#
#// or
#
#$stmt = $db->prepare('SELECT * FROM users WHERE email = :email AND status=:status');
#$stmt->execute(['email' => $email, 'status' => $status]);
#$user = $stmt->fetch();

#$data = [ 1 => 1000, 5 =>  300, 9 =>  200 ];
#$stmt = $db->prepare('UPDATE users SET bonus = bonus + ? WHERE id = ?');
#foreach ($data as $id => $bonus) {
#    $stmt->execute([$bonus, $id]);
#}

#$sql = "UPDATE users SET name = ? WHERE id = ?";
#$db->prepare($sql)->execute([$name, $id]);
#
#$stmt = $db->prepare("DELETE FROM goods WHERE category = ?");
#$stmt->execute([$cat]);
#$deleted = $stmt->fetchColumn();
