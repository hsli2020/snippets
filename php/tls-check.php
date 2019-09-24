<?php

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, "https://www.howsmyssl.com/a/check");
curl_setopt($ch, CURLOPT_SSLVERSION, 6);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
$response = curl_exec($ch);
curl_close($ch);

$tlsVer = json_decode($response, true);
print_r($tlsVer);

echo "Your TLS version is: ";
echo ($tlsVer['tls_version'] ? $tlsVer['tls_version'] : 'no TLS support');
echo "\n";