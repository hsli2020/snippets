<?php

哪个名字更好？

$user = User::make($pnum);
$user = User::init($pnum);
$user = User::factory($pnum);
$user = User::create($pnum);
$user = User::instance($pnum);

/**
 * @SuppressWarnings(PHPMD.NPathComplexity)
 * @SuppressWarnings(PHPMD.CyclomaticComplexity)
 */

$user = User::factory($info['pnum']);
Jobs::factory()->queue('profile', array('photos' => $action_info), $user); 
Jobs::factory()->queue('profile', array(), User::factory($pnum));
Jobs::factory()->queue('profile', 
    array(
        'notify_new_member'     => $values['notify_newmember'],
        'notify_favorite_login' => $values['notify_login'],
        'notify_new_mail'       => $values['notify_newmail']
    ),
    User::factory($pnum)
);
Jobs::factory()->queue('chat',
    array(
        'type'      => 7, 
        'message'   => $msg, 
        'createdon' => $createdon
    ), 
    User::factory($pnumfrom), 
    User::factory($pnum)
);

# ===========================================================
Sanitizer

If the swf_id is expected to be numeric, sanitize it using the intval() function. 
If it is expected to be a string, use the json_encode function.

<?php

class Request
{
    protected $parameters;
    protected $sanitize_commands;

    public function __construct()
    {
        $this->sanitize_commands = array(
            'clean'     => function ($var) { return htmlentities($var, ENT_QUOTES); },
            'boolean'   => function ($var) { return ((bool) $var); },
            'int'       => function ($var) { return intval($var); },
            'raw'       => function ($var) { return $var; },
            'striphtml' => function ($var) { return htmlentities(strip_tags($var), ENT_QUOTES); },
            'urlencode' => function ($var) { return urlencode($var); },
        );
    }

    public function getParameter($name, $sanitize = null)
    {
        if (!array_key_exists($name, $this->parameters))
            return false;

        $result = $this->parameters[$name];
        if ($sanitize!== null) {
            // Allow sanitization of inputs when the parameter is being fetched.
            if (array_key_exists($sanitize, $this->sanitize_commands)) {
                $result = $this->sanitize_commands[$sanitize]($result);
            }
        }

        return $result;
    }
}
# ===========================================================
define('EOL', PHP_EOL);

function pr()
{
    $output = '';
    foreach (func_get_args() as $var) {
        $output .= var_export($var, true) . EOL;
    }

    $patterns = [
        "/array \(/"    => "[",
        "/\),\n/"       => "],\n",
        "/\)\n/"        => "]\n",
        "/\)$/"         => "]",
        "/=> \n(\s+)/"  => "=> ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

function dpr()
{
    ob_start();
      call_user_func_array("var_dump", func_get_args());
      $output = ob_get_contents();
    ob_end_clean();

    echo preg_replace("/=>\n(\s+)/", "=> ", $output);
}
#-------------------------------------------------------------------------------
function t()
{
    $args = func_get_args();
    if (empty($args)) return '';

    $translate = 'ucwords'; // 'gettext' for real world
    $args = array_map($translate, $args);

    $format = array_shift($args);

    return vsprintf($format, $args);
}

echo t('%s is a city of %s', 'toronto', 'canada');
echo "\n";
echo t('is this a good idea?');
#-------------------------------------------------------------------------------
function setExtra($extra, $bit, $val)
{
    $arr = array();

    /**
     * for empty string, explode returns array(0 => '');
     * seems explode never return an empty array (?)
     */
    if (!empty($extra))
        $arr = explode('|', $extra);

    for ($i=0; $i<$bit; $i++) {
        if (!isset($arr[$i]))
            $arr[$i] = '0';
    }
    $arr[$bit] = $val;

    return implode('|', $arr);
}

$extra = '';
$extra = setExtra($extra, 5, 1);
echo $extra, EOL;
#-------------------------------------------------------------------------------
class User
{
  protected $id;

  function __construct($id) { $this->id = $id; }
  function say() { echo "My id is $this->id\n"; }
}

// This is ok (function & class share same name)

function User($id)
{
  static $cache = array();

  if (!isset($cache[$id])) {
    $cache[$id] = new User($id);
  }

  return $cache[$id];
}

User(123)->say();
User(234)->say();
#-------------------------------------------------------------------------------
function token()
{
	return md5(str_shuffle(chr(mt_rand(32, 126)) . uniqid() . microtime(TRUE)));
}
#-------------------------------------------------------------------------------
for ($i=0; $i<20; $i++) {
  $str = str_shuffle(str_repeat(uniqid(), 3));
  echo $str, EOL;
  echo md5($str), EOL;
  //echo md5(rand(10000,99999).uniqid().rand(10000,99999)), EOL;
}
#-------------------------------------------------------------------------------
for ($i=0; $i<10; $i++)
  echo time() . '-'
     . str_pad(rand(0, 99999), 5, '0', STR_PAD_LEFT) . '-'
     . str_pad(rand(0, 99999), 5, '0', STR_PAD_LEFT), EOL;
#-------------------------------------------------------------------------------
function colorize($text, $color, $bold = FALSE)
{
	$colors = array_flip(array(30 => 'gray', 'red', 'green', 'yellow', 'blue', 'purple', 'cyan', 'white', 'black'));
	return"\033[" . ($bold ? '1' : '0') . ';' . $colors[$color] . "m$text\033[0m";
}
echo colorize('Everything goes well', 'green'), EOL;
echo colorize('Everything goes well', 'yellow'), EOL;
echo colorize('Something went wrong', 'red'), EOL;
#-------------------------------------------------------------------------------
function mask_email($email, $mask_char='x')
{
    list($user, $domain) = explode('@', $email);
    $masked = str_repeat($mask_char, strlen($user));
   #$masked[0] = $email[0];
    return $masked.'@'.$domain;
}
echo mask_email('lihsca@gmail.com');
#-------------------------------------------------------------------------------
function maskEmailAddress($email, $maskChar='x')
{
    list($user, $domain) = explode('@', $email);
    $firstChar = $user[0];
    $masked = str_repeat($maskChar, strlen($user)-1);
    return $firstChar.$masked.'@'.$domain;
}

echo maskEmailAddress('lhs@gmail.com');
#-------------------------------------------------------------------------------
function t($text, $params=array())
{
    return strtr(gettext($text), $params);
}

echo t('hello, %name!', array('%name'=>'world')), EOL;
echo t('hello, PHP!'), EOL;
#-------------------------------------------------------------------------------
function t($text, $params=array())
{
  return str_replace(array_keys($params), array_values($params), $text);
}

$phrase = "You should eat fruits, vegetables, and fiber every day.";
$params = array("fruits"=>"pizza", "vegetables"=>"beer", "fiber"=>"ice cream");

echo t($phrase, $params);
#-------------------------------------------------------------------------------
function currencyFormat($amount, $currency, $fmt = null)
{
    $defaultFormat = "$ # @";  // "%symbol %amount %name"
  
    // if format specified
    if (!empty($fmt)) {
        $format = $fmt;
    } elseif (!empty($currency['format'])) {
        $format = $currency['format'];
    } else {
        $format = '';//$cfg['payment']['currency_format'];
        // if not present, use defaultFormat
        if (empty($format))
          $format = $defaultFormat;
    }
  
    $symbol = $currency['symbol'];
    $name   = $currency['name'];
  
    return strtr($format, array('$' => $symbol, '#' => $amount, '@' => $name));
}

echo currencyFormat(123, array('symbol' => 'kr', 'name' => 'SEK')), EOL;
echo currencyFormat(123, array('symbol' => '&#36;', 'name' => 'USD'), '@$ #'), EOL;
echo currencyFormat(123, array('symbol' => '$', 'name' => 'CAD', 'format' => '@$ #')), EOL;
#-------------------------------------------------------------------------------
function mkpath()
{
    $path = implode(DIRECTORY_SEPARATOR, func_get_args());
    return $path.DIRECTORY_SEPARATOR;
}

echo mkpath('a', 'b', 'c', 'd');
#-------------------------------------------------------------------------------
function _buildInsert($table, $data)
{
    $fields = array();
    $values = array();

    $template = "INSERT INTO `%s` (`%s`) VALUES ('%s')";

    foreach ($data as $field => $value)
    {
        $fields[] = $field;
        $values[] = $value; //$this->_quote($value);
    }

    $fields = join("`, `", $fields);
    $values = join("', '", $values);

    return sprintf($template, $table, $fields, $values);
}

echo _buildInsert('tablename', array('username'=>'john', 'password'=>'abcd')), EOL;
#-------------------------------------------------------------------------------
function format_date($time)
{
    //$t=time()-$time;
    $t=$time;
    $f=array(
        '31536000'=>'年',
        '2592000'=>'个月',
        '604800'=>'星期',
        '86400'=>'天',
        '3600'=>'小时',
        '60'=>'分钟',
        '1'=>'秒'
    );

    foreach ($f as $k=>$v)    {
        if (0 != $c=floor($t/(int)$k)) {
            return $c.$v.'前';
        }
    }
}
echo format_date(72000);
#-------------------------------------------------------------------------------
// 生成验证码的方法可以优化一下：
function getCode($len) {
   $str = "23456789ABCDEFGHIJKLMNQRSTUVWXYZ";
   $code = substr(str_shuffle($str), 0, $len);

   return $code;
}
for ($i=0; $i<10; $i++) {
    echo getCode(5), EOL;
}
#------------------------------------------------------------------------------
function gen_uuid_1($len=8) {

    $hex = md5("yourSaltHere" . uniqid("", true));

    $pack = pack('H*', $hex);
    $tmp =  base64_encode($pack);

    $uid = preg_replace("#(*UTF8)[^A-Za-z0-9]#", "", $tmp);

    $len = max(4, min(128, $len));

    while (strlen($uid) < $len)
        $uid .= gen_uuid_1(22);

    return substr($uid, 0, $len);
}

function gen_uuid_2($len=8)
{
    $hex = md5("your_random_salt_here_31415" . uniqid("", true));

    $pack = pack('H*', $hex);

    $uid = base64_encode($pack);        // max 22 chars

    $uid = preg_replace("/[^A-Za-z0-9]/", "", $uid);    // mixed case
    //$uid = ereg_replace("[^A-Z0-9]", "", strtoupper($uid));    // uppercase only

    if ($len<4)
        $len=4;
    if ($len>128)
        $len=128;                       // prevent silliness, can remove

    while (strlen($uid)<$len)
        $uid = $uid . gen_uuid_2(22);     // append until length achieved

    return substr($uid, 0, $len);
}
echo gen_uuid_1(20), EOL;
echo gen_uuid_2(20), EOL;
#------------------------------------------------------------------------------
function gen_uid($l=10){
    return substr(str_shuffle("0123456789abcdefghijklmnopqrstuvwxyz"), 0, $l);
}
echo gen_uid(10), EOL;
#------------------------------------------------------------------------------
echo sha1(uniqid('', true)), EOL;
echo sha1(md5(uniqid('', true))), EOL;

$x = md5(uniqid(mt_rand(), true));
$y = str_replace('.', '', uniqid(mt_rand(), true));
echo $x, EOL, $y, EOL;
#------------------------------------------------------------------------------
echo sprintf("%'-10s %s %'-10s", '-', date('Y-m-d'), '-'), EOL;
echo sprintf('Hello, %1$s. How are you today, %1$s?', 'Michael'), EOL;
#------------------------------------------------------------------------------
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
    $keys = array();
    $arrays = func_get_args();

    foreach ($arrays as $array) {
        $keys = array_merge($keys, array_keys($array));
    }
    $keys = array_unique($keys);

    $result = array();
    foreach ($keys as $key) {
        $merged = array();
        foreach ($arrays as $array) {
            if (array_key_exists($key, $array)) {
                $merged = array_merge($merged, $array[$key]);
            }
        }
        $result[$key] = $merged;
    }
    return $result;
}

// $x = array_merge_recursive($a, $b, $c);
// pr($x);
// $x = array_merge_simple($a, $b, $c);
// pr($x);
#------------------------------------------------------------------------------
$a = array(100 => array('red', 'blue', 'white'));
$b = array(100 => array('apple', 'banana', 'pear'));
$x = array_merge_recursive($a, $b);
pr($x);

$a = array('abc' => array('red', 'blue', 'white'));
$b = array('abc' => array('apple', 'banana', 'pear'));
$x = array_merge_recursive($a, $b);
pr($x);
#-------------------------------------------------------------------------------
function getEmailDeliveryId()
{
    $uid = uniqid('', TRUE);
    $arr = explode('.', $uid);
    return hexdec($arr[0].substr(dechex($arr[1]),-2));
}
echo getEmailDeliveryId(),EOL;
#-------------------------------------------------------------------------------
function format_bytes($size)
{
    $arr = array(' B', ' KB', ' MB', ' GB', ' TB');
    for ($f = 0; $size >= 1024 && $f < 4; $f++) {
        $size /= 1024;
    }
    return round($size, 2).$arr[$f];
}
echo format_bytes('123456789'), EOL;

function format_bytes1($size)
{
    $arr = array(' B', ' KB', ' MB', ' GB', ' TB');
    $log = floor(log($size, 1024));
    return round($size / pow(1024, $log), 2) . $arr[$log];
}
echo format_bytes1(123456789), EOL;
#-------------------------------------------------------------------------------
// 驼峰字符串转换成下划线样式
$str = 'openAPI';
echo $str, EOL;
echo preg_replace('/((?<=[a-z])(?=[A-Z]))/', '_', $str), EOL;
#-------------------------------------------------------------------------------
function getVars($s)
{
    if (preg_match_all("/{{\s*(.*?)\s*}}/si",$s,$m)) {
      return array_unique($m[1]); // some variables appear more than once.
    }
    return array();
}
var_export(getVars('{{name}} is {{ -name  }}, {{ fistname }}, {{ lastname }}'));
#-------------------------------------------------------------------------------
function startsWith($haystack, $needle)
{
     $length = strlen($needle);
     return (substr($haystack, 0, $length) === $needle);
}

function endsWith($haystack, $needle)
{
    $length = strlen($needle);
    if ($length == 0) {
        return true;
    }

    return (substr($haystack, -$length) === $needle);
}
#-------------------------------------------------------------------------------
$required = ['id', 'username', 'password'];
$actual = ['username', 'Password', 'country', 'lang'];

// $missing = array_diff($required, $actual);
// if (!empty($missing)) {
//     echo "required argument missing: [ '", implode($missing, "', '"), "' ]", EOL;
// }

function checkRequiredArgs($required, $actual)
{
    $missing = array_diff($required, $actual);
    if (!empty($missing)) {
        throw new Exception("Required argument missing: '" . implode($missing, "', '"). "'");
    }
}
// try {
//     checkRequiredArgs($required, $actual);
// } catch(Exception $e) {
//     echo $e->getMessage();
// }
#-------------------------------------------------------------------------------
/**
 * return  true for [ 'true' | '1' | 'on' | 'yes' ],
 *        false for everything else
 */
function parseBool($string)
function stringToBool($string)
{
    return filter_var($string, FILTER_VALIDATE_BOOLEAN);
}
// parseBool(9);     // false !!!
// parseBool('ok');  // false !!!
#-------------------------------------------------------------------------------
function getDataFunc($data)
{
    return function($key, $default=null) use ($data)
    {
        $val = $default;
        if (isset($data[$key]))
            $val = $data[$key];
        return $val;
    };
}

$pageVar     = getDataFunc($Data['app']['page']);
$sessionVar  = getDataFunc($Data['app']['session']);
$identityVar = getDataFunc($Data['app']['identity']);

pr($pageVar('coupon'));
pr($pageVar('nil'));
pr($pageVar('empty'));
pr($pageVar('bool'));
pr($pageVar('not', 'abc'));
#-------------------------------------------------------------------------------
function objectToArray($d) {
    if (is_object($d)) {
        // Gets the properties of the given object
        // with get_object_vars function
        $d = get_object_vars($d);
    }

    if (is_array($d)) {
        /*
        * Return array converted to object
        * Using __FUNCTION__ (Magic constant)
        * for recursive call
        */
        return array_map(__FUNCTION__, $d);
    } else {
        // Return array
        return $d;
    }
}

function arrayToObject($d) {
    if (is_array($d)) {
        /*
        * Return array converted to object
        * Using __FUNCTION__ (Magic constant)
        * for recursive call
        */
        return (object) array_map(__FUNCTION__, $d);
    } else {
        // Return object
        return $d;
    }
}

// $init = new stdClass;
// $init->foo = "Test data";
// $init->bar = new stdClass;
// $init->bar->baaz = "Testing";
// $init->bar->fooz = new stdClass;
// $init->bar->fooz->baz = "Testing again";
// $init->foox = "Just test";
// 
// $array = objectToArray($init);
// $object = arrayToObject($array);
// 
// pr($init);
// pr($array);
// pr($object);

// $json = json_encode($object);
// pr($json);
// $array = json_decode($json, true);
// pr($array);
//
// $arr =  (array) $Obj;
#-------------------------------------------------------------------------------
function getMaxUploadSize()
{
    $val = ini_get('upload_max_filesize');

#   $val = trim($val);
#   $last = strtolower($val[strlen($val)-1]);
    $last = strtolower(substr(trim($val),  -1));

    $factors =  array(
        'g' => 1024 * 1024 * 1024,
        'm' => 1024 * 1024,
        'k' => 1024,
    );

    return $val * $factors[$last];
}
#-------------------------------------------------------------------------------
# ===========================================================
class strongpass {

    public function check($password){
	$response = "OK";
        if(strlen($password) < 8){
            $response = "Password must be at least 8 characters";
        } else if(is_numeric($password)){
            $response = "Password must contain at least one letter";
	} else if(!preg_match('#[0-9]#', $password)){
	    $response = "Password must contain at least one number";
	}
	/* 
	Additional checks you can accomplish as homework 
        - Make sure there is at least one lowercase letter
        - Make sure there is at least one uppercase letter
        - Make sure there is at least one symbol character
        */
        return $response;
    }

}
# ===========================================================
function pr($var)
{
    $output = var_export($var, true);
    echo preg_replace("/=> \n(\s+)/", "=> ", $output), EOL;
}

function pr($var)
{
    $output = var_export($var, true);

    $output = str_replace("array (", "[", $output);
    $output = preg_replace("/\),\n/", "],\n", $output);
    $output = preg_replace("/\)$/", "]", $output);
    $output = preg_replace("/=> \n(\s+)/", "=> ", $output);

    echo $output, EOL;
}

function pr($var)
{
    $output = var_export($var, true);

    $patterns = [
        "/array \(/"    => "[",
        "/\),\n/"       => "],\n",
        "/\)$/"         => "]",
        "/=> \n(\s+)/"  => "=> ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output), EOL;
}

/**
 * return  true for 'true' | '1' | 'on' | 'yes',
 *        false for everything else
 */
function parseBool($string)
function stringToBool($string)
{
    return filter_var($string, FILTER_VALIDATE_BOOLEAN);
}


function dpr($var)
{
    ob_start();
      var_dump($var);
      $output = ob_get_contents();
    ob_end_clean();

    $output = preg_replace("/=>\n(\s+)/", "=> ", $output);
    echo $output, PHP_EOL;
}

function dpr()
{
    ob_start();
      call_user_func_array("var_dump", func_get_args());
      $output = ob_get_contents();
    ob_end_clean();

    $output = preg_replace("/=>\n(\s+)/", "=> ", $output);
    echo $output, PHP_EOL;
}
# ===========================================================
function fpr()
{
    static $first = true;

    $filename = '/tmp/trace.log';

    if ($first) {
        $first = false;
        $str = sprintf("%'-30s %s %'-30s\n", '-', date('Y-m-d H:i:s'), '-');
        error_log($str, 3, $filename);
    }

    $args = func_get_args();
    foreach ($args as $var) {
        $str = print_r($var, true);
        error_log($str."\n", 3, $filename);
    }
}
# ===========================================================
<?php

function pr($label, $var='')
{
    echo "<pre>\n";
    echo "<b>$label</b>\n";
    #print_r($var);
    $str = var_export($var, true);
    echo preg_replace("/=> \n(\s+)/", "=> ", $str), "\n";
    echo "</pre>\n";
}

function dpr()
{
    echo "<pre>\n";
    $args = func_get_args();
    foreach ($args as $var) {
        #print_r($var);
        $str = var_export($var, true);
        echo preg_replace("/=> \n(\s+)/", "=> ", $str), "\n";
    }
    echo "</pre>\n";
}

function fpr()
{
    static $first = true;

    $filename = '/tmp/ztrace.log';

    if ($first) {
        $first = false;
       #$str = str_repeat('-', 30).' '.date('Y-m-d H:i:s').' '.str_repeat('-', 30)."\n";
        $str = sprintf("%'-30s %s %'-30s\n", '-', date('Y-m-d H:i:s'), '-');
        $str .= "\tHTTP_HOST    = ".$_SERVER['HTTP_HOST']."\n";
#       $str .= "\tSERVER_NAME  = ".$_SERVER['SERVER_NAME']."\n";
        $str .= "\tREQUEST_URI  = ".$_SERVER['REQUEST_URI']."\n";
        $str .= "\tQUERY_STRING = ".$_SERVER['QUERY_STRING']."\n";
#       $str .= "\tUSER_AGENT   = ".$_SERVER['HTTP_USER_AGENT']."\n";
#       if (isset($_SERVER['HTTP_REFERER']))
#           $str .= "\tHTTP_REFERER = ".$_SERVER['HTTP_REFERER']."\n";
        $str .= "\n";
        error_log($str, 3, $filename);
    }

    $args = func_get_args();
    foreach ($args as $var) {
        $str = var_export($var, true);
        $str = preg_replace("/=> \n(\s+)/", "=> ", $str);
        error_log($str."\n", 3, $filename);
    }
}

function ftr($msg)
{
    fpr($msg);

    $trace = debug_backtrace();
    foreach ($trace as $entry) {
        if (isset($entry['file'])) {
            fpr($entry['file'] .':'. $entry['line']);
        }
    }
}
# ===========================================================
# Web page can't be in iframe

<?php header('X-Frame-Options: GOFORIT'); ?>
X-Frame-Options SAMEORIGIN, GOFORIT
# ===========================================================

/**
*url解析分发类
**/
class Route{
    private $_module;
private $_controller;
private $_action;
private $_requestParam = array();
private $_rules = array();

public function __construct(){
    $this->_module = '';
$this->_controller = C('defaultController');
$this->_action = C('defaultAction');
$this->_rules = C('urlManager=>rules');
}

    /**
*获取模型名称
**/
public function getModule()
{
return $this->_module;
}
    
/**
*获取控制器名称
**/
public function getController()
{
return $this->_controller;
}

    /**
*获取动作名称
**/
public function getAction()
{
return $this->_action;
}

    /**
*设置请求参数数组，解析二级域名
**/
    private function parse()
{
//先检查二级域名
$modules = C('modules');
if($modules) //只有设置了模块才检查
{
$host = $_SERVER['HTTP_HOST'];
if(C('domain') && $pos = strpos($host,C('domain')))
{
$preffix = substr($host,0,$pos-1);
if($pos = strrpos($preffix,'.'))
{
$preffix = substr($preffix,$pos);
}
if(in_array($preffix,$modules))
{
$this->_module = $preffix;
}
}
}
}



/**
*  解析普通参数
**/
private function parseParam($controllerObj)
{
$num=count($this->_requestParam);
if($num == 0)
return;
for($i=0,$num=count($this->_requestParam);$i<$num;$i+=2)
{
            $controllerObj->setParam($this->_requestParam[$i],$this->_requestParam[$i+1]);
}
}

    /**
*  url解析规则
*  1. 'variable/varialbe:regexp>'=>'controller/action',
*  2. 'variable/variable'=>'controller/action',
*  3. 'variable/variable/:varialbe|regexp'=>'controller/action',
*  4. 'variable/:variable|regexp/varialbe'=>'controller/action',
*  5. 'variable/variable/*'=>'controller/action',
*  desc :开头的是变量，|后面的正则是变量匹配类型，如果没有，则是任意变量
**/
private function rule()
{
$parseParam = array();
$requestParamNum = count($this->_requestParam);
foreach($this->_rules as $rule=>$path)
{
$ruleParam = explode('/',$rule);
$ruleParamNum = count($ruleParam);
            if(($ruleParam[$ruleParamNum-1] == '*' && $requestParamNum >= $ruleParamNum-1) || $requestParamNum == $ruleParamNum)//rule满足初始条件
{
                foreach($ruleParam as $key=>$value)
{
if($value == '*')
{
break;
}
if(substr($value,0,1) == ':')
{
$value = substr($value,1);
    if(strpos($value,'|'))
    {
       list($value,$regexp)=explode('|',$value);
   if(!preg_match("/$regexp/",$this->_requestParam[$key]))
   {
                               $parseParam = array();
                               break;
   }
    }
                        $parseParam[] = $value; //参数名
    $parseParam[] = $this->_requestParam[$key]; //参数值
}
else if(strcasecmp($value,$this->_requestParam[$key]))
{
                         $parseParam = array();
  break;
}
}
if($parseParam)
{
$pathParam = explode('/',$path);
break;
}
}
}

if($parseParam && $key == $ruleParamNum-1) //匹配了
{
if($ruleParam[$ruleParamNum-1] == '*' && $requestParamNum >= $ruleParamNum)
{
                   for(;$key<$requestParamNum;$key++)
   $parseParam[] = $this->_requestParam[$key];
}
                $this->_requestParam = array_merge($pathParam,$parseParam);
unset($parseParam);
}
}

    /**
*  执行控制器方法
**/
public function run()
{
$this->parse();
        if($_SERVER['REQUEST_URI'])
        {
$this->_requestParam = explode('/',trim($_SERVER['REQUEST_URI'],'/'));
if($this->_rules)//存在路由规则
{
$this->rule(); //本次请求满足路由规则
}
if($modules && $this->_module == '' && in_array($this->_requestParam[0],$modules))
{
$this->_module = $this->_requestParam[0];
array_shift($this->_requestParam);
}
//print_r($this->_requestParam);
$controllerPath = APP_PATH.DS.($this->_module?'modules'.DS.$this->_module.DS:'').'controllers'.DS;
if(file_exists($controllerPath.$this->_requestParam[0].'Controller.class.php'))
{
                $this->_controller = $this->_requestParam[0];
array_shift($this->_requestParam);
}
}

$controllerClass = $this->_controller.'Controller';
        import($controllerClass,$controllerPath,'class');
$controllerObj = new $controllerClass($this);
    $this->_requestParam[0];

if($this->_requestParam[0] && method_exists($controllerObj,$this->_requestParam[0].'Action'))
{
            $this->_action = $this->_requestParam[0];
array_shift($this->_requestParam);
}
$this->parseParam($controllerObj); //对象时引用方式

try{
$action = $this->_action.'Action';
             $controllerObj->$action();
}
catch(Exception $e)
{
$this->rediretUrl('/');
}
}

    /**
*  url重定向
**/
public function rediret($url)
{
header('location:'.$url);
}

    /**
* 生成url方法，需要逆解析rule规则
**/
public function url($controller,$action,$param,$module = '')
{
         $urlPath = $controller.'/'.$action;
if($module)
             $urlPath = $module .'/'. $urlPath;
$parseParam = array();
if($this->_rules)
{
$ruleArr = array();
     foreach($this->_rules as $rule=>$path)
{
if($path ==  $urlPath)
{
$ruleArr[] = $rule;
}
}
if($ruleArr)
{
foreach($ruleArr as $rule)
{
                     $ruleParam = explode('/',$rule);
foreach($ruleParam as $key=>$value)
{
if(substr($value,0,1) == ':')
{
                             $value = substr($value,1);
if(strpos($value,'|'))
{
     list($value,$regexp)=explode('|',$value);
}
else
{
$regexp = '';
}
if(isset($param[$value]) && (!$regexp || preg_match("/$regexp/",$param[$value])))
{
                                 $parseParam[] = $param[$value];
unset($param[$value]);
}
else
{
                                 $parseParam = array();
break;
}
}
else
{
$parseParam[] = $value;
}
}
if($parseParam)
break;
} //foreach($ruleArr as $rule)
if($parseParam)
                    $urlPath = implode('/',$parseParam);
} //if($ruleArr)
} //if($this->_rules)
if($param)
{
foreach($param as $key=>$value)
{
                 $urlPath .= '/'.$key.'/'.$value;
}
}
return $urlPath;
}
}
# ===========================================================
<?php

class SuperGlobal
{
    public $data = array();

    public function __construct($data)
    {
        $data = $this->clean($data);
    }

    public function get($key, $default = null)
    {
        return isset($this->data[$key]) ? $this->data[$key] : $default;
    }

    protected function clean($data)
    {
        if (is_array($data)) {
            foreach ($data as $key => $value) {
                unset($data[$key]);
                $data[$this->clean($key)] = $this->clean($value);
            }
        } else { 
            $data = htmlspecialchars($data, ENT_COMPAT, 'UTF-8');
        }

        return $data;
    }
}

class Request
{
    public $get = array();  // better name: query
    public $post = array();
    public $request = array(); // better name: params
    public $cookie = array();
    public $files = array();
    public $server = array();

    public function __construct()
    {
        $this->get     = new SuperGlobal($_GET);
        $this->post    = new SuperGlobal($_POST);
        $this->request = new SuperGlobal($_REQUEST);
        $this->cookie  = new SuperGlobal($_COOKIE);
        $this->files   = new SuperGlobal($_FILES);
        $this->server  = new SuperGlobal($_SERVER);
    }
}
# ===========================================================
<?php
/**
*视图文件
*
**/
class View {
private $_variables = array();//参数列表
private $_useLayout = true;//使用公共魔板
private $_templateType = 'phtml';
private $_route;
private $_layout;

public function __construct($route)
{
$this->_route = $route;
$this->_layout = APP_PATH.DS.'layouts'.DS.'main.'.$this->_templateType;
}

public function __set($name,$value)
{
$this->_variables[$name] = $value;
}

public function __get($name)
{
return $this->_variables[$name];
}

public function noLayout()
{
        $this->_useLayout = false;
}

public function setTemplateType($templatType)
{
        $this->_templateType = $templatType;
}
    
public function setLayout($layout,$path='')
{
       $path == '' && $path = APP_PATH.DS.'layouts';
       $this->_layout = $path.DS.$layout.'.'.$this->_templateType;
}
public function render($template = '')
{
        if(!$template)
           $template = $this->_route->getAction();
$module = $this->_route->getModule();
$templateFile = APP_PATH.DS.($module?'modules'.DS.$module.DS:'').'views'.DS.$this->_route->getController().DS.$template.'.'.$this->_templateType;
$this->_variables && extract($this->_variables,EXTR_OVERWRITE);
if($this->_useLayout)
{
include $this->_layout;
}
else
{
include $templateFile;
}
}
}

# ===========================================================
# AbTest (Why it doesn't work properly?)

$abTest = new AbTest('name', [50, 50]);
if ($abTest->getGroup($userId) == 1)
if ($abTest->getGroup($userId) == 2)

<?php

class AbTest
{
    protected $testName;
    protected $percentages;
    protected $minUserId;
    protected $maxUserId;

    public function __construct($testName, $percentages = array(50, 50))
    {
        $this->minUserId = null;
        $this->maxUserId = null;

        $this->testName = $testName;
        $this->percentages = $percentages;

        if (array_sum($percentages) != 100) {
            throw new InvalidArgumentException('Invalid groups setting');
        }
    }

    public function setMinUserId($userId)
    {
        $this->minUserId = $userId;
    }

    public function setMaxUserId($userId)
    {
        $this->maxUserId = $userId;
    }

    public function getGroup($userId)
    {
        if ((isset($this->minUserId) && ($userId < $this->minUserId)) ||
            (isset($this->maxUserId) && ($userId > $this->maxUserId))) {
            return 0;
        }

        $hash = substr($this->hash($userId), -3) * 0.1;

        $percent = 0;
        foreach ($this->percentages as $group => $perc) {
            $percent += $perc;
            if ($hash <= $percent) {
                return $group + 1;
            }
        }

        return 0;  // should never reach here, but just in case
    }

    protected function hash($userId)
    {
        return sprintf('%u', crc32($this->testName . $userId));
    }
}

include "AbTest-New.php";
echo PHP_INT_MAX, EOL;
$abtest = new AbTest('The_Test_Name', [40, 20, 40]);
$arr = [0, 0, 0, 0];
#for ($id=11000; $id<12000; $id++) {
for ($id=100000; $id<200000; $id++) {
  $g = $abtest->getGroup($id);
  $arr[$g]++;
}
print_r($arr);
# ===========================================================
PSR-0	Autoloading Standard
PSR-1	Basic Coding Standard	
PSR-2	Coding Style Guide
PSR-3	Logger Interface
PSR-4	Autoloading Standard
PSR-5	PHPDoc Standard
PSR-6	Caching Interface
PSR-7	HTTP Message Interface
PSR-8	Huggable Interface
PSR-9	Security Disclosure
PSR-10	Security Advisories

PSR-0 (now deprecated) and PSR-4 are autoloading standards that actually map your namespaces to real folders.

PSR-0
vendor/<VendorName>/<ProjectName>/src/<NamespaceVendor>/<NamespaceProject>/File.php

PSR-4 (See that they removed the namespaces folders? Because you already reference that in composer.json
vendor/<VendorName>/<ProjectName>/src/File.php
# ===========================================================
http://www.php-fig.org/psr/psr-0/
http://www.php-fig.org/psr/psr-1/
http://www.php-fig.org/psr/psr-2/
http://www.php-fig.org/psr/psr-3/
http://www.php-fig.org/psr/psr-4/

$_POST = json_decode(file_get_contents('php://input'), true);
# ===========================================================
Laravel Snippets

通过一个Config类访问多个配置文件

$timezone = Config::get('app.timezone');
$timezone = Config::get('app.timezone', 'UTC');

Config::set('database.default', 'sqlite');

app/config/app.php
app/config/auth.php
app/config/cache.php
app/config/compile.php
app/config/database.php
app/config/mail.php
app/config/queue.php
app/config/remote.php
app/config/services.php
app/config/session.php
app/config/view.php
app/config/workbench.php

app/config/local/app.php
app/config/local/database.php

app/config/testing/cache.php
app/config/testing/session.php
# ===========================================================
CakePHP

class AppController extends Controller { }
class AppModel extends Model { }
class AppHelper extends Helper { }

views/layouts/default.ctp

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<?php echo $this->Html->charset(); ?>
	<title>
		<?php __('CakePHP: the rapid development php framework:'); ?>
		<?php echo $title_for_layout; ?>
	</title>
	<?php
		echo $this->Html->meta('icon');
		echo $this->Html->css('cake.generic');
		echo $scripts_for_layout;
	?>
</head>
<body>
	<div id="container">
		<div id="header">
			<h1><?php echo $this->Html->link(__('CakePHP: the rapid development php framework', true), 'http://cakephp.org'); ?></h1>
		</div>
		<div id="content">
			<?php echo $this->Session->flash(); ?>
			<?php echo $content_for_layout; ?>
		</div>
		<div id="footer">
			<?php echo $this->Html->link(
					$this->Html->image('cake.power.gif', array('alt'=> __('CakePHP: the rapid development php framework', true), 'border' => '0')),
          'http://www.cakephp.org/', array('target' => '_blank', 'escape' => false)
				);
			?>
		</div>
	</div>
	<?php echo $this->element('sql_dump'); ?>
</body>
</html>

CREATE TABLE i18n (
	id int(10) NOT NULL auto_increment,
	locale varchar(6) NOT NULL,
	model varchar(255) NOT NULL,
	foreign_key int(10) NOT NULL,
	field varchar(255) NOT NULL,
	content mediumtext,
	PRIMARY KEY	(id),
	INDEX locale (locale),
	INDEX model (model),
	INDEX row_id (foreign_key),
	INDEX field (field)
);

CREATE TABLE cake_sessions (
  id varchar(255) NOT NULL default '',
  data text,
  expires int(11) default NULL,
  PRIMARY KEY  (id)
);

Router::connect('/', array('controller' => 'pages', 'action' => 'display', 'home'));
Router::connect('/pages/*', array('controller' => 'pages', 'action' => 'display'));

	Configure::write('debug', 2);
	Configure::write('log', true);
	Configure::write('App.encoding', 'UTF-8');

	//Configure::write('App.baseUrl', env('SCRIPT_NAME'));
	//Configure::write('Routing.prefixes', array('admin'));
	//Configure::write('Cache.disable', true);
	//Configure::write('Cache.check', true);

	define('LOG_ERROR', 2);

	Configure::write('Session.save', 'php');

	//Configure::write('Session.model', 'Session');
	//Configure::write('Session.table', 'cake_sessions');
	//Configure::write('Session.database', 'default');

	Configure::write('Session.cookie', 'CAKEPHP');
	Configure::write('Session.timeout', '120');
	Configure::write('Session.start', true);
	Configure::write('Session.checkAgent', true);
	Configure::write('Security.level', 'medium');
	Configure::write('Security.salt', 'DYhG93b0qyJfIxfs2guVoUubWwvniR2G0FgaC9mi');
	Configure::write('Security.cipherSeed', '76859309657453542496749683645');

	//Configure::write('Asset.timestamp', true);
	//Configure::write('Asset.filter.css', 'css.php');
	//Configure::write('Asset.filter.js', 'custom_javascript_output_filter.php');

	Configure::write('Acl.classname', 'DbAcl');
	Configure::write('Acl.database', 'default');

	//date_default_timezone_set('UTC');

	Cache::config('default', array('engine' => 'File'));

class FileLog
{
  var $_path = null;

  function FileLog($options = array())
  {
    $options += array('path' => LOGS);
    $this->_path = $options['path'];
  }

  function write($type, $message)
  {
    $debugTypes = array('notice', 'info', 'debug');

    if ($type == 'error' || $type == 'warning') {
      $filename = $this->_path  . 'error.log';
    } elseif (in_array($type, $debugTypes)) {
      $filename = $this->_path . 'debug.log';
    } else {
      $filename = $this->_path . $type . '.log';
    }
    $output = date('Y-m-d H:i:s') . ' ' . ucfirst($type) . ': ' . $message . "\n";
    $log = new File($filename, true);
    if ($log->writable()) {
      return $log->append($output);
    }
  }
}
# ===========================================================
A PHP callable is a PHP variable that can be used by the call_user_func() function and 
returns true when passed to the is_callable() function. It can be 
-- a \Closure instance, 
-- an object implementing an __invoke method (which is what closures are in fact), 
-- a string representing a function, 
-- an array representing an object method or a class method.


function gbk_to_utf8($str){
    return mb_convert_encoding($str, 'utf-8', 'gbk');
}
 
function utf8_to_gbk($str){
    return mb_convert_encoding($str, 'gbk', 'utf-8');
}
# ===========================================================
#echo base64_encode(hex2bin(md5(''))), 
#echo base64_encode(md5('', true));

$myArray = array(
    'key1' => 'value1',
    'key2' => array(
        'subkey' => 'subkeyval'
    ),
    'key3' => 'value3',
    'key4' => array(
        'subkey4' => array(
            'subsubkey4' => 'subsubkeyval4',
            'subsubkey5' => 'subsubkeyval5',
        ),
        'subkey5' => 'subkeyval5'
    )
);

function array_flatten($array)                                                                                             
{                                                                                                                          
    $iter = new \RecursiveIteratorIterator(new \RecursiveArrayIterator($array));

    $result = array();
    foreach ($iter as $leafValue) {
        $keys = array();
        foreach (range(0, $iter->getDepth()) as $depth) {
            $keys[] = $iter->getSubIterator($depth)->key();
        }
        $result[join('.', $keys)] = $leafValue;
    }

    return $result;
}

print_r(array_flatten($myArray));

# ===========================================================
use Model to do 'ON DUPLICATE KEY UPDATE'

app/models/Test.php
<?php

namespace App\Models;

use Phalcon\Mvc\Model;

class Test extends Model
{
    public $sku;
    public $price;
    public $qty;
    public $isnew;
    public $updatedon;

    public function getSource() { return 'test';  }
}

// tt.php
<?php

include __DIR__ . '/public/init.php';

use App\Models\Test;

$t = new Test();

$t = Test::findFirst("sku='SKU-4'");
if ($t) {
    #$t->qty = 444;
    #$t->isnew = 0;
    $t->save(['qty' => 444, 'isnew' => 0]);
} else {
    $t = new Test();
    #$t->sku = 'SKU-4';
    #$t->price = 44.0;
    #$t->qty = 44;
    $t->save(['sku' => 'SKU-4', 'price' => 44.0, 'qty' => 44]);
}
# ===========================================================
OpenCart

$ ag 'new Action'
$ ag --php "\bforward\b"
$ ag 'load->controller'【老版本】

* OpenCart的Action类是理解OpenCart结构的关键点，在这里加入追踪语句。

* OpenCart的结构或设计思路上的一个缺点(或特点)
每个页面对应一个controller中的action，这固然没错。但是，将组成页面中的各个区域，如header、footer、sidebar、menu等也都当做controller，主controller通过this->load->controller('common/header')的方式加载子controller，这种方式似乎不妥，总感觉存在安全问题，毕竟，那些区域并不是页面，不能具有url，不能通过url访问。
我的想法是，这些都应该是Block类(或Region/Box类)system/engine/block.php的派生类，位于catalog/block/以及admin/block目录中，通过this->load->block('path')加载。
($this->block('path')->render())
(在新版本中，this->load->controller被getChild代替，离我的想法近了一步，但仍然存在问题，child不应该在controller目录下)

* 查看opencart中的api实现方法 【老版本】

* opencart每处理一个请求，都要调入很多文件(30多个)，有些是不必要的，这会带来性能损失，应该如何解决。

* opencart中如何将页面静态化，提高加载速度，因为商品变化不会很频繁。

* opencart中缺少的东西
   - 灵活的router
   - validator类
   - 一个类似QueryBuilder的类

* 在opencart的controller中发现如下目录，究竟有何作用？
   - extension
   - design
   - module

* 查一下这个类 (Pimple最新版使用了这个类)
        $this->factories = new \SplObjectStorage();
        $this->protected = new \SplObjectStorage();
# ===========================================================
