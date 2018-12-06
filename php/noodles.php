<?php

$noodles = 1;

class noodles
{
    var $noodles = 2;

    function noodles()
    {
        $noodles['noodles'] = 'noodles';
    }
}

function noodles() {
    return new noodles;
}

$noodles = noodles();
var_dump($noodles);
