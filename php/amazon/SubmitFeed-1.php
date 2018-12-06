<?php
/*
I have a simple script that updates just one products price every 30 minutes.
This was working fine until sometime in November 2014 when it just stopped
working. Nothing had been changed, the script had not been changed, absolutely
nothing had been changed!

Now, it says the signature does not match. Here is my script:
*/
$myprice = "get lowest price from amazon here";
$write = fopen("tempprice.xml", "w");
fwrite($write, "sku price keyring ".$myprice);
fclose($write);

$file = 'tempprice.xml';
$fo = fopen($file, 'r');

$httpHeader=array();
$httpHeader[]='Transfer-Encoding: chunked';
$httpHeader[]='Content-Type: text/xml';
$httpHeader[]='Content-MD5: ' . base64_encode(md5_file($file, true));
$httpHeader[]='Expect:';
$httpHeader[]='Accept:';

$curl_options=array(
    CURLOPT_UPLOAD         => true,
    CURLOPT_INFILE         => $fo,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST           => true,
    CURLOPT_PORT           => 443, //#
    CURLOPT_SSLVERSION     => 3,   //#
    CURLOPT_SSL_VERIFYHOST => false,
    CURLOPT_SSL_VERIFYPEER => false,
    CURLOPT_FOLLOWLOCATION => 1,
    CURLOPT_PROTOCOLS      => CURLPROTO_HTTPS,
    CURLINFO_HEADER_OUT    => TRUE,
    CURLOPT_HTTPHEADER     => $httpHeader,
    CURLOPT_CUSTOMREQUEST  => 'POST',
    CURLOPT_VERBOSE        => true,
    CURLOPT_HEADER         => true,
);

$param = array();
$param['AWSAccessKeyId']         = 'ACCESS_KEY_ID';
$param['Action']                 = 'SubmitFeed';
$param['SellerId']               = 'SELLER_ID';
$param['SignatureMethod']        = 'HmacSHA256';
$param['SignatureVersion']       = '2';
$param['Timestamp']              = gmdate("Y-m-d\TH:i:s.\0\0\0\Z", time());
$param['Version']                = '2009-01-01';
$param['MarketplaceIdList.Id.1'] = 'A1F83G8C2ARO7P'; // UK
$param['FeedType']               = "POST_FLAT_FILE_PRICEANDQUANTITYONLY_UPDATE_DATA";

$secret = 'SECRET_KEY';

$url = array();
foreach ($param as $key => $val) {
    $key = str_replace("%7E", "~", rawurlencode($key));
    $val = str_replace("%7E", "~", rawurlencode($val));
    $url[] = "{$key}={$val}";
}

sort($url);

$arr = implode('&', $url);

$sign = 'POST' . "\n";
$sign .= 'mws-eu.amazonservices.com' . "\n";
$sign .= '/Feeds/2009-01-01' . "\n";
$sign .= $arr;

$signature = hash_hmac("sha256", $sign, $secret, true);
$signature = urlencode(base64_encode($signature));
$link = "https://mws-eu.amazonservices.com/Feeds/2009-01-01?";
$link .= $arr . "&Signature=" . $signature;

$ch = curl_init($link);
curl_setopt_array($ch,$curl_options);
$response=curl_exec($ch);
$info = curl_getinfo($ch);
curl_close($ch);
/*
Can anybody shed some light on why a perfectly working script just stops working 
without intervention?

I have figured it out and I don’t know why this works but if I comment out the two lines:
    //CURLOPT_PORT=&gt;443,
    //CURLOPT_SSLVERSION=&gt;3,
IT WORKS!

So it was my servers security certificate that was causing the issue because it was 
possible outdated?
*/
?>
