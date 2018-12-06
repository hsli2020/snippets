<?php

// token comes from userid (same user always get same token in one day)
class CsrfToken_v1
{
    public static function genToken($userid)
    {
        $date = date('Y-m-d');
        $salt = 's@1t$alt$@1t';

        return md5($userid . $date . $salt);
    }

    public static function checkToken($userid, $token)
    {
        return self::genToken($userid) == $token;
    }
}

// token comes from userid+formkey
class CsrfToken_v2
{
    public static function getFormKey()
    {
        return substr(md5(uniqid(mt_rand(), true)), 0, 16);
    }

    public static function genToken($userid, $formkey)
    {
        $date = date('Y-m-d');
        $salt = 's@1t$alt$@1t';

        return md5($userid . $date . $formkey . $salt);
    }

    public static function checkToken($userid, $formkey, $token)
    {
        return self::genToken($userid, $formkey) == $token;
    }
}

// token comes from formkey, userid is not needed
class CsrfToken_v3
{
    public static function getFormKey()
    {
        return substr(md5(uniqid(mt_rand(), true)), 0, 16);
    }

    public static function genToken($formkey)
    {
        $date = date('Y-m-d');
        $salt = 's@1t$alt$@1t';

        return md5($date . $formkey . $salt);
    }

    public static function checkToken($formkey, $token)
    {
        return self::genToken($formkey) == $token;
    }
}

define('EOL', PHP_EOL);
date_default_timezone_set('America/New_York');

$formkey = CsrfToken_v3::getFormKey();
$token  = CsrfToken_v3::genToken($formkey);

// form submitted
echo $formkey, EOL;
echo $token, EOL;

if (CsrfToken_v3::checkToken($formkey, $token)) {
    echo 'good token', EOL;
} else {
    echo 'bad token', EOL;
}
