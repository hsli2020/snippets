<?php

#url = 'https://view.news.qq.com/l/history_new/dsj/list20130813152200.htm';
$url = 'https://www.wxnmh.com/user-42989.htm';
$html = httpGet($url);
file_put_contents('dsj-2018-1.html', $html);

$maxpg = 35;
for ($p=6; $p<=$maxpg; $p++) {
#   $url = "https://view.news.qq.com/l/history_new/dsj/list20130813152200_$p.htm";
    $url = "https://www.wxnmh.com/user-42989-$p.htm";
    echo $url, PHP_EOL;
    $html = httpGet($url);
    file_put_contents("dsj-2018-$p.html", $html);
    sleep(2);
}

function httpGet($url)
{
    $ua = 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.13) Gecko/20080311 Firefox/2.0.0.13';
    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_USERAGENT, $ua);
    $output = curl_exec($ch);

    if(!curl_exec($ch)){
        echo 'Error: "' . curl_error($ch) . '" - Code: ' . curl_errno($ch), PHP_EOL;
    }
    curl_close($ch);      
    return $output;
}
