var fs = require('fs');

var gitignore = "*\n!.gitignore";
var emptyVolt = '{% extends "layouts/base.volt" %}\n\n{% block main %}\n{% endblock %}';

fs.writeFileSync('.htaccess',
`Options +FollowSymLinks -MultiViews -Indexes

DirectoryIndex index.html index.php

<IfModule mod_rewrite.c>
    RewriteEngine on
    RewriteRule  ^$ public/    [L]
    RewriteRule  (.*) public/$1 [L]
</IfModule>`);

fs.writeFileSync('.gitignore',
`Thumbs.db
Desktop.ini
.DS_Store
composer.lock
app/cache/
app/logs/
vendor/
*.tmp
zzz.log
filelist`);

fs.writeFileSync('composer.json',
`{
    "require": {
        "php": ">=5.4",
        "ext-phalcon": ">=2.0.9",
        "phalcon/devtools": "dev-master"
    },
    "autoload": {
        "psr-4": {
            "Supplier\\\\": "src/Supplier/",
            "Utility\\\\":  "src/Utility/" 
        }
    }
}`);

fs.mkdirSync('app');
fs.mkdirSync('src');

fs.mkdirSync('app/controllers');

fs.writeFileSync('app/controllers/ControllerBase.php',
`<?php
namespace App\\Controllers;

use Phalcon\\Mvc\\Controller;
use Phalcon\\Mvc\\Dispatcher;

/**
 * ControllerBase
 * This is the base controller for all controllers in the application
 */
class ControllerBase extends Controller
{
    /**
     * Execute before the router so we can determine if this is a private controller, 
     * and must be authenticated, or a public controller that is open to all.
     *
     * @param Dispatcher $dispatcher
     * @return boolean
     */
    public function beforeExecuteRoute(Dispatcher $dispatcher)
    {
        $controllerName = $dispatcher->getControllerName();

        // Only check permissions on private controllers
        if ($this->acl->isPrivate($controllerName)) {

            // Get the current identity
            $identity = $this->auth->getIdentity();

            // If there is no identity available the user is redirected to index/index
            if (!is_array($identity)) {

                $this->flash->notice("You don't have access to this module: private");

                $dispatcher->forward(array(
                    'controller' => 'index',
                    'action' => 'index'
                ));
                return false;
            }

            // Check if the user have permission to the current option
            $actionName = $dispatcher->getActionName();
            if (!$this->acl->isAllowed($identity['profile'], $controllerName, $actionName)) {

                $this->flash->notice("You don't have access to this module: $controllerName:$actionName");

                if ($this->acl->isAllowed($identity['profile'], $controllerName, 'index')) {
                    $dispatcher->forward(array(
                        'controller' => $controllerName,
                        'action' => 'index'
                    ));
                } else {
                    $dispatcher->forward(array(
                        'controller' => 'user_control',
                        'action' => 'index'
                    ));
                }

                return false;
            }
        }
    }
}`);

fs.writeFileSync('app/controllers/IndexController.php',
`<?php
namespace App\\Controllers;

/**
 * Display the default index page.
 */
class IndexController extends ControllerBase
{
    /**
     * Default action. Set the public layout (layouts/public.volt)
     */
    public function indexAction()
    {
        $this->view->setVar('logged_in', is_array($this->auth->getIdentity()));
    }
}`);

fs.writeFileSync('app/controllers/TestController.php',
`<?php
namespace App\\Controllers;

class TestController extends ControllerBase
{
    public function indexAction()
    {
    }
}`);

fs.writeFileSync('app/controllers/UserController.php',
`<?php
namespace App\\Controllers;

class UserController extends ControllerBase
{
    public function indexAction()
    {
    }
}`);

fs.writeFileSync('app/controllers/AboutController.php',
`<?php
namespace App\\Controllers;

class AboutController extends ControllerBase
{
    public function indexAction()
    {
    }
}`);

fs.mkdirSync('app/models');
fs.writeFileSync('app/models/Orders.php',
`<?php
namespace App\\Models;

use Phalcon\\Mvc\\Model;

/**
 * App\\Models\\Orders
 */
class Orders extends Model
{
    public $id;
    public $channel;
    public $date;
    public $orderId;
    public $mgnOrderId;
    public $express;
    public $buyer;
    public $address;
    public $city;
    public $province;
    public $postalcode;
    public $country;
    public $phone;
    public $email;
    public $sku;
    public $price;
    public $qty;
    public $shipping;
    public $mgnInvoiceId;

    public function initialize()
    {
        $this->setSource('all_mgn_orders');
    }

    public function columnMap()
    {
        // Keys are the real names in the table and
        // the values their names in the application

        return array(
            'id'             => 'id',
            'channel'        => 'channel',
            'date'           => 'date',
            'order_id'       => 'orderId',
            'mgn_order_id'   => 'mgnOrderId',
            'express'        => 'express',
            'buyer'          => 'buyer',
            'address'        => 'address',
            'city'           => 'city',
            'province'       => 'province',
            'postalcode'     => 'postalcode',
            'country'        => 'country',
            'phone'          => 'phone',
            'email'          => 'email',
            'skus_sold'      => 'sku',
            'sku_price'      => 'price',
            'skus_qty'       => 'qty',
            'shipping'       => 'shipping',
            'mgn_invoice_id' => 'mgnInvoiceId'
        );
    }
}`);

fs.mkdirSync('app/views');
fs.writeFileSync('app/views/index.volt', '{{ content() }}');

fs.mkdirSync('app/views/layouts');
fs.writeFileSync('app/views/layouts/base.volt',
`<!DOCTYPE html>
<html lang="en-us">
<head>
  <meta charset="utf-8">
  <!--<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">-->
  <title>{% block title %}BTE Intranet{% if pageTitle is defined %} &bull; {{ pageTitle }}{% endif %}{% endblock %}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">

  {% block cssfile %}
    {{ stylesheet_link('/lib/bootstrap/3.3.7/css/bootstrap.min.css') }}
    {{ stylesheet_link('/assets/css/style.css') }}
  {% endblock %}

  <style type="text/css">
    {% block csscode %}{% endblock %}
  </style>

  <!-- FAVICONS -->
  <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">
  <link rel="icon" href="/favicon.ico" type="image/x-icon">

  <!-- GOOGLE FONT -->
  <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Open+Sans:400italic,700italic,300,400,700">
</head>
<body>
  {% include "partials/navigation.volt" %}
  {% block sidebar %}{% endblock %}

  <div class="container main-container">
    <?php $this->flashSession->output(); ?>
    {% block main %}{% endblock %}
  </div>

  {# loading icon placholder #}
  <div id="loading" style="display:none;"></div>
  <div id='toast' style='display:none'></div>

  {% block jsfile %}
    {{ javascript_include('/lib/jquery/jquery-3.1.0.min.js') }}
    {{ javascript_include('/lib/bootstrap/3.3.7/js/bootstrap.min.js') }}
    {{ javascript_include('/lib/layer/layer.js') }}
    {{ javascript_include('/assets/js/script.js') }}
  {% endblock %}

  <script type="text/javascript">
    {% block jscode %}{% endblock %}
    $(document).ready(function() {
      {% block docready %}{% endblock %}
    });
  </script>
</body>
</html>`);

fs.writeFileSync('app/views/layouts/public.volt', emptyVolt);
fs.writeFileSync('app/views/layouts/private.volt', emptyVolt);

fs.mkdirSync('app/views/partials');
fs.writeFileSync('app/views/partials/navigation.volt', '');
fs.writeFileSync('app/views/partials/sidebar.volt', '');
fs.writeFileSync('app/views/partials/header.volt', '');
fs.writeFileSync('app/views/partials/footer.volt', '');
fs.mkdirSync('app/views/index');
fs.writeFileSync('app/views/index/index.volt', emptyVolt);
fs.mkdirSync('app/views/test');
fs.writeFileSync('app/views/test/index.volt', emptyVolt);
fs.mkdirSync('app/views/about');
fs.writeFileSync('app/views/about/index.volt', emptyVolt);

fs.mkdirSync('app/config');
fs.writeFileSync('app/config/config.php',
`<?php

use Phalcon\\Config;
use Phalcon\\Logger;

return new Config([
    'database' => [
        'adapter' => 'Mysql',
        'host' => '127.0.0.1',
        'username' => 'root',
        'password' => '',
        'dbname' => 'bte'
    ],
    'application' => [
        'controllersDir' => APP_DIR . '/controllers/',
        'modelsDir' => APP_DIR . '/models/',
        'formsDir' => APP_DIR . '/forms/',
        'viewsDir' => APP_DIR . '/views/',
        'libraryDir' => APP_DIR . '/library/',
        'pluginsDir' => APP_DIR . '/plugins/',
        'cacheDir' => APP_DIR . '/cache/',
        'baseUri' => '/',
        'publicUrl' => 'http://www.sitename.com',
        'cryptSalt' => 'eEAfR|_&G&f,+vU]:jFr!!A&+71w1Ms9~8_4L!<@[N@DyaIP_2My|:+.u>/6m,$D'
    ],
    'logger' => [
        'path'     => APP_DIR . '/logs/',
        'format'   => '%date% [%type%] %message%',
        'date'     => 'Y-m-d H:i:s',
        'logLevel' => Logger::DEBUG,
        'filename' => 'app.log',
    ],
]);`);

fs.writeFileSync('app/config/loader.php',
`<?php

use Phalcon\\Loader;

$loader = new Loader();

/**
 * We're a registering a set of directories taken from the configuration file
 */
$loader->registerNamespaces([
    'App\\Models'      => $config->application->modelsDir,
    'App\\Controllers' => $config->application->controllersDir,
    'App\\Forms'       => $config->application->formsDir,
    'App'             => $config->application->libraryDir
]);

$loader->register();

// Use composer autoloader to load vendor classes
require_once __DIR__ . '/../../vendor/autoload.php';`);

fs.writeFileSync('app/config/routes.php',
`<?php
/*
 * Define custom routes. File gets included in the router service definition.
 */
$router = new Phalcon\\Mvc\\Router();

$router->add('/ajax/order/detail', [
    'controller' => 'ajax',
    'action' => 'orderDetail'
])->via(array("POST"));

$router->add('/ajax/make/purchase', [
    'controller' => 'ajax',
    'action' => 'makePurchase'
])->via(array("POST"));

$router->add('/ajax/price/avail', [
    'controller' => 'ajax',
    'action' => 'priceAvail'
])->via(array("POST"));

// aliases

$router->add('/dropship', [
    'controller' => 'purchase',
    'action' => 'index'
]);

return $router;`);

fs.writeFileSync('app/config/services.php',
`<?php

use Phalcon\\DI\\FactoryDefault;
use Phalcon\\Mvc\\View;
use Phalcon\\Crypt;
use Phalcon\\Mvc\\Dispatcher;
use Phalcon\\Mvc\\Url as UrlResolver;
use Phalcon\\Db\\Adapter\\Pdo\\Mysql as DbAdapter;
use Phalcon\\Mvc\\View\\Engine\\Volt as VoltEngine;
use Phalcon\\Mvc\\Model\\Metadata\\Files as MetaDataAdapter;
use Phalcon\\Session\\Adapter\\Files as SessionAdapter;
use Phalcon\\Flash\\Direct as Flash;
use Phalcon\\Logger\\Adapter\\File as FileLogger;
use Phalcon\\Logger\\Formatter\\Line as FormatterLine;
use Phalcon\\Logger;
use Phalcon\\Events\\Manager as EventsManager;

use App\\Auth\\Auth;
use App\\Acl\\Acl;
use App\\Mail\\Mail;

/**
 * The FactoryDefault Dependency Injector automatically register the right services
 * providing a full stack framework
 */
$di = new FactoryDefault();

/**
 * Register the global configuration as config
 */
$di->set('config', $config);

/**
 * The URL component is used to generate all kind of urls in the application
 */
$di->set('url', function () use ($config) {
    $url = new UrlResolver();
    $url->setBaseUri($config->application->baseUri);
    return $url;
}, true);

/**
 * Setting up the view component
 */
$di->set('view', function () use ($config) {

    $view = new View();

    $view->setViewsDir($config->application->viewsDir);

    $view->registerEngines(array(
        '.volt' => function ($view, $di) use ($config) {

            $volt = new VoltEngine($view, $di);

            $volt->setOptions(array(
                'compiledPath' => $config->application->cacheDir . 'volt/',
                'compiledSeparator' => '_',
                'compiledPath' => function($templatePath) use ($config) {
                    return $config->application->cacheDir . 'volt/' . md5($templatePath) . '.php';
                },
            ));

            return $volt;
        }
    ));

    return $view;
}, true);

/**
 * Database connection is created based in the parameters defined in the configuration file
 */
$di->set('db', function () use ($config) {
    $eventsManager = new EventsManager();

    $logger = new FileLogger(APP_DIR . "/logs/db.log");

    // Listen all the database events
    $eventsManager->attach('db', function ($event, $connection) use ($logger) {
        if ($event->getType() == 'beforeQuery') {
            $logger->log($connection->getSQLStatement(), Logger::INFO);
        }
    });

    $connection = new DbAdapter(array(
        'host' => $config->database->host,
        'username' => $config->database->username,
        'password' => $config->database->password,
        'dbname' => $config->database->dbname
    ));

    // Assign the eventsManager to the db adapter instance
    $connection->setEventsManager($eventsManager);

    return $connection;
});

/**
 * If the configuration specify the use of metadata adapter use it or use memory otherwise
 */
$di->set('modelsMetadata', function () use ($config) {
    return new MetaDataAdapter(array(
        'metaDataDir' => $config->application->cacheDir . 'metaData/'
    ));
});

/**
 * Start the session the first time some component request the session service
 */
$di->set('session', function () {
    $session = new SessionAdapter();
    $session->start();
    return $session;
});

/**
 * Crypt service
 */
$di->set('crypt', function () use ($config) {
    $crypt = new Crypt();
    $crypt->setKey($config->application->cryptSalt);
    return $crypt;
});

/**
 * Dispatcher use a default namespace
 */
$di->set('dispatcher', function () {
    $dispatcher = new Dispatcher();
    $dispatcher->setDefaultNamespace('App\\Controllers');
    return $dispatcher;
});

/**
 * Loading routes from the routes.php file
 */
$di->set('router', function () {
    return require __DIR__ . '/routes.php';
});

/**
 * Flash service with custom CSS classes
 */
$di->set('flash', function () {
    return new Flash(array(
        'error' => 'alert alert-danger',
        'success' => 'alert alert-success',
        'notice' => 'alert alert-info',
        'warning' => 'alert alert-warning'
    ));
});

/**
 * Custom authentication component
 */
$di->set('auth', function () {
    return new Auth();
});

/**
 * Mail service uses AmazonSES
 */
$di->set('mail', function () {
    return new Mail();
});

/**
 * Access Control List
 */
$di->set('acl', function () {
    return new Acl();
});

/**
 * Logger service
 */
$di->set('logger', function ($filename = null, $format = null) use ($config) {
    $format   = $format ?: $config->get('logger')->format;
    $filename = trim($filename ?: $config->get('logger')->filename, '\\\\/');
    $path     = rtrim($config->get('logger')->path, '\\\\/') . DIRECTORY_SEPARATOR;

    $formatter = new FormatterLine($format, $config->get('logger')->date);
    $logger    = new FileLogger($path . $filename);

    $logger->setFormatter($formatter);
    $logger->setLogLevel($config->get('logger')->logLevel);

    return $logger;
});`);

fs.mkdirSync('app/forms');
fs.writeFileSync('app/forms/LoginForm.php',
`<?php
namespace App\\Forms;

use Phalcon\\Forms\\Form;
use Phalcon\\Forms\\Element\\Text;
use Phalcon\\Forms\\Element\\Password;
use Phalcon\\Forms\\Element\\Submit;
use Phalcon\\Forms\\Element\\Check;
use Phalcon\\Forms\\Element\\Hidden;
use Phalcon\\Validation\\Validator\\PresenceOf;
use Phalcon\\Validation\\Validator\\Email;
use Phalcon\\Validation\\Validator\\Identical;

class LoginForm extends Form
{
    public function initialize()
    {
        // Email
        $email = new Text('email', array(
            'placeholder' => 'Email'
        ));

        $email->addValidators(array(
            new PresenceOf(array(
                'message' => 'The e-mail is required'
            )),
            new Email(array(
                'message' => 'The e-mail is not valid'
            ))
        ));

        $this->add($email);

        // Password
        $password = new Password('password', array(
            'placeholder' => 'Password'
        ));

        $password->addValidator(new PresenceOf(array(
            'message' => 'The password is required'
        )));

        $password->clear();

        $this->add($password);

        // Remember
        $remember = new Check('remember', array(
            'value' => 'yes'
        ));

        $remember->setLabel('Remember me');

        $this->add($remember);

        // CSRF
        $csrf = new Hidden('csrf');

        $csrf->addValidator(new Identical(array(
            'value' => $this->security->getSessionToken(),
            'message' => 'CSRF validation failed'
        )));

        $csrf->clear();

        $this->add($csrf);

        $this->add(new Submit('go', array(
            'class' => 'btn btn-success'
        )));
    }
}`);

//fs.writeFileSync('app/forms/SignUpForm.php', emptyForm);
//fs.writeFileSync('app/forms/UsersForm.php', emptyForm);

fs.mkdirSync('app/library');
fs.mkdirSync('app/library/Acl');
fs.writeFileSync('app/library/Acl/Acl.php', '');
fs.mkdirSync('app/library/Auth');
fs.writeFileSync('app/library/Auth/Auth.php', '');
fs.writeFileSync('app/library/Auth/Exception.php', '');
fs.mkdirSync('app/library/Mail');
fs.writeFileSync('app/library/Mail/Mail.php', '');
fs.writeFileSync('app/library/Mail/Exception.php', '');

fs.mkdirSync('app/logs');
fs.mkdirSync('app/logs/acl');
fs.writeFileSync('app/logs/acl/.gitignore', gitignore);

fs.mkdirSync('app/cache');
fs.mkdirSync('app/cache/acl');
fs.writeFileSync('app/cache/acl/.gitignore', gitignore);
fs.mkdirSync('app/cache/metadata');
fs.writeFileSync('app/cache/metadata/.gitignore', gitignore);
fs.mkdirSync('app/cache/swift');
fs.writeFileSync('app/cache/swift/.gitignore', gitignore);
fs.mkdirSync('app/cache/volt');
fs.writeFileSync('app/cache/volt/.gitignore', gitignore);

fs.mkdirSync('public');

fs.writeFileSync('public/.htaccess',
`AddDefaultCharset UTF-8

<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteRule ^(.*)$ index.php?_url=/$1 [QSA,L]
</IfModule>`);

fs.writeFileSync('public/index.php',
`<?php

error_reporting(E_ALL);

try {

    /**
     * Define some useful constants
     */
    define('BASE_DIR', dirname(__DIR__));
    define('APP_DIR', BASE_DIR . '/app');

    #include 'trace.php';

    /**
     * Read the configuration
     */
    $config = include APP_DIR . '/config/config.php';

    /**
     * Read auto-loader
     */
    include APP_DIR . '/config/loader.php';

    /**
     * Read services
     */
    include APP_DIR . '/config/services.php';

    /**
     * Handle the request
     */
    $application = new \\Phalcon\\Mvc\\Application($di);

    echo $application->handle()->getContent();

} catch (Exception $e) {
    echo $e->getMessage(), '<br>';
    echo nl2br(htmlentities($e->getTraceAsString()));
}`);
fs.mkdirSync('public/lib');
fs.mkdirSync('public/lib/bootstrap');
fs.mkdirSync('public/lib/bootstrap/css');
fs.mkdirSync('public/lib/bootstrap/fonts');
fs.mkdirSync('public/lib/bootstrap/js');
fs.mkdirSync('public/lib/jquery');
fs.mkdirSync('public/lib/jquery/plugins');
fs.mkdirSync('public/assets');
fs.mkdirSync('public/assets/css');
fs.mkdirSync('public/assets/img');
fs.mkdirSync('public/assets/js');
