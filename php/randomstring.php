<?php

function randomString($type = 'alnum', $len = 8)
{
    $types = array(
        'basic'   => function() { return mt_rand(); },
        'alpha'   => 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'alnum'   => '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'numeric' => '0123456789',
        'nozero'  => '123456789',
		'unique'  => function() { return md5(uniqid(mt_rand())); },
		'md5'     => function() { return md5(uniqid(mt_rand())); },
		'encrypt' => function() { return sha1(uniqid(mt_rand(), TRUE)); },
		'sha1'	  => function() { return sha1(uniqid(mt_rand(), TRUE)); },
    );

    if (isset($types[$type])) {
       $pool = $types[$type];
       if (is_callable($pool)) {
            return $pool();
       }

       if (is_string($pool)) {
            $str = '';
            for ($i=0; $i < $len; $i++) {
                $str .= substr($pool, mt_rand(0, strlen($pool) -1), 1);
            }
            return $str;
       }
    }

    return '';
}

