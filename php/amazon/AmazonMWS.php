<?php

// https://stackoverflow.com/questions/40874090/amazon-mws-listorders-from-scratchpad-to-request

class AmazonMWS
{
    private $secretKey = '';

    private $parameters = array();

    /**
     * Constructor for the AmazonMWS class.
     * Initializes constants.
     */
    public function __construct() 
    {
        $this->secretKey = Constant::get('SECRET_KEY');

        $this->parameters['AWSAccessKeyId']     = Constant::get('AWSAccessKeyId');
        $this->parameters['MarketplaceId.Id.1'] = Constant::get('MarketplaceId.Id.1');
        $this->parameters['SellerId']           = Constant::get('SellerId');
        $this->parameters['SignatureMethod']    = Constant::get('SignatureMethod');
        $this->parameters['SignatureVersion']   = Constant::get('SignatureVersion');
    }

    public function setListOrders()
    {
        $this->parameters['Action'] = 'ListOrders';
        $this->parameters['Version'] = '2013-09-01';
        $this->parameters['Timestamp'] = $this->getTimestamp();

        // this part should change and depend on the method/parameter.. for now just for testing

        $this->parameters['CreatedAfter'] = '2015-11-01';
    }

    public function listOrders()
    {
        $request = "https://mws.amazonservices.com/Orders/2013-09-01?";
        $request .= $this->getParameterString($this->parameters) . "&Signature=" . $this->calculateSignature($this->calculateStringToSign($this->parameters));

        echo $request;

        return Curl::fetchSSL($request);
    }

    /**
     * Calculates String to sign.
     * 
     * @param array $parameters request parameters
     * @return String to sign
     */
    protected function calculateStringToSign(array $parameters)
    {
        $stringToSign  = 'GET' . "\n";
        $stringToSign .= 'mws.amazonservices.com' . "\n";
        $stringToSign .= '/Orders/2013-09-01' . "\n";
        $stringToSign .= $this->getParameterString($parameters);

        return $stringToSign;
    }

    /**
     * Gets the query parameters as a String sorted in natural-byte order.
     * 
     * @param array $parameters request parameters
     * @return String of parameters
     */
    protected function getParameterString(array $parameters)
    {
        $url = array();
        foreach ($parameters as $key => $val) {
            $key = $this->urlEncode($key);
            $val = $this->urlEncode($val);
            $url[] = "{$key}={$val}";
        }
        sort($url);

        $parameterString = implode('&', $url);

        return $parameterString;
    }

    /**
     * Computes RFC 2104-compliant HMAC signature.
     *
     * @param String to sign
     */
    protected function calculateSignature($stringToSign)
    {
        $signature = hash_hmac("sha256", $stringToSign, $this->secretKey, true);
        return urlencode(base64_encode($signature));
    }

    /**
     * URL encodes a string.
     */
    protected function urlEncode($string)
    {
        return str_replace("%7E", "~", rawurlencode($string));
    }

    /**
     * Gets the current date as ISO 8601 timestamp
     */
    protected function getTimestamp()
    {
        return gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time());
    }
}