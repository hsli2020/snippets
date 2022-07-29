<?php

PHP Snippets

/** 
 * Send a POST requst using cURL 
 * @param string $url to request 
 * @param array $post values to send 
 * @param array $options for cURL 
 * @return string 
 */ 
function curl_post($url, array $post = NULL, array $options = array()) 
{ 
    $defaults = array( 
        CURLOPT_POST => 1, 
        CURLOPT_HEADER => 0, 
        CURLOPT_URL => $url, 
        CURLOPT_FRESH_CONNECT => 1, 
        CURLOPT_RETURNTRANSFER => 1, 
        CURLOPT_FORBID_REUSE => 1, 
        CURLOPT_TIMEOUT => 4, 
        CURLOPT_POSTFIELDS => http_build_query($post) 
    ); 

    $ch = curl_init(); 
    curl_setopt_array($ch, ($options + $defaults)); 
    if( ! $result = curl_exec($ch)) 
    { 
        trigger_error(curl_error($ch)); 
    } 
    curl_close($ch); 
    return $result; 
    } 

# ===========================================================
DealTap

php text to image word wrap

http://stackoverflow.com/questions/9870287/is-there-a-word-wrap-function-for-gd2-in-php

$words = explode(" ",$text);
$wnum = count($words);
$line = '';
$text='';
for($i=0; $i<$wnum; $i++){
  $line .= $words[$i];
  $dimensions = imagettfbbox($font_size, 0, $font_file, $line);
  $lineWidth = $dimensions[2] - $dimensions[0];
  if ($lineWidth > $maxwidth) {
    $text.=($text != '' ? '|'.$words[$i].' ' : $words[$i].' ');
    $line = $words[$i].' ';
  }
  else {
    $text.=$words[$i].' ';
    $line.=' ';
  }
}
# ===========================================================
DealTap

app-ui/public/assets/js/app/controls/fields/properties/StringPropertyField.js

fpr($pdf->getPageWidth(), $pdf->getPageHeight());
	210.00014444444
	297.00008333333


A4 measures 210 × 297 millimeters or 8.27 × 11.69 inches.
In PostScript, its dimensions are rounded off to 595 × 842 points.
A point equals 1/72 of inch, that is to say about 0.35 mm (an inch being 2.54 cm). 


$pdf = new TCPDF(PDF_PAGE_ORIENTATION, PDF_UNIT, array(251.27, 325.24), true, 'UTF-8', false);


public static function getPageSizeFromFormat($format) {
switch (strtoupper($format)) {
    // ISO 216 A Series + 2 SIS 014711 extensions
    case 'A0' : {$pf = array( 2383.937, 3370.394); break;}
    case 'A1' : {$pf = array( 1683.780, 2383.937); break;}
    case 'A2' : {$pf = array( 1190.551, 1683.780); break;}
    case 'A3' : {$pf = array(  841.890, 1190.551); break;}
    case 'A4' : {$pf = array(  595.276,  841.890); break;}
    case 'A5' : {$pf = array(  419.528,  595.276); break;}
    case 'A6' : {$pf = array(  297.638,  419.528); break;}


// Image method signature:
// Image($file, $x='', $y='', $w=0, $h=0, $type='', $link='', $align='', $resize=false, $dpi=300, $palign='', $ismask=false, $imgmask=false, $border=0, $fitbox=false, $hidden=false, $fitonpage=false)


There are 72 points per inch; if it is sufficient to assume 96 pixels per inch, the formula is rather simple:

points = pixels * 72 / 96


There are 72 points in an inch (that is what a point is, 1/72 of an inch)
on a system set for 150dpi, there are 150 pixels per inch.
1 in = 72pt = 150px (for 150dpi setting)

To convert point point to pixel it is not that difficult. Here are the two formulas you can use to do the conversion.


To convert pixel to point:
points = pixel * 72 / 96

To convert point to pixel:
pixels = point * 96 / 72

1 px = 0.75 point; 1 point = 1.333333 px

alias f='find . -type f | grep'
alias ff='find . -type f | vi-'
alias mig='/media/sf_src/app-core/vendor/bin/phinx migrate -c /media/sf_src/app-core/phinx.yml'
alias pgd='pg_dump --schema-only > /media/sf_src/files/schema-full'
alias topdf='libreoffice --headless --invisible --convert-to pdf:writer_pdf_Export'
alias ..='cd ..'
alias ...='cd ../..'
alias ....='cd ../../..'
alias .....='cd ../../../..'

/src/code/app-ui/public/assets/js/app/services/AuthService.js :Bearer
/src/code/app-core/app/system/services/AuthService.php :Bearer

sudo /etc/init.d/nginx restart
sudo /etc/init.d/postgresql restart
# ===========================================================
PHP Snippets

    public function getHumanCreatedAt()
    {
        $diff = time() - $this->created_at;
        if ($diff > (86400 * 30)) {
            return date('M \'y', $this->created_at);
        } else {
            if ($diff > 86400) {
                return ((int)($diff / 86400)) . 'd ago';
            } else {
                if ($diff > 3600) {
                    return ((int)($diff / 3600)) . 'h ago';
                } else {
                    return ((int)($diff / 60)) . 'm ago';
                }
            }
        }
    }

	/**
	 * Convert an object to an array
	 * @param \stdClass $object
	 * @return array
	 *
	 * @see http://stackoverflow.com/questions/18576762/php-stdclass-to-array
	 */
	protected function object_to_array($object) {
		return json_decode(json_encode($object), true);
	}
# ===========================================================
DealTap

http://api.localhost/api/dev/seed?pk=6DK2c8C5yDjY7usG&qt=1

http://api.hanson.dealtap.ca/api/dev/invalidate-roles?pk=6DK2c8C5yDjY7usG&qt=1465308939

http://api.localhost//api/dev/invalidate-roles?pk=6DK2c8C5yDjY7usG&qt=1

alias phinx='/src/code/app-core/vendor/bin/phinx -c /src/code/app-core/phinx.yml'
alias upcfg='cp /src/code/app-core/app/config/samples/config.dev.php /src/code/app-core/app/config/config.php'
alias mig='/src/code/app-core/vendor/bin/phinx -c /src/code/app-core/phinx.yml'
alias up='(cd /src/code/app-core && git fetch origin dev:dev && cd /src/code/app-ui && git pull)'
alias nodemod="ll /usr/local/lib/node_modules/"

function v {
  if test -z "$1"
  then
    find . -type f -o -type d \( -name .git \) -prune
    return
  fi 
  vi -p $(find . -type f -o -type d \( -name .git \) -prune | grep $@)
}

<?php


$router = $di->get("router");


foreach ($application->getModules() as $key => $module) {
    $namespace = str_replace('Module','Controllers', $module["className"]);
    $router->add('/'.$key.'/:params', array(
        'namespace' => $namespace,
        'module' => $key,
        'controller' => 'index',
        'action' => 'index',
        'params' => 1
    ))->setName($key);
    $router->add('/'.$key.'/:controller/:params', array(
        'namespace' => $namespace,
        'module' => $key,
        'controller' => 1,
        'action' => 'index',
        'params' => 2
    ));
    $router->add('/'.$key.'/:controller/:action/:params', array(
        'namespace' => $namespace,
        'module' => $key,
        'controller' => 1,
        'action' => 2,
        'params' => 3
    ));
}


$di->set("router", $router);
# ===========================================================
PHP Snippet

    /**
     * Calculates password strength score
     *
     * @param   string $value - password
     * @return  int (1 = very weak, 2 = weak, 3 = medium, 4+ = strong)
     */
    private function countScore($value)
    {
        $score = 0;
        $hasLower = preg_match('![a-z]!', $value);
        $hasUpper = preg_match('![A-Z]!', $value);
        $hasNumber = preg_match('![0-9]!', $value);


        if ($hasLower && $hasUpper) {
            ++$score;
        }
        if (($hasNumber && $hasLower) || ($hasNumber && $hasUpper)) {
            ++$score;
        }
        if (preg_match('![^0-9a-zA-Z]!', $value)) {
            ++$score;
        }


        $length = mb_strlen($value);


        if ($length >= 16) {
            $score += 2;
        } elseif ($length >= 8) {
            ++$score;
        } elseif ($length <= 4 && $score > 1) {
            --$score;
        } elseif ($length > 0 && $score === 0) {
            ++$score;
        }


        return $score;
    }
# ===========================================================
deleteOldFiles

/** define the directory **/
$dir = "images/temp/";
$dir = "./";

/*** cycle through all files in the directory ***/
foreach (glob($dir."*") as $file) {

    /*** if file is 24 hours (86400 seconds) old then delete it ***/
    if (filemtime($file) < time() - 86400) {
        //unlink($file);
        echo $file, "\n";
    }
}

//files older than 24hours
$interval = strtotime('-24 hours');

foreach (glob($dir."*") as $file) {
    //delete if older
    if (filemtime($file) <= $interval) {
        //unlink($file);
        echo $file, "\n";
    }
}
# ===========================================================

const EOL = PHP_EOL;
/*
$root = 'app-core/app/system';

$iter = new RecursiveIteratorIterator(
    new RecursiveDirectoryIterator($root, RecursiveDirectoryIterator::SKIP_DOTS),
    RecursiveIteratorIterator::SELF_FIRST,
    RecursiveIteratorIterator::CATCH_GET_CHILD // Ignore "Permission denied"
);

$filetypes = array("php"); 
foreach ($iter as $path => $dir) {
    $filetype = pathinfo($dir, PATHINFO_EXTENSION); 
    if (in_array(strtolower($filetype), $filetypes)) { 
        echo 'include "', str_replace('\\', '/', $path), '";', EOL;
        //include $path;
    }
}

foreach (get_declared_classes() as $className) {
    echo $className, EOL;
}

ReflectionClass::export('Phalcon\Mvc\Model');
*/

//$r = new \ReflectionClass('Phalcon\Config\Adapter\Ini');
//echo $r->getName(), ' (', $r->getParentClass()->getName(), ') [', join(', ', $r->getInterfaceNames()), ']', EOL;

$tree = [];

$classes = get_declared_classes();
foreach($classes as $class) {
    if (strncmp($class, 'Phalcon', 7) != 0) {
        continue;
    }

    $rf = new \ReflectionClass($class);

    $className = $rf->getName();

    if ($rf->getInterfaceNames()) {
        $className .= ' ['. join(', ', $rf->getInterfaceNames()). ']';
    }

    if ($rf->getParentClass()) {
        $tree[$rf->getParentClass()->getName()][] = $className;
    } else {
        $tree[$className] = [];
    }
}

foreach($tree as $class => $subclasses) {
    echo ($subclasses) ? '+ ' : '- ';
    echo $class, EOL;
    foreach($subclasses as $subclass) {
        echo "\t- ", $subclass, EOL;
    }
}

# ===========================================================
const EOL = PHP_EOL;
const EOL = (PHP_SAPI == 'cli') ? PHP_EOL : '<br />';

define('CLI', (PHP_SAPI == 'cli') ? true : false);
define('EOL', CLI ? PHP_EOL : '<br />');

define('EOL', (PHP_SAPI == 'cli') ? PHP_EOL : '<br />');

date_default_timezone_set('America/New_York');

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


function getParameter($name)
{
    $tokens  = explode('.', $name);
   #$context = $this->parameters;
    $context = [
        'db' => [
            'master' => [
                'host' => 'master-host',
                'port' => 'master-port',
            ],
            'slave' => [
                'host' => 'slave-host',
                'port' => 'slave-port',
            ],
        ],
    ];

    while (null !== ($token = array_shift($tokens))) {
        if (!isset($context[$token])) {
            return null;
        }

        $context = $context[$token];
    }

    return $context;
}


/** 
 * Send a POST requst using cURL 
 * @param string $url to request 
 * @param array $post values to send 
 * @param array $options for cURL 
 * @return string 
 */ 
function curl_post($url, array $post = NULL, array $options = array()) 
{ 
    $defaults = array( 
        CURLOPT_POST => 1, 
        CURLOPT_HEADER => 0, 
        CURLOPT_URL => $url, 
        CURLOPT_FRESH_CONNECT => 1, 
        CURLOPT_RETURNTRANSFER => 1, 
        CURLOPT_FORBID_REUSE => 1, 
        CURLOPT_TIMEOUT => 4, 
        CURLOPT_POSTFIELDS => http_build_query($post) 
    ); 

    $ch = curl_init(); 

    curl_setopt_array($ch, ($options + $defaults)); 
    if (!$result = curl_exec($ch)) { 
        trigger_error(curl_error($ch)); 
    } 
    curl_close($ch); 

    return $result; 
} 

function randomString($type = 'alnum', $len = 8)
{
    $types = array(
        'basic'   => function() { return mt_rand(); },
        'alpha'   => 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'alnum'   => '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
        'numeric' => '0123456789',
        'nozero'  => '123456789',
        'unique'  => function() { return  md5(uniqid(mt_rand(), TRUE)); },
        'md5'     => function() { return  md5(uniqid(mt_rand(), TRUE)); },
        'encrypt' => function() { return sha1(uniqid(mt_rand(), TRUE)); },
        'sha1'    => function() { return sha1(uniqid(mt_rand(), TRUE)); },
    );

    if (isset($types[$type])) {
        $pool = $types[$type];
        if (is_callable($pool)) {
            return $pool();
        }

        if (is_string($pool)) {
            $str = '';
            for ($i=0; $i < $len; $i++) {
                $str .= substr($pool, mt_rand(0, strlen($pool) - 1), 1);
            }
            return $str;
        }
    }

    return '';
}

# echo randomString('basic'), EOL;
# echo randomString('alpha'), EOL;
# echo randomString('alnum'), EOL;
# echo randomString('numeric'), EOL;
# echo randomString('nozero'), EOL;
# echo randomString('unique'), EOL;
# echo randomString('md5'), EOL;
# echo randomString('encrypt'), EOL;
# echo randomString('sha1'), EOL;

function randomString0($strLen = 32)
{
    // Create our character arrays
    $chrs = array_merge(range('a', 'z'), range('A', 'Z'), range(0, 9));

    // Just to make the output even more random
    shuffle($chrs);

    // Create a holder for our string
    $randStr = '';

    // Now loop through the desired number of characters for our string
    for ($i=0; $i<$strLen; $i++) {
        $randStr .= $chrs[mt_rand(0, (count($chrs) - 1))];
    }

    return $randStr;
}

function deleteOldFiles($dir)
{
    $dir = "./";
    $interval = strtotime('-24 hours');

    foreach (glob($dir."*") as $file) {
        /** if file is 24 hours (86400 seconds) old then delete it **/
        //if (filemtime($file) < time() - 86400) {
        if (filemtime($file) < $interval) {
            echo $file, EOL;
            //unlink($file);
        }
    }
}

function removeOldFiles($folderName)
{
    $folderName = "./";
    $interval = strtotime('-5 days');
    foreach (new DirectoryIterator($folderName) as $fileInfo) {
        if ($fileInfo->isDir() || $fileInfo->isLink()) {
            continue;
        }
        #if (time() - $fileInfo->getCTime() >= 2*24*60*60) {
        if ($fileInfo->getCTime() < $interval) {
            echo $fileInfo->getRealPath(), EOL;
            #unlink($fileInfo->getRealPath());
        }
    }
}

class Bar {
    public $x;
    public $y;
    public $z;

    protected function init($args) {
        $vars = get_object_vars($this);
        foreach ($vars as $name => $value) {
            if (isset($args[$name])) {
                $this->$name = $args[$name];
            }
        }
    }

    public function __construct($args) {
        $this->init($args);
    }

    public function output() { vd($this); }
}

# ===========================================================
http://hashids.org/php/

public static function stripComments($source)
{
    if (!function_exists('token_get_all')) {
        return $source;
    }
    $rawChunk ='';
    $output ='';
    $tokens = token_get_all($source);
    $ignoreSpace = false;
    for (reset($tokens); false !== $token = current($tokens); next($tokens)) {
        if (is_string($token)) {
            $rawChunk .= $token;
        } elseif (T_START_HEREDOC === $token[0]) {
            $output .= $rawChunk.$token[1];
            do {
                $token = next($tokens);
                $output .= $token[1];
            } while ($token[0] !== T_END_HEREDOC);
            $rawChunk ='';
        } elseif (T_WHITESPACE === $token[0]) {
            if ($ignoreSpace) {
                $ignoreSpace = false;
                continue;
            }
            $rawChunk .= preg_replace(array('/\n{2,}/S'),"\n", $token[1]);
        } elseif (in_array($token[0], array(T_COMMENT, T_DOC_COMMENT))) {
            $ignoreSpace = true;
        } else {
            $rawChunk .= $token[1];
            if (T_OPEN_TAG === $token[0]) {
                $ignoreSpace = true;
            }
        }
    }
    $output .= $rawChunk;
    return $output;
}

public function getRootDir()
{
    if (null === $this->rootDir) {
        $r = new \ReflectionObject($this);
        $this->rootDir = str_replace('\\','/', dirname($r->getFileName()));
    }
    return $this->rootDir;
}

    /**
     * Utility method that "unindents" the given $code when all its lines start
     * with a tabulation of four white spaces.
     *
     * @param  string $code
     *
     * @return string
     */
    private function unindentCode($code)
    {
        $formattedCode = $code;
        $codeLines = explode("\n", $code);

        $indentedLines = array_filter($codeLines, function ($lineOfCode) {
            return '' === $lineOfCode || '    ' === substr($lineOfCode, 0, 4);
        });

        if (count($indentedLines) === count($codeLines)) {
            $formattedCode = array_map(function ($lineOfCode) { return substr($lineOfCode, 4); }, $codeLines);
            $formattedCode = implode("\n", $formattedCode);
        }

        return $formattedCode;
    }
private function isAbsolutePath($file)
{
    if ($file[0] ==='/'|| $file[0] ==='\\'|| (strlen($file) > 3 && ctype_alpha($file[0])
        && $file[1] ===':'&& ($file[2] ==='\\'|| $file[2] ==='/')
    )
    || null !== parse_url($file, PHP_URL_SCHEME)
    ) {
        return true;
    }
    return false;
}

public static function getRelativePath($basePath, $targetPath)
{
    if ($basePath === $targetPath) {
        return'';
    }
    $sourceDirs = explode('/', isset($basePath[0]) &&'/'=== $basePath[0] ? substr($basePath, 1) : $basePath);
    $targetDirs = explode('/', isset($targetPath[0]) &&'/'=== $targetPath[0] ? substr($targetPath, 1) : $targetPath);
    array_pop($sourceDirs);
    $targetFile = array_pop($targetDirs);
    foreach ($sourceDirs as $i => $dir) {
        if (isset($targetDirs[$i]) && $dir === $targetDirs[$i]) {
            unset($sourceDirs[$i], $targetDirs[$i]);
        } else {
            break;
        }
    }
    $targetDirs[] = $targetFile;
    $path = str_repeat('../', count($sourceDirs)).implode('/', $targetDirs);
    return''=== $path ||'/'=== $path[0]
        || false !== ($colonPos = strpos($path,':')) && ($colonPos < ($slashPos = strpos($path,'/')) || false === $slashPos)
        ? "./$path" : $path;
}
# ===========================================================
function varToString($var)
{
        if (is_object($var)) {
            return sprintf('Object(%s)', get_class($var));
        }
        if (is_array($var)) {
            $a = array();
            foreach ($var as $k => $v) {
                $a[] = sprintf('%s => %s', $k, varToString($v));
            }
            return sprintf('Array(%s)', implode(', ', $a));
        }
        if (is_resource($var)) {
            return sprintf('Resource(%s)', get_resource_type($var));
        }
        if (null === $var) {
            return'null';
        }
        if (false === $var) {
            return'false';
        }
        if (true === $var) {
            return'true';
        }
        return (string) $var;
}

function slugify($string)
{
    return trim(preg_replace('/[^a-z0-9]+/', '-', strtolower(strip_tags($string))), '-');
}

function camelize($id)
{
    return preg_replace_callback('/(^|_|\.)+(.)/', function ($match) { return ('.' === $match[1] ? '_' : '').strtoupper($match[2]); }, $id);
}
    
function underscore($id)
{
    return strtolower(preg_replace(array('/([A-Z]+)([A-Z][a-z])/', '/([a-z\d])([A-Z])/'), array('\\1_\\2', '\\1_\\2'), strtr($id, '_', '.')));
}
# ===========================================================
function has_color_support()
{
    static $support;

    if (null === $support) {
        if (DIRECTORY_SEPARATOR == '\\') {
            $support = false !== getenv('ANSICON') || 'ON' === getenv('ConEmuANSI');
        } else {
            $support = function_exists('posix_isatty') && @posix_isatty(STDOUT);
        }
    }

    return $support;
}

function echo_style($style, $message)
{
    // ANSI color codes
    $styles = array(
        'reset' => "\033[0m",
        'red' => "\033[31m",
        'green' => "\033[32m",
        'yellow' => "\033[33m",
        'error' => "\033[37;41m",
        'success' => "\033[37;42m",
        'title' => "\033[34m",
    );
    $supports = has_color_support();

    echo($supports ? $styles[$style] : '').$message.($supports ? $styles['reset'] : '');
}

function echo_block($style, $title, $message)
{
    $message = ' '.trim($message).' ';
    $width = strlen($message);

    echo PHP_EOL.PHP_EOL;

    echo_style($style, str_repeat(' ', $width).PHP_EOL);
    echo_style($style, str_pad(' ['.$title.']',  $width, ' ', STR_PAD_RIGHT).PHP_EOL);
    echo_style($style, str_pad($message,  $width, ' ', STR_PAD_RIGHT).PHP_EOL);
    echo_style($style, str_repeat(' ', $width).PHP_EOL);
}

function echo_title($title, $style = null)
{
    $style = $style ?: 'title';

    echo PHP_EOL;
    echo_style($style, $title.PHP_EOL);
    echo_style($style, str_repeat('~', strlen($title)).PHP_EOL);
    echo PHP_EOL;
}

echo_title('Symfony2 Requirements Checker');

echo_style('green', '  '.$iniPath);
echo_style('warning', '  WARNING: No configuration file (php.ini) used by PHP!');
echo_style('red', 'E');
echo_style('green', '.');

echo_style('yellow', 'W');
echo_style('green', '.');

echo_block('success', 'OK', 'Your system is ready to run Symfony2 projects', true);
echo_block('error', 'ERROR', 'Your system is not ready to run Symfony2 projects', true);
echo_title('Fix the following mandatory requirements', 'red');

echo_title('Optional recommendations to improve your setup', 'yellow');
echo_style('title', 'Note');
echo_style('title', '~~~~');
echo_style('yellow', 'web/check.php');
# ===========================================================
$url = 'http://blog.wenxuecity.com/blog/frontend.php?page=0&act=articleList&blogId=35538&catId=56487';

function parseUrl($url)
{
    $parts = parse_url($url);

    if (isset($parts['path'])) {
        $parts['path'] = trim($parts['path'], '/');
        $segments = explode('/', $parts['path']);
        foreach ($segments as $key => $val) {
            $parts['path:'.($key+1)] = $val;
        }
    }

    if (isset($parts['query'])) {
        parse_str($parts['query'], $vars);
        unset($parts['query']);

        #$parts['query'] = $vars;
        foreach ($vars as $key => $val) {
            $parts['query.'.$key] = $val;
        }
    }

    $parts['md5'] = implode("/", str_split(substr(md5($url), 0, 6), 2)) . '/' . substr(md5($url), 6);
    #$parts['sha1'] = sha1($url);
    #$parts['sha256'] = hash('sha256', $url);
    $parts['domain'] = $parts['scheme'] . '://' . $parts['host'];
    $parts['date'] = date('Ymd');
    $parts['time'] = date('His');
    $parts['rand'] = rand(100, 1000);

    return $parts;
}

function urlToFilename($url, $fmask)
{
    $info = parseUrl($url);
    pr($info);

    $fname = $fmask;
    foreach ($info as $name => $value) {
        $fname = str_replace('{'.$name.'}', $value, $fname);
    }

    return $fname;
}

$fmask = '{host}/{path}/{query.blogId}/{query.catId}-{query.page}.html';
// pr(urlToFilename($url, $fmask));

$fmask = '{host}/{path:1}/{md5}.html';
$fmask = '{host}/{date}/{time}-{rand}.html';
pr(urlToFilename($url, $fmask));

$url = 'http://blog.wenxuecity.com/myblog/35538/56487.html';
$fmask = '{host}/{path:1}-{path:2}-{path:3}';
// pr(urlToFilename($url, $fmask));


easy to format html

$html = file_get_contents('myblog-35538-201510-238466.html');
$html = preg_replace("/>\s+</", ">\n<", $html);
$html = str_replace("><", ">\n<", $html);
echo html_entity_decode(strip_tags($html));
# ===========================================================
trait MagicSetter
{
    /**
     * Magic set method
     *
     * @param string $name
     * @param mixed $value
     *
     * @throws \RuntimeException
     */
    public function __set($name, $value)
    {
        $methodName = 'set'.ucfirst($name);
        if (method_exists($this, $methodName)) {
            $this->$methodName($value);
        } elseif (property_exists($this, $name)) {
            $this->$name = $value;
        } else {
            throw new \RuntimeException("Unable to set the property '$name' on this class: " . get_class($this));
        }
        return $this;
    }
}

trait MagicGetter
{
    /**
     * Magic get method
     *
     * @param string $name
     *
     * @return mixed
     */
    public function __get($name)
    {
        $methodName = 'get'.ucfirst($name);
        if (method_exists($this, $methodName)) {
            return $this->$methodName();
        }

        if (property_exists($this, $name)) {
            return $this->$name;
        }
        return null;
    }
}

class Bar
{
    use MagicSetter;
    use MagicGetter;

    private function setName($name) { $this->name = $name; }
}

$bar = new Bar();
$bar->name = 'John';

function runningInConsole()
{
    return php_sapi_name() == 'cli';
}

function startsWith($haystack, $needle)
{
    return $needle === "" || strrpos($haystack, $needle, -strlen($haystack)) !== FALSE;
}
// pr(startsWith('command', 'com'));
// pr(startsWith('aommand', 'com'));

$id = '123asd';

function getCacheFilename($id)
{
    $directory = '/app/cache';
    $extension = '.cache.php';

    $path = implode(str_split(md5($id), 12), DIRECTORY_SEPARATOR);
    $path = $directory . DIRECTORY_SEPARATOR . $path;

    return $path . DIRECTORY_SEPARATOR . $id . $extension;
}
// echo getCacheFilename($id), EOL;

function getDir($id, $levels_deep = 3)
{
    $filehash = md5($id);
    echo $filehash, CR;
    $dirname  = implode("/", str_split(substr($filehash, 0, $levels_deep*2), 2))
        . '/' . substr($filehash, $levels_deep*2);
    return $dirname;
}
// echo getDir($id, 3), EOL;

/**
 * Generate an MD5 hash string from the contents of a directory.
 *
 * @param string $directory
 * @return boolean|string
 */
function hashDirectory($directory)
{
    if (! is_dir($directory)) {
        return false;
    }
 
    $files = array();
    $dir = dir($directory);
 
    while (false !== ($file = $dir->read())) {
        if ($file != '.' and $file != '..') {
            if (is_dir($directory . '/' . $file)) {
                $files[] = hashDirectory($directory . '/' . $file);
            } else {
                $files[] = md5_file($directory . '/' . $file);
            }
        }
    }
 
    $dir->close();
 
    return md5(implode('', $files));
}
// echo hashDirectory('..');

function snake($value, $delimiter = '_')
{
    if (! ctype_lower($value)) {
        $value = strtolower(preg_replace('/(.)(?=[A-Z])/', '$1'.$delimiter, $value));

        $value = preg_replace('/\s+/', '', $value);
    }

    return $value;
}
// echo snake('EventDispatcher', '-'), EOL;

/**
 * @param  mix $value
 * @return string
 */
function encode($value)
{
    return urlencode(base64_encode(json_encode($value)));
}

/**
 * @param  string $value
 * @return mix
 */
function decode($value)
{
    return json_decode(base64_decode(urldecode($value)), true);
}
// pr(encode(1));
// pr(encode(0));
// $en = encode($arr); pr($en);
// $de = decode($en);  pr($de);
# ===========================================================
function quickRandom($length = 10)
{
    $pool = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    return substr(str_shuffle(str_repeat($pool, $length)), 0, $length);
}
// echo quickRandom(), EOL;

function randomString($length = 10)
{
    $seed = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    if ($length > strlen($seed)) {
        $seed = str_repeat($seed, $length / strlen($seed) + 1);
    }
    return substr(str_shuffle($seed), 0, $length);
}
// var_dump(randomString(100));

function strContains($haystack, $needles)
{
    foreach ((array) $needles as $needle) {
        if ($needle != '' && strpos($haystack, $needle) !== false) {
            return true;
        }
    }

    return false;
}
// pr( strContains('repeat', ['rep', 'eat']) );

function strEndsWith($haystack, $needles)
{
    foreach ((array) $needles as $needle) {
        if ((string) $needle === substr($haystack, -strlen($needle))) {
            return true;
        }
    }

    return false;
}
// pr( strEndsWith('repeat', ['rep', 'eat']) );

function strStartsWith($haystack, $needles)
{
    foreach ((array) $needles as $needle) {
        if ($needle != '' && strpos($haystack, $needle) === 0) {
            return true;
        }
    }

    return false;
}
// pr( strEndsWith('repeat', ['rep', 'eat']) );

// echo mb_convert_case('how are you doing', MB_CASE_TITLE, 'UTF-8');

function strMatch($pattern, $value)
{
    if ($pattern == $value) {
        return true;
    }

    $pattern = preg_quote($pattern, '#');

    // Asterisks are translated into zero-or-more regular expression wildcards
    // to make it convenient to check if the strings starts with the given
    // pattern such as "library/*", making any string check convenient.
    $pattern = str_replace('\*', '.*', $pattern).'\z';

    return (bool) preg_match('#^'.$pattern.'#', $value);
}
pr( strMatch('r*p*t', 'repeat') );

function cacheFile($key)
{
    $directory = '/storage/cache';

    $parts = array_slice(str_split($hash = md5($key), 2), 0, 2);

    return $directory.'/'.implode('/', $parts).'/'.$hash;
}
echo cacheFile('app/public/home'), EOL;

function only($array, $keys)
{
    return array_intersect_key($array, array_flip((array) $keys));
}

$arr = ['name'=>'john', 'pass'=>'12345', 'email'=>'123@mail.com', 'token'=>'asdf'];

pr(only($arr, ['name', 'pass', 'token']));

function except($array, $keys)
{
    return array_diff_key($array, array_flip((array)$keys));
}
pr(except($arr, ['pass', 'token']));
# ===========================================================
// only allow adults to signup
if ($this->params['dob'] > strtotime('18 years ago')) {
    $this->error_msg = "You must be 18 years of age or older to signup";
    return;
}
# ===========================================================
private function varToString($var)
{
    if (is_object($var)) {
        return sprintf('Object(%s)', get_class($var));
    }
    if (is_array($var)) {
        $a = array();
        foreach ($var as $k => $v) {
            $a[] = sprintf('%s => %s', $k, $this->varToString($v));
        }
        return sprintf('Array(%s)', implode(', ', $a));
    }
    if (is_resource($var)) {
        return sprintf('Resource(%s)', get_resource_type($var));
    }
    if (null === $var) {
        return'null';
    }
    if (false === $var) {
        return'false';
    }
    if (true === $var) {
        return'true';
    }
    return (string) $var;
}
# ===========================================================
PHPUnit Annotations

/**
 * @requires PHP 5.4.6
 * @requires PHPUnit 3.7
 * @requires function Imagick::readImage
 * @requires extension pdo_sqlite
 *
 * @group math
 *
 * @backupGlobals disabled
 * @backupStaticAttributes disabled
 * @depends testXXX
 * @dataProvider providerXY
 */
# ===========================================================
