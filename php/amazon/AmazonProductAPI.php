<?php

$AWS_ACCESS_KEY_ID = "[myaccesskeyhere]";
$AWS_SECRET_ACCESS_KEY = "[mysecretkeyhere]";

$base_url = "http://ecs.amazonaws.com/onca/xml?";
$url_params = array('Operation'=>"ItemSearch",'Service'=>"AWSECommerceService",
 'AWSAccessKeyId'=>$AWS_ACCESS_KEY_ID,'AssociateTag'=>"yourtag-10",
 'Version'=>"2006-09-11",'Availability'=>"Available",'Condition'=>"All",
 'ItemPage'=>"1",'ResponseGroup'=>"Images,ItemAttributes,EditorialReview",
 'Keywords'=>"Amazon");

// Add the Timestamp
$url_params['Timestamp'] = gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time());

// Sort the URL parameters
$url_parts = array();
foreach(array_keys($url_params) as $key)
    $url_parts[] = $key."=".$url_params[$key];
sort($url_parts);

// Construct the string to sign
$string_to_sign = "GET\necs.amazonaws.com\n/onca/xml\n".implode("&",$url_parts);
$string_to_sign = str_replace('+','%20',$string_to_sign);
$string_to_sign = str_replace(':','%3A',$string_to_sign);
$string_to_sign = str_replace(';',urlencode(';'),$string_to_sign);

// Sign the request
$signature = hash_hmac("sha256",$string_to_sign,$AWS_SECRET_ACCESS_KEY,TRUE);

// Base64 encode the signature and make it URL safe
$signature = base64_encode($signature);
$signature = str_replace('+','%2B',$signature);
$signature = str_replace('=','%3D',$signature);

$url_string = implode("&",$url_parts);
$url = $base_url.$url_string."&Signature=".$signature;
print $url;

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL,$url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_TIMEOUT, 15);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);

$xml_response = curl_exec($ch);
echo $xml_response;