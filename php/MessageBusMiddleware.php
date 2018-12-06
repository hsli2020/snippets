<?php
# common/vendor/simple-bus/message-bus/src/Bus/MessageBus.php

namespace SimpleBus\Message\Bus;

interface MessageBus
{
    /**
     * @param object $message
     * @return void
     */
    public function handle($message);
}

<?php
# common/vendor/simple-bus/message-bus/src/Bus/Middleware/MessageBusMiddleware.php

namespace SimpleBus\Message\Bus\Middleware;

interface MessageBusMiddleware
{
    /**
     * The provided $next callable should be called whenever the next middleware should start handling the message.
     * Its only argument should be a Message object (usually the same as the originally provided message).
     *
     * @param object $message
     * @param callable $next
     * @return void
     */
    public function handle($message, callable $next);
}

<?php
# common/vendor/simple-bus/message-bus/src/Bus/Middleware/MessageBusSupportingMiddleware.php

namespace SimpleBus\Message\Bus\Middleware;

use SimpleBus\Message\Bus\MessageBus;

class MessageBusSupportingMiddleware implements MessageBus
{
    /**
     * @var MessageBusMiddleware[]
     */
    private $middlewares = [];

    public function __construct(array $middlewares = [])
    {
        foreach ($middlewares as $middleware) {
            $this->appendMiddleware($middleware);
        }
    }

    /**
     * Appends new middleware for this message bus. Should only be used at configuration time.
     *
     * @private
     * @param MessageBusMiddleware $middleware
     * @return void
     */
    public function appendMiddleware(MessageBusMiddleware $middleware)
    {
        $this->middlewares[] = $middleware;
    }

    /**
     * Returns a list of middlewares. Should only be used for introspection.
     *
     * @private
     * @return array<MessageBusMiddleware>
     */
    public function getMiddlewares()
    {
        return $this->middlewares;
    }

    /**
     * Prepends new middleware for this message bus. Should only be used at configuration time.
     *
     * @private
     * @param MessageBusMiddleware $middleware
     * @return void
     */
    public function prependMiddleware(MessageBusMiddleware $middleware)
    {
        array_unshift($this->middlewares, $middleware);
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
