<?php

abstract class BasicHandler
{
    /**
     * @var BasicHandler
     */
    protected $successor = null;
 
    /**
     * Sets a successor handler.
     * 
     * @param BasicHandler $handler
     */
    public function setSuccessor(BasicHandler $handler)
    {
        $this->successor = $handler;
    }
 
    /**
     * Handles the request and/or redirect the request
     * to the successor.
     *
     * @param mixed $request
     *
     * @return mixed
     */
    abstract public function handle($request);
}


class FirstHandler extends BasicHandler
{
    public function handle($request)
    {
        // provide a response, call the next successor
        echo __METHOD__, "('$request');", PHP_EOL;
        if ($this->successor)
            $this->successor->handle($request);
    }
}

class SecondHandler extends BasicHandler
{
    public function handle($request)
    {
        // provide a response, call the next successor
        echo __METHOD__, "('$request');", PHP_EOL;
        if ($this->successor)
            $this->successor->handle($request);
    }
}

class ThirdHandler extends BasicHandler
{
    public function handle($request)
    {
        // provide a response, call the next successor
        echo __METHOD__, "('$request');", PHP_EOL;
        if ($this->successor)
            $this->successor->handle($request);
    }
}

$firstHandler = new FirstHandler();
$secondHandler = new SecondHandler();
$thirdHandler = new ThirdHandler();
 
$firstHandler->setSuccessor($secondHandler);
$secondHandler->setSuccessor($thirdHandler);
 
$request = 'REQUEST';
$result = $firstHandler->handle($request);
