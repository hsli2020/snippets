<?php

$valid_passwords = array ("bteadmin" => "btepass");
$valid_users = array_keys($valid_passwords);

$user = $_SERVER['PHP_AUTH_USER'];
$pass = $_SERVER['PHP_AUTH_PW'];

$validated = (in_array($user, $valid_users)) && ($pass == $valid_passwords[$user]);

$fresh = false;
if (!isset($_COOKIE['BTEAUTH'])) {
  $fresh = true;
  setcookie("BTEAUTH", md5(uniqid()), time()+60);
}
if ($fresh || !$validated) {
  header('WWW-Authenticate: Basic realm="My Realm"');
  header('HTTP/1.0 401 Unauthorized');
  die("Not authorized");
}

if ($validated) {
    // If arrives here, is a valid user.
    echo "<p>Welcome $user.</p>";
    var_dump($_COOKIE);
    var_dump($_SERVER);
}
