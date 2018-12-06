<?php

class MagicMethods
{
    public function __construct()
    {
        parent::__construct();

        // PHP 5 allows developers to declare constructor methods for classes. Classes which
        // have a constructor method call this method on each newly-created object, so it is
        // suitable for any initialization that the object may need before it is used.
    }

    public function __destruct()
    {
        parent::__destruct();

        // PHP 5 introduces a destructor concept similar to that of other object-oriented 
        // languages, such as C++. The destructor method will be called as soon as there are 
        // no other references to a particular object, or in any order during the shutdown sequence.
    }

    public function __call($name, $arguments)
    {
        // __call() is triggered when invoking inaccessible methods in an object context.
    }

    public static function __callStatic($name, $arguments)
    {
        // __callStatic() is triggered when invoking inaccessible methods in a static context.
    }

    public function __get($name)
    {
        // __get() is utilized for reading data from inaccessible properties.
    }

    public function __set($name, $value)
    {
        // __set() is run when writing data to inaccessible properties.
    }

    public function __isset($name)
    {
        // __isset() is triggered by calling isset() or empty() on inaccessible properties.
    }

    public function __unset($name)
    {
        // __unset() is invoked when unset() is used on inaccessible properties.
    }

    public function __sleep()
    {
        // serialize() checks if your class has a function with the magic name __sleep().
        // If so, that function is executed prior to any serialization. It can clean up the
        // object and is supposed to return an array with the names of all variables of that
        // object that should be serialized. If the method doesn't return anything then NULL
        // is serialized and E_NOTICE is issued. 
    }

    public function __wakeup()
    {
        // Conversely, unserialize() checks for the presence of a function with the magic name
        // __wakeup(). If present, this function can reconstruct any resources that the object
        // may have.
    }

    public function __toString()
    {
        // The __toString() method allows a class to decide how it will react when it is treated
        // like a string. For example, what echo $obj; will print. This method must return a string,
        // as otherwise a fatal E_RECOVERABLE_ERROR level error is emitted.
    }

    public function __invoke()
    {
        //  The __invoke() method is called when a script tries to call an object as a function.
    }

    public function __set_state($properties)
    {
    }

    public function __clone()
    {
    }

    public function __debugInfo()
    {
        // This method is called by var_dump() when dumping an object to get the properties
        // that should be shown. If the method isn't defined on an object, then all public,
        // protected and private properties will be shown. 
    }
}
