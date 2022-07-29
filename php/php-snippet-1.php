<?php

public function initService($task): void
{
    $this->service = match ($task->type) {
        'type_a' => new ThisService(),
        'type_b' => new ThatService(),
        default  => new DefaultService(),
    };
}

# ===========================================================

class Foo
{
    protected function __construct($arg)
    {
        $this->arg = $arg;
    }

    public static function init($arg) // new/create/query/from
    {
        return new self($arg);
    }

    public function hello()
    {
        echo "Hello";
        return $this;
    }

    public function world()
    {
        echo "world";
        return $this;
    }

    public function comma()
    {
        echo ", ";
        return $this;
    }

    public function fullstop()
    {
        echo ".";
        return $this;
    }

    public function newline()
    {
        echo "\n";
        return $this;
    }
}

Foo::init('greeting')
    ->hello()
    ->comma()
    ->world()
    ->fullstop()
    ->newline();

# ===========================================================

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


# ===========================================================

final class File
{
    public static function exists($filename)
    {
        return file_exists($filename);
    }

    public static function read($filename)
    {
        return file_get_contents($filename);
    }

    public static function write($filename, $content)
    {
        return file_put_contents($filename, $content);
    }

    public static function size($filename)
    {
        return filesize($filename);
    }

    public static function rename($oldname, $newname)
    {
        return rename($oldname, $newname);
    }

    public static function copy($oldname, $newname)
    {
        return copy($oldname, $newname);
    }

    public static function delete($filename)
    {
        return unlink($filename);
    }
}
