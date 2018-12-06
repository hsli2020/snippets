<?php

namespace Phalcon\Di;

/**
 * Phalcon\Di\Injectable
 *
 * This class allows to access services in the services container by just only accessing
 * a public property with the same name of a registered service
 *
 * $dispatcher          \Phalcon\Mvc\Dispatcher
 *                      \Phalcon\Mvc\DispatcherInterface
 *
 * $router              \Phalcon\Mvc\Router
 *                      \Phalcon\Mvc\RouterInterface
 *
 * $url                 \Phalcon\Mvc\Url
 *                      \Phalcon\Mvc\UrlInterface
 *
 * $request             \Phalcon\Http\Request
 *                      \Phalcon\Http\RequestInterface
 *
 * $response            \Phalcon\Http\Response
 *                      \Phalcon\Http\ResponseInterface
 *
 * $cookies             \Phalcon\Http\Response\Cookies
 *                      \Phalcon\Http\Response\CookiesInterface
 *
 * $filter              \Phalcon\Filter
 *                      \Phalcon\FilterInterface
 *
 * $flash               \Phalcon\Flash\Direct
 *
 * $flashSession        \Phalcon\Flash\Session
 *
 * $session             \Phalcon\Session\Adapter\Files
 *                      \Phalcon\Session\Adapter
 *                      \Phalcon\Session\AdapterInterface
 *
 * $eventsManager       \Phalcon\Events\Manager
 *                      \Phalcon\Events\ManagerInterface
 *
 * $db                  \Phalcon\Db\AdapterInterface
 *
 * $security            \Phalcon\Security
 *
 * $crypt               \Phalcon\Crypt
 *                      \Phalcon\CryptInterface
 *
 * $tag                 \Phalcon\Tag
 *
 * $escaper             \Phalcon\Escaper
 *                      \Phalcon\EscaperInterface
 *
 * $annotations         \Phalcon\Annotations\Adapter\Memory
 *                      \Phalcon\Annotations\Adapter
 *
 * $modelsManager       \Phalcon\Mvc\Model\Manager
 *                      \Phalcon\Mvc\Model\ManagerInterface
 *
 * $modelsMetadata      \Phalcon\Mvc\Model\MetaData\Memory
 *                      \Phalcon\Mvc\Model\MetadataInterface
 *
 * $transactionManager  \Phalcon\Mvc\Model\Transaction\Manager
 *                      \Phalcon\Mvc\Model\Transaction\ManagerInterface
 *
 * $assets              \Phalcon\Assets\Manager
 *
 * $di                  \Phalcon\Di
 *                      \Phalcon\DiInterface
 *
 * $persistent          \Phalcon\Session\Bag
 *                      \Phalcon\Session\BagInterface
 *
 * $view                \Phalcon\Mvc\View
 *                      \Phalcon\Mvc\ViewInterface
 */
abstract class Injectable implements \Phalcon\Di\InjectionAwareInterface,
                                     \Phalcon\Events\EventsAwareInterface
{
    /** @var \Phalcon\DiInterface */
    protected $_dependencyInjector;

    /** @var \Phalcon\Events\ManagerInterface */
    protected $_eventsManager;

    public function setDI(\Phalcon\DiInterface $dependencyInjector);
    public function getDI();

    public function setEventsManager(\Phalcon\Events\ManagerInterface $eventsManager);
    public function getEventsManager();

    public function __get($propertyName);
}

class Phalcon\Cli\Task           extends \Phalcon\Di\Injectable
class Phalcon\Forms\Form         extends \Phalcon\Di\Injectable
class Phalcon\Mvc\Application    extends \Phalcon\Di\Injectable
class Phalcon\Mvc\Controller     extends \Phalcon\Di\Injectable
class Phalcon\Mvc\Micro          extends \Phalcon\Di\Injectable
class Phalcon\Mvc\User\Component extends \Phalcon\Di\Injectable
class Phalcon\Mvc\User\Module    extends \Phalcon\Di\Injectable
class Phalcon\Mvc\User\Plugin    extends \Phalcon\Di\Injectable
class Phalcon\Mvc\View\Engine    extends \Phalcon\Di\Injectable
class Phalcon\Mvc\View\Simple    extends \Phalcon\Di\Injectable
class Phalcon\Mvc\View           extends \Phalcon\Di\Injectable
class Phalcon\Validation         extends \Phalcon\Di\Injectable

$this->dispatcher->              $this->security->
$this->router->                  $this->crypt->
$this->url->                     $this->tag->
$this->request->                 $this->escaper->
$this->response->                $this->annotations->
$this->cookies->                 $this->modelsManager->
$this->filter->                  $this->modelsMetadata->
$this->flash->                   $this->transactionManager->
$this->flashSession->            $this->assets->
$this->session->                 $this->di->
$this->eventsManager->           $this->persistent->
$this->db->                      $this->view->
