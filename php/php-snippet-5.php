<?php

YII 分库分表扩展（支持主从）

YII 分库分表扩展，
支持一主多从配置， 强制主库
支持自定义的分库分表方法

http://git.oschina.net/664712890/YII-sharded-ext

// 根据uid进行分库分表
$uid = 10;
$model = UserContact::model($uid); // 覆盖了原有参数，改为传入分库分表所依据的值
$data = $model->findAll();

$data = $model->dbConnection->createCommand()
 ->select("*")
 ->from($model->tableName())
 ->where('uid>1')
 ->limit(10)
 ->queryAll();

$db = Yii::app()->dbConnectionManager->sharded($uid); // 调用sharded方法 设置库和表
$data = $db->createCommand()
 ->select("*")
 ->from('user_contact_'.$db->shardedTableKey)
 ->where('uid>1')
 ->limit(10)
 ->queryAll();
# ===========================================================
Magento Orders Export

<?php
require_once("app/Mage.php");
Mage::app();
Mage::app()->getStore()->setId(Mage_Core_Model_App::ADMIN_STORE_ID);
$orders = Mage::getModel('sales/order')->getCollection();
$orders->setPage(1, 3000);
$fp = fopen('file.csv', 'w');
foreach($orders as $order) {
	$fields = array($order->getBillingAddress()->getFirstname(),$order->getBillingAddress()->getLastname(),$order->getBillingAddress()->getEmail());
	fputcsv($fp, $fields);
}
?>
http://www.hummingbirduk.com/export-customer-emails-to-csv-in-magento/
http://inchoo.net/ecommerce/magento/tracing-magento-from-export-csv-to-save-file-ok-button/

// update item in magento
public static function update($sku, $qty)
{
    // get product and update quantity
    $product = Mage::getModel('catalog/product')->loadByAttribute('sku', $sku);

    // if the product does not exist
    if ($product) {
        // update stock quantities
        $stock = Mage::getModel('cataloginventory/stock_item')->loadByProduct($product->getId());
        $stock->setData('qty', $qty);

        try {
            $stock->save();
            fwrite(STDOUT, "    :: UPDATED SKU $sku W/ QTY: $qty \n");
        } catch(Exception $e) {
            Mage::log($e->getMessage());
        }
    }

    return true;
}

    public static function orderCollection() {
        $orders = Mage::getModel('sales/order')->getCollection()
            ->addFieldToFilter('status', 'processing')
            // products that have not been submitted for packaging
            // ->addFieldToFilter('created_at', array(
            //     'from'     => strtotime('-1 day', time()),
            //     'to'       => time(),
            //     'datetime' => true
            // ))
        ;

        return $orders;
    }

    public static function orderData($order) {
        return $order->getData();
    }

    public static function setOrderToPending($order_id) {
        // get order
        $order = Mage::getModel('sales/order')->loadByIncrementId($order_id);

        // set status to submitted (processing: submitted)
        $order->setStatus('submitted');
        $order->save();
    }

    public static function orderShippingAddress($obj) {
        $address = Mage::getModel('sales/order_address')
            ->load($obj->getShippingAddressId())
            ->getData();

        return $address;
    }

    public static function orderItems($obj) {
        $items = $obj->getAllVisibleItems();
        $data = array();
        foreach ($items as $item) {
            $data[] = $item->getData();
        }

        return $data;
    }

    public static function orderShippingMethod($obj) {
        $shipping = array(
            'agent'   => $obj->getShippingMethod(),
            'service' => $obj->getShippingDescription()
        );

        return $shipping;
    }

    public static function update($order_id, $tracking_number) {

        // load order
        $order = Mage::getModel('sales/order')->loadByIncrementId($order_id);

        // create shipment
        try {

            // create a shipment
            $shipment = $order->prepareShipment();

            if ($shipment) {

                $shipment->register();
                $order->setIsInProcess(true);

                // generate tracking
                $tracking = Mage::getModel('sales/order_shipment_track')
                    ->setCarrierCode($order->getShippingCarrier()->getCarrierCode())
                    ->setTitle($order->getShippingDescription())
                    ->setNumber($tracking_number);

                // add tracking
                $shipment->addTrack($tracking);

                // save updated transaction
                $transactionSave = Mage::getModel('core/resource_transaction')
                    ->addObject($shipment)
                    ->addObject($shipment->getOrder())
                    ->save();

                // send confirmation email
                $shipment->sendEmail(true);
            }
        } catch (Exception $e) {
            Mage::log($e->getMessage());
        }

        // set state/status and save to complete the full life-cycle of the order
        $order->setData('state', Mage_Sales_Model_Order::STATE_COMPLETE);
        $order->setData('status', Mage_Sales_Model_Order::STATE_COMPLETE);
        $order->save();

        return true;
    }
# ===========================================================
Bing Image

	$str=file_get_contents('http://cn.bing.com/HPImageArchive.aspx?idx=0&n=1');
	if(preg_match("/<url>(.+?)<\/url>/ies",$str,$matches)){
		$imgurl='http://cn.bing.com'.$matches[1];
	}
	if($imgurl){
        echo $imgurl, "\n";
		// header('Content-Type: image/JPEG');
		// @ob_end_clean();
		// @readfile($imgurl);
		// @flush(); @ob_flush();
		// exit();
	} else {
		exit('error');
	}


# ===========================================================
PHP debug trace

function pre($var, $label='') // dpr pr
{
    echo "<pre>\n";
    if (!empty($label))
        echo "<b>$label = </b>";
    var_export($var); // print_r($var);
    echo "\n</pre>\n";
}

// trace(__FILE__);
// trace(__METHOD__);
// trace($var, 'return value is');
// trace($var, __METHOD__);
// trace($var, __METHOD__, __LINE__);

function trace($var, $func='', $ln=0) // dtr
{
    static $_first = true;

    $str = '';

    if ($_first) {
        $_first = false;
        $str = str_repeat('-', 80)."\n".date('Y-m-d H:i:s')."\n\n";
    }

    if (!empty($func) && !is_bool($func)) {
        $str .= $func;
        if (!empty($ln)) $str .= " $ln\n";
    }

    if (is_array($var) || is_object($var))
        $str .= print_r($var, true)."\n";
    else
        $str .= $var."\n";

    //$str = rtrim($str)."\n\n";
    if ($func === true) $str .= "\n";

    file_put_contents('/data/source/trace.log', $str, FILE_APPEND);
    #file_put_contents('/data/source/trace.log', $str); // overwrite
}

function dpr($var, $func='', $ln=0) // dtr
{
    static $_buffer = '';
    static $_first = true;

    if ($_first) {
        $_first = false;
        register_shutdown_function('dump', '#END#');
    }

    if ($var !== '#END#') {
        if (!empty($func)) {
            $_buffer .= $func;
            if (!empty($ln)) $_buffer .= " # ($ln)";
            $_buffer .= "\n";
        }

        if (is_array($var) || is_object($var))
            $_buffer .= print_r($var, true);
        else
            $_buffer .= $var;

        $_buffer = rtrim($_buffer)."\n\n";
    } 
    else {
        if (empty($_buffer)) return;

        $body = "TIMESTAMP: ".date('Y-m-d H:i:s')."\n".$_buffer."\n";
        file_put_contents('/data/source/trace.log', $body);
    }
}

function dump($var, $label='')
{
    echo "<pre>\n";
    if ($label) echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

function h2($s='')
{
  echo "<h2>$s</h2>\n";
}
# ===========================================================
PHPUnit

        $this->accountProvider->expects($this->once())
             ->method('find')
             ->with($pnum)
             ->will($this->returnValue($this->getEngagerMock()));

$this->once() 是说这个函数至少要被调用一次
$this->any() 是说这个函数可以被调用任意多次(或不被调用)
$this->at(1) 指定在这个函数被第二次调用时
# ===========================================================
http://ca3.php.net/manual/en/class.sphinxclient.php

$route['search/orders/(:any)'] = 'search/orders/$1';
$route['search/order/(:any)'] = 'search/orders/$1';

$route['orders/(:any)'] = 'search/orders/$1';
$route['order/(:any)'] = 'search/orders/$1';

$route['search/shipment/(:any)'] = 'search/shipment/$1';
$route['shipment/(:any)'] = 'search/shipment/$1';

<div class="top-fix-bar">
<div class="top-fix-inner">
<div class="top-fix-container">
<div class="top-fix-wrap">
</div>
</div>
</div>
</div>
# ===========================================================
Silex
$app->escape($string);
$app->error(function (\Exception $e, $code) { })
$app->view(function (array $controllerResult) { });
$app->redirect('/hello');
$app->handle($subRequest, HttpKernelInterface::SUB_REQUEST);
$app->json($error, 404);
$app->json($user);
$app->stream(...);
$app->abort(404);
$app->sendFile('/base/path/' . $path);

$app->before(function (Request $request, Application $app) { });
$app->after(function (Request $request, Response $response) { });
$app->finish(function (Request $request, Response $response) { });

Silex Core services
================

$app['request']
$app['routes']
$app['controllers']
$app['dispatcher']
$app['resolver']
$app['kernel']
$app['request_context']
$app['exception_handler']
$app['logger']

Silex Core parameters
===================

$app['request.http_port']
$app['request.https_port']
$app['locale']
$app['debug']
$app['charset']

Silex Traits
============

class MyApplication extends Silex\Application
{
    use Application\TwigTrait;
    use Application\SecurityTrait;
    use Application\FormTrait;
    use Application\UrlGeneratorTrait;
    use Application\SwiftmailerTrait;
    use Application\MonologTrait;
    use Application\TranslationTrait;
}

Silex provides two types of providers defined by two interfaces:
======================================================
  - ServiceProviderInterface for services
  - ControllerProviderInterface for controllers

interface Silex\ServiceProviderInterface
{
    public function register(Application $app);
    public function boot(Application $app);
}

interface Silex\ControllerProviderInterface
{
    public function connect(Application $app);
}

Silex\Provider\DoctrineServiceProvider
Silex\Provider\MonologServiceProvider
Silex\Provider\SessionServiceProvider
Silex\Provider\SerializerServiceProvider
Silex\Provider\SwiftmailerServiceProvider 
Silex\Provider\TwigServiceProvider
Silex\Provider\TranslationServiceProvider
Silex\Provider\UrlGeneratorServiceProvider
Silex\Provider\ValidatorServiceProvider
Silex\Provider\HttpCacheServiceProvider
Silex\Provider\FormServiceProvider
Silex\Provider\SecurityServiceProvider
Silex\Provider\RememberMeServiceProvider
Silex\Provider\ServiceControllerServiceProvider

Following Symfony components are used by Silex:
• HttpFoundation: For Request and Response.
• HttpKernel: Because we need a heart.
• Routing: For matching defined routes.
• EventDispatcher: For hooking into the HttpKernel.

silex定义route的三种方式
- closure (dependency injection)
- "namespace\\class::action"
- ControllerProviderInterface (Ashley is using this)

https://github.com/silexphp/Silex/wiki/Third-Party-ServiceProviders (downloaded)
https://github.com/silexphp/Silex/wiki/Third-Party-ServiceProviders-for-Silex-2.x

# ===========================================================
# By using ‘==’, php converts ‘1e1’ to ‘10’ because it treats ‘e1’ as ‘times 10’. etc
$ php -r "var_dump('1e1' === '10');"
bool(false)

$ php -r "var_dump('1e1' == '10');"
bool(true)

# another example, both sides are converted to value ‘0’ if using ‘==’.
$ php -r "var_dump('0e222222' === '0e4444444');"
bool(false)

$ php -r "var_dump('0e222222' == '0e4444444');"
bool(true)
# ===========================================================
how to write unittest for the following function

function addHashExtension($file) // addVersionIdToFilename($file)
{
    $info = pathinfo($file);

    $version = '';
    if (file_exists($file)) {
#       $version = substr(md5_file($file), 0, 8);
#       $version = hexdec(substr(md5_file($file), 0, 6));
        $version = substr(filectime($file), 2, 8);
        $info['filename'] .= ".$version";
    }

#   return "{$info['dirname']}/{$info['filename']}.{$version}.{$info['extension']}";
    return "{$info['dirname']}/{$info['filename']}.{$info['extension']}";
}

// $file = './tt.php';
// echo addHashExtension($file), EOL;
// 
// $file = '/Users/hansonli/Downloads/exif.jpg';
// echo addHashExtension($file), EOL;
// 
// $file = '/Users/hansonli/Downloads/xyz.exif.jpg';
// echo addHashExtension($file), EOL;
// 
// $file = '/tmp/nosuch.php';
// echo addHashExtension($file), EOL;
# ===========================================================
$ php -r "print_r(get_declared_classes());"
$ php -r "print_r(get_declared_interfaces());"
$ php -r "print_r(get_defined_functions());"

$ sudo pecl install redis

SessionUser/CurrentUser/TargetUser

git git@code.csdn.net:chaoticjoy/wechat-php-sdk.git
git https://github.com/netputer/wechat-php-sdk


$ which memcached
$ which pecl
/usr/local/bin/memcached
/usr/local/bin/pecl

$ ll /usr/local/bin/memcached
$ ll /usr/local/bin/pecl

$ brew search memcached
$ brew search redis
# ===========================================================
function getEmailDeliveryId() {
    $uid = uniqid('', TRUE);
    $arr = explode('.', $uid);
    return hexdec($arr[0].substr(dechex($arr[1]),-2));
}

Practically Guaranteed Unique 

$token = hash('sha1', uniqid(mt_rand(), true) . php_uname('n'));

$randstr = str_shuffle('abcdefghijklmnopqrstuvwxyz0123456789');
# ===========================================================
<?php
/**
 * @link http://www.yiiframework.com/
 * @copyright Copyright (c) 2008 Yii Software LLC
 * @license http://www.yiiframework.com/license/
 */

namespace yii\base;

/**
 * ArrayAccessTrait provides the implementation for [[\IteratorAggregate]], [[\ArrayAccess]] and [[\Countable]].
 *
 * Note that ArrayAccessTrait requires the class using it contain a property named `data` which should be an array.
 * The data will be exposed by ArrayAccessTrait to support accessing the class object like an array.
 *
 * @property array $data
 *
 * @author Qiang Xue <qiang.xue@gmail.com>
 * @since 2.0
 */
trait ArrayAccessTrait
{
    /**
     * Returns an iterator for traversing the data.
     * This method is required by the SPL interface `IteratorAggregate`.
     * It will be implicitly called when you use `foreach` to traverse the collection.
     * @return \ArrayIterator an iterator for traversing the cookies in the collection.
     */
    public function getIterator()
    {
        return new \ArrayIterator($this->data);
    }

    /**
     * Returns the number of data items.
     * This method is required by Countable interface.
     * @return integer number of data elements.
     */
    public function count()
    {
        return count($this->data);
    }

    /**
     * This method is required by the interface ArrayAccess.
     * @param mixed $offset the offset to check on
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->data[$offset]);
    }

    /**
     * This method is required by the interface ArrayAccess.
     * @param integer $offset the offset to retrieve element.
     * @return mixed the element at the offset, null if no element is found at the offset
     */
    public function offsetGet($offset)
    {
        return isset($this->data[$offset]) ? $this->data[$offset] : null;
    }

    /**
     * This method is required by the interface ArrayAccess.
     * @param integer $offset the offset to set element
     * @param mixed $item the element value
     */
    public function offsetSet($offset, $item)
    {
        $this->data[$offset] = $item;
    }

    /**
     * This method is required by the interface ArrayAccess.
     * @param mixed $offset the offset to unset element
     */
    public function offsetUnset($offset)
    {
        unset($this->data[$offset]);
    }
}
# ===========================================================
GetUrlVar

upload/catalog/view/javascript/crossdomain.php

<?php 
header('P3P: CP="CAO COR CURa ADMa DEVa OUR IND ONL COM DEM PRE"');

if (isset($_GET['session_id'])) {
	setcookie(session_name(), $_GET['session_id'], 0, getenv('HTTP_HOST'));
}
?>

function getURLVar(key) {
	var value = [];
	
	var query = String(document.location).split('?');
	
	if (query[1]) {
		var part = query[1].split('&');

		for (i = 0; i < part.length; i++) {
			var data = part[i].split('=');
			
			if (data[0] && data[1]) {
				value[data[0]] = data[1];
			}
		}
		
		if (value[key]) {
			return value[key];
		} else {
			return '';
		}
	}
} 
# ===========================================================
PHP Utils

<?php

define('EOL', PHP_EOL);

function isPost() // not good
{
    return ($_SERVER['REQUEST_METHOD'] == 'POST');
}

function isGet() // not good
{
    return ($_SERVER['REQUEST_METHOD'] == 'GET');
}

function requestMethod() // better
function getRequestMethod() // better
{
    return ($_SERVER['REQUEST_METHOD']);
}

// Check to see if it is an ajax request
function isAjax()
{
    return isset($_SERVER['HTTP_X_REQUESTED_WITH']) &&
        $_SERVER['HTTP_X_REQUESTED_WITH'] == 'XMLHttpRequest');
}

function println($string)
{
    echo $string, PHP_EOL;
}

function postRequest($url, $data)
{
    $options = array(
        'http' => array(
            'method'  => 'POST',
            'header'  => "Content-type: application/x-www-form-urlencoded",
            'content' => $data,
        )
    );
    $context = stream_context_create($options);
    $result = file_get_contents($url, NULL, $context);
    return $result;
}

function curlPost($url, $data)
{
    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);

    $result = curl_exec($ch);
    curl_close($ch);

    return $result;
}

function getPartNums()
{
    if (!isset($_POST['partnums']))
        return false;

    $partnums = preg_split("/[\s,]+/", $_POST['partnums']);
    if (empty($partnums))
        return false;

    return $partnums;
}

function getSku($partNum) // CleanPartNum
{
    $sections = explode('-', $partNum);

    if (strncasecmp($sections[0], 'FP', 2) == 0 || strncasecmp($sections[0], 'FL', 2) == 0)
        array_shift($sections);

    $supplier = $sections[0];

    if ($supplier == 'BTC')
        $sections[0] = 'BTE';
    elseif ($supplier != 'BTU')
        array_shift($sections);

    return array($supplier, implode('-', $sections));
}

function render($__file, $__vars) // not good
{
    extract($__vars);
    ob_start();
    include(TEMPLATE_DIR . $__file);
    $contents = ob_get_contents();
    ob_end_clean();
    return $contents;
}

function renderView($__file, $__data) // not good
{
    ob_start();
    extract($__data);
    include("views/$__file.tpl");
    $content = ob_get_contents();
    ob_end_clean();
    return $content;
}

function render($__file, $__data, $__layout = '') // not good (bad)
{
    $content = renderView($__file, $__data);
    if (empty($__layout))
        $__layout = 'layout';
//  extract($__data);
    include("views/$__layout.tpl");
}

function render($__file, $__data = array(), $__layout = '') // better
{
    ob_start();
    extract($__data);
    include("views/$__file.tpl");
    $__content = ob_get_contents();
    ob_end_clean();
    
    if ($__layout !== false) {
        if (empty($__layout)) {
            $__layout = 'layout';
        }
        include("views/$__layout.tpl");
    } else {
        // don't need layout
        echo $__content;
    }
}

public function h($str)
{
    return htmlentities($str, ENT_NOQUOTES, "UTF-8");
}
# ===========================================================
<?php
const EOL = PHP_EOL;

require 'Database.php';

$db = Db::connection(QA4DB);

# $query =  "SHOW DATABASES";
# $databases = $db->query($query)->fetchAll(PDO::FETCH_ASSOC);
# $databases = array_map('reset', $databases);
# print_r($databases);

foreach(array("am", "amadmin", "amarchive", "aminno", "ammem", "amsite", "google", "processor_pool", "research") as $database) {
    echo '-- Database: ', $database, EOL;

    $query =  "SHOW TABLES FROM $database";
    $tables = $db->query($query)->fetchAll(PDO::FETCH_ASSOC);
    $tables = array_map('reset', $tables);
    #print_r($tables);

    foreach($tables as $table) {
        $query =  "SHOW CREATE TABLE $database.$table";
        $info = $db->query($query)->fetchAll(PDO::FETCH_ASSOC);
       #echo '-- Table: ', $info[0]['Table'], EOL;
        echo $info[0]['Create Table'], EOL, EOL;
    }
}
# ===========================================================
UTC time and local time

echo date('Y-m-d H:i:s'), PHP_EOL; # local
echo gmdate('Y-m-d H:i:s'), PHP_EOL; # UTC

$offset = date("Z");

# local to UTC
$local_ts = time();
$utc_ts = $local_ts - $offset;

$utc_time = date("Y-m-d H:i:s", $utc_ts);
echo $utc_time;

###

date_default_timezone_set('America/New_York');
$utc_offset = date('Z');
echo date('Y-m-d H:i:s', time() - $utc_offset);

# UTC to local
echo date("Y-m-d H:i:s", strtotime(gmdate('Y-m-d H:i:s').' UTC')), PHP_EOL;

###

$today = gmmktime(0, 0, 0);
for ($i = 0; $i < 24*3600/300; $i++) {
echo date('Y-m-d H:i:s', $today + $i*300), PHP_EOL;
}
# ===========================================================
DealTap

class Enums {
	public static function getConst($name) {
		//$rc = new ReflectionClass(__CLASS__);
		$rc = new ReflectionClass(get_called_class());
		$consts = $rc->getConstants();

		if (!$name) {
			return $consts;
		}

		return isset($consts[$name]) ? $consts[$name] : null;
	}
}

class RoleTag extends Enums {
	const BUYERS = 'the one who want to buy something';
	public static function say() { echo __METHOD__, EOL; }
}

#$a = 'RoleTag';
#$a::say();

#$a='BUYERS';
#$refl = new ReflectionClass('RoleTag');
#$consts = $refl->getConstants();
#echo $consts[$a], EOL;

$a='BUYERS';
echo RoleTag::getConst($a), EOL;


const EOL = (PHP_SAPI == 'cli') ? PHP_EOL : '<br/>';

date_default_timezone_set('America/New_York');

function pr()
{
    $output = '';
    foreach (func_get_args() as $var) {
        $output .= var_export($var, true) . PHP_EOL;
    }

    $patterns = [
        "/array \(/"    => "[",
        "/\),\n/"       => "],\n",
        "/\)\n/"        => "]\n",
        "/^\)/m"        => "]",
        "/=> \n(\s+)/"  => "=> ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

function vd()
{
    ob_start();
    $var = func_get_args(); 
    call_user_func_array('var_dump', $var);
    $output = ob_get_contents();
    ob_end_clean();

    $patterns = [
        "/\"\]/"      => "\"",
        "/\[\"/"      => "\"",
        "/=>\n(\s+)/" => " => ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

// function pr($var) { var_export($var); }
# ===========================================================
DealTap

trait ArgsToProps
{
    protected function args_to_props($args) {
    }
}

class NestedEntryItem {
	public $index;

	public function __construct($args = []) {
		$properties = array_keys(get_object_vars($this));
		foreach ($properties as $name) {
			if (isset($args[$name])) {
				$this->$name = $args[$name];
			}
		}
		$this->index = $this->index ? $this->index : uniqid(); 
	}

	public function  to_db() {
		$db = new stdClass();

		$properties = array_keys(get_object_vars($this));
		foreach ($properties as $name) {
			$db->$name = $this->$name;
		}

		return $db;
	}

	public function to_json() {
		$json = new stdClass();

		$properties = array_keys(get_object_vars($this));
		foreach ($properties as $name) {
			$json->$name = $this->$name;
		}

		return $json;
	}
}

class Contact extends NestedEntryItem {
	public $name;
	public $email;
	public $phone;
}

$contact = new Contact([
	'name' => 'someone', 
	'email' => 'some@email.com', 
	'pass' => 'secret', // not in class properties
	'phone' => '800-900-1011',
]);
#pr($contact);
#pr($contact->to_db());
#pr($contact->to_json());


const EOL = (PHP_SAPI == 'cli') ? PHP_EOL : '<br/>';

date_default_timezone_set('America/New_York');

function pr()
{
    $output = '';
    foreach (func_get_args() as $var) {
        $output .= var_export($var, true) . PHP_EOL;
    }

    $patterns = [
        "/array \(/"    => "[",
        "/\),\n/"       => "],\n",
        "/\)\n/"        => "]\n",
        "/^\)/m"        => "]",
        "/=> \n(\s+)/"  => "=> ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

function vd()
{
    ob_start();
    $var = func_get_args(); 
    call_user_func_array('var_dump', $var);
    $output = ob_get_contents();
    ob_end_clean();

    $patterns = [
        "/\"\]/"      => "\"",
        "/\[\"/"      => "\"",
        "/=>\n(\s+)/" => " => ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

// function pr($var) { var_export($var); }
# ===========================================================
