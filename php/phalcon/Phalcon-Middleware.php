// Event.php
<?php  // https://github.com/SidRoberts/phalcon-authmiddleware/blob/4.0.x/src/Event.php

namespace Sid\Phalcon\AuthMiddleware;

use Phalcon\Mvc\DispatcherInterface;
use Phalcon\Di\Injectable;

class Event extends Injectable
{
    /**
     * @throws Exception
     */
    public function beforeExecuteRoute(\Phalcon\Events\Event $event, DispatcherInterface $dispatcher, $data) : bool
    {
        $methodAnnotations = $this->annotations->getMethod(
            $dispatcher->getHandlerClass(),
            $dispatcher->getActiveMethod()
        );

        if (!$methodAnnotations->has("AuthMiddleware")) {
            return true;
        }

        foreach ($methodAnnotations->getAll("AuthMiddleware") as $annotation) {
            $class = $annotation->getArgument(0);

            $authMiddleware = new $class();

            if (!($authMiddleware instanceof MiddlewareInterface)) {
                throw new Exception(
                    "Not an auth middleware."
                );
            }

            $result = $authMiddleware->authenticate();
            
            /*
             * Multi-middleware mode
             */
            if ($result === false) {
                return $result;
            }
        }

        return true;
    }
}

<?php // IndexController.php

namespace Tests;

use Phalcon\Mvc\Controller;

class IndexController extends Controller
{
    /**
     * @AuthMiddleware("Tests\Middleware1")
     */
    public function indexAction()
    {
        return "Hello world";
    }

    /**
     * @AuthMiddleware("Tests\Middleware2")
     */
    public function index2Action()
    {
        return "Hello world";
    }
    
    /**
     * @AuthMiddleware("Tests\Middleware1")
     * @AuthMiddleware("Tests\Middleware2")
     */
    public function index3Action()
    {
        return "Accepted all";
    }

    /**
     * @AuthMiddleware("Tests\Middleware2")
     * @AuthMiddleware("Tests\Middleware3")
     */
    public function index4Action()
    {
        return "Accepted all";
    }

    public function noMiddlewareAction()
    {
        return "Hello world";
    }

    /**
     * @AuthMiddleware("Tests\IndexController")
     */
    public function notProperMiddlewareAction()
    {
        return "This won't work as an Exception should get thrown";
    }
}

<?php // Middleware1.php

namespace Tests;

use Phalcon\Di\Injectable;
use Sid\Phalcon\AuthMiddleware\MiddlewareInterface;

class Middleware1 extends Injectable implements MiddlewareInterface
{
    public function authenticate() : bool
    {
        $this->dispatcher->setReturnedValue("Goodbye cruel world");
        return false;
    }
}

<?php // Middleware2.php

namespace Tests;

use Phalcon\Di\Injectable;
use Sid\Phalcon\AuthMiddleware\MiddlewareInterface;

class Middleware2 extends Injectable implements MiddlewareInterface
{
    public function authenticate() : bool
    {
        $this->dispatcher->setReturnedValue("Goodbye cruel world");
        return true;
    }
}

<?php // Middleware3.php

namespace Tests;

use Phalcon\Di\Injectable;
use Sid\Phalcon\AuthMiddleware\MiddlewareInterface;

class Middleware3 extends Injectable implements MiddlewareInterface
{
    public function authenticate() : bool
    {
        $this->dispatcher->setReturnedValue("Goodbye cruel world");
        return true;
    }
}

<?php // MiddlewareCest.php

namespace Tests;

use Phalcon\Di;
use Phalcon\Mvc\Dispatcher;

class MiddlewareCest
{
    public function _before()
    {
        Di::reset();

        $di = new \Phalcon\Di\FactoryDefault();

        $di->set("dispatcher", function () {
                $dispatcher = new Dispatcher();

                $eventsManager = new \Phalcon\Events\Manager();
                $eventsManager->attach("dispatch",
                    new \Sid\Phalcon\AuthMiddleware\Event()
                );

                $dispatcher->setEventsManager($eventsManager);
                $dispatcher->setDefaultNamespace("Tests\\");

                return $dispatcher;
            },
            true
        );

        $this->dispatcher = $di->get("dispatcher");
    }

    public function middlewareIsAbleToInterfereWhenReturningTrue(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("index");

        $dispatcher->dispatch();

        $I->assertEquals("Goodbye cruel world", $dispatcher->getReturnedValue());
    }

    public function middlewareDoesNotInterfereWhenReturningFalse(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("index2");

        $dispatcher->dispatch();

        $I->assertEquals("Hello world", $dispatcher->getReturnedValue());
    }

    public function dispatcherWorksAsNormalWithoutAnyMiddleware(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("noMiddleware");

        $dispatcher->dispatch();

        $I->assertEquals("Hello world", $dispatcher->getReturnedValue());
    }

    public function anExceptionIsThrownIfWePassSomethingThatIsntProperMiddleware(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("notProperMiddleware");

        $I->expectThrowable(
            \Sid\Phalcon\AuthMiddleware\Exception::class,
            function () use ($dispatcher) {
                $dispatcher->dispatch();
            }
        );
    }

    public function multiMiddlewareModeFirstCase(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("index3");

        $dispatcher->dispatch();

        $I->assertNotEquals("Accepted all", $dispatcher->getReturnedValue());
    }

    public function multiMiddlewareModeSecondCase(UnitTester $I)
    {
        $dispatcher = $this->dispatcher;

        $dispatcher->setControllerName("index");
        $dispatcher->setActionName("index4");

        $dispatcher->dispatch();

        $I->assertEquals("Accepted all", $dispatcher->getReturnedValue());
    }
}
