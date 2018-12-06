<?php

class SessionKey
{
    const MAX_SESSION_LIFE_TIME = 3600;

    protected static $pass = '1234pswd5678';

    public static function generate($userid)
    {
        $chksum = self::getChecksum();
#       $string = json_encode(array($userid, time(), $chksum));
        $string = implode(',', array($userid, time(), $chksum));

        $cipher = new Cipher(self::$pass);
        return $cipher->encrypt($string);
    }

    public static function getUserId($sesskey = null)
    {
        if (empty($sesskey)) {
            $sesskey = $_COOKIE['uid'];
        }

        $cipher = new Cipher(self::$pass);
        $string = $cipher->decrypt($sesskey);

#       list($userid, $time, $chksum) = json_decode($string);
        list($userid, $time, $chksum) = explode(',', $string);

        if (time() - $time > self::MAX_SESSION_LIFE_TIME) {
            return false;
        }

        if ($chksum != self::getChecksum()) {
            return false;
        }

        return $userid;
    }

    protected static function getChecksum()
    {
        $chksum = 2425597808; // crc32($_SERVER['REMOTE_ADDR'].$_SERVER['HTTP_USER_AGENT']);
        return $chksum;
    }
}

// @codingStandardsIgnoreStart
class Cipher
{
    private $securekey;
    private $iv;

    public function __construct($textkey)
    {
        $this->securekey = hash('sha256', $textkey, true);
        $this->iv = mcrypt_create_iv(32);
    }

    public function encrypt($input)
    {
        return base64_encode(mcrypt_encrypt(
            MCRYPT_RIJNDAEL_256,
            $this->securekey,
            $input,
            MCRYPT_MODE_ECB,
            $this->iv
        ));
    }

    public function decrypt($input)
    {
        return trim(mcrypt_decrypt(
            MCRYPT_RIJNDAEL_256,
            $this->securekey,
            base64_decode($input),
            MCRYPT_MODE_ECB,
            $this->iv
        ));
    }
}
// @codingStandardsIgnoreEnd

//*
const EOL = PHP_EOL;

$userid  = 87654321;
$sesskey = SessionKey::generate($userid);
echo $sesskey, EOL;

$sesskeys = [
    '3UpW5mTJHxBML1xoBoyEmhHjxVxBWCPELIJfMRHmY/E=',
    'Sw6Cm1nKEOnyw6ZnlFYsiwmknkzBKg7mI0aPruhfRwM=',
];

foreach ($sesskeys as $sesskey) {
    $uid  = SessionKey::getUserId($sesskey);
    echo $uid, EOL;
}
//*/
