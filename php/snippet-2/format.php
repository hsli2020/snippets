<?php

const EOL = PHP_EOL;

#http://stackoverflow.com/questions/4708248/formatting-phone-numbers-in-php
function formatPhoneNumber($number)
{
    $number = str_replace(['-', '.', ' ', '(', ')'], '', $number);

    if (preg_match('/^\+\d(\d{3})(\d{3})(\d{4})$/', $number, $matches)) {
        $result = $matches[1]. '-' .$matches[2]. '-' .$matches[3];
        return $result;
    }

    if (preg_match('/^(\d{3})(\d{3})(\d{4})$/', $number, $matches)) {
        $result = $matches[1]. '-' .$matches[2]. '-' .$matches[3];
        return $result;
    }
}

$data = '123-456-7890';
$data = '+1 123-456-7890';
echo formatPhoneNumber($data), EOL;

function formatCanadaPostalcode($code)
{
    $code = str_replace(' ', '', strtoupper($code));
    if (preg_match('/[A-Za-z]\d[A-Za-z]\d[A-Za-z]\d/', $code)) {
        return substr($code, 0, 3). ' ' .substr($code, 3);
    }
    return $code;
}

$code = 'M2H3BC';
$code = ' M2H  3B5 ';
echo formatCanadaPostalcode($code), EOL;
