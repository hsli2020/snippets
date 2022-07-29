<?php

CakePHP

class PostsController extends AppController {
    public $uses = array('Users');       // for Models
    public $helpers = array('Html', 'Form', 'Session');  // for Helpers
    public $components = array('Session', 'Auth', 'Acl');  // for Components
}
# ===========================================================
<?php

// input misspelled word
$input = 'ccarrot';
echo "Input word: $input\n\n";

// array of words to check against
$words  = array('apple','pineapple','banana','orange',
                'radish','carrot','pea','bean','potato', '');

// no shortest distance found, yet
$shortest = -1;

// loop through words to find the closest
foreach ($words as $word) {

    // calculate the distance between the input word,
    // and the current word
    // $lev = levenshtein($input, $word);
    $percent = 0;
    $lev = similar_text($input, $word, $percent);
    $percent = round($percent, 2);

    // check for an exact match
    // if ($lev == 0) {
    if ($percent == 100) {

        // closest word is this one (exact match)
        $closest = $word;
        $shortest = 100;

        // break out of the loop; we've found an exact match
        break;
    }
    
    // if this distance is less than the next found shortest
    // distance, OR if a next shortest word has not yet been found
    if ($percent >= $shortest || $shortest < 0) {
        // set the closest match, and shortest distance
        $closest  = $word;
        $shortest = $percent;
    }
    
    echo "$input\t$word\t\t$percent\n";
}

echo "\n";
if ($shortest == 100) {
    echo "Exact match found: $closest\n";
} else {
    echo "Did you mean: $closest?\n";
}

echo "\n";
echo "shortest=$shortest\n";

# ===========================================================
SNIPPET from CakePHP

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

	function h($text, $charset = null) {
		if (is_array($text)) {
			return array_map('h', $text);
		}

		static $defaultCharset = 'UTF-8';

		if ($charset) {
			return htmlspecialchars($text, ENT_QUOTES, $charset);
		} else {
			return htmlspecialchars($text, ENT_QUOTES, $defaultCharset);
		}
	}
# ===========================================================
app/.htaccess
  <IfModule mod_rewrite.c>
      RewriteEngine on
      RewriteRule    ^$    webroot/    [L]
      RewriteRule    (.*) webroot/$1    [L]
   </IfModule>

app/import/blah/.htaccess
  <IfModule mod_rewrite.c>
     RewriteEngine on
     RewriteBase /
     RewriteCond %{REQUEST_URI} ^/export/(.*)$
     RewriteRule ^.*$ - [L]
  </IfModule>
  <IfModule mod_rewrite.c>
     RewriteEngine on
     RewriteRule    ^$ app/webroot/    [L]
     RewriteRule    (.*) app/webroot/$1 [L]
     php_value max_execution_time 360
  </IfModule>
#save file instead of opening in browser
AddType application/octet-stream .po

app/webroot/.htaccess
  <IfModule mod_rewrite.c>
      RewriteEngine On
      RewriteCond %{REQUEST_FILENAME} !-d
      RewriteCond %{REQUEST_FILENAME} !-f
      RewriteRule ^(.*)$ index.php?url=$1 [QSA,L]
  </IfModule>

app/webroot/ckeditor/.htaccess
#
# On some specific Linux installations you could face problems with Firefox.
# It could give you errors when loading the editor saying that some illegal
# characters were found (three strange chars in the beginning of the file).
# This could happen if you map the .js or .css files to PHP, for example.
#
# Those characters are the Byte Order Mask (BOM) of the Unicode encoded files.
# All FCKeditor files are Unicode encoded.

  AddType application/x-javascript .js
  AddType text/css .css

# If PHP is mapped to handle XML files, you could have some issues. The
# following will disable it.

  AddType text/xml .xml

cake/console/templates/skel/.htaccess
  <IfModule mod_rewrite.c>
      RewriteEngine on
      RewriteRule    ^$    webroot/    [L]
      RewriteRule    (.*) webroot/$1    [L]
   </IfModule>

./cake/console/templates/skel/webroot/.htaccess
  <IfModule mod_rewrite.c>
      RewriteEngine On
      RewriteCond %{REQUEST_FILENAME} !-d
      RewriteCond %{REQUEST_FILENAME} !-f
      RewriteRule ^(.*)$ index.php?url=$1 [QSA,L]
  </IfModule>
# ===========================================================
If you would like to convert numbers into just the uppercase alphabet base and 
vice-versa (e.g. the column names in a Microsoft Windows Excel sheet..A-Z, AA-ZZ, 
AAA-ZZZ, ...), the following functions will do this.

/**
 * Converts an integer into the alphabet base (A-Z).
 *
 * @param int $n This is the number to convert.
 * @return string The converted number.
 * @author Theriault
 * 
 */
function num2alpha($n) {
    $r = '';
    for ($i = 1; $n >= 0 && $i < 10; $i++) {
        $r = chr(0x41 + ($n % pow(26, $i) / pow(26, $i - 1))) . $r;
        $n -= pow(26, $i);
    }
    return $r;
}
/**
 * Converts an alphabetic string into an integer.
 *
 * @param int $n This is the number to convert.
 * @return string The converted number.
 * @author Theriault
 * 
 */
function alpha2num($a) {
    $r = 0;
    $l = strlen($a);
    for ($i = 0; $i < $l; $i++) {
        $r += pow(26, $i) * (ord($a[$l - $i - 1]) - 0x40);
    }
    return $r - 1;
}

Microsoft Windows Excel stops at IV (255), but this function can handle much larger. 
However, English words will start to form after a while and some may be offensive, 
so be careful.
# ===========================================================
CodeIgniter Router

application/config/routes.php 

$route['default_controller'] = "welcome";
$route['404_override'] = '';

// custom routes
$route['register'] = 'users/register';
$route['login'] = 'users/login';
$route['logout'] = 'users/logout';
$route['search'] = 'users/search';
$route['profile'] = 'users/profile';
$route['settings'] = 'users/settings';
$route['friend/(:num)'] = 'users/friend/$1';
$route['unfriend/(:num)'] = 'users/unfriend/$1';
$route['users/edit/(:num)'] = 'users/edit/$1';
$route['users/delete/(:num)'] = 'users/delete/$1';
$route['users/undelete/(:num)'] = 'users/undelete/$1';
$route['fonts/(:any)'] = 'files/fonts/$1';
$route['thumbnails/(:num)'] = 'files/thumbnails/$1';
$route['files/delete/(:num)'] = 'files/delete/$1';
$route['files/undelete/(:num)'] = 'files/undelete/$1';

INI route configuration

route[register] = users/register
route[login] = users/login
route[logout] = users/logout
...
route[thumbnails/:id] = files/thumbnails/:id
route[files/delete/:id] = files/delete/:id
route[files/undelete/:id] = files/undelete/:id

require webroot('app/config/routes.php');
#require appdir('config/routes.php');
# ===========================================================
function mb($str, $left, $right)
{
    // if操作
    $str = preg_replace( "/".$left."if([^{]+?)".$right."/", "<?php if \\1 { ?>", $str );
    $str = preg_replace( "/".$left."else".$right."/", "<?php } else { ?>", $str );
    $str = preg_replace( "/".$left."elseif([^{]+?)".$right."/", "<?php } elseif \\1 { ?>", $str );

    // foreach操作
    $str = preg_replace("/".$left."foreach([^{]+?)".$right."/","<?php foreach \\1 { ?>",$str);
    $str = preg_replace("/".$left."\/foreach".$right."/","<?php } ?>",$str);

    // for操作
    $str = preg_replace("/".$left."for([^{]+?)".$right."/","<?php for \\1 { ?>",$str);
    $str = preg_replace("/".$left."\/for".$right."/","<?php } ?>",$str);

    // 输出变量
    $str = preg_replace( "/".$left."(\\$[a-zA-Z_\x7f-\xff][a-zA-Z0-9_$\x7f-\xff\[\]\'\']*)".$right."/", "<?php echo \\1;?>", $str );

    // 常量输出
    $str = preg_replace( "/".$left."([A-Z_\x7f-\xff][A-Z0-9_\x7f-\xff]*)".$right."/s", "<?php echo \\1;?>", $str );

    // 标签解析
    $str = preg_replace ( "/".$left."\/if".$right."/", "<?php } ?>", $str );

    $pattern = array('/'.$left.'/', '/'.$right.'/');
    $replacement = array('<?php ', ' ?>');

    return preg_replace($pattern, $replacement, $str);
}
# ===========================================================
CakePHP ACL

在CakePHP中，当添加或修改用户和组时，CakePHP会自动添加或修改ARO，这是因为在User Model类和Group Model类中加了下面一句：

 var $actsAs = array('Acl' => array('type' => 'requester'));

所以，没有必要这么做

      // add ARO for new group
      $aro = $this->Acl->Aro;
      $aro->save(array(
        'parent_id'   => null, 
        'model'       => 'Group',
        'foreign_key' => $this->Group->id,
        'alias'       => $this->data['Group']['name']
        //'alias'     => 'Group:'.$this->Group->id
      ));
# ===========================================================
<?php

// You can build an events system as simple or complex as you want it.

/**
 * Attach (or remove) multiple callbacks to an event and trigger those callbacks when 
 * that event is called.
 *
 * @param string $event name
 * @param mixed $value the optional value to pass to each callback
 * @param mixed $callback the method or function to call - FALSE to remove all callbacks for event
 */
function event($event, $value = NULL, $callback = NULL)
{
    static $events;

    // Adding or removing a callback?
    if ($callback !== NULL)
    {
        if ($callback) {
            $events[$event][] = $callback;
        }
        else {
            unset($events[$event]);
        }
    }
    elseif (isset($events[$event])) { // Fire a callback
        foreach($events[$event] as $function) {
            $value = call_user_func($function, $value);
        }
        return $value;
    }
}

// 1) Add an event
event('filter_text', NULL, function($text) { return htmlspecialchars($text); });
// add more as needed
event('filter_text', NULL, function($text) { return nl2br($text); });
// OR like this
//event('filter_text', NULL, 'nl2br');

// 2) Then call it like this
$text = event('filter_text', $_POST['text']);

// 3) Or remove all callbacks for that event like this
event('filter_text', null, false);

# ===========================================================
CakePHP schema

    [_schema] => Array
        (
            [username] => Array
                (
                    [type] => string
                    [null] => 1
                    [default] => 
                    [length] => 100
                    [collate] => utf8_general_ci
                    [charset] => utf8
                )

            [email] => Array
                (
                    [type] => string
                    [null] => 1
                    [default] => 
                    [length] => 100
                    [collate] => utf8_general_ci
                    [charset] => utf8
                )

            [password] => Array
                (
                    [type] => string
                    [null] => 1
                    [default] => 
                    [length] => 100
                    [collate] => utf8_general_ci
                    [charset] => utf8
                )

            [language_id] => Array
                (
                    [type] => integer
                    [null] => 1
                    [default] => 
                    [length] => 11
                    [key] => index
                )

            [id] => Array
                (
                    [type] => integer
                    [null] => 
                    [default] => 
                    [length] => 11
                    [key] => primary
                )

            [group_id] => Array
                (
                    [type] => integer
                    [null] => 1
                    [default] => 
                    [length] => 11
                )

            [firstname] => Array
                (
                    [type] => string
                    [null] => 1
                    [default] => 
                    [length] => 50
                    [collate] => utf8_general_ci
                    [charset] => utf8
                )

            [lastname] => Array
                (
                    [type] => string
                    [null] => 1
                    [default] => 
                    [length] => 50
                    [collate] => utf8_general_ci
                    [charset] => utf8
                )

            [disable] => Array
                (
                    [type] => boolean
                    [null] => 1
                    [default] => 0
                    [length] => 1
                )
        )

# ===========================================================
PHP设想

PHP中用＃作单行注释有些多余，用//作单行注释已经足够了。完全可以为＃派更好的用场。

PHP中用.作字符串相加有些失误，这样迫使类方法调用只能用->，不太完美。
理想的做法是：用＃作字符串相加，用.作方法调用和命名空间等的分割符。
$str = 'abc' # 'def';
$obj.method();

# ===========================================================
multi level config, override up level

config.php
<?php

return array(
  'debug' => false,
  'cache' => 24*60*60,

  'db_login' => 'user',
  'db_password' => 'password',
  'db_server' => '127.0.0.1',
  'db_database' => 'database',
  'db_engine' => 'mysql',
);

config.qa.php - override something in config.php

<?php

$base_conf = include __DIR__ . '/config.php';
$test_conf =  array(
  'debug' => true,
  'cache' => 0,

  'db_login' => 'testuser',
  'db_password' => 'testpassword',
  'db_server' => '127.0.0.1',
  'db_database' => 'testdatabase',
  'db_engine' => 'mysql',
);

return array_merge($base_conf, $test_conf);
# ===========================================================
<?php
// php curl 模拟ftp上传
 
function upload($dir,$src,$dest)
{
    $ch = curl_init();
    $fp = fopen($src, 'r');
    curl_setopt($ch, CURLOPT_URL, 'ftp://user:pwd@host/interpretation/'.$dir .'/'. $dest);
    curl_setopt($ch, CURLOPT_UPLOAD, 1);
    curl_setopt($ch, CURLOPT_INFILE, $fp);
    curl_setopt($ch, CURLOPT_INFILESIZE, filesize($src));
    curl_exec ($ch);
    $error_no = curl_errno($ch);
    curl_close ($ch);
    if ($error_no != 0) {
      return 0;
    } else {
      return 1;
    }
}


//php简单缩略图类|image.class.php

使用方法：
$img = new iamge;
$img->resize('dstimg.jpg', 'srcimg.jpg', 300, 400);
说明：这个是按照比例缩放，dstimg.jpg是目标文件，srcimg.jpg是源文件，后面的是目标文件的宽和高
$img->thumb('dstimg.jpg', 'scrimg.jpg', 300, 300);
说明：这个是按照比例缩略图，比如常用在头像缩略图的时候，dstimg.jpg是目标文件，srcimg.jpg是源文件，后面的是目标文件的宽和高
这个是针对GD库才这样麻烦的，如果采用Imagick的话，就只需要两个函数就实现上面的功能，去查下文档就找到了。

<?php
class image{
     
    public function resize($dstImg, $srcImg, $dstW, $dstH){
        list($srcW, $srcH) = getimagesize($srcImg);
        $scale = min($dstW/$srcW, $dstH/$srcH);
        $newW = round($srcW * $scale);
        $newH = round($srcH * $scale);
        $newImg = imagecreatetruecolor($newW, $newH);
        $srcImg = imagecreatefromjpeg($srcImg);
        imagecopyresampled($newImg, $srcImg, 0, 0, 0, 0, $newW, $newH, $srcW, $srcH);
        imagejpeg($newImg, $dstImg);
    }
     
    public function thumb($dstImg, $srcImg, $dstW, $dstH){
        list($srcW, $srcH) = getimagesize($srcImg);
        $scale = max($dstW/$srcW, $dstH/$srcH);
        $newW = round($dstW/$scale);
        $newH = round($dstH/$scale);
        $x = ($srcW - $newW)/2;
        $y = ($srcH - $newH)/2;
        $newImg = imagecreatetruecolor($dstW, $dstH);
        $srcImg = imagecreatefromjpeg($srcImg);
        imagecopyresampled($newImg, $srcImg, 0, 0, $x, $y, $dstW, $dstH, $newW, $newH);
        imagejpeg($newImg, $dstImg);
    }
}
 
//为了解决不同图片格式的问题
function createFromType($type, $srcImg){
    $function = "imagecreatefrom$type";
    return $function($srcImg);
}
# ===========================================================
<?php if (!defined('BASEPATH')) exit('No direct script access allowed');

@@

  // The name of THIS file
  define('SELF',     pathinfo(__FILE__, PATHINFO_BASENAME));
  define('BASEPATH', str_replace('\\', '/', $system_path));
  define('SYSDIR',   trim(strrchr(trim(BASEPATH, '/'), '/'), '/'));
  define('APPPATH',  $application_folder.'/');
  define('VIEWPATH', APPPATH.'views/' );

  // Path to the front controller (this file)
  define('FCPATH', str_replace(SELF, '', __FILE__));

  // Call the requested method.
  // Any URI segments present (besides the class/function) will be passed to the method for convenience
  call_user_func_array(array($class, $method), array_slice($uri->segments, 2));
  call_user_func_array(array(&$controller, $method), array_slice($routes->fetch(true), 2));

@@

  if (isset($_SERVER['HTTP_HOST'])) {
    $base_url = (!empty($_SERVER['HTTPS']) && strtolower($_SERVER['HTTPS']) !== 'off') ? 'https' : 'http';
    $base_url .= '://'.$_SERVER['HTTP_HOST'].str_replace(basename($_SERVER['SCRIPT_NAME']), '', $_SERVER['SCRIPT_NAME']);
  }
  else {
    $base_url = 'http://localhost/';
  }
# ===========================================================
function BadIdeaID() { return uniqid() . '_' . md5(mt_rand()); }
echo uniqid(php_uname('n'), true);
$better_token = uniqid(md5(rand()), true);
$uniqueId= time().'-'.mt_rand();
$s = uniqid(time(), true);

$id = rand(1000, 9999) . rand(1000, 9999); // mt_rand()
// this is better than rand(1000,9999): str_pad(mt_rand(0, 9999), 4, 0);

$id = uniqid(rand(), true);
$u = sprintf("%010s-%s", time(), uniqid(true));
$u = time().'-'.uniqid(true);

function generate_uid(){
 return md5(mktime()."-".rand()."-".rand());
}

while($i < $characters) 
{ 
    $code .= substr($possible, mt_rand(0, strlen($possible)-1), 1);
    $i++;
}
# ===========================================================
/**
 * Load the Core System Files
 */
// glob() is much slower so we use opendir...
if ($dh = opendir(INCLUDES_DIR)) {
  while (($file = readdir($dh)) !== false) {
    if (preg_match("/.php/", $file)) {
      $file = INCLUDES_DIR. $file;
      //Include the file
      require_once($file);
    }
  }
  closedir($dh);
} else {
  die('Couldn\'t load the system files');
}


define('START_MEMORY_USAGE', memory_get_usage());
# ===========================================================
set_include_path(get_include_path() . PATH_SEPARATOR . '../../.');    

function __autoload($className)
{
    $filePath = str_replace('_', DIRECTORY_SEPARATOR, $className) . '.php';
    $includePaths = explode(PATH_SEPARATOR, get_include_path());
    foreach ($includePaths as $includePath) {
        if (file_exists($includePath . DIRECTORY_SEPARATOR . $filePath)) {
            require_once $filePath;
            return;
        }
    }
}
# ===========================================================
CakePHP ACL

// User/id 语法
$this->Acl->allow(array('User'=>array('id'=>55)), 'Lang-1'); // en_US
$this->Acl->allow(array('User'=>array('id'=>55)), 'Lang-2'); // es_US
$this->Acl->allow(array('User'=>array('id'=>55)), 'Lang-3'); // de_DE

// model/foreign_key 语法
$this->Acl->allow(array('model'=>'User', 'foreigin_key'=>55), 'Lang-1');

// 检查权限
if ($this->Acl->check(array('User'=>array('id'=>55)), 'Lang-1'))
  echo "Allowed";
else
  echo "Denied";

// 应充分利用aros.alias，能带来方便
UPDATE aros SET alias=CONCAT(model,':',foreign_key);
# ===========================================================
