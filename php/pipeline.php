<?php

// Unleashing the Power of Pipelines: A Custom PHP Implementation
// Unchain yourself from Laravel’s ecosystem and enjoy pipelines anywhere!

// Before we dive into the code, let me give you a quick overview of what pipelines are and 
// why you should care. In a nutshell, pipelines allow you to sequentially pass a value 
// through a series of callable “pipes” (read: middleware, filters, or processors). 
// Each pipe can modify the value before passing it along to the next one in the chain. 
// This pattern is particularly useful for handling requests, processing data, or applying 
// transformations in a clean, maintainable, and testable way.

class Pipeline
{
    protected $passable;
    protected $pipes;
    protected $method = 'handle';

    public static function send($passable)
    {
        $pipeline = new static;

        $pipeline->passable = $passable;

        return $pipeline;
    }

    public function through(array $pipes)
    {
        $this->pipes = $pipes;

        return $this;
    }

    public function then(Closure $destination)
    {
        $pipeline = array_reduce(
            array_reverse($this->pipes),
            $this->carry(),
            function ($passable) use ($destination) {
                return $destination($passable);
            }
        );

        return $pipeline($this->passable);
    }

    public function thenReturn()
    {
        return $this->then(function ($passable) {
            return $passable;
        });
    }

    protected function carry()
    {
        return function ($stack, $pipe) {
            return function ($passable) use ($stack, $pipe) {
                if (is_callable($pipe)) {
                    return $pipe($passable, $stack);
                } elseif (is_object($pipe)) {
                    return $pipe->{$this->method}($passable, $stack);
                } elseif (is_string($pipe) && class_exists($pipe)) {
                    $pipeInstance = new $pipe;
                    return $pipeInstance->{$this->method}($passable, $stack);
                } else {
                    throw new \InvalidArgumentException('Invalid pipe type.');
                }
            };
        };
    }
}

// To use Pipeline, all you need to do is:
$result = Pipeline::send($input)
    ->through($pipes)
    ->thenReturn();

// Imagine we have a string that needs to be processed through a series of transformations: 
// trim the input, capitalize the first letter of each word, and add an exclamation mark at 
// the end. Here’s how we can do it using our custom Pipeline class:

interface PipeInterface
{
    public function handle($passable, $next);
}

class TrimPipe implements PipeInterface
{
    public function handle($passable, $next)
    {
        $trimmed = trim($passable);
        return $next($trimmed);
    }
}

class CapitalizePipe implements PipeInterface
{
    public function handle($passable, $next)
    {
        $capitalized = ucwords(strtolower($passable));
        return $next($capitalized);
    }
}

class AddExclamationPipe implements PipeInterface
{
    public function handle($passable, $next)
    {
        $withExclamation = $passable . '!';
        return $next($withExclamation);
    }
}

class RemoveExtraSpacesPipe implements PipeInterface
{
    public function handle($passable, $next)
    {
        $withoutExtraSpaces = preg_replace('/\s+/', ' ', $passable);
        return $next($withoutExtraSpaces);
    }
}

// Now that we have our pipes ready, 
// let’s pass our input string through them using Mario’s Pipeline:

$input = "   thE quiCk BroWn   fOx    ";
$pipes = [
    TrimPipe::class,
    CapitalizePipe::class,
    RemoveExtraSpacesPipe::class,
    AddExclamationPipe::class,
];

$result = Pipeline::send($input)
    ->through($pipes)
    ->thenReturn();

echo $result; // Output: "The Quick Brown Fox!"