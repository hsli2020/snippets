<?php

# ===========================================================
CakePHP CSS

<?php $this->Html->css('myplugin/css/forms','stylesheet', array('inline' => false ) ); ?>

Will output in header: (will add the proper <link> element to <head>)
<link rel="stylesheet" type="text/css" href="/myplugin/css/forms.css" />

<?php echo $this->Html->script('scripts', array('inline' => false )); ?>

# ===========================================================

  function forceDownload($fullpath)
  {
    $filename=basename($fullpath);

    header('Content-Type: "application/octet-stream"');
    header('Content-Disposition: attachment; filename="'.$filename.'"');
    header("Content-Transfer-Encoding: binary");
    header('Expires: 0');
    header('Pragma: no-cache');
    header("Content-Length: ".filesize($fullpath));
    readfile($fullpath);
    //unlink($fullpath);
    die();
  }

# ===========================================================

function doshuffle()
{
  $ord = str_shuffle('12345678');
  $num = mt_rand(1000, 9999);
  $str = $num = $num . sprintf("%04s", 10000 - $num);

  for ($i=0; $i<strlen($ord); $i++) {
    $c = $ord[$i];
    $str[$c-1] = $num[$i];
  }
  //echo $ord.'-'.$str, EOL;
  return $ord.$str;
}

function unshuffle($str)
{
  $ord = substr($str, 0, 8);
  $ret = $date = substr($str, 8);

  for ($i=0; $i<strlen($ord); $i++) {
    $x = $ord[$i];
    $c = $date[$x-1];
    $ret[$i] = $c;
  }
  return $ret;
}

for ($i=0; $i<20; $i++) {
  $s = doshuffle();
  $d = unshuffle($s);

  echo $s, ' ', $d, ' ';
  echo substr($d,0,4) + substr($d,4), EOL;
}

# ===========================================================

function fpr($var, $label='')
{
  $filename = '/tmp/hsli.log';

  if (!empty($label)) {
    error_log($label.' : ', 3, $filename);
    #file_put_contents($filename, $label.' : ', FILE_APPEND);
  }

  $str = print_r($var, true)."\n";
  error_log($str, 3, $filename);
  #file_put_contents($filename, $str, FILE_APPEND);
}

fpr(array(1,2,3,4), 'name');

# not good version
function fpr()
{
  $filename = '/tmp/hsli.log';

  $numargs = func_num_args();
  $args = func_get_args();

  //if ($numargs == 0) return;

  if ($numargs == 1) {
    $var = $args[0];
    $str = print_r($var, true)."\n";
    error_log($str, 3, $filename);
  } else if ($numargs == 2) {
    $label = $args[0];
    $var = $args[1];

    if (!empty($label)) {
      error_log($label.' : ', 3, $filename);
    }

    $str = print_r($var, true)."\n";
    error_log($str, 3, $filename);
  } else {
    foreach ($args as $var) {
      $str = print_r($var, true)."\n";
      error_log($str, 3, $filename);
    }
  }
}

# best version
function fpr()
{
  $filename = '/tmp/cake.log';

  $numargs = func_num_args();
  $args = func_get_args();

  // if ($numargs == 0) return;

  if ($numargs == 2) {
    $label = array_shift($args);
    if (!empty($label)) {
      error_log($label.' : ', 3, $filename);
    }
  }

  foreach ($args as $var) {
    $str = print_r($var, true)."\n";
    error_log($str, 3, $filename);
  }
}

function dpr($var, $label='', $file=false)
{
  $output = function($str, $isLabel) use ($file) {
    if ($file) {
      if ($isLabel)
        $str = $str." : ";
      else
        $str = $str."\n";

      $filename = '/tmp/hsli.log';
      error_log($str, 3, $filename);
    } else {
      if ($isLabel)
        $str = "<b>".$str."</b>\n";
      else
        $str = "<pre>\n".$str."\n</pre>\n";

      echo $str;
    }
  };

  if (!empty($label)) {
    $output($label, true);
  }

  $output(print_r($var, true), false);
}

# ===========================================================

function doshuffle()
{
  $ord = str_shuffle('12345678');
  $str = $date = date('Ymd');

  for ($i=0; $i<strlen($ord); $i++) {
    $c = $ord[$i];
    $str[$c-1] = $date[$i];
  }
  //echo $ord.'-'.$str, EOL;
  return $ord.$str;
}

function unshuffle($str)
{
  $ord = substr($str, 0, 8);
  $ret = $date = substr($str, 8);

  for ($i=0; $i<strlen($ord); $i++) {
    $x = $ord[$i];
    $c = $date[$x-1];
    $ret[$i] = $c;
  }
  return $ret;
}

for ($i=0; $i<80; $i++) {
  $s = doshuffle();
  echo $s, ' ', unshuffle($s), EOL;
}

# ===========================================================

$vid = strtoupper(md5($_SERVER['UNIQUE_ID'].uniqid(rand()).microtime().'hd&#73d)#:"['));

# ===========================================================

  public function approvePhoto() {
    $query = "select id,display from aminno_member_photo where pnum = {$this->user->getPnum()} and approved = 'QUEUED'";
    
    $result = Db::connection(ASHLEY_DB_SLAVE_AMINNO_URI)->query($query);
    if ($result === false) {
      return false;
    }

    while ($row = $result->fetch(PDO::FETCH_ASSOC)) {
      $approve = $row['display'] == 'PRIVATE' ? 'PRIVATE' : 'PUBLIC';
      MemberPhotos::updatePhotoApproval($this->user->getPnum(),$row['id'], $approve);
    } 
  }

# ===========================================================

  public function save($data)
  {
    $this->validate($data);
    $data['created'] = date('Y-m-d H:i:s');

    $stmt = $this->db->prepare(
      "INSERT INTO $this->__table (`label`, `description`, `meta`, `body`, `created`)"
      ." VALUES (:label, :description, :meta, :body, :created)"
    );

    $stmt->bindParam(':label', $data['label']);
    $stmt->bindParam(':description', $data['description']);
    $stmt->bindParam(':meta', $data['meta']);
    $stmt->bindParam(':body', $data['body']);
    $stmt->bindParam(':created', $data['created']);

    if ( $stmt->execute() ) {
      $data['id'] = $this->db->lastInsertId();
      $data['revisions'] = $this->getRevisions($data['label']);
      return $data;
    }
    return false;
  }

# ===========================================================

    header("Content-type: image/png");
    $im = @imagecreate(1, 1);
    imagecolorallocate($im, 255, 255, 255);
    imagepng($im);
    imagedestroy($im);

# ===========================================================

  protected function runhook($data)
  {
    if (!is_array($data))
      return FALSE;

    if (!isset($data['filepath'], $data['filename']))
      return FALSE;

    $filepath = APPPATH.$data['filepath'].'/'.$data['filename'];

    if (!file_exists($filepath))
      return FALSE;

    $class    = FALSE;
    $function = FALSE;
    $params   = '';

    if (!empty($data['class']))
      $class = $data['class'];

    if (!empty($data['function']))
      $function = $data['function'];

    if (isset($data['params']))
      $params = $data['params'];

    if ($class === FALSE && $function === FALSE)
      return FALSE;

    // Call the requested class and/or function

    if ($class !== FALSE) {
      if (!class_exists($class))
        require($filepath);

      $hook = new $class;
      $hook->$function($params);
    }
    else {
      if (!function_exists($function))
        require($filepath);

      $function($params);
    }
    return TRUE;
  }

# ===========================================================

  public function valid_ip($ip)
  {
    return (bool) filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_IPV4);
  }

  function html_escape($var)
  {
    return is_array($var)
      ? array_map('html_escape', $var)
      : htmlspecialchars($var, ENT_QUOTES, config_item('charset'));
  }

if (isset($_SERVER['HTTP_X_REQUESTED_WITH']) && $_SERVER['HTTP_X_REQUESTED_WITH']=="XMLHttpRequest") {
  define('AJAX_REQUEST', 1);
} else {
  define('AJAX_REQUEST', 0);
}

# ===========================================================

//Require the config file for this site name
require(SITE_NAME. '/config.php');

//Require the config file for this server
require_once(SITE_NAME. '/server.php');

//Include the common file
require_once(INCLUDES_DIR. 'common.php');

# ===========================================================

  mb_internal_encoding('UTF-8');

  public function convert_to_utf8($str, $encoding)
  {
    if (function_exists('iconv'))
      return @iconv($encoding, 'UTF-8', $str);
    elseif (function_exists('mb_convert_encoding'))
      return @mb_convert_encoding($str, 'UTF-8', $encoding);

    return FALSE;
  }
  protected function _is_ascii($str)
  {
    return (preg_match('/[^\x00-\x7F]/S', $str) === 0);
  }
  public function clean_string($str)
  {
    if (_is_ascii($str) === FALSE)
      $str = @iconv('UTF-8', 'UTF-8//IGNORE', $str);

    return $str;
  }

  if (@preg_match('/./u', 'é') === 1    // PCRE must support UTF-8
    && function_exists('iconv')     // iconv must be installed
    && @ini_get('mbstring.func_overload') != 1  // Multibyte string function overloading cannot be enabled
    )
    define('UTF8_ENABLED', TRUE);

# ===========================================================

//Set the unique name of the current page
$var = preg_replace("/([^a-z0-9_\-\.]+)/i", '', $_SERVER["REQUEST_URI"]);
define('PAGE_NAME', ($var ? $var : 'index'));
--------------------------------------------------------------------------------
//Set the unique name of the current page
$var = preg_replace("/([^a-z0-9_\-\.]+)/i", '', $_SERVER["REQUEST_URI"]);
define('PAGE_NAME', ($var ? $var : 'index'));
--------------------------------------------------------------------------------
// Get the Site Name: www.site.com -also protects from XSS/CSFR attacks
preg_match('/(?=([a-z]+:\/\/)?)(([a-z0-9\-]{1,70}\.){1,5}([a-z]{2,4}))|localhost/i',
($_SERVER["SERVER_NAME"] ? $_SERVER["SERVER_NAME"] : $_SERVER['HTTP_HOST']), $matches);
define('SITE_NAME', $matches[0]);

# ===========================================================

/**
 * Enable or Disable caching for this site
 * Set to FALSE to disable cacheing
 * Set to a number (in seconds) to enable:
 * 60 * 2 = 2 minutes
 */
define('CACHING', false);

/* FILE SYSTEM PATHS */

// Absolute file system path to the root (WEBROOT/ROOT_DIR?)
define('SITE_DIR', rtrim(realpath(dirname(__FILE__). "/../"), '/\\'). '/');
define('THEME_DIR', SITE_DIR. 'themes/'. $config['theme']. '/');
define('INCLUDES_DIR', SITE_DIR. "includes/");
define('LIBRARIES_DIR', SITE_DIR. "libraries/");
define('MODELS_DIR', SITE_DIR. "models/");
define('FUNCTIONS_DIR', SITE_DIR. "functions/");
define('UPLOAD_DIR', SITE_DIR. 'uploads/');
define('CACHE_DIR', SITE_DIR. 'cache/');

/* URL ADDRESS PATHS */

// Absolute URL path to the system root
// Leave blank unless this site is in a subfolder.
define('SITE_PATH', '/MicroMVC/');
define('THEME_PATH', SITE_PATH. 'themes/'. $config['theme']. '/');
define('UPLOAD_PATH', SITE_PATH. 'uploads/');
define('CACHE_PATH', SITE_PATH. 'cache/');

define('DEBUG_MODE', true);
define('APPMODE', 'dev');
define('APPENV', 'dev');

define('WEB_ROOT', rtrim(realpath(dirname(__FILE__). "/../"), '/\\'). '/');
define('SITE_DIR', rtrim(realpath(dirname(__FILE__). "/../"), '/\\'). '/');
define('THEME_DIR', SITE_DIR. 'themes/'. $config['theme']. '/');
define('INCLUDES_DIR', SITE_DIR. "includes/");
define('LIBRARIES_DIR', SITE_DIR. "libraries/");
define('MODELS_DIR', SITE_DIR. "models/");
define('FUNCTIONS_DIR', SITE_DIR. "functions/");
define('UPLOAD_DIR', SITE_DIR. 'uploads/');
define('CACHE_DIR', SITE_DIR. 'cache/');

# ===========================================================

  //Try and find a temporary directory
  function _tmp_dir(){
    if(function_exists("sys_get_temp_dir")){
      return sys_get_temp_dir();
    }elseif(is_writable('/tmp')){
      return '/tmp';
    }elseif(is_writable('c:\Windows\Temp')){
      return 'c:\Windows\Temp';
    }else{
      //Default to the current directory
      return './hsapi_cache';
    }
  }

# ===========================================================

$content = file_get_contents($url);
$content = iconv("gbk", "utf-8", $content);

使用mb_convert_encoding()函数进行转换编码,iconv有些字符会转换失败.

# ===========================================================

用PHP做模板文件，在缓存中生成的内容是HTML。
用独立的模板(Smarty、Twig等)系统，在缓存中能生成PHP代码。
难道这就是独立模板系统存在的原因？

# ===========================================================

var winks = {
    "Wink Message": "[Wink Message]",
    "What Gets Me Hot": "(What Gets Me Hot)",
};

function translate(str) {
    var x;
    for (x in winks) {
        str = str.replace(new RegExp(x, 'g'), winks[x]);
    }
}

# ===========================================================

ZF中非常喜欢使用三个继承层次的类结构定义，以Zend_View类为例
Zend_View_Interface
    Zend_View_Abstract
        Zend_View

# ===========================================================
REGEX for DATE

array(
    'year' => '[12][0-9]{3}',
    'month' => '0[1-9]|1[012]',
    'day' => '0[1-9]|[12][0-9]|3[01]'
)

# ===========================================================

Initialize ACL with this command. 
This command will create table ‘acos’, ‘aros’ and ‘aros_acos’.

$ cake schema run create DbAcl

This is permission tree I will create.

- controllers
    - notes
        - add
        - edit
        - delete
    - users
        - list
        - add
        - edit
        - delete
    - groups
        - list
        - add
        - edit
        - delete

Execute this following code to create ACO list.

$ cake acl create aco root controllers
$ cake acl create aco controllers Users
$ cake acl create aco Users add
$ cake acl create aco Users edit
$ cake acl create aco Users delete
$ cake acl create aco controllers Groups
$ cake acl create aco Groups add
$ cake acl create aco Groups edit
$ cake acl create aco Groups delete
$ cake acl create aco controllers Notes
$ cake acl create aco Notes add
$ cake acl create aco Notes edit
$ cake acl create aco Notes delete

After we have ARO and ACO, it’s time to connect both of them. 
We can connect a user or a group to list of permissions.

Grant permission Group 1 (Admin group) to all controller actions.
$ cake acl grant Group.1 controllers all

Grant permission Group 2 (Manager group) to all notes controller actions.
$ cake acl grant Group.2 Notes all

# ===========================================================
问题：

1、users表中的group_id字段是必须的吗？起什么作用？如果这个字段是必须的，那是否表明每个用户只能属于一个组，而不能属于多个组？

2、一个用户能否属于多个组？从aros表的结构看似乎是可以的，但users表的group_id字段似乎又表明不可以。(user表和group表与acl有关吗？)

3、ACL只能控制到controller和action级别吗？能否用来控制其它对象？能否做到更精细的权限控制？ACOs could be anything you want to control

4、怎样在不同的地方(MVC)进行权限判断？特别是针对上面的2和3该怎么做？

# ===========================================================

ACL or Access Control List is security concept about permission. It’s contain a list of permission to access some of resources. It specify who and what is allowed to access something.

In CakePHP ACL is used to specify users access to a controller, or specific action of a controller. Before we get to the how to section, there are a few terms we have to understand first.

- Access Request Object (ARO) is defined who want to access or request object. In simple term, it will be a user or group. This list of user or group is stored in table called ‘aros’.

- Access Control Object (ACO) is defined object that is being protected from a user or a group. In other word ACO related with which controller or action of controller you want to protect. ACO is stored in table called ‘acos’.

‘acos’ and ‘aros’ table is created automatically when you initialize the DB Acl tables. Along with those table, there is table called ‘aros_acos’ which specify relationship between ARO and ACO. In english it will be mean which user has permission to a controller or action of controller.

Essentially, ACL is what is used to decide when an ARO can have access to an ACO.

ACL is most usually implemented in a tree structure. There is usually a tree of AROs and a tree of ACOs. By organizing your objects in trees, permissions can still be dealt out in a granular fashion, while still maintaining a good grip on the big picture. 

Cake’s first ACL implementation was based on INI files stored in the Cake installation. While it’s useful and stable, we recommend that you use the database backed ACL solution, mostly because of its ability to create new ACOs and AROs on the fly. We meant it for usage in simple applications - and especially for those folks who might not be using a database for some reason.

By default, CakePHP’s ACL is database-driven. To enable INI-based ACL, you’ll need to tell CakePHP what system you’re using by updating the following lines in app/Config/core.php

// Change these lines:
Configure::write(’Acl.classname’, ’DbAcl’);
Configure::write(’Acl.database’, ’default’);

// To look like this:
Configure::write(’Acl.classname’, ’IniAcl’); 
//Configure::write(’Acl.database’, ’default’);

ARO/ACO permissions are specified in /app/Config/acl.ini.php. The basic idea is that AROs are specified in an INI section that has three properties: groups, allow, and deny.

• groups: names of ARO groups this ARO is a member of.
• allow: names of ACOs this ARO has access to
• deny: names of ACOs this ARO should be denied access to

ACOs are specified in INI sections that only include the allow and deny properties.

each ACO automatically contains four properties related to CRUD (create, read, update, and delete) actions.

you can add a column in the aros_acos database table (prefixed with _ - for example _admin) and use it alongside the defaults.

# ===========================================================

$user = $this->Auth->user();
pr($user);

you can know a user is logged in if Auth’s user() array is not empty.

Now, Auth is a component, meant to be used in a controller. We need to know from a view if a user is logged in, preferably via a helper. How can we use a component inside a helper? CakePHP is so awesome and flexible that it is possible.

<?php
class AccessHelper extends Helper{   
    var $helpers = array("Session");   
   
    function isLoggedin(){   
        App::import('Component', 'Auth');   
        $auth = new AuthComponent();   
        $auth->Session = $this->Session;   
        $user = $auth->user();   
        return !empty($user);   
    }   
?>   

Save this snippet in views/helpers as access.php. 

Helpers can include other helpers, just like $components can. 

For Auth and Acl to work properly we need to associate our users and groups to rows in the Acl tables. 


# ===========================================================
CakePHP ACL

CREATE TABLE acos (
  id INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  parent_id INTEGER(10) DEFAULT NULL,
  model VARCHAR(255) DEFAULT '',
  foreign_key INTEGER(10) UNSIGNED DEFAULT NULL,
  alias VARCHAR(255) DEFAULT '',
  lft INTEGER(10) DEFAULT NULL,
  rght INTEGER(10) DEFAULT NULL,
  PRIMARY KEY  (id)
);

CREATE TABLE aros_acos (  // permissions to User & Group
  id INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  aro_id INTEGER(10) UNSIGNED NOT NULL,
  aco_id INTEGER(10) UNSIGNED NOT NULL,
  _create CHAR(2) NOT NULL DEFAULT 0,
  _read CHAR(2) NOT NULL DEFAULT 0,
  _update CHAR(2) NOT NULL DEFAULT 0,
  _delete CHAR(2) NOT NULL DEFAULT 0,
  PRIMARY KEY(id)
);

CREATE TABLE aros ( // User & Group
  id INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  parent_id INTEGER(10) DEFAULT NULL, // belong to Group
  model VARCHAR(255) DEFAULT '',  // User or Group? (class Name)
  foreign_key INTEGER(10) UNSIGNED DEFAULT NULL, // which User/Group
  alias VARCHAR(255) DEFAULT '',
  lft INTEGER(10) DEFAULT NULL,
  rght INTEGER(10) DEFAULT NULL,
  PRIMARY KEY  (id)
);

# ===========================================================

Each permission will be given to the group, not to the user; so if a user is promoted to Admin, he will be checked against the group permission – not his. These roles and child nodes (users) are called Access Requests Objects, or AROs.

Now, on the other side, we have the Access Control Objects, or ACOs. These are the objects to be controlled. Above I mentioned Posts and Users. Normally, these objects are directly linked with the models, so if we have a Post model, we will need an ACO for this model.

Each ACO has four basic permissions: create, read, update and delete. You can remember them with the keyword CRUD. There’s a fifth permission, the asterisk, that is a shortcut for full access.

By simply including the Auth component, all the actions are by default, denied for guests. We need to deny three actions for unauthorized users: create, edit and delete. In other terms, we’ll have to allow index and view.

# ===========================================================

preprocess is now called import

public function preprocess()
{
  $rawfile = DATA_DIR.'\\DH-ITEMLIST';
  $in = fopen($rawfile, 'r');
  if (empty($in)) {
    error("Failed to open file $rawfile");
    return;
  }

  $newfile = DATA_DIR.'\\pricelist-DH.csv';
  $out = fopen($newfile, 'w');
  if (empty($out)) {
    error("Failed to create file $newfile");
    return;
  }

  // There is no title in DH-ITEMLIST
  // fgetcsv($in, 0, '|');

  fputcsv($out, array('sku', 'price', 'instock'));

  while (($fields = fgetcsv($in, 0, '|')) !== false) {
    // Just In Case
    if (count($fields) < 3)
      continue;

    // get useful fields
    $sku     = $fields[4];
    $price   = tidyPrice($fields[9]);
    $inStock = $fields[1] <= 0 ? 'N' : 'Y';
    $rebate  = $fields[2];

    // end rebate earlier, restore price to regular price
    if ($rebate == 'y') {
      $today = strtotime(date('m/d/y'));
      $rebateEndDate = strtotime($fields[3]);
      $days = ($rebateEndDate - $today) / (60*60*24);
      if ($days <= 2)
        $price += tidyPrice($fields[10]);
    }

    $data = array($sku, $price, $inStock);
    fputcsv($out, $data);
  }
  println("File $newfile created");
  fclose($in);
  fclose($out);
}


tidyPrice should return -1 when price is not a valid number (bad format)

# ===========================================================

	header("Date: " . date("D, j M Y G:i:s ", $templateModified) . 'GMT');
	header("Content-Type: text/css");
	header("Expires: " . gmdate("D, d M Y H:i:s", time() + DAY) . " GMT");
	header("Cache-Control: max-age=86400, must-revalidate"); // HTTP/1.1
	header("Pragma: cache");        // HTTP/1.0
	print $output;

ini_set('include_path',
  ini_get('include_path') 
  . PATH_SEPARATOR . CAKE_CORE_INCLUDE_PATH
  . PATH_SEPARATOR . ROOT . DS . APP_DIR . DS
);

	define('SECOND', 1);
	define('MINUTE', 60);
	define('HOUR', 3600);
	define('DAY', 86400);
	define('WEEK', 604800);
	define('MONTH', 2592000);
	define('YEAR', 31536000);

	function getMicrotime() {
		list($usec, $sec) = explode(' ', microtime());
		return ((float)$usec + (float)$sec);
	}

# ===========================================================

function backupLogFile($filename)
{
    $maxNum = 5;

    //if (file_exists($filename) && date('Ymd', filemtime($filename)) == date('Ymd')) {
    if (file_exists($filename) && filesize($filename) >= 100*1024*1024) {

        if (file_exists("$filename.$maxNum")) {
            unlink("$filename.$maxNum");
        }

        for ($i=$maxNum; $i>1; $i--) {
            $oldname = $filename.".".($i-1);
            $newname = $filename.".".$i;

            if (file_exists($oldname))
                rename($oldname, $newname);
        }

        rename($filename, "$filename.1");
    }
}

backupLogFile("tt.log");
touch("tt.log");

# ===========================================================
function vpost($url,$data,$cookie){ // 模拟提交数据函数
    $curl = curl_init(); // 启动一个CURL会话
    curl_setopt($curl, CURLOPT_URL, $url); // 要访问的地址
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, 0); // 对认证证书来源的检查
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 1); // 从证书中检查SSL加密算法是否存在
    curl_setopt($curl, CURLOPT_USERAGENT, $_SERVER['HTTP_USER_AGENT']); // 模拟用户使用的浏览器
    curl_setopt($curl, CURLOPT_COOKIE, $cookie);
    curl_setopt($curl, CURLOPT_REFERER,'https://www.baidu.com');// 设置Referer
    curl_setopt($curl, CURLOPT_POST, 1); // 发送一个常规的Post请求
    curl_setopt($curl, CURLOPT_POSTFIELDS, $data); // Post提交的数据包
    curl_setopt($curl, CURLOPT_TIMEOUT, 30); // 设置超时限制防止死循环
    curl_setopt($curl, CURLOPT_HEADER, 0); // 显示返回的Header区域内容
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1); // 获取的信息以文件流的形式返回
    $tmpInfo = curl_exec($curl); // 执行操作
    if (curl_errno($curl)) {
       echo 'Errno'.curl_error($curl);//捕抓异常
    }
    curl_close($curl); // 关闭CURL会话
    return $tmpInfo; // 返回数据
}
# ===========================================================
    /**
     * Get UTF-8 CJK character ranges
     *
     * @return array of UTF-8 CJK character ranges
     */
    protected function getCJKRanges()
    {
        return array(
            "[\x{2E80}-\x{2EFF}]",      # CJK Radicals Supplement
            "[\x{2F00}-\x{2FDF}]",      # Kangxi Radicals
            "[\x{2FF0}-\x{2FFF}]",      # Ideographic Description Characters
            "[\x{3000}-\x{303F}]",      # CJK Symbols and Punctuation
            "[\x{3040}-\x{309F}]",      # Hiragana
            "[\x{30A0}-\x{30FF}]",      # Katakana
            "[\x{3100}-\x{312F}]",      # Bopomofo
            "[\x{3130}-\x{318F}]",      # Hangul Compatibility Jamo
            "[\x{3190}-\x{319F}]",      # Kanbun
            "[\x{31A0}-\x{31BF}]",      # Bopomofo Extended
            "[\x{31F0}-\x{31FF}]",      # Katakana Phonetic Extensions
            "[\x{3200}-\x{32FF}]",      # Enclosed CJK Letters and Months
            "[\x{3300}-\x{33FF}]",      # CJK Compatibility
            "[\x{3400}-\x{4DBF}]",      # CJK Unified Ideographs Extension A
            "[\x{4DC0}-\x{4DFF}]",      # Yijing Hexagram Symbols
            "[\x{4E00}-\x{9FFF}]",      # CJK Unified Ideographs
            "[\x{A000}-\x{A48F}]",      # Yi Syllables
            "[\x{A490}-\x{A4CF}]",      # Yi Radicals
            "[\x{AC00}-\x{D7AF}]",      # Hangul Syllables
            "[\x{F900}-\x{FAFF}]",      # CJK Compatibility Ideographs
            "[\x{FE30}-\x{FE4F}]",      # CJK Compatibility Forms
            "[\x{1D300}-\x{1D35F}]",    # Tai Xuan Jing Symbols
            "[\x{20000}-\x{2A6DF}]",    # CJK Unified Ideographs Extension B
            "[\x{2F800}-\x{2FA1F}]"     # CJK Compatibility Ideographs Supplement
        );
    }
# ===========================================================
    private static function isPublicIp($ip)
    {
        return filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE);
    }
# ===========================================================
remap of Controller

    /**
     * remaps the coming requested page if the permission is valid
     * 
     * @param string $method the required method to be executed
     * @param array $params the rest of URI string segments as array
     */
    public function _remap($method, $params = array())
    {
        if (!perm_chck($this->perm) )
            show_error(lang('system_permission_denied'));
        
        $this->page = array_key_exists($method,$this->pages)
                        ? $this->pages[$method] : $method;
                        
        if (method_exists($this, $method))
        {
            return call_user_func_array(array($this, $method), $params);
        }
        show_404();
    }
# ===========================================================
CodeIgniter constants

SELF=index.php
BASEPATH=C:/xampp/htdocs/bte/system/
FCPATH=C:\xampp\htdocs\bte\
SYSDIR=system
APPPATH=app/
# ===========================================================
Magento Classes

Mage_Core_Model_Resource

Mage_Sales_Model_Order
Mage_Sales_Model_Order_Item

Mage_Sales_Model_Resource_Order
Mage_Sales_Model_Resource_Order_Collection

Mage_Sales_Model_Resource_Order_Item
Mage_Sales_Model_Resource_Order_Item_Collection
# ===========================================================
function _getVariablesInContent($s) 
{
    if (preg_match_all("/{:(.*?):}/si", $s, $m)) {
       return array_unique($m[1]); // some variables appear more than once.
    }
    return array();
}
print_r(_getVariablesInContent("{:a1:}{:b2:}{:c3:}{:d4:}{:e5:}{:a1:}"));

$pattern[] = '{:' . $name . ':}';
$replacement[] = $value;
$content = str_replace($pattern, $replacement, $content);

function errlog($var) { 
    error_log(print_r($var, true), 3, "/tmp/errlog"); 
}
# ===========================================================
RewriteEngine On
RewriteRule ^tracker.jpg$ index.php?tracker

RewriteEngine On
RewriteRule .* - [E=HTTP_IF_MODIFIED_SINCE:%{HTTP:If-Modified-Since}]
RewriteRule .* - [E=HTTP_IF_NONE_MATCH:%{HTTP:If-None-Match}]
# ===========================================================
var_dump() 能自适应命令行和web环境，在不同环境显示不同格式，在命令行显示纯文本，在web显示html，用来调试比print_r和var_export都方便。
# ===========================================================
cake dir tree

./app
./app/config
./app/vendors

./app/models
./app/views
./app/controllers

./app/libs
./app/locale/en/LC_MESSAGES
./app/plugins
./app/import

./app/tmp/cache
./app/tmp/logs
./app/tmp/sessions

./app/webroot
./app/webroot/css
./app/webroot/files
./app/webroot/img
./app/webroot/js

./cake  // ./core or ./sys
./cake/config
./cake/libs
./cake/libs/log
./cake/libs/cache
./cake/libs/model
./cake/libs/view
./cake/libs/controller

./export
./plugins
./vendors
# ===========================================================
    public static $statusTexts = array(
        100 => 'Continue',
        101 => 'Switching Protocols',
        102 => 'Processing',            // RFC2518
        200 => 'OK',
        201 => 'Created',
        202 => 'Accepted',
        203 => 'Non-Authoritative Information',
        204 => 'No Content',
        205 => 'Reset Content',
        206 => 'Partial Content',
        207 => 'Multi-Status',          // RFC4918
        208 => 'Already Reported',      // RFC5842
        226 => 'IM Used',               // RFC3229
        300 => 'Multiple Choices',
        301 => 'Moved Permanently',
        302 => 'Found',
        303 => 'See Other',
        304 => 'Not Modified',
        305 => 'Use Proxy',
        306 => 'Reserved',
        307 => 'Temporary Redirect',
        308 => 'Permanent Redirect',    // RFC-reschke-http-status-308-07
        400 => 'Bad Request',
        401 => 'Unauthorized',
        402 => 'Payment Required',
        403 => 'Forbidden',
        404 => 'Not Found',
        405 => 'Method Not Allowed',
        406 => 'Not Acceptable',
        407 => 'Proxy Authentication Required',
        408 => 'Request Timeout',
        409 => 'Conflict',
        410 => 'Gone',
        411 => 'Length Required',
        412 => 'Precondition Failed',
        413 => 'Request Entity Too Large',
        414 => 'Request-URI Too Long',
        415 => 'Unsupported Media Type',
        416 => 'Requested Range Not Satisfiable',
        417 => 'Expectation Failed',
        418 => 'I\'m a teapot',                                               // RFC2324
        422 => 'Unprocessable Entity',                                        // RFC4918
        423 => 'Locked',                                                      // RFC4918
        424 => 'Failed Dependency',                                           // RFC4918
        425 => 'Reserved for WebDAV advanced collections expired proposal',   // RFC2817
        426 => 'Upgrade Required',                                            // RFC2817
        428 => 'Precondition Required',                                       // RFC6585
        429 => 'Too Many Requests',                                           // RFC6585
        431 => 'Request Header Fields Too Large',                             // RFC6585
        500 => 'Internal Server Error',
        501 => 'Not Implemented',
        502 => 'Bad Gateway',
        503 => 'Service Unavailable',
        504 => 'Gateway Timeout',
        505 => 'HTTP Version Not Supported',
        506 => 'Variant Also Negotiates (Experimental)',                      // RFC2295
        507 => 'Insufficient Storage',                                        // RFC4918
        508 => 'Loop Detected',                                               // RFC5842
        510 => 'Not Extended',                                                // RFC2774
        511 => 'Network Authentication Required',                             // RFC6585
    );

    protected $_mimeTypes = array(
        'txt' => 'text/plain',
        'html' => 'text/html',
        'xhtml' => 'application/xhtml+xml',
        'xml' => 'application/xml',
        'css' => 'text/css',
        'js' => 'application/javascript',
        'json' => 'application/json',
        'csv' => 'text/csv',

        // images
        'png' => 'image/png',
        'jpe' => 'image/jpeg',
        'jpeg' => 'image/jpeg',
        'jpg' => 'image/jpeg',
        'gif' => 'image/gif',
        'bmp' => 'image/bmp',
        'ico' => 'image/vnd.microsoft.icon',
        'tiff' => 'image/tiff',
        'tif' => 'image/tiff',
        'svg' => 'image/svg+xml',
        'svgz' => 'image/svg+xml',

        // archives
        'zip' => 'application/zip',
        'rar' => 'application/x-rar-compressed',

        // adobe
        'pdf' => 'application/pdf'
    );

    public function isMobile()
    {
        $op = strtolower($_SERVER['HTTP_X_OPERAMINI_PHONE']);
        $ua = strtolower($_SERVER['HTTP_USER_AGENT']);
        $ac = strtolower($_SERVER['HTTP_ACCEPT']);

        return (
            strpos($ac, 'application/vnd.wap.xhtml+xml') !== false
            || strpos($ac, 'text/vnd.wap.wml') !== false
            || $op != ''
            || strpos($ua, 'iphone') !== false
            || strpos($ua, 'android') !== false
            || strpos($ua, 'iemobile') !== false 
            || strpos($ua, 'kindle') !== false
            || strpos($ua, 'sony') !== false 
            || strpos($ua, 'symbian') !== false 
            || strpos($ua, 'nokia') !== false 
            || strpos($ua, 'samsung') !== false 
            || strpos($ua, 'mobile') !== false
            || strpos($ua, 'windows ce') !== false
            || strpos($ua, 'epoc') !== false
            || strpos($ua, 'opera mini') !== false
            || strpos($ua, 'nitro') !== false
            || strpos($ua, 'j2me') !== false
            || strpos($ua, 'midp-') !== false
            || strpos($ua, 'cldc-') !== false
            || strpos($ua, 'netfront') !== false
            || strpos($ua, 'mot') !== false
            || strpos($ua, 'up.browser') !== false
            || strpos($ua, 'up.link') !== false
            || strpos($ua, 'audiovox') !== false
            || strpos($ua, 'blackberry') !== false
            || strpos($ua, 'ericsson,') !== false
            || strpos($ua, 'panasonic') !== false
            || strpos($ua, 'philips') !== false
            || strpos($ua, 'sanyo') !== false
            || strpos($ua, 'sharp') !== false
            || strpos($ua, 'sie-') !== false
            || strpos($ua, 'portalmmm') !== false
            || strpos($ua, 'blazer') !== false
            || strpos($ua, 'avantgo') !== false
            || strpos($ua, 'danger') !== false
            || strpos($ua, 'palm') !== false
            || strpos($ua, 'series60') !== false
            || strpos($ua, 'palmsource') !== false
            || strpos($ua, 'pocketpc') !== false
            || strpos($ua, 'smartphone') !== false
            || strpos($ua, 'rover') !== false
            || strpos($ua, 'ipaq') !== false
            || strpos($ua, 'au-mic,') !== false
            || strpos($ua, 'alcatel') !== false
            || strpos($ua, 'ericy') !== false
            || strpos($ua, 'up.link') !== false
            || strpos($ua, 'vodafone/') !== false
            || strpos($ua, 'wap1.') !== false
            || strpos($ua, 'wap2.') !== false
        );
    }
# ===========================================================
function code62($x) {
    $show = '';
    while($x > 0) {
        $s = $x % 62;
        if ($s > 35) {
            $s = chr($s+61);
        } elseif ($s > 9 && $s <=35) {
            $s = chr($s + 55);
        }
        $show .= $s;
        $x = floor($x/62);
    }
    return $show;
}
   
function shorturl($url) {
    $url = crc32($url);
    $result = sprintf("%u", $url);
    //return $url;
    //return $result;
    return code62($result);
}
 
echo shorturl("http://www.website.com");

base_convert($i, 10, 36);
echo base_convert(md5($key), 16, 36);
echo base_convert(sha1($key), 16, 36);

$val = base_convert(md5($url), 16, 10);
echo code62(substr($val, 0, 20)), EOL;
# ===========================================================
function validEmailAddress($mail) {
  return (bool)filter_var($mail, FILTER_VALIDATE_EMAIL);
}
# ===========================================================
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
# ===========================================================
Api::get('ashley.user')

$memberCache = PINF::CALL('GetForMember',$Pnum,
    array(AM__MEMBER_GENDER,AM__MEMBER_ISHOST));

两种方式有那么一点点相似之处
# ===========================================================
SiteConfig
SiteInfo
Redis::connect($options = array())
# ===========================================================
user.account = /User/Account
user.profile = /User/Profile
# ===========================================================
/**
 * Color output text for the CLI
 *
 * @param string $text to color
 * @param string $color of text
 * @param string $background color
 */
function colorize($text, $color, $bold = FALSE)
{
	// Standard CLI colors
	$colors = array_flip(array(30 => 'gray', 'red', 'green', 'yellow', 'blue', 'purple', 'cyan', 'white', 'black'));

	// Escape string with color information
	return"\033[" . ($bold ? '1' : '0') . ';' . $colors[$color] . "m$text\033[0m";
}

/**
 * Format the given string using the current system locale
 * Basically, it's sprintf on i18n steroids.
 *
 * @param string $string to parse
 * @param array $params to insert
 * @return string
 */
function __($string, array $params = NULL)
{
	return msgfmt_format_message(setlocale(LC_ALL, 0), $string, $params);
}

/**
 * Make a request to the given URL using cURL.
 *
 * @param string $url to request
 * @param array $options for cURL object
 * @return object
 */
function curl_request($url, array $options = NULL)
{
	$ch = curl_init($url);

	$defaults = array(
		CURLOPT_HEADER => 0,
		CURLOPT_RETURNTRANSFER => 1,
		CURLOPT_TIMEOUT => 5,
	);

	// Connection options override defaults if given
	curl_setopt_array($ch, (array) $options + $defaults);

	// Create a response object
	$object = new stdClass;

	// Get additional request info
	$object->response = curl_exec($ch);
	$object->error_code = curl_errno($ch);
	$object->error = curl_error($ch);
	$object->info = curl_getinfo($ch);

	curl_close($ch);

	return $object;
}

/**
 * Attach (or remove) multiple callbacks to an event and trigger those callbacks when that event is called.
 *
 * @param string $event name
 * @param mixed $value the optional value to pass to each callback
 * @param mixed $callback the method or function to call - FALSE to remove all callbacks for event
 */
function event($event, $value = NULL, $callback = NULL)
{
	static $events;

	// Adding or removing a callback?
	if($callback !== NULL)
	{
		if($callback) {
			$events[$event][] = $callback;
		}
		else {
			unset($events[$event]);
		}
	}
	elseif(isset($events[$event])) // Fire a callback
	{
		foreach($events[$event] as $function) {
			$value = call_user_func($function, $value);
		}
		return $value;
	}
}


/**
 * Fetch a config value from a module configuration file
 *
 * @param string $file name of the config
 * @param boolean $clear to clear the config object
 * @return object
 */
function config($file = 'Config', $clear = FALSE)
{
	static $configs = array();

	if($clear) {
		unset($configs[$file]);
		return;
	}

	if(empty($configs[$file])) {
		//$configs[$file] = new \Core\Config($file);
		require(SP . 'Config/' . $file . EXT);
		$configs[$file] = (object) $config;
		//print dump($configs);
	}

	return $configs[$file];
}
# ===========================================================
interface Doctrine\DBAL\Driver\Connection
{
    function prepare($prepareString);
    function query();
    function quote($input, $type=\PDO::PARAM_STR);
    function exec($statement);
    function lastInsertId($name = null);
    function beginTransaction();
    function commit();
    function rollBack();
    function errorCode();
    function errorInfo();
}
# ===========================================================
两种定义常量的方式

const COUNTRY_US = 'US';
Country::US

const CURRENCY_USD = 'USD';
Currency::USD

const GENDER_MALE = 1;
const GENDER_FEMALE = 2;
const GENDER_MALE_STR = 'M';
const GENDER_FEMALE_STR = 'F';

class Gender
{
    const MALE = 1; // GENDER_MALE
    const FEMALE = 2; // GENDER_FEMALE

    const MALE_STR = 'M'; // GENDER_MALE_STR
    const FEMALE_STR = 'F'; // GENDER_FEMALE_STR
}

// getGenderFromSeeking($seeking)

class Seeking
{
}

const COUNTRY_US = 1;
const COUNTRY_CA = 2;

class Country
{
}

const CURRENCY_USD = 1;
const CURRENCY_CAD = 2;

class Currency
{
}
# ===========================================================
Api::get('ashley.account.current')->getPnum();
$app['ashley.account.current']->getPnum();

类似这样的也存在一个缺陷，无法像CodeIgniter的Controller类那样，利用__get()转发所有的成员变量的访问到全局容器中。
$this->currentAccount->pnum

也就是说，'ashley.account.current'这样的方式不是十分好。
(不过这个缺点不算很严重，很容易避免，只需添加一个类成员即可
$this->currentAccount = Api::get('ashley.account.current');
)
# ===========================================================

protected $directory; // The cache directory
protected $extension; // The cache file extension

protected function getFilename($id)
{
    $hash = hash('sha256', $id);
    $path = implode(str_split($hash, 16), DIRECTORY_SEPARATOR);
    $path = $this->directory . DIRECTORY_SEPARATOR . $path;
    $id   = preg_replace('@[\\\/:"*?<>|]+@', '', $id);

    return $path . DIRECTORY_SEPARATOR . $id . $this->extension;
}
# ===========================================================
CodeIgniter的Router

过于简单了
	/**
	 *  Parse Routes
	 *
	 * This function matches any routes that may exist in
	 * the config/routes.php file against the URI to
	 * determine if the class/method need to be remapped.
	 *
	 * @access	private
	 * @return	void
	 */
	function _parse_routes()
	{
		// Turn the segment array into a URI string
		$uri = implode('/', $this->uri->segments);

		// Is there a literal match?  If so we're done
		if (isset($this->routes[$uri]))
		{
			return $this->_set_request(explode('/', $this->routes[$uri]));
		}

		// Loop through the route array looking for wild-cards
		foreach ($this->routes as $key => $val)
		{
			// Convert wild-cards to RegEx
			$key = str_replace(':any', '.+', str_replace(':num', '[0-9]+', $key));

			// Does the RegEx match?
			if (preg_match('#^'.$key.'$#', $uri))
			{
				// Do we have a back-reference?
				if (strpos($val, '$') !== FALSE AND strpos($key, '(') !== FALSE)
				{
					$val = preg_replace('#^'.$key.'$#', $val, $uri);
				}

				return $this->_set_request(explode('/', $val));
			}
		}

		// If we got this far it means we didn't encounter a
		// matching route so we'll set the site default route
		$this->_set_request($this->uri->segments);
	}
# ===========================================================
Safe UserID in Cookie

$iv     = "1234567812345678";
$pass   = '1234pswd5678';
$method = 'AES-256-CBC';
$method = 'aes128';

$string = json_encode(array(1234567, time(), crc32($_SERVER['REMOTE_ADDR'].$_SERVER['HTTP_USER_AGENT'])));
$string = '[1234567,1409258856,2425597808]';

#serialize generate bigger data than json_encode
#echo serialize(array(time(),'name'=>'user')), EOL;

echo $string, EOL;

$string = openssl_encrypt($string, $method, md5($pass), 0, $iv);
echo $string, EOL;

$string = openssl_decrypt($string, $method, md5($pass), 0, $iv);
echo $string, EOL;
# ===========================================================
array_merge_simple

$a = array(
    2000006 => array (
        'id' => 2000006,
        'credits' => 0,
        'flat_rate' => 0,
    )
);

$b = array (
    2000006 => array (
        'id' => 2000006,
        'accept_collect' => 1,
        'view_all_photos' => 0,
        'im_enabled' => 1,
    )
);

$c = array (
    2000006 => array(
        'id' => 2000006,
        'global_restricted' => 1,
        'broadcast_restricted' => 1,
    )
);

function array_merge_simple()
{
    $arrays = func_get_args();

    $keys = array();
    foreach ($arrays as $array) {
        $keys = array_merge($keys, array_keys($array));
    }
    $keys = array_unique($keys);

    $result = array();
    foreach ($keys as $key) {
        $merged = array();
        foreach ($arrays as $array) {
            if (is_array($array) && array_key_exists($key, $array)) {
                $merged = array_merge($merged, $array[$key]);
            }
        }
        $result[$key] = $merged;
    }
    return $result;
}

$x = array_merge_recursive($a, $b, $c);
$x = array_merge_simple($a, $b, $c);
print_r($x);

// string-key vs number-key

$a = array(100 => array('red', 'blue', 'white'));
$b = array(100 => array('apple', 'banana', 'pear'));
$x = array_merge_recursive($a, $b);
pr($x);

$a = array('abc' => array('red', 'blue', 'white'));
$b = array('abc' => array('apple', 'banana', 'pear'));
$x = array_merge_recursive($a, $b);
pr($x);
# ===========================================================
Update & Join

    $query = "UPDATE cost c
              JOIN (SELECT utm_campaign, utm_content, utm_term, sum(clicks) as total_clicks
                    FROM cost WHERE utm_id = 0
                    GROUP BY utm_campaign, utm_content, utm_term
                    HAVING total_clicks > 0
                    ORDER BY total_clicks DESC LIMIT 1000) as temp
                 ON c.utm_campaign = temp.utm_campaign
                AND c.utm_term = temp.utm_term
                AND c.utm_content = temp.utm_content
               JOIN aminno.utm u
                 ON u.hash = unhex(md5(lower(concat(c.utm_term,'|',c.utm_content,'|',c.utm_campaign))))
                SET c.utm_id = u.id
              WHERE c.utm_id = 0 AND c.clicks > 0";

    $db = Db::connection(ASHLEY_DB_STATS_WRITE_URI);

    $result = $db->query($query);
# ===========================================================
Login & addUser & setTemplate

public function login($username, $password) {
    $username = $this->db->escape($username);
    $password = $this->db->escape($password);

    $result = $this->db->query(
        "SELECT * FROM user
          WHERE username = '$username'
            AND password = SHA1(CONCAT(salt, '$password'))
            AND status = 1";
	);
}

public function addUser($data) {
    $username = $this->db->escape($data['username']);
    $password = $this->db->escape($data['password']);
    $salt = substr(md5(uniqid(rand(), true)), 0, 9);

    $this->db->query(
        "INSERT INTO user
            SET username = '$username', 
                    salt = '$salt', 
                password = '" . sha1($salt . $password) . "'" . );
	);
}

    protected function setTemplate($template) {
		$theme = $this->config->get('config_template');
		if (file_exists(DIR_TEMPLATE . "$theme/template/$template.tpl")) {
			$this->template = "$theme/template/$template.tpl";
		} else {
			$this->template = 'default/template/$template.tpl';
		}
    }
# ===========================================================
Controller::render + layout (still not good)

	protected function render() {
		foreach ($this->children as $child) {
			$this->data[basename($child)] = $this->getChild($child);
		}

		if (file_exists(DIR_TEMPLATE . $this->template)) {
			extract($this->data);

			ob_start();
			require(DIR_TEMPLATE . $this->template);
			$this->output = ob_get_contents();
			ob_end_clean();

$this->data['XXX'] = $this->document->getXXX();  // ???

            if (file_exists(DIR_LAYOUT . $this->layout)) {
                ob_start();
                $content = $this->output;
                require(DIR_LAYOUT . $this->layout);
                $this->output = ob_get_contents();
                ob_end_clean();
            }

			return $this->output;
		} else {
			trigger_error('Error: Could not load template ' . DIR_TEMPLATE . $this->template . '!');
			exit();				
		}
	}
# ===========================================================
Template Header & Variables

<!DOCTYPE html>
<html dir="<?php echo $direction; ?>" lang="<?php echo $lang; ?>">
<head>
<meta charset="UTF-8" />
<title><?php echo $title; ?></title>
<base href="<?php echo $base; ?>" />

<?php if ($description) { ?>
<meta name="description" content="<?php echo $description; ?>" />
<?php } ?>

<?php if ($keywords) { ?>
<meta name="keywords" content="<?php echo $keywords; ?>" />
<?php } ?>

<?php foreach ($links as $link) { ?>
<link href="<?php echo $link['href']; ?>" rel="<?php echo $link['rel']; ?>" />
<?php } ?>

<link rel="stylesheet" type="text/css" href="view/stylesheet/stylesheet.css" />

<?php foreach ($styles as $style) { ?>
<link rel="<?php echo $style['rel']; ?>" type="text/css" href="<?php echo $style['href']; ?>" media="<?php echo $style['media']; ?>" />
<?php } ?>

<script type="text/javascript" src="view/javascript/jquery/jquery-1.7.1.min.js"></script>
<script type="text/javascript" src="view/javascript/common.js"></script>

<?php foreach ($scripts as $script) { ?>
<script type="text/javascript" src="<?php echo $script; ?>"></script>
<?php } ?>

</head>
<body>
<div id="container">
  <div id="header"></div>
</div>
<div id="footer"><?php echo $text_footer; ?></div>
</body>
</html>


		$this->data['title'] = $this->document->getTitle();
		$this->data['description'] = $this->document->getDescription();
		$this->data['links'] = $this->document->getLinks();
		$this->data['styles'] = $this->document->getStyles();
		$this->data['scripts'] = $this->document->getScripts();		
		
		$this->data['base'] = HTTP_SERVER;
# ===========================================================
Parse DSN

/* parse the URL from the DSN string
 *  Database settings can be passed as discreet
 *  parameters or as a data source name in the first
 *  parameter. DSNs must have this prototype:
 *  $dsn = 'driver://username:password@hostname/database';
 */
$params = 'driver://username:password@hostname/database?port=3306&stricton=true';

if (($dns = @parse_url($params)) === FALSE) {
    echo 'Invalid DB Connection String';
}

$info = array(
    'dbdriver'	=> $dns['scheme'],
    'hostname'	=> (isset($dns['host'])) ? rawurldecode($dns['host']) : '',
    'username'	=> (isset($dns['user'])) ? rawurldecode($dns['user']) : '',
    'password'	=> (isset($dns['pass'])) ? rawurldecode($dns['pass']) : '',
    'database'	=> (isset($dns['path'])) ? rawurldecode(substr($dns['path'], 1)) : ''
);

// were additional config items set?
if (isset($dns['query'])) {
    parse_str($dns['query'], $extra);

    foreach ($extra as $key => $val) {
        // booleans please
        if (strtoupper($val) == "TRUE") {
            $val = TRUE;
        }
        elseif (strtoupper($val) == "FALSE") {
            $val = FALSE;
        }

        $info[$key] = $val;
    }
}
print_r($dns);
print_r($info);
# ===========================================================
这样的一个分类数组怎么用ui li无限分类表现出来，或是转换为多维数组表现
Array
(
    [list1] => Array
    (
        [0] => Array ( [code] => - [name] => 设备1)
        [1] => Array ( [code] => 0 [name] => 设备2)
        [2] => Array ( [code] => 01 [name] => 设备3)
        [3] => Array ( [code] => 010 [name] => 设备4)
        [4] => Array ( [code] => 01001 [name] => 设备5)
        [5] => Array ( [code] => 01002 [name] => 设备5)
        [6] => Array ( [code] => 01099 [name] => 设备5)
        [7] => Array ( [code] => 011 [name] => 设备4)
        [8] => Array ( [code] => 01101 [name] => 设备5)
        [9] => Array ( [code] => 01102 [name] => 设备5)
        [10] => Array ( [code] => 01103 [name] => 设备5)
        [11] => Array ( [code] => 01104 [name] => 设备5)
        [16] => Array ( [code] => 012 [name] => 设备4)
        [17] => Array ( [code] => 01201 [name] => 设备5)
        [18] => Array ( [code] => 01202 [name] => 设备5)
        [19] => Array ( [code] => 01299 [name] => 设备5)
    )
)
# ===========================================================
Hello, World in JavaScript

<html>
<head><title>Hello, World in JavaScript</title></head>
<body>
<p id="greeting"></p>
<script type="text/javascript">
  var isIE = document.attachEvent;
  var addListener = isIE
      ? function(e, t, fn) {e.attachEvent('on' + t, fn);}
      : function(e, t, fn) {e.addEventListener(t, fn, false);};

  addListener(window, 'load', function() {
    var greeting = document.getElementById('greeting');
    if (isIE) {
      greeting.innerText = 'Hello, World';
    } else {
      greeting.textContent = 'Hello, World';
    }
  });
</script>
</body>
</html>
# ===========================================================
A Simple AngularJS App

<html ng-app>
<head>
  <title>Your Shopping Cart</title>
</head>
<body ng-controller='CartController'>
<h1>Your Shopping Cart</h1>
<div ng-repeat='item in items'>
  <span>{{item.title}}</span>
  <input ng-model='item.quantity'>
  <span>{{item.price | currency}}</span>
  <span>{{item.price * item.quantity | currency}}</span>
  <button ng-click="remove($index)">Remove</button>
</div>
<script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.0.8/angular.min.js"></script>
<script>
  function CartController($scope) {
    $scope.items = [
      {title: 'Paint pots', quantity: 8, price: 3.95},
      {title: 'Polka dots', quantity: 17, price: 12.95},
      {title: 'Pebbles', quantity: 5, price: 6.95}
    ];

    $scope.remove = function(index) {
      $scope.items.splice(index, 1);
    };
  }
</script>
</body>
</html>


<!doctype html>
<html ng-app>
  <head>
    <script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.0.8/angular.min.js"></script>
  </head>
  <body>
    <table ng-controller='AlbumController'>
      <tr ng-repeat='track in album'>
        <td>{{$index + 1}}</td>
        <td>{{track.name}}</td>
        <td>{{track.duration}}</td>
      </tr>
    </table>
    <script>
      var album = [{name:'Southwest Serenade', duration: '2:34'},
                   {name:'Northern Light Waltz', duration: '3:21'},
                   {name:'Eastern Tango', duration: '17:45'}];

      function AlbumController($scope) {
        $scope.album = album;
      }
    </script>
  </body>
</html>

<!doctype html>
<html ng-app>
  <head>
    <script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.0.8/angular.min.js"></script>
    <style type="text/css">
      .selected {
        background-color: lightgreen;
      }
    </style>
  </head>
  <body>
    <table ng-controller='RestaurantTableController'>
      <tr ng-repeat='restaurant in directory' ng-click='selectRestaurant($index)'
          ng-class='{selected: $index==selectedRow}'>
        <td>{{restaurant.name}}</td>
        <td>{{restaurant.cuisine}}</td>
      </tr>
    </table>
    <script>
      function RestaurantTableController($scope) {
        $scope.directory = [{name:'The Handsome Heifer', cuisine:'BBQ'},
          {name:'Green\'s Green Greens', cuisine:'Salads'},
          {name:'House of Fine Fish', cuisine:'Seafood'}];

        $scope.selectRestaurant = function(row) {
          $scope.selectedRow = row;
        };
      }
    </script>
  </body>
</html>

<!doctype html>
<html ng-app>
  <head>
    <script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.0.8/angular.min.js"></script>
  </head>
  <body>
    <div ng-controller="CountController">
      {{count}}
      <button ng-click="count=3">Count Three</button>
      <button ng-click='setCount()'>Set count to three</button>
    </div>
    <script>
    function CountController($scope) {
      $scope.setCount = function() {
        $scope.count = 3;
      };
    }
    </script>
  </body>
</html>
# ===========================================================
似曾相识的写法

$orders = Mage::getModel('mage/orders');

$account = Api::get('ashley.account')->find($pnum);

$account = Ashley::getModel('account/account', $pnum);
$account = Ashley::getModel('account/account')->find($pnum);
# ===========================================================
    protected function substitutedVars($html)
    {
        $pixelVariables = array_keys($this->config['general']['pixel_variables']);

        $formattedVars = array();
        foreach ($pixelVariables as $name) {
            $key = '{{' . $name . '}}';
            $formattedVars[$key] = isset($this->options[$name]) ? $this->options[$name] : '';
        }

        $formattedVars = $this->mapToVendorValue($formattedVars);

        return str_replace(array_keys($formattedVars), array_values($formattedVars), $html);
    }
# ===========================================================

# ===========================================================

# ===========================================================
