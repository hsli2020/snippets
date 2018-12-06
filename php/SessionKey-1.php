<?php

class SessionKey
{
    const MAX_SESSION_LIFE_TIME = 3600;

    protected static $iv     = "1234567812345678";
    protected static $pass   = '1234pswd5678';
    protected static $method = 'aes128';
#   protected static $method = 'AES-256-CBC';

    public static function generate($userid)
    {
        $chksum = self::getChecksum();
        $json = json_encode(array($userid, time(), $chksum));

        #serialize generate bigger data than json_encode
        #echo serialize(array($userid, time(), $chksum)), EOL;
        
        return openssl_encrypt($json, self::$method, md5(self::$pass), 0, self::$iv);
    }

    public static function getUserId($sesskey = null)
    {
        if (empty($sesskey))
            $sesskey = $_COOKIE['uid'];

        $json = openssl_decrypt($sesskey, self::$method, md5(self::$pass), 0, self::$iv);
        list($userid, $time, $chksum) = json_decode($json);

        if (time() - $time > self::MAX_SESSION_LIFE_TIME)
            return false;

        if ($chksum !== self::getChecksum())
            return false;

        return $userid;
    }

    protected static function getChecksum()
    {
        $chksum = 2425597808; // crc32($_SERVER['REMOTE_ADDR'].$_SERVER['HTTP_USER_AGENT']);
        return $chksum;
    }
}

const EOL = PHP_EOL;

$sesskey = SessionKey::generate(87654321);
echo $sesskey, EOL;

$userid = SessionKey::getUserId($sesskey);
echo $userid, EOL;
