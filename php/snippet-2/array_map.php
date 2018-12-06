<?php

$default['Accept']                   = "application/xml";
$default['WM_SVC.NAME']              = "Walmart Marketplace";
$default['WM_CONSUMER.ID']           = "consumerID";
$default['WM_SEC.TIMESTAMP']         = time();
$default['WM_SEC.AUTH_SIGNATURE']    = "WalmartAuthSignature";
$default['WM_CONSUMER.CHANNEL.TYPE'] = "channelType";
$default['WM_QOS.CORRELATION_ID']    = "wm-bte-cli";

$header['Content-Type'] = "application/json";

$headers = array_merge($default, $header);

$headers = array_map(
    function($key, $val) {
       return $key.': '.$val; 
    }, array_keys($headers), $headers);

print_r($headers);
