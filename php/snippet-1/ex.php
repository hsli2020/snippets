<?php

class BaseClient
{
    public function purchase($ordr)
    {
        $requestClass = static::PurchaseRequest;
        $request = new $requestClass;
        $request->send();

        $responseClass = static::PurchaseResponse;
        $response = new $responseClass;
        $response->parse();
    }
}

##

class S1_Client extends BaseClient
{
    const PurchaseRequest  = S1_PurchaseRequest::class;
    const PurchaseResponse = S1_PurchaseResponse::class;
}

class S1_PurchaseRequest
{
    public function send() { echo __METHOD__, PHP_EOL; }
}

class S1_PurchaseResponse
{
    public function parse() { echo __METHOD__, PHP_EOL; }
}

##

class S2_Client extends BaseClient
{
    const PurchaseRequest  = S2_PurchaseRequest::class;
    const PurchaseResponse = S2_PurchaseResponse::class;
}

class S2_PurchaseRequest
{
    public function send() { echo __METHOD__, PHP_EOL; }
}

class S2_PurchaseResponse
{
    public function parse() { echo __METHOD__, PHP_EOL; }
}


$cl = new S1_Client();
$cl->purchase('');

$cl = new S2_Client();
$cl->purchase('');
