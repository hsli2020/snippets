<?php

class Pimple implements ArrayAccess
{
    protected $values = array();

    public function __construct(array $values = array())
    {
        $this->values = $values;
    }

    public function offsetSet($id, $value) { $this->values[$id] = $value; }
    public function offsetExists($id) { return array_key_exists($id, $this->values); }
    public function offsetUnset($id) { unset($this->values[$id]); }

    public function offsetGet($id)
    {
        if (!array_key_exists($id, $this->values)) {
            throw new InvalidArgumentException(sprintf('"%s" is not defined.', $id));
        }

        $isFactory = is_object($this->values[$id]) && method_exists($this->values[$id], '__invoke');

        return $isFactory ? $this->values[$id]($this) : $this->values[$id];
    }

    public static function share($callable)
    {
        if (!is_object($callable) || !method_exists($callable, '__invoke')) {
            throw new InvalidArgumentException('not a Closure or invokable object.');
        }

        return function ($c) use ($callable) {
            static $object;

            if (null === $object) {
                $object = $callable($c);
            }

            return $object;
        };
    }

    public static function protect($callable)
    {
        if (!is_object($callable) || !method_exists($callable, '__invoke')) {
            throw new InvalidArgumentException('not a Closure or invokable object.');
        }

        return function ($c) use ($callable) {
            return $callable;
        };
    }

    public function raw($id)
    {
        if (!array_key_exists($id, $this->values)) {
            throw new InvalidArgumentException(sprintf('"%s" is not defined.', $id));
        }

        return $this->values[$id];
    }

    public function extend($id, $callable)
    {
        if (!array_key_exists($id, $this->values)) {
            throw new InvalidArgumentException(sprintf('"%s" is not defined.', $id));
        }

        if (!is_object($this->values[$id]) || !method_exists($this->values[$id], '__invoke')) {
            throw new InvalidArgumentException(sprintf('"%s" not an object definition.', $id));
        }

        if (!is_object($callable) || !method_exists($callable, '__invoke')) {
            throw new InvalidArgumentException('not a Closure or invokable object.');
        }

        $factory = $this->values[$id];

        return $this->values[$id] = function ($c) use ($callable, $factory) {
            return $callable($factory($c), $c);
        };
    }

    public function keys() { return array_keys($this->values); }
}
