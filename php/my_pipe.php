<?php

function pipe($subject) {
    return new Pipe($subject);
}

class Pipe implements \Stringable, \ArrayAccess, \IteratorAggregate
{
    use Transparency;

    function __construct($target)
    {
        $this->target = $target;
    }

    function __invoke(...$params) {
        if (empty($params)) return $this->target;

        [ $before, $through, $after ] = [ [], null, [] ];

        foreach ($params as $key => $param) {
            if (! $through) {
                if (is_callable($param)) $through = $param;

                else $before[$key] = $param;
            } else {
                $after[$key] = $param;
            }
        }

        $params = [ ...$before, $this->target, ...$after ];

        $this->target = $through(...$params);

        return $this;
    }
}

trait Transparency
{
    public $target;

    function __toString()
    {
        return (string) $this->target;
    }

    function offsetExists(mixed $offset): bool
    {
        return isset($this->target[$offset]);
    }

    function offsetGet(mixed $offset): mixed
    {
        return $this->target[$offset];
    }

    function offsetSet(mixed $offset, mixed $value): void
    {
        $this->target[$offset] = $value;
    }

    function offsetUnset(mixed $offset): void
    {
        unset($this->target[$offset]);
    }

    function getIterator(): \Traversable
    {
        return (function () {
            foreach ($this->target as $key => $value) {
                yield $key => $value;
            }
        })();
    }
}

