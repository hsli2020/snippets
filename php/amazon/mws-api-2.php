<?php
/*
Your main problem is, that you use GET instead of POST. This version of your code works:
*/
$param = array();
$param['AWSAccessKeyId']   = '<YOUR-ACCESS-KEY>';
$param['Action']           = 'GetLowestPricedOffersForASIN';
$param['MarketplaceId']    = 'A2EUQ1WTGCTBG2';
$param['SellerId']         = '<YOUR-SELLER-ID>';
$param['ASIN']             = 'B002BYQIWM';
$param['ItemCondition']    = 'New'; 
$param['SignatureMethod']  = 'HmacSHA256';  
$param['SignatureVersion'] = '2'; 
$param['Timestamp']        = gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time());
$param['Version']          = '2011-10-01'; 
$secret = '<YOUR-SECRET-KEY>';
$url = array();
foreach ($param as $key => $val) {
    $key = str_replace("%7E", "~", rawurlencode($key));
    $val = str_replace("%7E", "~", rawurlencode($val));
    $url[] = "{$key}={$val}";
}

sort($url);

$arr   = implode('&', $url);

$sign  = 'POST' . "\n";
$sign .= 'mws.amazonservices.com' . "\n";
$sign .= '/Products/2011-10-01' . "\n";
$sign .= $arr;

$signature = hash_hmac("sha256", $sign, $secret, true);
$s64 = base64_encode($signature);

$signature = urlencode($s64);
$link  = "https://mws.amazonservices.com/Products/2011-10-01";
$arr .= "&Signature=" . $signature;

$ch = curl_init($link);
curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded', 'Accept:'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
curl_setopt($ch, CURLOPT_POST, 1); 
curl_setopt($ch, CURLOPT_POSTFIELDS, $arr); 
$response = curl_exec($ch);
$info = curl_getinfo($ch);
curl_close($ch);
echo $response;

/*
I have used curl_setopt($ch, CURLOPT_VERBOSE, true); to debug the response from server.
Your code did not produced any http body to output, but this http header HTTP/1.1 405
Method Not Allowed. Changing to POST solved your problem.
*/
/*
Hi, Daniel, thanks for your answer, but I have tried using POST method, with POST method
I get this error: Sender SignatureDoesNotMatch The request signature we calculated does
not match the signature you provided.
*/
/*
Please try my code. It is working. Signature is valid and the XML with the competitor
offers is returned.
*/
