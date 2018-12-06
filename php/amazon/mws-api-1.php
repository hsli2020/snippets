<?php
/* Amazon MWS API using PHP shows no Result

I have been using amazon mws API for developing a project to find Lowest Priced offers,
It works fine with action ListMatchingProducts and GetMatchingProduct but when it comes
to GetLowestPricedOffersForASIN, it shows no result in XML output

"This XML file does not appear to have any style information associated with it.
The document tree is shown below."

My PHP File Here: */
$param = array();
$param['AWSAccessKeyId']   = '<YOUR-ACCESS-ID>';
$param['Action']           = 'GetLowestPricedOffersForASIN';
$param['SellerId']         = '<YOUR-SELLER-ID>';
$param['SignatureMethod']  = 'HmacSHA256';  
$param['SignatureVersion'] = '2'; 
$param['Timestamp']        = gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time());
$param['Version']          = '2011-10-01'; 
$param['MarketplaceId']    = 'A2EUQ1WTGCTBG2'; 
$param['ItemCondition']    = 'used';
$param['ASIN']			   = '0439139600';
$secret = '<YOUR-SECRET-ID>';
$url = array();
foreach ($param as $key => $val) {
	$key = str_replace("%7E", "~", rawurlencode($key));
	$val = str_replace("%7E", "~", rawurlencode($val));
	$url[] = "{$key}={$val}";
}

sort($url);

$arr   = implode('&', $url);

$sign  = 'GET' . "\n";
$sign .= 'mws.amazonservices.com' . "\n";
$sign .= '/Products/2011-10-01' . "\n";
$sign .= $arr;

$signature = hash_hmac("sha256", $sign, $secret, true);
$signature = urlencode(base64_encode($signature));

$link  = "https://mws.amazonservices.com/Products/2011-10-01?";
$link .= $arr . "&Signature=" . $signature;
$ch = curl_init($link);
curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-type: application/xml'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE); 
$response = curl_exec($ch);
$info = curl_getinfo($ch);
curl_close($ch);
echo $response;
?>
