<?php

// Descartes

$a = [ 'a1', 'a2' ];
$b = [ 'b1', 'b2', 'b3' ];
$c = [ 'c1', 'c2' ];

function method1()
{
    $t = func_get_args();

    if (func_num_args() == 1) {
        $t0 = $t[0];
        return call_user_func_array(__FUNCTION__, $t0);
    }

    $a = array_shift($t);
    if (!is_array($a)) {
        $a = array($a);
    }

    $a = array_chunk($a, 1);
    do {
        $r = array();

        $b = array_shift($t);
        if (!is_array($b)) {
            $b = array($b);
        }

        foreach ($a as $p) {
            foreach (array_chunk($b, 1) as $q) {
                $r[] = array_merge($p, $q);
            }
        }

        $a = $r;
    } while ($t);

    return $r;
}

//print_r(method1($a, $b, $c));

function method2($array)
{
    $a = array_shift($array);
    $b = array_shift($array);

    $result = array();

    foreach ($a as $av) {
        foreach($b as $bv) {
            $result[] = $av.'-'.$bv;
        }
    }

    array_unshift($array, $result);

    if (count($array) > 1) {
        $result = method2($array);
    }

    return $result;
}

print_r(method2([$a, $b, $c]));

function method3($arr, $tmp = array())
{
    foreach (array_shift($arr) as $v) {
        $tmp[] = $v;

        if ($arr) {
            method3($arr, $tmp);
        }
        else {
            $GLOBALS["res"][] = $tmp;
        }

        array_pop($tmp);
    }
}

method3([$a, $b, $c]);
//print_r($res);
