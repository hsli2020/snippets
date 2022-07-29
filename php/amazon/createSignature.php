<?php

// $secretKey is the AWS secret key from IAM user
private function createSignature($secretKey) 
{
	// Derived according to docs: 
	// https://docs.aws.amazon.com/general/latest/gr/sigv4-create-string-to-sign.html
	$stringToSign = $this->createStringToSign(); 

	$algorithm = 'sha256';
	$kSecret = "AWS4".$secretKey;

	// Date derived per amazon documentation eg. 20201020 for 20 Oct 2020
	$date = $this->credentials->date; 
	
	// Region for request eg. us-east-1
	$region = $this->credentials->region; 
	
	// Service eg. execute-api
	$service = $this->credentials->service; 
	
	// Signature termination string eg. aws4_request
	$terminationString = $this->credentials->terminationString;

	$kDate = hash_hmac($algorithm, $date, $kSecret, true);
	$kRegion = hash_hmac($algorithm, $region, $kDate, true);
	$kService = hash_hmac($algorithm, $service, $kRegion, true);
	$kSigning = hash_hmac($algorithm, $terminationString, $kService, true);

	// Without fourth parameter passed as true, returns lowercase hexits as called for by docs
	$signature = hash_hmac($algorithm, $stringToSign, $kSigning);
	
	// Trimming maybe not necessary here but can't hurt.
	return trim($signature);
}
