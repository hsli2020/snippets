<?php

// How to Build a JSON Web Token in PHP 

# header.payload.signature

// Create token header as a JSON string
$header = json_encode(['typ' => 'JWT', 'alg' => 'HS256']);

// Create token payload as a JSON string
$payload = json_encode(['user_id' => 123]);

// Encode Header to Base64Url String
$base64Header = str_replace(['+', '/', '='], ['-', '_', ''], base64_encode($header));

// Encode Payload to Base64Url String
$base64Payload = str_replace(['+', '/', '='], ['-', '_', ''], base64_encode($payload));

// Create Signature Hash
$signature = hash_hmac('sha256', $base64Header . "." . $base64Payload, 'abC123!', true);

// Encode Signature to Base64Url String
$base64Signature = str_replace(['+', '/', '='], ['-', '_', ''], base64_encode($signature));

// Create JWT
$jwt = $base64Header . "." . $base64Payload . "." . $base64Signature;

// Output
// eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxMjN9.NYlecdiqVuRg0XkWvjFvpLvglmfR1ZT7f8HeDDEoSx8

echo $jwt, PHP_EOL;

list($b64hdr, $b64pyl, $b64sig) = explode('.', $jwt);
$hdr = base64_decode($b64hdr);
echo $hdr, PHP_EOL;

$pyl = base64_decode($b64pyl);
echo $pyl, PHP_EOL;

$sig = base64_decode($b64sig);

$signature = hash_hmac('sha256', $b64hdr . "." . $b64pyl, 'abC123!', true);
$b64signature = str_replace(['+', '/', '='], ['-', '_', ''], base64_encode($signature));
if ($b64signature == $b64sig) {
    echo "OK\n";
}
