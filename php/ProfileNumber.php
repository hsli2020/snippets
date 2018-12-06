<?php

const EOL = PHP_EOL;

class Bijective
{
#   public $dictionary = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    public $dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

    public function __construct()
    {
        $this->dictionary = str_split($this->dictionary);
    }

    public function encode($i)
    {
        if ($i == 0)
            return $this->dictionary[0];

        $result = '';
        $base = count($this->dictionary);

        while ($i > 0)
        {
            $result[] = $this->dictionary[($i % $base)];
            $i = floor($i / $base);
        }

        $result = array_reverse($result);

        return join("", $result);
    }

    public function decode($input)
    {
        $i = 0;
        $base = count($this->dictionary);

        $input = str_split($input);

        foreach($input as $char)
        {
            $pos = array_search($char, $this->dictionary);

            $i = $i * $base + $pos;
        }

        return $i;
    }
}

// $bi = new Bijective();
// $en = $bi->encode(4987654321);
// $de = $bi->decode($en);
// echo $en, EOL;
// echo $de, EOL;
#------------------------------------------------------------------------------
class ProfileNumberUtil
{
#   public $dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    protected static $dictionary = "PFNRE7K5AC6D98YLVXBM1SJ4Q032WIOTHZGU";

    public static function encode($i)
    {
        if ($i == 0)
            return 0;

        $dictionary = str_split(self::$dictionary);
        $checksum = array_sum(str_split($i));
        $base = count($dictionary);

        $result = '';
        $i ^= 0x7FFFFFFF;

        while ($i > 0) {
            $result[] = $dictionary[($i % $base)];
            $i = floor($i / $base);
        }

        $result = array_reverse($result);

#       return join("", $result) . str_pad($checksum, 2, '0', STR_PAD_LEFT);
        return join("", $result) . sprintf('%02X', $checksum);
    }

    public static function decode($input)
    {
        $i = 0;

        $dictionary = str_split(self::$dictionary);
        $base = count($dictionary);

        $checksum = substr($input, -2);
        $input = substr($input, 0, -2);
        $input = str_split($input);

        foreach($input as $char) {
            $pos = array_search($char, $dictionary);
            $i = $i * $base + $pos;
        }

        $i ^= 0x7FFFFFFF;

        if (array_sum(str_split($i)) != hexdec($checksum)) {
            return 0;
        }

        return $i;
    }
}

$en = ProfileNumberUtil::encode(9765432); echo $en, ' ==> ';
$de = ProfileNumberUtil::decode($en); echo $de, EOL;
#------------------------------------------------------------------------------
class ProfileNumber
{
#   public $dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    public $dictionary = "PFNRE7K5AC6D98YLVXBM1SJ4Q032WIOTHZGU";

    public function __construct()
    {
        $this->dictionary = str_split($this->dictionary);
    }

    public function encode($i)
    {
        if ($i == 0)
            return $this->dictionary[0];

        $checksum = array_sum(str_split($i));
        $base = count($this->dictionary);

        $result = '';
        $i ^= 0x7FFFFFFF;

        while ($i > 0) {
            $result[] = $this->dictionary[($i % $base)];
            $i = floor($i / $base);
        }

        $result = array_reverse($result);

#       return join("", $result) . str_pad($checksum, 2, '0', STR_PAD_LEFT);
        return join("", $result) . sprintf('%02X', $checksum);
    }

    public function decode($input)
    {
        $i = 0;
        $base = count($this->dictionary);

        $checksum = substr($input, -2);
        $input = substr($input, 0, -2);
        $input = str_split($input);

        foreach($input as $char) {
            $pos = array_search($char, $this->dictionary);
            $i = $i * $base + $pos;
        }

        $i ^= 0x7FFFFFFF;

        if (array_sum(str_split($i)) != hexdec($checksum)) {
            return 0;
        }

        return $i;
    }
}

$pn = new ProfileNumber();

$en = $pn->encode(87654321); echo $en, ' ==> ';
$de = $pn->decode($en); echo $de, EOL;

$en = $pn->encode(7654321); echo $en, ' ==> ';
$de = $pn->decode($en); echo $de, EOL;

$en = $pn->encode(322); echo $en, ' ==> ';
$de = $pn->decode($en); echo $de, EOL;

$en = $pn->encode(323); echo $en, ' ==> ';
$de = $pn->decode($en); echo $de, EOL;
