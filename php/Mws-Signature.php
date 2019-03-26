<?php

# http://www.voidcn.com/article/p-yxbwnmkv-bta.html
# https://stackoverflow.com/questions/29679646/issues-calculating-signature-for-amazon-marketplace-api

namespace App\Marketplace\Amazon;

class Signature
{
    /**
     * The signed string.
     *
     * @var string
     */
    protected $signedString;

    /**
     * Create a new signature instance.
     *
     * @param  string  $url
     * @param  array   $data
     * @param  string  $secretAccessKey
     */
    public function __construct($url, array $parameters, $secretAccessKey)
    {
        $stringToSign = $this->calculateStringToSign($url, $parameters);

        $this->signedString = $this->sign($stringToSign, $secretAccessKey);
    }

    /**
     * Calculate the string to sign.
     *
     * @param  string  $url
     * @param  array   $parameters
     * @return string
     */
    protected function calculateStringToSign($url, array $parameters)
    {
        $url = parse_url($url);

        $string = "POST\n";
        $string .= $url['host']."\n";
        $string .= $url['path']."\n";
        $string .= $this->getParametersAsString($parameters);

        return $string;
    }

    /**
     * Computes RFC 2104-compliant HMAC signature.
     *
     * @param  string  $data
     * @param  string  $secretAccessKey
     * @return string
     */
    protected function sign($data, $secretAccessKey)
    {
        return base64_encode(hash_hmac('sha256', $data, $secretAccessKey, true));
    }

    /**
     * Convert paremeters to URL-encoded query string.
     *
     * @param  array  $parameters
     * @return string
     */
    protected function getParametersAsString(array $parameters)
    {
        uksort($parameters, 'strcmp');

        $queryParameters = [];

        foreach ($parameters as $key => $value) {
            $key = rawurlencode($key);
            $value = rawurlencode($value);

            $queryParameters[] = sprintf('%s=%s', $key, $value);
        }

        return implode('&', $queryParameters);
    }

    /**
     * The string representation of this signature.
     *
     * @return string
     */
    public function __toString()
    {
        return $this->signedString;
    }
}

$version = '2011-07-01';

$url = 'https://mws.amazonservices.com/Sellers/'.$version;

$timestamp = gmdate('c', time());

$parameters = [
    'AWSAccessKeyId' => $command->accessKeyId,
    'Action' => 'GetAuthToken',
    'SellerId' => $command->sellerId,
    'SignatureMethod' => 'HmacSHA256',
    'SignatureVersion' => 2,
    'Timestamp' => $timestamp,
    'Version' => $version,
];

$signature = new Signature($url, $parameters, $command->secretAccessKey);

$parameters['Signature'] = strval($signature);

try {
    $response = $this->client->post($url, [
        'headers' => [
            'User-Agent' => 'my-app-name',
        ],
        'body' => $parameters,
    ]);

    dd($response->getBody());
} catch (\Exception $e) {
    dd(strval($e->getResponse()));
}