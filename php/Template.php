<?php

class Template
{
    protected $dir;
    protected $vars;
 
    public function __construct($dir = "") {
        $this->dir = (substr($dir, -1) == "/") ? $dir : $dir . "/";
        $this->vars = array();
    }
 
    public function __set($var, $value) {
        $this->vars[$var] = $value;
    }
 
    public function __get($var) {
        return $this->vars[$var];
    }
 
    public function __isset($var) {
        return isset($this->vars[$var]);
    }
 
    public function set() {
        $args = func_get_args();
        if (func_num_args() == 2) {
            $this->__set($args[0], $args[1]);
        }
        else {
            foreach ($args[0] as $var => $value) {
                $this->__set($var, $value);
            }
        }
    }
 
    public function out($template, $asString = false) {
        ob_start();
    #   extract($this->vars); // I added this
        require $this->dir . $template . ".php";
        $content = ob_get_clean();
 
        if ($asString) {
            return $content;
        }
        else {
            echo $content;
        }
    }
}

$t = new Template('/Users/hansonli/tmp');
// setting a value as if it were a property
$t->greeting = "Hello World!";
// setting a value with set()
$t->set("number", 42);
// setting multiple values with set()
$t->set(array(
    "foo" => "zip",
    "bar" => "zap"
));
// render template
$t->out("mytpl");
/*
<!DOCTYPE html>
<html lang="en">
 <head>
  <meta charset="utf-8">
 </head>
 <body>
  <div role="main">
   <h1><?=$this->greeting;?></h1>
   <h2><?=$this->number;?></h2>
   <h3><?=$this->foo;?></h3>
   <h4><?=$this->bar;?></h4>
  </div>
 </body>
</html>
*/
