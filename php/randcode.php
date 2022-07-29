<?php

echo generateRandomCode(6), "\n";
#echo randomString(6), "\n";

function generateRandomCode(int $length)
{
    $min = (int) pow(10, $length - 1);
    $max = (int) pow(10, $length) - 1;

    return (string) random_int($min, $max);
}

function randomString($length)
{
	$allowedCharacters = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ';
#   $allowedCharacters = '23456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-_=+[]{};:<>,.?/';
#   $allowedCharacters = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';

	// Handle other formats (e.g., alphanumeric, mixed characters).
	$randomString = '';

	for ($i = 0; $i < $length; $i++) {
		$randomIndex   = random_int(0, mb_strlen($allowedCharacters) - 1);
		$randomString .= $allowedCharacters[$randomIndex];
	}

	return $randomString;
}
