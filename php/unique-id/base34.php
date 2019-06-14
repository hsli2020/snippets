<?php

# https://stackoverflow.com/questions/1467581/how-to-generate-unique-id-in-mysql

# SELECT TO_BASE64(RANDOM_BYTES(16));

$a = 4200000019;
$a = date('Ymd').rand(10000, 99999);
$b = base_convert($a, 10, 34); // 2oevc0b
$c = str_replace(['0', '1'], ['y', 'z'], $b);
echo strtoupper($c), PHP_EOL; // 2OEVCYB

$b = str_replace(['y', 'z'], ['0', '1'], $c);
$a = base_convert($b, 34, 10);
echo $a, PHP_EOL;

/*
You may like the way that we do it. I wanted a reversible unique code that looked
"random" -a fairly common problem.

    We take an input number such as 1,942.
    Left pad it into a string: "0000001942"
    Put the last two digits onto the front: "4200000019"
    Convert that into a number: 4,200,000,019

We now have a number that varies wildly between calls and is guaranteed to be less
than 10,000,000,000. Not a bad start.

    Convert that number to a Base 34 string: "2oevc0b"
    Replace any zeros with 'y' and any ones with 'z': "2oevcyb"
    Upshift: "2OEVCYB"

The reason for choosing base 34 is so that we don't worry about 0/O and 1/l collisions.
Now you have a short random-looking key that you can use to look up a LONG database identifier.
*/
