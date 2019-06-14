<?php

//https://blog.csdn.net/hizzana/article/details/53212323 

namespace App\Support\Facades {
    class Facades {
        public function __call($name, $params) {
            return call_user_func_array([$this, $name], $params);
        } 
        public static function __callStatic($name, $params) {
            //echo __METHOD__, PHP_EOL;
            return call_user_func_array([new static(), $name], $params);
        }  
    }
    
    class Cache extends Facades {
        protected function fn($a, $b) {
            echo "function parameters: ${a} and ${b}\n";    
        }
        protected function static_fn($a, $b) {
            echo "static function parameters: ${a} and ${b}\n";      
        }
    }
}

namespace {
    class Autoload {
        public $aliases;
        public function __construct($aliases = []) {
            $this->aliases = $aliases;
        }
        public function register() {
            spl_autoload_register([$this, 'load'], true, true);
            return $this;
        }
        public function load($alias) {
            if (isset($this->aliases[$alias])) {
                return class_alias($this->aliases[$alias], $alias);
            }    
        }
    }
    
    $aliases = [
        'Cache' => App\Support\Facades\Cache::class,
    ];

    $autoloader = (new Autoload($aliases))->register();

    Cache::fn(3,6);
    Cache::static_fn(4,7);
}
