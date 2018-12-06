<?php
/*
    Sample usage:

    Looking up an ASIN is very simple:
    
    $bettermws = new bettermws('your_merchant_id', 'your_marketplace_id', 'your_secret_key');
    $pricing = $bettermws->itemlookup($asin);

*/

class bettermws
{
    private $merchant_id;
    private $marketplace_id;
    private $secret;

    private $ua = 'BetterMWS PHP Class Version 0.1';

    private $endpoint = 'http://api.bettermws.com/v1';

    public function __construct($merchant_id, $marketplace_id, $secret)
    {
        $this->merchant_id      = $merchant_id;
        $this->marketplace_id   = $marketplace_id;
        $this->secret           = $secret;
    }



    public function itemlookup($asin, $persist = null)
    {
        $persist    = (int)$persist;

        $url        = $this->endpoint."/lookup?merchant_id={$this->merchant_id}&marketplace_id={$this->marketplace_id}&asin={$asin}";
        $url       .= '&format=sphp';
        $url       .= "&ts=".time();
        $url       .= $persist ? "&persist={$persist}" : '';
        $signature  = md5($this->secret . $url);
        $url       .= "&sig={$signature}";
        $response =  $this->sendRequest($url);
        return unserialize($response);
    }


    public function bulkpersist($asins, $persist = null)
    {
        if (!is_array($asins)) {
            throw new Exception('bettermws::bulkpersist() requires first argument as an array');
        }
        $persist = (int)$persist;
        if ($persist < 0 || $persist > 168) {
            $persist = 2;
        }

        $url = $this->endpoint.'/bulkpersist';

        $params = array(
            'merchant_id'       => $this->merchant_id,
            'marketplace_id'    => $this->marketplace_id,
            'persist'           => $persist,
            'ts'                => time(),
            'asins'             => implode(',', $asins),
            'format'            => 'sphp',
        );

        $post_data = http_build_query($params);
        $string_to_sign = $this->secret . $url . '?'. $post_data;
        $signature = md5($string_to_sign);
        $post_data .= "&sig={$signature}";

        $response = $this->sendRequest($url, $post_data);
        return unserialize($response);

    }


    private function sendRequest($url, $post_data = '')
    {
        $ch = curl_init();
        if ($post_data) {
            curl_setopt($ch, CURLOPT_POST,              true);
            curl_setopt($ch, CURLOPT_POSTFIELDS,        $post_data);
        }
        curl_setopt($ch, CURLOPT_URL,               $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER,    true);
        // curl_setopt($ch, CURLOPT_HEADER,         true);
        // curl_setopt($ch, CURLINFO_HEADER_OUT,    true);      // for use with debugging
        // curl_setopt($ch, CURLOPT_VERBOSE,        true);      // for use with debugging
        curl_setopt($ch, CURLOPT_TIMEOUT,           30);
        curl_setopt($ch, CURLOPT_USERAGENT,         $this->ua);

        $result = curl_exec($ch);
        curl_close($ch);
        return $result;
    }

}

