<?php

abstract class AdvancedHandler
{
    /**
     * @var AdvancedHandle
     */
    protected $successor = null;
 
    /**
     * Sets a successor handler,
     * in case the class is not able to satisfy the request.
     *
     * @param AdvancedHandler $handler
     */
    final public function setSuccessor(AdvancedHandler $handler)
    {
        if ($this->successor === null) {
            $this->successor = $handler;
        } else {
            $this->successor->setSuccessor($handler);
        }
    }
 
    /**
     * Handles the request or redirect the request
     * to the successor, if the process response is null.
     *
     * @param string|array $data
     *
     * @return string
     */
    final public function handle($request)
    {
        $response = $this->process($request);
        if (($response === null) && ($this->successor !== null)) {
            $response = $this->successor->handle($request);
        }
 
        return $response;
    }
 
    /**
     * Processes the request.
     * This is the only method a child can implements, 
     * with the constraint to return null to dispatch the request to next successor.
     * 
     * @param $request
     *
     * @return null|mixed
     */
    abstract protected function process($request);
}

class FirstHandler extends AdvancedHandler
{
    public function process($request)
    {
        echo __METHOD__, "('$request');", PHP_EOL;
    }
}
 
class SecondHandler extends AdvancedHandler
{
    public function process($request)
    {
        echo __METHOD__, "('$request');", PHP_EOL;
    }
}
 
class ThirdHandler extends AdvancedHandler
{
    public function process($request)
    {
        echo __METHOD__, "('$request');", PHP_EOL;
    }
}
 
$firstHandler = new FirstHandler();
$secondHandler = new SecondHandler();
$thirdHandler = new ThirdHandler();
 
// the code below sets all successors through the first handler
$firstHandler->setSuccessor($secondHandler);
$firstHandler->setSuccessor($thirdHandler);
 
$request = 'REQUEST';
$result = $firstHandler->handle($request);
