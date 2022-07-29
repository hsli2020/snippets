<?php

/* First-Class-Callable

this code...

    $myFunction = strtoupper(...);

... is equivalent to:

    $myFunction = function(...$arguments) {
        return strtoupper(...$argument);
    }

Let's use it.

    $myFunction('a') // returns 'A';
*/

// Old way
$arr = array_map(
    function($letter) { return strtoupper($letter); }, 
    ['a', 'b', 'c']
);
print_r($arr);

// New way
$arr = array_map(
    strtoupper(...), 
    ['a', 'b', 'c']
);
print_r($arr);

// Class method
class MyClass
{
    public function execute()
    {
        return array_map(
            $this->doubleString(...),
            ['a', 'b', 'c']
        );
    }
    
    public function doubleString(string $string): string
    {
        return $string . $string;
    }
}

// returns an array with 'aa', 'bb, and 'cc'.
$a = (new MyClass)->execute();
print_r($a);
