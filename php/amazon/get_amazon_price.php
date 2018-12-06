<?php
/*
Use the Amazon API to Get Product Prices
========================================

You can use the Amazon Advertising API (AWS) to fetch the current prices, item
description, customer reviews, product images, quantity available and other
details of any product listed on various Amazon websites.

The getAmazonPrice() method takes the region (like ?com? for US or ?es? for Spain) 
and the 10-digit ASIN of any product. The response is returned as JSON that the 
Amazon Price Tracker can easily parse with Google Scripts. You cannot however use 
this technique for fetching prices of Kindle Books as the Amazon API doesn?t support 
them yet.
*/

// Region code and Product ASIN
$response = getAmazonPrice("com", "B00KQPGRRE");

function getAmazonPrice($region, $asin) {

	$xml = aws_signed_request($region, array(
		"Operation" => "ItemLookup",
		"ItemId" => $asin,
		"IncludeReviewsSummary" => False,
		"ResponseGroup" => "Medium,OfferSummary",
	));

	$item = $xml->Items->Item;
	$title = htmlentities((string) $item->ItemAttributes->Title);
	$url = htmlentities((string) $item->DetailPageURL);
	$image = htmlentities((string) $item->MediumImage->URL);
	$price = htmlentities((string) $item->OfferSummary->LowestNewPrice->Amount);
	$code = htmlentities((string) $item->OfferSummary->LowestNewPrice->CurrencyCode);
	$qty = htmlentities((string) $item->OfferSummary->TotalNew);

	if ($qty !== "0") {
		$response = array(
			"code" => $code,
			"price" => number_format((float) ($price / 100), 2, '.', ''),
			"image" => $image,
			"url" => $url,
			"title" => $title
		);
	}

	return $response;
}

function getPage($url) {

	$curl = curl_init($url);
	curl_setopt($curl, CURLOPT_FAILONERROR, true);
	curl_setopt($curl, CURLOPT_FOLLOWLOCATION, true);
	curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
	curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, false);
	curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, false);
	$html = curl_exec($curl);
	curl_close($curl);
	return $html;
}

function aws_signed_request($region, $params) {

	$public_key = "PUBLIC_KEY";
	$private_key = "PRIVATE_KEY";

	$method = "GET";
	$host = "ecs.amazonaws." . $region;
	$host = "webservices.amazon." . $region;
	$uri = "/onca/xml";

	$params["Service"] = "AWSECommerceService";
	$params["AssociateTag"] = "affiliate-20"; // Put your Affiliate Code here
	$params["AWSAccessKeyId"] = $public_key;
	$params["Timestamp"] = gmdate("Y-m-d\TH:i:s\Z");
	$params["Version"] = "2011-08-01";

	ksort($params);

	$canonicalized_query = array();
	foreach ($params as $param => $value) {
		$param = str_replace("%7E", "~", rawurlencode($param));
		$value = str_replace("%7E", "~", rawurlencode($value));
		$canonicalized_query[] = $param . "=" . $value;
	}

	$canonicalized_query = implode("&", $canonicalized_query);

	$string_to_sign = $method . "\n" . $host . "\n" . $uri . "\n" . $canonicalized_query;
	$signature = base64_encode(hash_hmac("sha256", $string_to_sign, $private_key, True));
	$signature = str_replace("%7E", "~", rawurlencode($signature));

	$request = "http://" . $host . $uri . "?" . $canonicalized_query . "&Signature=" . $signature;
	$response = getPage($request);

	var_dump($response);

	$pxml = @simplexml_load_string($response);
	if ($pxml === False) {
		return False;// no xml
	} else {
		return $pxml;
	}
}
/*
    You'll have to get your public and private access keys from the AWS Management Console
    and replace them in the PHP snippet. You also need to put your Amazon Affiliate ID for
    the country specific Amazon Affiliate program. For instance, the id for Amazon US would
    be different from Amazon UK though you can use the same Access Keys with the code.
*/
?>
