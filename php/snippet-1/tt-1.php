<?php

const EOL = PHP_EOL;

date_default_timezone_set('America/New_York');

function pr()
{
    $output = '';
    foreach (func_get_args() as $var) {
        $output .= var_export($var, true) . EOL;
    }

    $patterns = [
        "/array \(/"    => "[",
        "/\),\n/"       => "],\n",
        "/\)\n/"        => "]\n",
        "/\)$/"         => "]",
        "/=> \n(\s+)/"  => "=> ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

// function pr($var) { var_export($var); }

function debugResponse($url, $payload, array $headers)
{
    // debug for xterm-compliant cli
    if (php_sapi_name() === 'cli') {
        echo PHP_EOL;
        echo "\e[1;33m<<<\e[0;33m [{$url}] \e[1;33m<<<\e[00m" . PHP_EOL;
        foreach ($headers as $key => $value) {
            echo "\e[1;34m{$key}: \e[0;34m{$value}\e[00m" . PHP_EOL;
        }
        echo PHP_EOL;
        echo ($json = json_decode($payload, true)) == NULL
                ? $payload : json_encode($json, JSON_PRETTY_PRINT);
        echo PHP_EOL;
    } else {  // debug for the rest
        echo "<<< [{$url}] <<<" . PHP_EOL;
        print_r(($json = json_decode($payload)) == NULL ? $payload : $json);
        echo PHP_EOL;
    }
}

function debugRequest($url, $payload, array $headers)
{
    // debug for xterm-compliant cli
    if (php_sapi_name() === 'cli') {
        echo PHP_EOL;
        echo "\e[1;33m>>>\e[0;33m [{$url}] \e[1;33m>>>\e[00m" . PHP_EOL;

        foreach ($headers as $value) {
            $matches = array();
            preg_match('#^(?P<key>[^:\s]+)\s*:\s*(?P<value>.*)$#S', $value, $matches);
            echo "\e[1;34m{$matches['key']}: \e[0;34m{$matches['value']}\e[00m" . PHP_EOL;
        }
        echo PHP_EOL;
        echo (is_array($payload))
            ? json_encode($payload, JSON_PRETTY_PRINT)
            : ((is_null($json = json_decode($payload)))
                ? $payload : json_encode($json, JSON_PRETTY_PRINT));

        echo PHP_EOL;
    } else {  // debug for the rest
        echo ">>> [{$url}] >>>" . PHP_EOL;
        $message = print_r(is_null($json = json_decode($payload)) ? $payload : $json, true);
        echo $message . (strpos($message, PHP_EOL) ? null : PHP_EOL);
    }
}

#debugResponse(
#    'http://www.google.com', 
#    '{ "first": 1, "second": 2, "third": 3, "letters": "abc" }',
#    [ 'username' => 'John', 'password' => '123456' ]
#);

#$arr = [ "first"=> 1, "second"=> 2, "third"=> 3, "abc" ];
#echo json_encode($arr, JSON_PRETTY_PRINT);

#debugRequest(
#    'http://www.google.com', 
#    '{ "first": 1, "second": 2, "third": 3, "letters": "abc" }',
#    [ 'username' => 'John', 'password' => '123456' ]
#);

function getParameter($name)
{
    $tokens  = explode('.', $name);
   #$context = $this->parameters;
    $context = [
        'db' => [
            'master' => [
                'host' => 'master-host',
                'port' => 'master-port',
            ],
            'slave' => [
                'host' => 'slave-host',
                'port' => 'slave-port',
            ],
        ],
    ];

    while (null !== ($token = array_shift($tokens))) {
        if (!isset($context[$token])) {
            return null;
        }

        $context = $context[$token];
    }

    return $context;
}

#echo getParameter('db.master.host'), EOL;
#echo getParameter('db.slave.port'), EOL;

/** 
 * Send a POST requst using cURL 
 * @param string $url to request 
 * @param array $post values to send 
 * @param array $options for cURL 
 * @return string 
 */ 
function curl_post($url, array $post = NULL, array $options = array()) 
{ 
    $defaults = array( 
        CURLOPT_POST => 1, 
        CURLOPT_HEADER => 0, 
        CURLOPT_URL => $url, 
        CURLOPT_FRESH_CONNECT => 1, 
        CURLOPT_RETURNTRANSFER => 1, 
        CURLOPT_FORBID_REUSE => 1, 
        CURLOPT_TIMEOUT => 4, 
        CURLOPT_POSTFIELDS => http_build_query($post) 
    ); 

    $ch = curl_init(); 

    curl_setopt_array($ch, ($options + $defaults)); 
    if (!$result = curl_exec($ch)) { 
        trigger_error(curl_error($ch)); 
    } 
    curl_close($ch); 

    return $result; 
} 

function randomString($type = 'alnum', $len = 8)
{
    $types = array(
        'basic'   => function() { return mt_rand(); },
        'alpha'   => 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'alnum'   => '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'numeric' => '0123456789',
        'nozero'  => '123456789',
        'unique'  => function() { return  md5(uniqid(mt_rand(), TRUE)); },
        'md5'     => function() { return  md5(uniqid(mt_rand(), TRUE)); },
        'encrypt' => function() { return sha1(uniqid(mt_rand(), TRUE)); },
        'sha1'    => function() { return sha1(uniqid(mt_rand(), TRUE)); },
    );

    if (isset($types[$type])) {
        $pool = $types[$type];
        if (is_callable($pool)) {
            return $pool();
        }

        if (is_string($pool)) {
            $str = '';
            for ($i=0; $i < $len; $i++) {
                $str .= substr($pool, mt_rand(0, strlen($pool) - 1), 1);
            }
            return $str;
        }
    }

    return '';
}

# echo randomString('basic'), EOL;
# echo randomString('alpha'), EOL;
# echo randomString('alnum'), EOL;
# echo randomString('numeric'), EOL;
# echo randomString('nozero'), EOL;
# echo randomString('unique'), EOL;
# echo randomString('md5'), EOL;
# echo randomString('encrypt'), EOL;
# echo randomString('sha1'), EOL;

function randomString0($strLen = 32)
{
    // Create our character arrays
    $chrs = array_merge(range('a', 'z'), range('A', 'Z'), range(0, 9));

    // Just to make the output even more random
    shuffle($chrs);

    // Create a holder for our string
    $randStr = '';

    // Now loop through the desired number of characters for our string
    for ($i=0; $i<$strLen; $i++) {
        $randStr .= $chrs[mt_rand(0, (count($chrs) - 1))];
    }

    return $randStr;
}

function getPrimes($max) { // ??
    $primes = array();
    for ($x = 2; $x <= $max; $x++) {
        $xIsPrime = TRUE;
        $sqrtX = sqrt($x);
        foreach ($primes as $prime) {
            if ($prime > $sqrtX || ((!($x % $prime)) && (!$xIsPrime = FALSE))) 
                break;
        }
        if ($xIsPrime) echo ($primes[] = $x)  . " ";
    }
    echo EOL;
    return $primes;
}

getPrimes(100);
# for ($x = 2; $x <= 100; $x++) {
#     if (isPrimeNumber($x)) echo $x, ' ';
# }

