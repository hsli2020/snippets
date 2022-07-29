// https://github.com/mmartinjoo/di-container

<?php  // app.php

use Container\Container\ContainerManual;
use Container\Container\ContainerAuto;
use Container\ExampleService;
use Container\File;
use Container\Logger;

include __DIR__ . '/vendor/autoload.php';

$container = new ContainerManual();
$container->set(File::class, fn (ContainerManual $c) => new File());
$container->set(Logger::class, fn (ContainerManual $c) => new Logger($c->get(File::class)));
$container->set(ExampleService::class, fn (ContainerManual $c) => new ExampleService($c->get(Logger::class)));
$service = $container->get(ExampleService::class);
$service->create();

$container = new ContainerAuto();
$service = $container->get(ExampleService::class);
$service->create();

?><?php  // File.php

namespace Container;

class File
{
    public function write(string $message)
    {
        echo "Writing '$message' to a file\r\n";
    }
}

?><?php  // Logger.php

namespace Container;

class Logger
{
    public function __construct(private readonly File $file)
    {
    }

    public function log(string $message): void
    {
        echo "Logging\r\n";
        $this->file->write($message);
    }
}

?><?php  // ExampleService.php

namespace Container;

class ExampleService
{
    public function __construct(private readonly Logger $logger)
    {
    }

    public function create()
    {
        $this->logger->log('Creating stuff...');
    }
}

?><?php  // ContainerManual.php

namespace Container\Container;

class ContainerManual
{
    private array $bindings;

    public function set(string $abstract, callable $factory): void
    {
        $this->bindings[$abstract] = $factory;
    }

    public function get(string $abstract): mixed
    {
        return $this->bindings[$abstract]($this);
    }
}

?><?php  // ContainerAuto.php

namespace Container\Container;

use ReflectionClass;
use ReflectionParameter;

class ContainerAuto
{
    private array $bindings;

    public function set(string $abstract, callable $factory): void
    {
        $this->bindings[$abstract] = $factory;
    }

    public function get(string $abstract)
    {
        if (isset($this->bindings[$abstract])) {
            return $this->bindings[$abstract]($this);
        }

        $constructor = (new ReflectionClass($abstract))
            ->getConstructor();

        if ($constructor === null) {
            return new $abstract;
        }

        $parameters = $constructor->getParameters();
        if (count($parameters) === 0) {
            return new $abstract;
        }

        $dependencies = array_map(
            fn (ReflectionParameter $parameter) => $parameter->getType()->getName(),
            $parameters
        );

        $resolvedDependencies = array_map(fn (string $dependency) =>
			$this->get($dependency), $dependencies
		);

        return new $abstract(...$resolvedDependencies);
    }
}
