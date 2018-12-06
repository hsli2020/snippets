<?php
$methods = openssl_get_cipher_methods();

var_export($methods);

$textToEncrypt = "he who doesn't do anything, doesn't go wrong -- Zeev Suraski";
$secretKey = "glop";

echo "\n";
foreach ($methods as $method) {
    $iv  = substr(hash('sha256', $secretKey), 0, openssl_cipher_iv_length($method));
    $encrypted = openssl_encrypt($textToEncrypt, $method, $secretKey, 0, $iv);
    $decrypted = openssl_decrypt($encrypted, $method, $secretKey, 0, $iv);
    echo $method, "\n", $encrypted, "\n", $decrypted, "\n\n";
}
