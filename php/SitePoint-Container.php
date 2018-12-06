<?php

namespace SitePoint\Container;

use SitePoint\Container\Exception\ContainerException;
use SitePoint\Container\Exception\ParameterNotFoundException;
use SitePoint\Container\Exception\ServiceNotFoundException;
use SitePoint\Container\Reference\ParameterReference;
use SitePoint\Container\Reference\ServiceReference;

/**
 * A very simple dependency injection container.
 */
class Container implements ContainerInterface
{
    /**
     * @var array
     */
    private $services;
    private $parameters;
    private $serviceStore;

    /**
     * Constructor for the container.
     *
     * Entries into the $services array must be an associative array with a
     * 'class' key and an optional 'arguments' key. Where present the arguments
     * will be passed to the class constructor. If an argument is an instance of
     * ContainerService the argument will be replaced with the corresponding
     * service from the container before the class is instantiated. If an
     * argument is an instance of ContainerParameter the argument will be
     * replaced with the corresponding parameter from the container before the
     * class is instantiated.
     *
     * @param array $services   The service definitions.
     * @param array $parameters The parameter definitions.
     */
    public function __construct(array $services = [], array $parameters = [])
    {
        $this->services     = $services;
        $this->parameters   = $parameters;
        $this->serviceStore = [];
    }

    public function get($name)
    {
        if (!$this->has($name)) {
            throw new ServiceNotFoundException('Service not found: '.$name);
        }

        // If we haven't created it, create it and save to store
        if (!isset($this->serviceStore[$name])) {
            $this->serviceStore[$name] = $this->createService($name);
        }

        // Return service from store
        return $this->serviceStore[$name];
    }

    public function has($name) { return isset($this->services[$name]); }

    public function getParameter($name)
    {
        $tokens  = explode('.', $name);
        $context = $this->parameters;

        while (null !== ($token = array_shift($tokens))) {
            if (!isset($context[$token])) {
                throw new ParameterNotFoundException('Parameter not found: '.$name);
            }

            $context = $context[$token];
        }

        return $context;
    }

    public function hasParameter($name)
    {
        try {
            $this->getParameter($name);
        } catch (ParameterNotFoundException $exception) {
            return false;
        }

        return true;
    }

    /**
     * Attempt to create a service.
     *
     * @param string $name The service name.
     *
     * @return mixed The created service.
     *
     * @throws ContainerException On failure.
     */
    private function createService($name)
    {
        $entry = $this->services[$name];

        if (!is_array($entry) || !isset($entry['class'])) {
            throw new ContainerException($name.
                ' service entry must be an array containing a \'class\' key');
        } elseif (!class_exists($entry['class'])) {
            throw new ContainerException($name.
                ' service class does not exist: '.$entry['class']);
        }

        $arguments = isset($entry['arguments']) 
                   ? $this->resolveArguments($name, $entry['arguments']) 
                   : [];

        $reflector = new \ReflectionClass($entry['class']);
        $service = $reflector->newInstanceArgs($arguments);

        if (isset($entry['calls'])) {
            $this->initializeService($service, $name, $entry['calls']);
        }

        return $service;
    }

    /**
     * Resolve argument definitions into an array of arguments.
     *
     * @param string $name                The service name.
     * @param array  $argumentDefinitions The service arguments definition.
     *
     * @return array The service constructor arguments.
     *
     * @throws ContainerException On failure.
     */
    private function resolveArguments($name, array $argumentDefinitions)
    {
        $arguments = [];

        foreach ($argumentDefinitions as $argumentDefinition) {
            if ($argumentDefinition instanceof ServiceReference) {
                $argumentServiceName = $argumentDefinition->getName();

                if ($argumentServiceName === $name) {
                    throw new ContainerException($name.' service contains a circular reference');
                }

                $arguments[] = $this->get($argumentServiceName);
            } elseif ($argumentDefinition instanceof ParameterReference) {
                $argumentParameterName = $argumentDefinition->getName();

                $arguments[] = $this->getParameter($argumentParameterName);
            } else {
                $arguments[] = $argumentDefinition;
            }
        }

        return $arguments;
    }

    /**
     * Initialize a service using the call definitions.
     *
     * @param object $service         The service.
     * @param string $name            The service name.
     * @param array  $callDefinitions The service calls definition.
     *
     * @throws ContainerException On failure.
     */
    private function initializeService($service, $name, array $callDefinitions)
    {
        foreach ($callDefinitions as $callDefinition) {
            if (!is_array($callDefinition) || !isset($callDefinition['method'])) {
                throw new ContainerException($name.
                    ' service calls must be arrays containing a \'method\' key');
            } elseif (!is_callable([$service, $callDefinition['method']])) {
                throw new ContainerException($name.
                    ' service asks for call to uncallable method: '
                    .$callDefinition['method']);
            }

            $arguments = isset($callDefinition['arguments']) 
                       ? $this->resolveArguments($name, $callDefinition['arguments']) 
                       : [];

            call_user_func_array([$service, $callDefinition['method']], $arguments);
        }
    }
}

namespace SitePoint\Container\Reference;

abstract class AbstractReference
{
    private $name;

    public function __construct($name) { $this->name = $name; }
    public function getName() { return $this->name; }
}

class ParameterReference extends AbstractReference {}
class ServiceReference extends AbstractReference {}

A simple, easy to follow PHP dependency injection container.

## How to Use

Although it isn't required to do so, a good practice is to split up the configuration for 
our container. In this example we'll use three files to create our container for the 
Monolog component.

Another good practice is to use class and interface paths as service names. This provides 
a stricter naming convention that gives us more information about the services.

In the service definitions file, we define three services. All of the services require 
constructor injection arguments. Some of these arguments are imported from the container 
parameters and some are defined directly. The logger service also requires two calls to the 
`pushHandler` method, each with a different handler service imported.

<?php // config/services.php

// Value objects are used to reference parameters and services in the container
use SitePoint\Container\Reference\ParameterReference as PR;
use SitePoint\Container\Reference\ServiceReference as SR;

use Monolog\Logger;
use Monolog\Handler\NativeMailerHandler;
use Monolog\Handler\StreamHandler;
use Psr\Log\LoggerInterface;

return [
    StreamHandler::class => [
        'class' => StreamHandler::class,
        'arguments' => [
            new PR('logger.file'),
            Logger::DEBUG,
        ],
    ],
    NativeMailHandler::class => [
        'class' => NativeMailerHandler::class,
        'arguments' => [
            new PR('logger.mail.to_address'),
            new PR('logger.mail.subject'),
            new PR('logger.mail.from_address'),
            Logger::ERROR,
        ],
    ],
    LoggerInterface::class => [
        'class' => Logger::class,
        'arguments' => [ 'channel-name' ],
        'calls' => [
            [
                'method' => 'pushHandler',
                'arguments' => [
                    new SR(StreamHandler::class),
                ]
            ],
            [
                'method' => 'pushHandler',
                'arguments' => [
                    new SR(NativeMailHandler::class),
                ]
            ]
        ]
    ]
];

The parameters definitions file just returns an array of values. These are defined as an 
N-dimensional array, but they are accessed through references using the notation: 
'logger.file' or 'logger.mail.to_address'.

<?php // config/parameters.php

return [
    'logger' => [
        'file' => __DIR__.'/../app.log',
        'mail' => [
            'to_address' => 'webmaster@domain.com',
            'from_address' => 'alerts@domain.com',
            'subject' => 'App Logs',
        ],
    ],
];

The container file just extracts the service and parameter definitions and passes them to 
the `Container` class constructor.

<?php // config/container.php

use SitePoint\Container\Container;

$services   = include __DIR__.'/services.php';
$parameters = include __DIR__.'/parameters.php';

return new Container($services, $parameters);

Now we can obtain the container in our app and use the logger service.

<?php // app/file.php

use Psr\Log\LoggerInterface;

require_once __DIR__.'/../vendor/autoload.php';

$container = include __DIR__.'/../config/container.php';

$logger = $container->get(LoggerInterface::class);
$logger->debug('This will be logged to the file');
$logger->error('This will be logged to the file and the email');
