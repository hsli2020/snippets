<?php
    protected function sendRequest($path, $params, $options = [])
    {
        $params['AWSAccessKeyId']   = $this->config['AccessKey'];
        $params['SellerId']         = $this->config['SellerId'];
        $params['MarketplaceId']    = $this->config['MarketplaceId'];
        $params['SignatureMethod']  = 'HmacSHA256';
        $params['SignatureVersion'] = '2';
        $params['Timestamp']        = gmdate('Y-m-d\TH:i:s.\\0\\0\\0\\Z', time());

        $queryString = $this->makeQueryString($params);
        $signature = $this->sign($path, $queryString);

        // https://mws.amazonservices.com/Products/2011-10-01
        $serviceUrl = $this->config['ServiceUrl'];
        $url = "https://$serviceUrl$path";
        $data = $queryString . "&Signature=" . $signature;

        $this->log($this->method.' '.$url);
        $this->log($data);

        if ($this->method == 'GET') {
            $url = "$url?$data";
        }

        $ch = curl_init($url);

#       curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-type: application/xml'));
        curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: x-www-form-urlencoded'));
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);

        if ($this->method == 'POST') {
            curl_setopt($ch, CURLOPT_POST, true);
            curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
        }

        $response = curl_exec($ch);
        $info = curl_getinfo($ch);
        curl_close($ch);

        $this->log($this->formatXml($response));
#       $this->log($info);
        $this->log(str_repeat('-', 80));

        $response = preg_replace('/&(?!#?[a-z0-9]+;)/', '&amp;', $response);
        return simplexml_load_string($response);
    }

    protected function makeQueryString($params)
    {
        $arr = [];

        foreach ($params as $key => $val) {
            $key = str_replace("%7E", "~", rawurlencode($key));
            $val = str_replace("%7E", "~", rawurlencode($val));
            $arr[] = "{$key}={$val}";
        }

        sort($arr);
#       uksort($arr, 'strcmp');

        $str = implode('&', $arr);

        return $str;
    }

    protected function sign($path, $queryString)
    {
        $secretKey  = $this->config['SecretKey'];
        $serviceUrl = $this->config['ServiceUrl'];

        $sign  = $this->method . "\n";  // 'GET' | 'POST'
        $sign .= $serviceUrl . "\n";    // 'mws.amazonservices.com'
        $sign .= $path . "\n";          // '/Products/2011-10-01'
        $sign .= $queryString;

        $signature = hash_hmac("sha256", $sign, $secretKey, true);
        $signature = urlencode(base64_encode($signature));

        return $signature;
    }
