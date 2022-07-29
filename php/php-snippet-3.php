<?php

PHPUnit在运行某个测试方法前，会调用一个名叫setUp()的方法。 setUp()是创建测试所用对象的地方。
当测试方法运行结束后，不管是成功还是失败，都会调用另外一个名叫 tearDown() 的方法。tearDown() 
是清理测试所用对象的地方。

测试类的每个测试方法都会运行一次 setUp() 与 tearDown() 模板方法
(同时，每个测试方法都是在一个全新的测试类实例上运行的)。

class TemplateMethodsTest extends PHPUnit_Framework_TestCase
{
    public static function setUpBeforeClass()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    protected function setUp()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    protected function assertPreConditions()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    public function testOne()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
        $this->assertTrue(TRUE);
    }
 
    public function testTwo()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
        $this->assertTrue(FALSE);
    }
 
    protected function assertPostConditions()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    protected function tearDown()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    public static function tearDownAfterClass()
    {
        fwrite(STDOUT, __METHOD__ . "\n");
    }
 
    protected function onNotSuccessfulTest(Exception $e)
    {
        fwrite(STDOUT, __METHOD__ . "\n");
        throw $e;
    }
}

# ===========================================================
CakePHP paginate

    // TranslationController 
    $this->paginate = array(
        'limit' => $this->pageSize, 
        'order'=>array(
            'Translation.requester_rejected'=>'desc', 
            'Translation.requester_accepted' => 'asc', 
            'Translation.translator_id' => 'asc', 
            'Translation.translation_updatedon' => 'asc'
        ),
        'conditions'=>$conditions
    );
    $rows = $this->paginate('Translation');
    $this->set('data',$rows);

    // UserController
    $this->layout = 'default';
    $this->paginate = array(
        'limit' => $this->pageSize, 
        'order' => array('User.id' => 'asc')
    );
    $data = $this->paginate('User');
    $this->set('data', $data);

# ===========================================================

/**
 * Parse a db URI string into an array representation.
 *
 * @param string $uri
 * @return array 
 *   Array
 *   (
 *       [scheme] => mysql
 *       [host] => master.sql.amcluster.com
 *       [name] => aminno
 *       [user] => amroot
 *       [pass] => WugDoahi
 *   )
 */
function parse_db_uri($uri)
{
  /**
   * For a url http://username:password@hostname/path?arg=value#anchor
   * We will get info like this
   * Array
   * (
   *     [scheme] => http
   *     [host] => hostname
   *     [user] => username
   *     [pass] => password
   *     [path] => /path
   *     [query] => arg=value
   *     [fragment] => anchor
   * )
   */
  $info = parse_url($uri);

  // db URI must start with 'mysql://'
  if (!isset($info['scheme']) || $info['scheme'] != 'mysql')
    return false;

  // ['path'] is database name
  if (isset($info['path'])) {
    $info['name'] = ltrim($info['path'], '/');
    unset($info['path']);
  }

  // get username & password from query string 'user=amroot,pass=WugDoahi'
  if (isset($info['query'])) {
    $str = str_replace(',', '&', $info['query']);
    parse_str($str, $vars);
    $info += $vars;
    unset($info['query']);
  }

  // ['fragment'] is an anchor (like #top), we don't need it.
  if (isset($info['fragment']))
    unset($info['fragment']);

  return $info;
}

print_r(parse_db_uri('mysql://amroot:WugDoahi@master.sql.amcluster.com/*'));
print_r(parse_db_uri('mysql://master.sql.amcluster.com/aminno?user=amroot,pass=WugDoahi'));
print_r(parse_db_uri('mysql://amroot:WugDoahi@admindb/*'));
print_r(parse_db_uri('mysql://amroot:WugDoahi@admindb/aminno'));
print_r(parse_db_uri('mysql://amroot:WugDoahi@master.sql.amcluster.com/aminno'));
print_r(parse_db_uri('mysql://master.sql.amcluster.com/aminno?user=amroot,pass=WugDoahi'));

# ===========================================================
Install Controller

<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');
/**
 * controller that install database to current default database connection 
 *
 * that controller contains one page that reads mysql.sql file
 * parse it and execute all  statments.
 *
 * @copyright  2011 Emad Elsaid a.k.a Blaze Boy
 * @license    http://www.gnu.org/licenses/gpl-2.0.txt   GPL License 2.0
 * @link       https://github.com/blazeeboy/Codeigniter-Egypt
 */ 
class Install extends CI_Controller
{
    public function __construct()
    {
        parent::__construct();

        $this->load->database();
        $this->load->helper('url');
    }
    
    public function index()
    {
        if (count($this->db->list_tables()) == 0)
        {
            $script = explode( ';', file_get_contents('mysql.sql'));
            $script = array_map('trim', $script);
            $script = array_filter($script, 'count');

            foreach($script as $line)
            {
                if ($line != '')
                    $this->db->query($line);
            }
        }

        redirect('');
    }
}

# ===========================================================
CakePHP relations

    [__associations] => Array
    (
        [0] => belongsTo
        [1] => hasOne
        [2] => hasMany
        [3] => hasAndBelongsToMany
    )

    [__associationKeys] => Array
    (
        [belongsTo] => Array
        (
            [0] => className
            [1] => foreignKey
            [2] => conditions
            [3] => fields
            [4] => order
            [5] => counterCache
        )

        [hasOne] => Array
        (
            [0] => className
            [1] => foreignKey
            [2] => conditions
            [3] => fields
            [4] => order
            [5] => dependent
        )

        [hasMany] => Array
        (
            [0] => className
            [1] => foreignKey
            [2] => conditions
            [3] => fields
            [4] => order
            [5] => limit
            [6] => offset
            [7] => dependent
            [8] => exclusive
            [9] => finderQuery
            [10] => counterQuery
        )

        [hasAndBelongsToMany] => Array
          (
              [0] => className
              [1] => joinTable
              [2] => with
              [3] => foreignKey
              [4] => associationForeignKey
              [5] => conditions
              [6] => fields
              [7] => order
              [8] => limit
              [9] => offset
              [10] => unique
              [11] => finderQuery
              [12] => deleteQuery
              [13] => insertQuery
          )
    )

# ===========================================================
如何在图片底部追加自己的

<?php
    $imgPath = 'http://itqiubai.com/imgs/bottom.jpg';
    $mainPath = 'http://img2.aili.com/201301/29/1359446507_22643900.jpg';

    $im = imagecreatefromjpeg($mainPath);
    list($width, $height, $type, $attr) = getimagesize($mainPath);

    $bottom = imagecreatefromjpeg($imgPath);
    list($width2, $height2, $type2, $attr2) = getimagesize($imgPath);

    $dest=imagecreatetruecolor($width,$height+$height2);
    imagecopy($dest,$im,0,0,0,0,$width,$height);
    imagecopy(
        $dest, //dst_im
        $bottom, //src_im
        ($width-$width2)/2, //dst_x
        $height, //dst_y
        0, //src_x
        0, //src_y
        $width, //src_w
        $height2 //src_h
    );
     
    header('Content-type: image/jpeg');
    imagejpeg($dest);
    imagedestroy($dest) ;


图片还原

/**
 * 裁切图片 高度冲裁
 * @param unknown $src_file   源文件
 * @param unknown $dst_file   新文件名
 * @param number $height      裁切掉的高度
 */
function my_image_resize($src_file, $dst_file, $height = 30) {
    if (! file_exists ( $src_file )) {
        echo $src_file . " is not exists !";
        exit ();
    }
    $type = exif_imagetype ( $src_file );
    $support_type = array (
            IMAGETYPE_JPEG,
            IMAGETYPE_PNG,
            IMAGETYPE_GIF
    );
    if (! in_array ( $type, $support_type, true )) {
        echo "this type of image does not support! only support jpg , gif or png";
        exit ();
    }
    switch ($type) {
        case IMAGETYPE_JPEG :
            $src_img = imagecreatefromjpeg ( $src_file );
            break;
        case IMAGETYPE_PNG :
            $src_img = imagecreatefrompng ( $src_file );
            break;
        case IMAGETYPE_GIF :
            $src_img = imagecreatefromgif ( $src_file );
            break;
        default :
            echo "Load image error!";
            exit ();
    }
    $w = imagesx ( $src_img );
    $h = imagesy ( $src_img );
    $ratio_w = $w;
    $ratio_h = $h - $height;
     
    $inter_img = imagecreatetruecolor ( $w, $h );
     
    imagecopyresampled ( $inter_img, $src_img, 0, 0, 0, 0, $w, $h, $w, $h );
    $new_img = imagecreatetruecolor ( $ratio_w, $ratio_h );
    imagecopy ( $new_img, $inter_img, 0, 0, 0, 0, $ratio_w, $ratio_h );
    switch ($type) {
        case IMAGETYPE_JPEG :
            imagejpeg ( $new_img, $dst_file, 100 );
            break;
        case IMAGETYPE_PNG :
            imagepng ( $new_img, $dst_file, 100 );
            break;
        case IMAGETYPE_GIF :
            imagegif ( $new_img, $dst_file, 100 );
            break;
        default :
            break;
    }
}

# ===========================================================
function sql_insert_into($table, array $bind)
{
    $cols = array();
    $vals = array();

    foreach ($bind as $col => $val) {
        $cols[] = $col;
        if (is_string($val))
            $val = "'".sqlite_escape_string($val)."'";
        $vals[] = $val;
    }
    
    $sql = 'INSERT INTO '.$table.' ('. 
            implode(',', $cols) .') VALUES ('.
            implode(',', $vals) .')';
    
    // echo $sql, PHP_EOL;
    return $sql;
}

function objectToArray( $object ) {
    if( !is_object( $object ) && !is_array( $object ) ) {
        return $object;
    }
    if( is_object( $object ) ) {
        $object = (array) $object;
    }
    return array_map( 'objectToArray', $object );
}

// another way to convert object to array
$array = json_decode(json_encode($object), true);

# ===========================================================
一句话添加数字千位符
// PHP
echo preg_replace('/(?<=\\d)(?=(?:\\d{3})+$)/', ',', '1234567890');
// Java
"1234567890".replaceAll("(?<=\\d)(?=(?:\\d{3})+$)", ",");
# ===========================================================
<?php

// trace(__FILE__);
// trace(__METHOD__);
// trace($var, 'return value is');
// trace($var, __METHOD__);
// trace($var, __METHOD__, __LINE__);

function trace($var, $func='', $ln=0)
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

function dpr($var, $func='', $ln=0)
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

function pr($var, $label='')
{
    echo "<pre>\n";
    if ($label) echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

# ===========================================================

public static function getInstance() {
   return new static(); // 能够继承
   return new self(); // 无法继承
}

# ===========================================================

function sendMail() {

  $email = 'lihsca@gmail.com';

  $csv = "pnum,email,username,password,fraud_status,membersince,profile_caption\n";

  // send the email

  $filename = "spam_suspect_list.csv";
  // $type = 'application/vnd.ms-excel';

  $semi_rand = md5(time());
  $mime_boundary = "==Multipart_Boundary_x{$semi_rand}x";
  $header = "MIME-Version: 1.0\n" .
            "Content-Type: multipart/mixed;\n" .
            " boundary=\"{$mime_boundary}\"";

  $subject = $body = "Spam suspect list on " . date('Y-m-d, l', strtotime('-1 day')) . '.';

  $message = "This is a multi-part message in MIME format.\n\n" .
             "--{$mime_boundary}\n" .
             "Content-Type: text/plain; charset=\"iso-8859-1\"\n" .
             "Content-Transfer-Encoding: 7bit\n\n" .
             $body. "\n\n" .
             "--{$mime_boundary}\n" .
             "Content-Type: application/vnd.ms-excel;\n" .
             "Content-Disposition: attachment; filename=\"$filename\"\n" .
             "Content-Transfer-Encoding: base64\n\n" .
             chunk_split(base64_encode($csv)) . "\n\n" .
             "--{$mime_boundary}--\n";

  mail($email,$subject, $message, $header);
}

sendMail();

# ===========================================================
$   php -r '
    require "app/config/config.php";

    $n = $config->database->name;
    $u = $config->database->username;
    $p = $config->database->password;

    echo `echo "CREATE DATABASE {$n}" | mysql -u {$u} -p {$p}`;
    echo `cat schemas/php_site.sql | mysql -u {$u} -p {$p} {$n}`;
    '
--------
    echo 'CREATE DATABASE forum' | mysql -u root
    cat schemas/forum.sql | mysql -u root forum

# ===========================================================
function t($text, $params=array())
{
  return strtr(gettext($text), $params); // strtr is slower than str_replace
  return str_replace(array_keys($params), array_values($params), gettext($text));
}

echo t('hello, %name!', array('%name'=>'world')), EOL;
echo t('hello, PHP!'), EOL;
# ===========================================================
CakePHP .htaccess

<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteRule ^(.*)$ index.php?url=$1 [QSA,L]
</IfModule>
# ===========================================================
// run syslog command to view syslog 

// open syslog, include the process ID and also send
// the log to standard error, and use a user defined
// logging mechanism
openlog("ashley", LOG_PID | LOG_PERROR, LOG_LOCAL0);
//openlog('ashley', LOG_ODELAY | LOG_PID, LOG_LOCAL0);

//syslog(LOG_WARNING, "Unauthorized client: {$_SERVER['REMOTE_ADDR']} ({$_SERVER['HTTP_USER_AGENT']})");
syslog(LOG_WARNING, "Unauthorized client.");
syslog(LOG_WARNING, "Close Log.");

closelog();
# ===========================================================
这是CakePHP中的做法
$beforeRender = new CakeEvent('View.beforeRender', $this, array($viewFileName));
$this->getEventManager()->dispatch(beforeRender);

$afterEvent = new CakeEvent('View.afterRender', $this, array($viewFileName));
$this->getEventManager()->dispatch(afterRender);

这是我的想法
interface EventListener
{
    public function handleEvent($event);
//  public function processEvent($event);
}

$beforeRender = new CakeEvent('View.beforeRender', $this, array($viewFileName));
$this->handleEvent($beforeRender);

$afterEvent = new CakeEvent('View.afterRender', $this, array($viewFileName));
$this->handleEvent($afterRender);

两种方法区别及优劣何在？


Chat.Views.Tools.GiftToolView = Chat.Views.ToolView.extend({
  id: 'chatGift',
  className: 'tool tool-gift',
  title: 'Send a Gift',

  events: {
    'click': 'onClick'
  },

  listeners: {
    'client:paused': 'onPause',
    'client:unpaused': 'onUnpause',
    'client:pause:hide': 'onHide',
    'client:pause:show': 'onShow',
    'client:timer:tick': 'onTick'
  },

  initialize: function(options) {

    Chat.Events.on({
      'tabs:add': this.onTabAdded,
      // TODO remove this, currently a hack for Screens.js
      'tabs:open': this.openTabs,
      'ui:no-user': this.closeTabs
    }, this);
  },

  onClick: function() {
    Chat.Events.trigger('tool:pause:click');
  },

  onTick: function() {
    Chat.Events.trigger('tool:gift:click');
  },

  onHide: function() {
    this.$el.hide();
  },

  onShow: function() {
    this.$el.show();
  }
});
# ===========================================================
<?php

function pr($label, $var)
{
    echo "<pre>\n";
    echo "<b>$label</b>\n";
    print_r($var);
    echo "</pre>\n";
}

function dpr()
{
    echo "<pre>\n";
    $args = func_get_args();
    foreach ($args as $var) {
        print_r($var);
    }
    echo "</pre>\n";
}

function fpr()
{
    $filename = '/data/source/trace.log';

    $args = func_get_args();
    foreach ($args as $var) {
        $str = print_r($var, true)."\n";
        error_log($str, 3, $filename);
    }
}

function fpr()
{
  $filename = '/tmp/cake.log';

  $numargs = func_num_args();
  $args = func_get_args();

  // if ($numargs == 0) return;

  if ($numargs == 2) {
    $label = array_shift($args);
    if (!empty($label)) {
      error_log($label.' = ', 3, $filename);
    }
  }

  foreach ($args as $var) {
    $str = print_r($var, true)."\n";
    error_log($str, 3, $filename);
  }
}
# ===========================================================
/**
 * strip the slashes that have been added to our POST/GET data!
 */
if (ini_get('magic_quotes_gpc')) {

  function array_clean(&$value) {
    $value = stripslashes($value);
  }
  //php 5+ only
  array_walk_recursive($_GET, 'array_clean');
  array_walk_recursive($_POST, 'array_clean');
  array_walk_recursive($_COOKIE, 'array_clean');
}
# ===========================================================
Upgrade PHP to v5.6.7

$ php -v
PHP 5.4.26 (cli) (built: Apr  3 2014 13:45:30) 
Copyright (c) 1997-2014 The PHP Group
Zend Engine v2.4.0, Copyright (c) 1998-2014 Zend Technologies
    with Xdebug v2.2.3, Copyright (c) 2002-2013, by Derick Rethans

$ which php
/usr/local/bin/php

$ ll /usr/local/bin/php
lrwxr-xr-x  1 hansonli  admin  30  3 Apr  2014 /usr/local/bin/php -> ../Cellar/php54/5.4.26/bin/php

$ brew update
$ brew instal php56
$ brew unlink php54
$ brew link php56
$ brew install php56-xdebug

$ rm /usr/local/bin/php
$ ln -s /usr/local/Cellar/php56/5.6.7/bin/php /usr/local/bin/php

$ php -v
PHP 5.6.7 (cli) (built: Mar 22 2015 19:03:55) 
Copyright (c) 1997-2015 The PHP Group
Zend Engine v2.6.0, Copyright (c) 1998-2015 Zend Technologies
    with Xdebug v2.3.2, Copyright (c) 2002-2015, by Derick Rethans
# ===========================================================
Phalcon\Acl
Phalcon\Acl\Adapter
Phalcon\Acl\AdapterInterface
Phalcon\Acl\Adapter\Memory
Phalcon\Acl\Exception
Phalcon\Acl\Resource
Phalcon\Acl\ResourceInterface
Phalcon\Acl\Role
Phalcon\Acl\RoleInterface
Phalcon\Annotations\Adapter
Phalcon\Annotations\AdapterInterface
Phalcon\Annotations\Adapter\Apc
Phalcon\Annotations\Adapter\Files
Phalcon\Annotations\Adapter\Memory
Phalcon\Annotations\Adapter\Xcache
Phalcon\Annotations\Annotation
Phalcon\Annotations\Collection
Phalcon\Annotations\Exception
Phalcon\Annotations\Reader
Phalcon\Annotations\ReaderInterface
Phalcon\Annotations\Reflection
Phalcon\Assets\Collection
Phalcon\Assets\Exception
Phalcon\Assets\FilterInterface
Phalcon\Assets\Filters\Cssmin
Phalcon\Assets\Filters\Jsmin
Phalcon\Assets\Filters\None
Phalcon\Assets\Manager
Phalcon\Assets\Resource
Phalcon\Assets\Resource\Css
Phalcon\Assets\Resource\Js
Phalcon\CLI\Console
Phalcon\CLI\Console\Exception
Phalcon\CLI\Dispatcher
Phalcon\CLI\Dispatcher\Exception
Phalcon\CLI\Router
Phalcon\CLI\Router\Exception
Phalcon\CLI\Task
Phalcon\Cache\Backend
Phalcon\Cache\BackendInterface
Phalcon\Cache\Backend\Apc
Phalcon\Cache\Backend\File
Phalcon\Cache\Backend\Memcache
Phalcon\Cache\Backend\Memory
Phalcon\Cache\Backend\Mongo
Phalcon\Cache\Backend\Xcache
Phalcon\Cache\Exception
Phalcon\Cache\FrontendInterface
Phalcon\Cache\Frontend\Base64
Phalcon\Cache\Frontend\Data
Phalcon\Cache\Frontend\Igbinary
Phalcon\Cache\Frontend\Json
Phalcon\Cache\Frontend\None
Phalcon\Cache\Frontend\Output
Phalcon\Cache\Multiple
Phalcon\Config
Phalcon\Config\Adapter\Ini
Phalcon\Config\Adapter\Json
Phalcon\Config\Exception
Phalcon\Crypt
Phalcon\CryptInterface
Phalcon\Crypt\Exception
Phalcon\DI
Phalcon\DI\Exception
Phalcon\DI\FactoryDefault
Phalcon\DI\FactoryDefault\CLI
Phalcon\DI\Injectable
Phalcon\DI\InjectionAwareInterface
Phalcon\DI\Service
Phalcon\DI\ServiceInterface
Phalcon\DI\Service\Builder
Phalcon\Db
Phalcon\Db\Adapter
Phalcon\Db\AdapterInterface
Phalcon\Db\Adapter\Pdo
Phalcon\Db\Adapter\Pdo\Mysql
Phalcon\Db\Adapter\Pdo\Oracle
Phalcon\Db\Adapter\Pdo\Postgresql
Phalcon\Db\Adapter\Pdo\Sqlite
Phalcon\Db\Column
Phalcon\Db\ColumnInterface
Phalcon\Db\Dialect
Phalcon\Db\DialectInterface
Phalcon\Db\Dialect\Mysql
Phalcon\Db\Dialect\Oracle
Phalcon\Db\Dialect\Postgresql
Phalcon\Db\Dialect\Sqlite
Phalcon\Db\Exception
Phalcon\Db\Index
Phalcon\Db\IndexInterface
Phalcon\Db\Profiler
Phalcon\Db\Profiler\Item
Phalcon\Db\RawValue
Phalcon\Db\Reference
Phalcon\Db\ReferenceInterface
Phalcon\Db\ResultInterface
Phalcon\Db\Result\Pdo
Phalcon\Debug
Phalcon\DiInterface
Phalcon\Dispatcher
Phalcon\DispatcherInterface
Phalcon\Escaper
Phalcon\EscaperInterface
Phalcon\Escaper\Exception
Phalcon\Events\Event
Phalcon\Events\EventsAwareInterface
Phalcon\Events\Exception
Phalcon\Events\Manager
Phalcon\Events\ManagerInterface
Phalcon\Exception
Phalcon\Filter
Phalcon\FilterInterface
Phalcon\Filter\Exception
Phalcon\Filter\UserFilterInterface
Phalcon\Flash
Phalcon\FlashInterface
Phalcon\Flash\Direct
Phalcon\Flash\Exception
Phalcon\Flash\Session
Phalcon\Forms\Element
Phalcon\Forms\ElementInterface
Phalcon\Forms\Element\Check
Phalcon\Forms\Element\Date
Phalcon\Forms\Element\Email
Phalcon\Forms\Element\File
Phalcon\Forms\Element\Hidden
Phalcon\Forms\Element\Numeric
Phalcon\Forms\Element\Password
Phalcon\Forms\Element\Select
Phalcon\Forms\Element\Submit
Phalcon\Forms\Element\Text
Phalcon\Forms\Element\TextArea
Phalcon\Forms\Exception
Phalcon\Forms\Form
Phalcon\Forms\Manager
Phalcon\Http\Cookie
Phalcon\Http\Cookie\Exception
Phalcon\Http\Request
Phalcon\Http\RequestInterface
Phalcon\Http\Request\Exception
Phalcon\Http\Request\File
Phalcon\Http\Request\FileInterface
Phalcon\Http\Response
Phalcon\Http\ResponseInterface
Phalcon\Http\Response\Cookies
Phalcon\Http\Response\CookiesInterface
Phalcon\Http\Response\Exception
Phalcon\Http\Response\Headers
Phalcon\Http\Response\HeadersInterface
Phalcon\Kernel
Phalcon\Loader
Phalcon\Loader\Exception
Phalcon\Logger
Phalcon\Logger\Adapter
Phalcon\Logger\AdapterInterface
Phalcon\Logger\Adapter\File
Phalcon\Logger\Adapter\Firephp
Phalcon\Logger\Adapter\Stream
Phalcon\Logger\Adapter\Syslog
Phalcon\Logger\Exception
Phalcon\Logger\Formatter
Phalcon\Logger\FormatterInterface
Phalcon\Logger\Formatter\Firephp
Phalcon\Logger\Formatter\Json
Phalcon\Logger\Formatter\Line
Phalcon\Logger\Formatter\Syslog
Phalcon\Logger\Item
Phalcon\Logger\Multiple
Phalcon\Mvc\Application
Phalcon\Mvc\Application\Exception
Phalcon\Mvc\Collection
Phalcon\Mvc\CollectionInterface
Phalcon\Mvc\Collection\Document
Phalcon\Mvc\Collection\Exception
Phalcon\Mvc\Collection\Manager
Phalcon\Mvc\Collection\ManagerInterface
Phalcon\Mvc\Controller
Phalcon\Mvc\ControllerInterface
Phalcon\Mvc\Dispatcher
Phalcon\Mvc\DispatcherInterface
Phalcon\Mvc\Dispatcher\Exception
Phalcon\Mvc\Micro
Phalcon\Mvc\Micro\Collection
Phalcon\Mvc\Micro\CollectionInterface
Phalcon\Mvc\Micro\Exception
Phalcon\Mvc\Micro\LazyLoader
Phalcon\Mvc\Micro\MiddlewareInterface
Phalcon\Mvc\Model
Phalcon\Mvc\ModelInterface
Phalcon\Mvc\Model\Behavior
Phalcon\Mvc\Model\BehaviorInterface
Phalcon\Mvc\Model\Behavior\SoftDelete
Phalcon\Mvc\Model\Behavior\Timestampable
Phalcon\Mvc\Model\Criteria
Phalcon\Mvc\Model\CriteriaInterface
Phalcon\Mvc\Model\Exception
Phalcon\Mvc\Model\Manager
Phalcon\Mvc\Model\ManagerInterface
Phalcon\Mvc\Model\Message
Phalcon\Mvc\Model\MessageInterface
Phalcon\Mvc\Model\MetaData
Phalcon\Mvc\Model\MetaDataInterface
Phalcon\Mvc\Model\MetaData\Apc
Phalcon\Mvc\Model\MetaData\Files
Phalcon\Mvc\Model\MetaData\Memory
Phalcon\Mvc\Model\MetaData\Session
Phalcon\Mvc\Model\MetaData\Strategy\Annotations
Phalcon\Mvc\Model\MetaData\Strategy\Introspection
Phalcon\Mvc\Model\MetaData\Xcache
Phalcon\Mvc\Model\Query
Phalcon\Mvc\Model\QueryInterface
Phalcon\Mvc\Model\Query\Builder
Phalcon\Mvc\Model\Query\BuilderInterface
Phalcon\Mvc\Model\Query\Lang
Phalcon\Mvc\Model\Query\Status
Phalcon\Mvc\Model\Query\StatusInterface
Phalcon\Mvc\Model\Relation
Phalcon\Mvc\Model\RelationInterface
Phalcon\Mvc\Model\ResultInterface
Phalcon\Mvc\Model\Resultset
Phalcon\Mvc\Model\ResultsetInterface
Phalcon\Mvc\Model\Resultset\Complex
Phalcon\Mvc\Model\Resultset\Simple
Phalcon\Mvc\Model\Row
Phalcon\Mvc\Model\Transaction
Phalcon\Mvc\Model\TransactionInterface
Phalcon\Mvc\Model\Transaction\Exception
Phalcon\Mvc\Model\Transaction\Failed
Phalcon\Mvc\Model\Transaction\Manager
Phalcon\Mvc\Model\Transaction\ManagerInterface
Phalcon\Mvc\Model\ValidationFailed
Phalcon\Mvc\Model\Validator
Phalcon\Mvc\Model\ValidatorInterface
Phalcon\Mvc\Model\Validator\Email
Phalcon\Mvc\Model\Validator\Exclusionin
Phalcon\Mvc\Model\Validator\Inclusionin
Phalcon\Mvc\Model\Validator\Numericality
Phalcon\Mvc\Model\Validator\PresenceOf
Phalcon\Mvc\Model\Validator\Regex
Phalcon\Mvc\Model\Validator\StringLength
Phalcon\Mvc\Model\Validator\Uniqueness
Phalcon\Mvc\Model\Validator\Url
Phalcon\Mvc\ModuleDefinitionInterface
Phalcon\Mvc\Router
Phalcon\Mvc\RouterInterface
Phalcon\Mvc\Router\Annotations
Phalcon\Mvc\Router\Exception
Phalcon\Mvc\Router\Group
Phalcon\Mvc\Router\Route
Phalcon\Mvc\Router\RouteInterface
Phalcon\Mvc\Url
Phalcon\Mvc\UrlInterface
Phalcon\Mvc\Url\Exception
Phalcon\Mvc\User\Component
Phalcon\Mvc\User\Module
Phalcon\Mvc\User\Plugin
Phalcon\Mvc\View
Phalcon\Mvc\ViewInterface
Phalcon\Mvc\View\Engine
Phalcon\Mvc\View\EngineInterface
Phalcon\Mvc\View\Engine\Php
Phalcon\Mvc\View\Engine\Volt
Phalcon\Mvc\View\Engine\Volt\Compiler
Phalcon\Mvc\View\Exception
Phalcon\Mvc\View\Simple
Phalcon\Paginator\AdapterInterface
Phalcon\Paginator\Adapter\Model
Phalcon\Paginator\Adapter\NativeArray
Phalcon\Paginator\Adapter\QueryBuilder
Phalcon\Paginator\Exception
Phalcon\Queue\Beanstalk
Phalcon\Queue\Beanstalk\Job
Phalcon\Security
Phalcon\Security\Exception
Phalcon\Session
Phalcon\Session\Adapter
Phalcon\Session\AdapterInterface
Phalcon\Session\Adapter\Files
Phalcon\Session\Bag
Phalcon\Session\BagInterface
Phalcon\Session\Exception
Phalcon\Tag
Phalcon\Tag\Exception
Phalcon\Tag\Select
Phalcon\Text
Phalcon\Translate\Adapter
Phalcon\Translate\AdapterInterface
Phalcon\Translate\Adapter\NativeArray
Phalcon\Translate\Exception
Phalcon\Validation
Phalcon\Validation\Exception
Phalcon\Validation\Message
Phalcon\Validation\Message\Group
Phalcon\Validation\Validator
Phalcon\Validation\ValidatorInterface
Phalcon\Validation\Validator\Between
Phalcon\Validation\Validator\Confirmation
Phalcon\Validation\Validator\Email
Phalcon\Validation\Validator\ExclusionIn
Phalcon\Validation\Validator\Identical
Phalcon\Validation\Validator\InclusionIn
Phalcon\Validation\Validator\PresenceOf
Phalcon\Validation\Validator\Regex
Phalcon\Validation\Validator\StringLength
Phalcon\Version
# ===========================================================
