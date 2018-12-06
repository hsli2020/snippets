<?php
private function _invoke(array $parameters)
{
    $parameters = $this->_addRequiredParameters($parameters);
    $retries = 0;
    for (;;) {
        $response = $this->_httpPost($parameters);
        $status = $response['Status'];
        if ($status == 200) {
            return array('ResponseBody' => $response['ResponseBody'],
              'ResponseHeaderMetadata' => $response['ResponseHeaderMetadata']);
        }
        if ($status == 500 && $this->_pauseOnRetry(++$retries)) {
            continue;
        }
        throw $this->_reportAnyErrors($response['ResponseBody'],
            $status, $response['ResponseHeaderMetadata']);
    }
}

private function _addRequiredParameters(array $parameters)
{
    $parameters['AWSAccessKeyId'] = $this->_awsAccessKeyId;
    $parameters['Timestamp'] = gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time());
    $parameters['Version'] = self::SERVICE_VERSION;
    $parameters['SignatureVersion'] = $this->_config['SignatureVersion'];
    if ($parameters['SignatureVersion'] > 1) {
        $parameters['SignatureMethod'] = $this->_config['SignatureMethod'];
    }
    $parameters['Signature'] = $this->_signParameters($parameters, $this->_awsSecretAccessKey);

    return $parameters;
}

private function _calculateStringToSignV2(array $parameters)
{
    $data = 'POST';  $data .= "\n";

    $endpoint = parse_url ($this->_config['ServiceURL']);
    $data .= $endpoint['host'];  $data .= "\n";

    $uri = array_key_exists('path', $endpoint) ? $endpoint['path'] : null;
    if (!isset ($uri)) {
        $uri = "/";
    }
    $uriencoded = implode("/", array_map(array($this, "_urlencode"), explode("/", $uri)));

    $data .= $uriencoded;  $data .= "\n";

    uksort($parameters, 'strcmp');
    $data .= $this->_getParametersAsString($parameters);

    return $data;
}

private function _getParametersAsString(array $parameters) {
    $queryParameters = array();
    foreach ($parameters as $key => $value) {
        $queryParameters[] = $key . '=' . $this->_urlencode($value);
    }
    return implode('&', $queryParameters);
}

private function _signParameters(array $parameters, $key) {
    $parameters['SignatureMethod'] = "HmacSHA1";
    $stringToSign = $this->_calculateStringToSignV2($parameters);
    return base64_encode(hash_hmac('sha256', $stringToSign, $key, true));
}

private function _urlencode($value) { return str_replace('%7E', '~', rawurlencode($value)); }

private function _httpPost(array $parameters)
{
    $config = $this->_config;
    $query = $this->_getParametersAsString($parameters);
    $url = parse_url ($config['ServiceURL']);
    $uri = array_key_exists('path', $url) ? $url['path'] : null;
    if (!isset ($uri)) { $uri = "/"; }

    switch ($url['scheme']) {
        case 'https':
            $scheme = 'https://';
            $port = isset($url['port']) ? $url['port'] : 443;
            break;
        default:
            $scheme = 'http://';
            $port = isset($url['port']) ? $url['port'] : 80;
    }

    $allHeaders = $config['Headers'];
    $allHeaders['Content-Type'] = "application/x-www-form-urlencoded; charset=utf-8";
    $allHeaders['Expect'] = null; // Don't expect 100 Continue
    $allHeadersStr = array();
    foreach($allHeaders as $name => $val) {
        $str = $name . ": ";
        if (isset($val)) { $str = $str . $val; }
        $allHeadersStr[] = $str;
    }

    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $scheme . $url['host'] . $uri);
    curl_setopt($ch, CURLOPT_PORT, $port);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 2);
    curl_setopt($ch, CURLOPT_USERAGENT, $this->_config['UserAgent']);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $query);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $allHeadersStr);
    curl_setopt($ch, CURLOPT_HEADER, true);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    $response = curl_exec($ch);

    if ($response === false) {
        $exProps["Message"] = curl_error($ch);
        $exProps["ErrorType"] = "HTTP";
        curl_close($ch);
        throw new FBAInventoryServiceMWS_Exception($exProps);
    }

    curl_close($ch);
    return $this->_extractHeadersAndBody($response);
}
