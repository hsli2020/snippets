<?php	# from keep.google.com

function isMobile()
{
    // 如果有HTTP_X_WAP_PROFILE则一定是移动设备
    if (isset ($_SERVER['HTTP_X_WAP_PROFILE'])) {
        return true;
    }

    // 如果via信息含有wap则一定是移动设备,部分服务商会屏蔽该信息
    if (isset ($_SERVER['HTTP_VIA'])) {
        // 找不到为flase,否则为true
        return stristr($_SERVER['HTTP_VIA'], "wap") ? true : false;
    }

    // 脑残法，判断手机发送的客户端标志,兼容性有待提高
    if (isset ($_SERVER['HTTP_USER_AGENT'])) {
        $clientkeywords = array(
            'nokia',
            'sony',
            'ericsson',
            'mot',
            'samsung',
            'htc',
            'sgh',
            'lg',
            'sharp',
            'sie-',
            'philips',
            'panasonic',
            'alcatel',
            'lenovo',
            'iphone',
            'ipod',
            'blackberry',
            'meizu',
            'android',
            'netfront',
            'symbian',
            'ucweb',
            'windowsce',
            'palm',
            'operamini',
            'operamobi',
            'openwave',
            'nexusone',
            'cldc',
            'midp',
            'wap',
            'mobile'
        );
        // 从HTTP_USER_AGENT中查找手机浏览器的关键字
        if (preg_match("/(" . implode('|', $clientkeywords) . ")/i", strtolower($_SERVER['HTTP_USER_AGENT']))) {
            return true;
        }
    }

    // 协议法，因为有可能不准确，放到最后判断
    if (isset ($_SERVER['HTTP_ACCEPT'])) {
        // 如果只支持wml并且不支持html那一定是移动设备
        // 如果支持wml和html但是wml在html之前则是移动设备
        if ((strpos($_SERVER['HTTP_ACCEPT'], 'vnd.wap.wml') !== false) && 
            (strpos($_SERVER['HTTP_ACCEPT'], 'text/html') === false || 
             (strpos($_SERVER['HTTP_ACCEPT'], 'vnd.wap.wml') < strpos($_SERVER['HTTP_ACCEPT'], 'text/html')))) {
            return true;
        }
    }
    return false;
}

//判断是否属手机
function is_mobile() {
    $user_agent = $_SERVER['HTTP_USER_AGENT'];
    $mobile_agents = array(
        "240x320","acer","acoon","acs-","abacho","ahong","airness","alcatel","amoi","android",
        "anywhereyougo.com","applewebkit/525","applewebkit/532","asus","audio","au-mic",
        "avantogo","becker","benq","bilbo","bird","blackberry","blazer","bleu","cdm-","compal",
        "coolpad","danger","dbtel","dopod","elaine","eric","etouch","fly ","fly_","fly-",
        "go.web","goodaccess","gradiente","grundig","haier","hedy","hitachi","htc","huawei",
        "hutchison","inno","ipad","ipaq","ipod","jbrowser","kddi","kgt","kwc","lenovo","lg ",
        "lg2","lg3","lg4","lg5","lg7","lg8","lg9","lg-","lge-","lge9","longcos","maemo",
        "mercator","meridian","micromax","midp","mini","mitsu","mmm","mmp","mobi","mot-",
        "moto","nec-","netfront","newgen","nexian","nf-browser","nintendo","nitro","nokia",
        "nook","novarra","obigo","palm","panasonic","pantech","philips","phone","pg-",
        "playstation","pocket","pt-","qc-","qtek","rover","sagem","sama","samu","sanyo",
        "samsung","sch-","scooter","sec-","sendo","sgh-","sharp","siemens","sie-","softbank",
        "sony","spice","sprint","spv","symbian","tablet","talkabout","tcl-","teleca","telit",
        "tianyu","tim-","toshiba","tsm","up.browser","utec","utstar","verykool","virgin",
        "vk-","voda","voxtel","vx","wap","wellco","wig browser","wii","windows ce","wireless",
        "xda","xde","zte");
    $is_mobile = false;
    foreach ($mobile_agents as $device) {
        if (stristr($user_agent, $device)) {
            $is_mobile = true;
            break;
        }
    }
    return $is_mobile;
}
//使用很简单
if( is_mobile() ){
  Your Code
}

# ===========================================================

这是一个类的构造函数

    public function __construct(
        Connection $dbWriteConnection,
        Connection $dbReadConnection,
        PimgProvider $pimgProvider,
        QueueProvider $queueProvider,
        MemcacheProvider $memcacheProvider,
        ImageConverter $imageConverter,
        ResqueJobProvider $resque
    ) {
        $this->dbWriteConnection = $dbWriteConnection;
        $this->dbReadConnection  = $dbReadConnection;
        $this->pimgProvider      = $pimgProvider;
        $this->queueProvider     = $queueProvider;
        $this->memcacheProvider  = $memcacheProvider;
        $this->imageConverter    = $imageConverter;
        $this->photoDBUtil       = new PhotoDBUtil($dbWriteConnection, $dbReadConnection);
        $this->resque            = $resque;
    }

每当看到这样的函数，我总有一个疑问：为什么要传递这样多的参数呢？这样做调用起来岂不是非常麻烦？
为什么不像这里的photoDBUtil那样，在构造函数内部创建对象呢？

稍微思考了一下，有些明白了。唯一合理的解释是，这里传入的每一个对象，创建起来都可能是非常复杂的，
需要更多其它信息的支持，比如数据库连接，要从配置文件中调入服务器参数，可能还要判断是处于'开发/调试/生产'
环境，分别调入不同的配置文件或参数。如果把这个创建过程都放在这个消费者的构造函数中，可以想见，
势必导致消费者的构造函数非常复杂，这显然是不合适的。

更重要的是，系统中还有许多这样的消费者类，它们都需要数据库连接这样的对象，如果都在消费者类的构造
函数中创建对象，系统中会存在大量重复的代码，这绝对应该避免。

至于那个PhotoDBUtil类，它本来就是专门为这个类服务的，所以在这里创建它是合适的。

# 这就是所谓的“Dependency Injection”模式

# ===========================================================

class Foo {

	var $stdin;
	var $stdout;
	var $stderr;

  public function __construct() {
     $this->stdin = fopen('php://stdin', 'r');
     $this->stdout = fopen('php://stdout', 'w');
     $this->stderr = fopen('php://stderr', 'w');
  }

	function stdout($string, $newline = true) {
		if ($newline) {
			return fwrite($this->stdout, $string . "\n");
		} else {
			return fwrite($this->stdout, $string);
		}
	}

	function stderr($string) {
		fwrite($this->stderr, $string);
	}

  public function show() {
    $this->stdout("Hello");
  }
}

$foo = new Foo();
$foo->show();

# ===========================================================

	function nl($multiplier = 1) {
		return str_repeat("\n", $multiplier);
	}

	function shortPath($file) {
		$shortPath = str_replace(ROOT, null, $file);
		$shortPath = str_replace('..' . DS, '', $shortPath);
		return str_replace(DS . DS, DS, $shortPath);
	}

# ===========================================================

<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteRule ^(.*)$ index.php?url=$1 [QSA,L]
</IfModule>


<IfModule mod_rewrite.c>
    RewriteEngine on
    RewriteRule    ^$    webroot/    [L]
    RewriteRule    (.*) webroot/$1    [L]
 </IfModule>

# ===========================================================

class Slug
{
    /**
     * Creates a slug to be used for pretty URLs
     *
     * @link http://cubiq.org/the-perfect-php-clean-url-generator
     * @param         $string
     * @param  array  $replace
     * @param  string $delimiter
     * @return mixed
     */
    public static function generate($string, $replace = array(), $delimiter = '-')
    {
        if (!extension_loaded('iconv')) {
            throw new \Phalcon\Exception('iconv module not loaded');
        }

        // Save the old locale and set the new locale to UTF-8
        $oldLocale = setlocale(LC_ALL, 'en_US.UTF-8');

        $clean = iconv('UTF-8', 'ASCII//TRANSLIT', $string);

        if (!empty($replace)) {
            $clean = str_replace((array) $replace, ' ', $clean);
        }

        $clean = preg_replace("/[^a-zA-Z0-9\/_|+ -]/", '', $clean);
        $clean = strtolower($clean);
        $clean = preg_replace("/[\/_|+ -]+/", $delimiter, $clean);
        $clean = trim($clean, $delimiter);

        // Revert back to the old locale
        setlocale(LC_ALL, $oldLocale);

        return $clean;
    }
}

# ===========================================================

Mobile_Detect 是一个轻量级的开源移动设备（手机）检测的 PHP Class，它使用 User-Agent 中的字符串，
并结合 HTTP Header，来检测移动设备环境。这个设备检测的 PHP 类库最强大的地方是，它有一个非常完整
的库，可以检测出所用的设备类型（包括操作类型，以及手机品牌等都能检测）和浏览器的详细信息。

官方主页：http://mobiledetect.net/
demo：http://demo.mobiledetect.net/

https://github.com/serbanghita/Mobile-Detect

//使用实例
 
include 'Mobile_Detect.php';
$detect = new Mobile_Detect();
 
// Check for any mobile device.
if ($detect->isMobile())
 
// Check for any tablet.
if($detect->isTablet())
 
// Check for any mobile device, excluding tablets.
if ($detect->isMobile() && !$detect->isTablet())
 
if ($detect->isMobile() && !$detect->isTablet())
 
// Alternative to $detect->isAndroidOS()
$detect->is('AndroidOS');
 
// Batch usage
foreach($userAgents as $userAgent){
  $detect->setUserAgent($userAgent);
  $isMobile = $detect->isMobile();
}
 
// Version check.
$detect->version('iPad'); // 4.3 (float)

# ===========================================================
Phalcon models

// Magic properties can be used to store a records and its related properties:

// Create an artist
$artist = new Artists();
$artist->name = 'Shinichi Osawa';
$artist->country = 'Japan';

// Create an album
$album = new Albums();
$album->name = 'The One';
$album->artist = $artist; //Assign the artist
$album->year = 2008;

//Save both records
$album->save();


// Saving a record and its related records in a has-many relation:

// Get an existing artist
$artist = Artists::findFirst('name = "Shinichi Osawa"');

// Create an album
$album = new Albums();
$album->name = 'The One';
$album->artist = $artist;

$songs = array();

// Create a first song
$songs[0] = new Songs();
$songs[0]->name = 'Star Guitar';
$songs[0]->duration = '5:54';

// Create a second song
$songs[1] = new Songs();
$songs[1]->name = 'Last Days';
$songs[1]->duration = '4:29';

// Assign the songs array
$album->songs = $songs;

// Save the album + its songs
$album->save();


# ===========================================================

总结在PHP中，类的组合方式

比如，我们有一个Cache类(或Database类)，很多地方都要用到，那么怎么样创建和调用这个Cache类？
常见的有哪些方法？(staticMethod，loadClass，Trait，Singleton，Factory，new，argToConstruction，
DenpendencyInjection/Registry)

可扩展的类设计(Adapter方式，如Cache、Database等；ZF中的ViewHelper，Ashley中的Assert)

怎样发现类？怎样合理地划分和组织类？

# ===========================================================

PHP DOM API

function getElementById($id)
{
    $xpath = new DOMXPath($this->domDocument);
    return $xpath->query("//*[@id='$id']")->item(0);
}


$doc = new DOMDocument;
$doc->preserveWhiteSpace = false;
$doc->load('book.xml');

$xpath = new DOMXPath($doc);

$tbody = $doc->getElementsByTagName('tbody')->item(0);

// our query is relative to the tbody node
$query = 'row/entry[. = "en"]';

$entries = $xpath->query($query, $tbody);

foreach ($entries as $entry) {
    echo "Found {$entry->previousSibling->previousSibling->nodeValue}," .
         " by {$entry->previousSibling->nodeValue}\n";
}


$doc = new DOMDocument;

// We don't want to bother with white spaces
$doc->preserveWhiteSpace = false;
$doc->Load('book.xml');

$xpath = new DOMXPath($doc);

// We starts from the root element
$query = '//book/chapter/para/informaltable/tgroup/tbody/row/entry[. = "en"]';

$entries = $xpath->query($query);

foreach ($entries as $entry) {
    echo "Found {$entry->previousSibling->previousSibling->nodeValue}," .
         " by {$entry->previousSibling->nodeValue}\n";
}

# ===========================================================

PHP Plates Template Engine

$templates = new League\Plates\Engine('/path/to/templates');
$templates = new Tem\Plates\Engine();  // a funny namespace name

有cache吗？（整页cache，部分cache）如何翻译？

http://platesphp.com/

# ===========================================================

function getHoroscope($dob)
{
    list (, $month, $day) = explode("-", $dob);
    $horoscopes = array (
        'Aquarius'    => array('from' =>  '1-20', 'to' =>  '2-20'),
        'Pisces'      => array('from' =>  '2-19', 'to' =>  '3-21'),
        'Aries'       => array('from' =>  '3-20', 'to' =>  '4-21'),
        'Taurus'      => array('from' =>  '4-20', 'to' =>  '5-22'),
        'Gemini'      => array('from' =>  '5-21', 'to' =>  '6-22'),
        'Cancer'      => array('from' =>  '6-21', 'to' =>  '7-24'),
        'Leo'         => array('from' =>  '7-23', 'to' =>  '8-24'),
        'Virgo'       => array('from' =>  '8-23', 'to' =>  '9-24'),
        'Libra'       => array('from' =>  '9-23', 'to' => '10-24'),
        'Scorpio'     => array('from' => '10-23', 'to' => '11-23'),
        'Sagittarius' => array('from' => '11-22', 'to' => '12-23'),
        'Capricorn'   => array('from' => '12-22', 'to' =>  '1-21'),
    );


    foreach ($horoscopes as $horo => $dates) {
        list ($from_month, $from_day) = explode("-", $dates['from']);
        list ($to_month, $to_day) = explode("-", $dates['to']);

        if (($month == $from_month && $day > $from_day) ||
            ($month == $to_month && $day < $to_day)) {
            return $horo;
        }
    }
}
echo getHoroscope('0000-11-9');

function getRealpathCacheSize()
{
    $size = ini_get('realpath_cache_size');
    $size = trim($size);
    $unit = strtolower(substr($size, -1, 1));
    switch ($unit) {
        case 'g':
            return $size * 1024 * 1024 * 1024;
        case 'm':
            return $size * 1024 * 1024;
        case 'k':
            return $size * 1024;
        default:
            return (int) $size;
    }
}
#echo getRealpathCacheSize();

function slugify($string)
{
    return trim(preg_replace('/[^a-z0-9]+/', '-', strtolower(strip_tags($string))), '-');
}
#echo slugify('--a@@b##c$$1%%2**3--');

function generateClientId()
{
    $md5 = md5(uniqid(rand(), true));
    $clientId = substr($md5, 0, 8) . '-' . substr($md5, 8, 4) . '-' . substr($md5, 12, 4) . '-' . substr($md5, 16, 4) . '-' . substr($md5, 20);
    return $clientId;
}

function getUserAgent()
{
    if (isset($_SERVER['HTTP_USER_AGENT'])) {
        return $_SERVER['HTTP_USER_AGENT'];
    }
    return null;
}

function getIpAddress()
{
    $ordered_choices = array(
        'HTTP_X_FORWARDED_FOR',
        'HTTP_X_REAL_IP',
        'HTTP_CLIENT_IP',
        'REMOTE_ADDR'
    );
    $invalid_ips = array('127.0.0.1', '::1');

    // check each server var in order
    // accepted ip must be non null and not in the invalid_ips list
    foreach ($ordered_choices as $var) {
        if (isset($_SERVER[$var])) {
            $ip = $_SERVER[$var];
            if ($ip && !in_array($ip, $invalid_ips)) {
                $ips = explode(',', $ip);
                return reset($ips);
            }
        }
    }

    return null;
}

# NullObject has any property and any method
class NullObject
{
    public function __set($name, $value)
    {
        echo __CLASS__, "->$name=$value", PHP_EOL;
        return $this;
    }

    public function __get($name)
    {
        echo __CLASS__, "->$name", PHP_EOL;
        return $this;
    }

    public function __isset($name)
    {
        return true;
    }

    public function __unset($name)
    {
    }

    public function __call($name, $arguments)
    {
        echo __CLASS__, "::$name()", PHP_EOL;
        return $this;
    }

    public static function __callStatic($name, $arguments)
    {
        echo __CLASS__, "::$name()", PHP_EOL;
    }

    public function __toString()
    {
        return '';
    }
}

// $obj = new NullObject();
// $obj->user->login()->logout();
// $obj->name = 'NULL';
// $obj->weight = '100KG';
// $obj::send();
// NullObject::hello();
// if (isset($obj->age)) echo "Age\n";
// echo $obj;

class NullArray implements \ArrayAccess //, \Iterator
{
    public function offsetSet($key, $value)
    {
    }

    public function offsetGet($key)
    {
        return null;
    }

    public function offsetExists($key)
    {
        return true;
    }

    public function offsetUnset($key)
    {
    }
}

// $arr = new NullArray();
// pr($arr);
// pr($arr['anykey']);
// pr($arr['anykey']['whatever']['none']);
// pr($arr['1']['2']['3']);
// pr(current($arr));

# ===========================================================

//const CONFIG_DIR = '~/tmp';
const CONFIG_DIR = '/Users/hansonli/tmp';

class Config
{
    static $_cache = array();

    public static function load($file)
    {
        if (array_key_exists($file, self::$_cache))
          return self::$_cache[$file];

        $filename = CONFIG_DIR.'/'.$file;

        $cfg = false;
        // file_exists doesn't support ~ in path like '~/tmp'
        if (file_exists($filename)) {
            $cfg = parse_ini_file($filename, true);
        }

        self::$_cache[$file] = $cfg;
        return $cfg;
    }
}


$cfg = Config::load('tt.ini');
print_r($cfg);

<?php

define('MYSQL_CHARSET', 'utf8');

class Database extends PDO
{
    private static $dblink = null;

    private function __construct() { }

    /**
     * connect to database server
     * $db = Database::connect();
     */
    public static function connect()
    {
        if (empty(self::$dblink)) {

            $protocol = 'mysql';
            $username = 'root';
            $password = '';
            $host     = 'localhost';
            $dbname   = 'btc';

            $options = array(PDO::MYSQL_ATTR_INIT_COMMAND => 'SET NAMES '.MYSQL_CHARSET);

            $db = new Database("$protocol:dbname=$dbname;host=$host;charset=".MYSQL_CHARSET,
                    $username, $password, $options);

            $db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

            self::$dblink = $db;
        }

        return self::$dblink;
    }
}

function render($__file, $__vars)
{
    extract($__vars);
    ob_start();
    include(TEMPLATE_DIR . $__file);
    $contents = ob_get_contents();
    ob_end_clean();
    return $contents;
}

public function h($str)
{
    return htmlentities($str, ENT_NOQUOTES, "UTF-8");
}

# ===========================================================
Phake 2.1.0 发布，PHP 单元测试框架

http://www.oschina.net/news/62280/phake-2-1-0

Phake 2.1.0 发布，此版本现已提供下载：
https://github.com/mlively/Phake/archive/v2.1.0.zip。

Phake 是 PHP 框架，提供模拟对象，双向测试和方法测试。

Phake 和其他 PHP 模拟框架（PHPUnit，PHPMock 等）最主要的不同之处是 Phake 
会使用一个验证策略来确保调用。也就是说，你正常调用代码的时候你就完成代码的测试，确定是否
是按计划调用了方法。这跟其他的 PHP 测试框架非常不同，在任何调用之前使用一个期望策略。

# ===========================================================
class SuperGlobal implements ArrayAccess;

$this->request->get['key']; 取值，若不存在，返回null，不会触发PHP警告
$this->request->get['key'] = 'value'; 赋值
$this->request->get($key, ''); 取值，若不存在，返回默认值
$this->request->post($key); 同上
$this->request->cookie($key); 同上

$this->request->param($key); // $_REQUEST=$_GET + $_POST + $_COOKIES

$this->response->json($data);

if ($this->session->get('captcha') != $this->request->post('captcha')) {
}

	public function captcha() {
		$this->load->library('captcha');

		$captcha = new Captcha();

		$this->session->data['captcha'] = $captcha->getCode();
		$this->session->set('captcha', $captcha->getCode());

		$captcha->showImage();
	}	

# ===========================================================
    public function createdAgo(\DateTime $dateTime)
    {
        $delta = time() - $dateTime->getTimestamp();
        if ($delta < 0)
            throw new \InvalidArgumentException("createdAgo is unable to handle dates in the future");

        $duration = "";
        if ($delta < 60)
        {
            // Seconds
            $time = $delta;
            $duration = $time . " second" . (($time > 1) ? "s" : "") . " ago";
        }
        else if ($delta <= 3600)
        {
            // Mins
            $time = floor($delta / 60);
            $duration = $time . " minute" . (($time > 1) ? "s" : "") . " ago";
        }
        else if ($delta <= 86400)
        {
            // Hours
            $time = floor($delta / 3600);
            $duration = $time . " hour" . (($time > 1) ? "s" : "") . " ago";
        }
        else
        {
            // Days
            $time = floor($delta / 86400);
            $duration = $time . " day" . (($time > 1) ? "s" : "") . " ago";
        }

        return $duration;
    }

public function slugify($text)
{
    // replace non letter or digits by -
    $text = preg_replace('#[^\\pL\d]+#u', '-', $text);

    // trim
    $text = trim($text, '-');

    // transliterate
    if (function_exists('iconv'))
    {
        $text = iconv('utf-8', 'us-ascii//TRANSLIT', $text);
    }

    // lowercase
    $text = strtolower($text);

    // remove unwanted characters
    $text = preg_replace('#[^-\w]+#', '', $text);

    if (empty($text))
    {
        return 'n-a';
    }

    return $text;
}

# ===========================================================
在Symfony的Controller中

$em = $this->getDoctrine()->getManager();

$blog = $em->getRepository('BloggerBlogBundle:Blog')->find($blog_id);
or
$blog = new Blog(); // Entity
$blog->title = '';

$em->persist($blog);
$em->flush();

# ===========================================================
我改过的MicroMVC Router，可用于opencart

<?php
/**
 * - Routes are matched from left-to-right.
 * - Regex can also be used to define routes if enclosed in "/.../"
 * - Each regex catch pattern (...) will be viewed as a parameter.
 * - The remaning (unmached) URL path will be passed as parameters.
 */
$config['routes'] = array(
    'homepage' => '\Controller\Index',
    '404' => '\Controller\Page404',
    'school' => '\Controller\School',

    'example/path' => '\Controller\Example\Hander',
    "forum/topic/view" => 'Forum\Controller\Forum\View',

    'example/([^/]+)' => '\Controller\Example\Param',
    '^(\w+)/recent/comments' => 'Comments\Controller\Recent',
);

class Dispatch
{
    public $routes;

    public function __construct(array $routes)
    {
        $this->routes = $routes;
    }

    /**
     * Parse the given URL path and return the correct controller and parameters.
     */
    public function route($path)
    {
        $path = trim($path, '/');

        if (isset($this->routes[$path])) {
            return array($this->routes[$path], array());
        }

        foreach($this->routes as $route => $controller) {

            if (substr($path, 0, strlen($route)) === $route) {
                $params = explode('/', trim(substr($path, strlen($route)), '/'));
                return array($controller, $params);
            }

            if (preg_match('#'.$route.'#', $path, $matches)) {
                $complete = array_shift($matches);

                $params = explode('/', trim(substr($path, strlen($complete)), '/'));

                if ($params[0]) {
                    foreach ($matches as $match) {
                        array_unshift($params, $match);
                    }
                }
                else {
                    $params = $matches;
                }

                return array($controller, $params);
            }
        }

        // path not found
        return array('error/not_found', array());
    }
}

$dispatch = new Dispatch($config['routes']);

$path = '404'; // ??
$path = 'school';
$path = "forum/topic/view";
$path = '/John_Doe4/recent/comments/3';
$path = '/forum/topic/view/45/Hello-World';
$path = 'example/path';

echo $path, "\n";
#list($controller, $params) = $dispatch->route($path);
$ret = $dispatch->route($path);
print_r($ret);

# ===========================================================

MicroMVC的路由

有空的时候研究一下这个，比较简洁实用
/**
 * URL Routing
 *
 * URLs are very important to the future usability of your site. Take
 * time to think about your structure in a way that is meaningful. Place
 * your most common page routes at the top for better performace.
 *
 * - Routes are matched from left-to-right.
 * - Regex can also be used to define routes if enclosed in "/.../"
 * - Each regex catch pattern (...) will be viewed as a parameter.
 * - The remaning (unmached) URL path will be passed as parameters.
 *
 ** Simple Example **
 * URL Path:	/forum/topic/view/45/Hello-World
 * Route:		"forum/topic/view" => 'Forum\Controller\Forum\View'
 * Result:		Forum\Controller\Forum\View->action('45', 'Hello-World');
 *
 ** Regex Example **
 * URL Path:	/John_Doe4/recent/comments/3
 * Route:		"/^(\w+)/recent/comments/' => 'Comments\Controller\Recent'
 * Result:		Comments\Controller\Recent->action($username = 'John_Doe4', $page = 3)
 */
$config = array();

$config['routes'] = array(
	''					=> '\Controller\Index',
	'404'				=> '\Controller\Page404',
	'school'			=> '\Controller\School',

	// Example paths
	//'example/path'		=> '\Controller\Example\Hander',
	//'example/([^/]+)'	=> '\Controller\Example\Param',
);

class Dispatch
{
	public $routes;

	public function __construct(array $routes)
	{
		$this->routes = $routes;
	}

	public function controller($path, $method)
	{
		// Parse the routes to find the correct controller
		list($params, $route, $controller) = $this->route($path);

		// Load and run action
		$controller = new $controller($route, $this);

		// We are ignoring TRACE & CONNECT
		$request_methods = array('GET', 'POST', 'PUT', 'DELETE', 'OPTIONS', 'HEAD');

		// Look for a RESTful method, or try the default run()
		if (!in_array($method, $request_methods) OR ! method_exists($controller, $method))
		{
			if (!method_exists($controller, 'run')) {
				throw new \Exception('Invalid Request Method.');
			}

			$method = 'run';
		}

		// Controller setup here
		$controller->initialize($method);

		if ($params) {
			call_user_func_array(array($controller, $method), $params);
		}
		else {
			$controller->$method();
		}

		// Return the controller instance
		return $controller;
	}

	/**
	 * Parse the given URL path and return the correct controller and parameters.
	 *
	 * @param string $path segment of URL
	 * @param array $routes to test against
	 * @return array
	 */
	public function route($path)
	{
		$path = trim($path, '/');

		// Default homepage route
		if ($path === '') {
			return array(array(), '', $this->routes['']);
		}

		// If this is not a valid, safe path (more complex params belong in GET/POST)
		if ($path AND ! preg_match('/^[\w\-~\/\.]{1,400}$/', $path)) {
			$path = '404';
		}

		foreach($this->routes as $route => $controller) {

			if (!$route) continue; // Skip homepage route

			// Is this a regex?
			if ($route{0} === '/') {
				if (preg_match($route, $path, $matches)) {
					$complete = array_shift($matches);

			// The following code tries to solve:
			// (Regex) "/^path/(\w+)/" + (Path) "path/word/other" = (Params) array(word, other)

					// Skip the regex match and continue from there
					$params = explode('/', trim(mb_substr($path, mb_strlen($complete)), '/'));

					if ($params[0]) {
						// Add captured group back into params
						foreach ($matches as $match) {
							array_unshift($params, $match);
						}
					}
					else {
						$params = $matches;
					}

					//print dump($params, $matches);
					return array($params, $complete, $controller);
				}
			}
			else {
				if (mb_substr($path, 0, mb_strlen($route)) === $route) {
					$params = explode('/', trim(mb_substr($path, mb_strlen($route)), '/'));
					return array($params, $route, $controller);
				}
			}
		}

		// Controller not found
		return array(array($path), $path, $this->routes['404']);
	}
}

# ===========================================================
OpenCart (done)

change this
    $this->load->model('install');
    $this->model_install->database($this->request->post);
to
    $model_install = $this->load->model('install');
    $model_install->database($this->request->post);

this is where to change

    public function model($model) {

        $object = new $class($this->registry);
        $this->registry->set('model_' . str_replace('/', '_', $model), $object);
        return $object;

# ===========================================================
