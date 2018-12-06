<?php

interface MiddlewareInterface
{
    public function handle($message, callable $next);
}

class MiddlewareManager
{
    private $middlewares = [];

    public function __construct(array $middlewares = [])
    {
        foreach ($middlewares as $middleware) {
            $this->appendMiddleware($middleware);
        }
    }

    public function appendMiddleware(MiddlewareInterface $middleware)
    {
        $this->middlewares[] = $middleware;
    }

    public function handle($message)
    {
        call_user_func($this->callableForNextMiddleware(0), $message);
    }

    private function callableForNextMiddleware($index)
    {
        if (!isset($this->middlewares[$index])) {
            return function() {};
        }

        $middleware = $this->middlewares[$index];

        return function($message) use ($middleware, $index) {
            $middleware->handle($message, $this->callableForNextMiddleware($index + 1));
        };
    }
}

class Middleware_1 implements MiddlewareInterface
{
    public function handle($message, callable $next)
    {
        echo __CLASS__, ' is handling ', $message, PHP_EOL;
        $next($message);
    }
}

class Middleware_2 implements MiddlewareInterface
{
    public function handle($message, callable $next)
    {
        echo __CLASS__, ' is handling ', $message, PHP_EOL;
        if (strpos($message, 'REQ') !== false)
            return;
        $next($message);
    }
}

class Middleware_3 implements MiddlewareInterface
{
    public function handle($message, callable $next)
    {
        echo __CLASS__, ' is handling ', $message, PHP_EOL;
        $next($message);
    }
}

$mgr = new MiddlewareManager();

$mgr->appendMiddleware(new Middleware_1());
$mgr->appendMiddleware(new Middleware_2());
$mgr->appendMiddleware(new Middleware_3());

$mgr->handle('#REQUEST#');
