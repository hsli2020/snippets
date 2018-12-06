<?php
require_once '../config.php';

function sign($url, $params) {
	$parsed_url = parse_url($url);
	$query = http_build_query_rfc3986($params);

	$request = array(
		'GET',
		$parsed_url['host'],
		$parsed_url['path'],
		$query
	);

	$signature = base64_encode(hash_hmac('sha256', implode("\n", $request), AWS_SECRET_KEY, true));
	return $signature;
}

function http_build_query_rfc3986($params) {
	ksort($params);
	$query = http_build_query($params);
	$query = strtr($query, array('%7E' => '~', '+' => '%20'));

	return $query;
}

function pr($var) {
	echo "<pre>\n";
	print_r($var);
	echo "</pre>\n";
}

function requestAPI($params) {
	$requestURI = 'http://ecs.amazonaws.jp/onca/xml?';

	$defaults = array(
		'Service' => 'AWSECommerceService',
		'AWSAccessKeyId' => AWS_ACCESS_KEY,
		'Operation' => 'ItemSearch',
		'SearchIndex' => 'Books',
		'Timestamp' => gmdate('Y-m-d\TH:i:s\Z'),
		'Version' => '2011-08-01',
		'Title' => '',
		'ResponseGroup' => 'ItemAttributes,Reviews',
		'AssociateTag' => '000000-22',
	);

	$params = array_merge($defaults, $params);
	$params['Signature'] = sign($requestURI, $params);

	$url = $requestURI . http_build_query_rfc3986($params);

	$result = request($url, true);

	return $result;
}

function request($url, $force = false, $opts = array()) {
	global $http_response_header;

	$headers = array(
		'http' => array(
			'method' => 'GET',
			'user_agent' => "Mozilla/5.0 (X11; Linux i686 on x86_64; rv:5.0) Gecko/20100101 Firefox/5.0",
			'timeout' => 60.0,
			'ignore_errors' => true,
		)
	);

	$context = stream_context_create(array_merge($headers, $opts));
	$result = file_get_contents($url, false, $context);

	if($force || strpos($http_response_header[0], '200') !== false) {
		return $result;
	}

	return false;
}

$keyword = "下町ロケット";



$result = requestAPI(array('Title' => $keyword));

if($result !== false) {
	$xml = simplexml_load_string($result);
	$asin = $xml->Items->Item->ASIN;

	$continue = true;
	$page = 1;
	$content = '';

	while($continue) {
		sleep(1);

		$lookup = requestAPI(array(
			'IdType' => 'ASIN',
			'ItemId' => $asin,
			'MerchantId' => 'Amazon',
			'ReviewPage' => $page,
			'ResponseGroup' => 'Reviews',
			'Title' => $keyword
		));

		$xml = simplexml_load_string($lookup);
		$has_reviews = $xml->Items->Item->CustomerReviews->HasReviews;
		if($has_reviews !== 'false') {
			$reviews_page = $xml->Items->Item->CustomerReviews->IFrameURL;
			$content[] = $reviews_page;
		} else {
			$continue = false;
		}
	}

	pr($content);
} else {
	pr($result);
}