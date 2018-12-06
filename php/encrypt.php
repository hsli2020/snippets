<?php

const EOL = PHP_EOL;

date_default_timezone_set('UTC');

# http://php.net/manual/en/function.mcrypt-encrypt.php

class Cipher
{
    private $securekey, $iv;

    function __construct($textkey)
    {
        $this->securekey = hash('sha256', $textkey, TRUE);
        $this->iv = mcrypt_create_iv(32);
    }

    function encrypt($input)
    {
        return base64_encode(mcrypt_encrypt(MCRYPT_RIJNDAEL_256, $this->securekey, $input, MCRYPT_MODE_ECB, $this->iv));
    }

    function decrypt($input)
    {
        return trim(mcrypt_decrypt(MCRYPT_RIJNDAEL_256, $this->securekey, base64_decode($input), MCRYPT_MODE_ECB, $this->iv));
    }
}

$data = "http://www.ashleymadison.com/";
$key  = "keepsecret";

$cipher = new Cipher($key);

$encryptedtext = $cipher->encrypt($data);
echo "$encryptedtext\n";

$decryptedtext = $cipher->decrypt($encryptedtext);
echo "$decryptedtext\n";

// var_dump($cipher);

#-------------------------------------------------------------------------------
//*
class Crypt
{
    public static function getiv($cipher='twofish', $mode='cfb')
    {
        $td = mcrypt_module_open($cipher, '', $mode, '');
        $iv = mcrypt_create_iv(mcrypt_enc_get_iv_size($td), MCRYPT_RAND);
        mcrypt_module_close($td);

        return $iv;
    }

    public static function encrypt($data, $key, $iv, $cipher='twofish', $mode='cfb')
    {
        $td = mcrypt_module_open($cipher, '', $mode, '');
        $key = substr($key, 0, mcrypt_enc_get_key_size($td));
        mcrypt_generic_init($td, $key, $iv);
        $crypted = mcrypt_generic($td, $data);
        mcrypt_generic_deinit($td);
        mcrypt_module_close($td);

        return $crypted;
    }

    public static function decrypt($data, $key, $iv, $cipher='twofish', $mode='cfb')
    {
        $td = mcrypt_module_open($cipher, '', $mode, '');
        $key = substr($key, 0, mcrypt_enc_get_key_size($td));
        mcrypt_generic_init($td, $key, $iv);
        $decrypted = rtrim(mdecrypt_generic($td, $data), "\0");
        mcrypt_generic_deinit($td);
        mcrypt_module_close($td);

        return $decrypted;
    }
}

$iv = Crypt::getiv();
$data = "http://www.ashleymadison.com";
$key  = "keepsecret";

$e = Crypt::encrypt($data, $key, $iv);
echo base64_encode($e), EOL;

$d = Crypt::decrypt($e, $key, $iv);
echo $d, EOL;
//*/
