<?php

class ProfileNum_v1
{
    protected $encryptMethod = "AES-256-CBC";
    protected $secretKey     = 'secret key'; // getenv('SECRET_KEY');
    protected $secretIV      = 'secret iv';  // getenv('SECRET_IV');

    public function encrypt($string)
    {
#       $key = hash('sha256', $this->secretKey);
#       $iv = substr(hash('sha256', $this->secretIV), 0, 16);

        $key = md5($this->secretKey);
        $iv = substr(md5($this->secretIV), 0, 16);

        return base64_encode(openssl_encrypt($string, $this->encryptMethod, $key, 0, $iv));
    }

    public function decrypt($string)
    {
#       $key = hash('sha256', $this->secretKey);
#       $iv  = substr(hash('sha256', $this->secretIV), 0, 16);

        $key = md5($this->secretKey);
        $iv = substr(md5($this->secretIV), 0, 16);

        return openssl_decrypt(base64_decode($string), $this->encryptMethod, $key, 0, $iv);
    }
}

class ProfileNum
{
    public function encrypt($string)
    {
        list($key, $iv, $method) = $this->getCipherInfo();

        $output = openssl_encrypt($string, $method, $key, 0, $iv);
        return strtr(rtrim($output, '='), '+/', '-_');
    }

    public function decrypt($string)
    {
        list($key, $iv, $method) = $this->getCipherInfo();

        $string = strtr($string, '-_', '+/');
        return openssl_decrypt($string, $method, $key, 0, $iv);
    }

    protected function getCipherInfo()
    {
        $encryptMethod = "AES-128-CBC";

        $secretKey = getenv('PNUM_SECRET_KEY') ? : 'pnum secret key';
        $secretIV  = getenv('PNUM_SECRET_IV')  ? : 'pnum secret iv';

        $key = hash('sha256', $secretKey);
        $iv  = substr(hash('sha256', $secretIV), 0, openssl_cipher_iv_length($encryptMethod));

        return array($key, $iv, $encryptMethod);
    }
}

$nums = range(0, 9);
$nums = range(234567810, 234567890);
$nums = [ 1234567890, 234567890, 34567890, 4567890, 567890, 67890, 7890, 890, 90, 0 ];

$profileNum = new ProfileNum();
foreach ($nums as $num) {
    $en = $profileNum->encrypt($num); echo $en, " => ";
    $de = $profileNum->decrypt($en);  echo $de, "\n";
}
$en = $profileNum->encrypt('234567890,1234567890,34567890');  echo $en, "\n";
