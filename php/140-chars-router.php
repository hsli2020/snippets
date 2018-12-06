<?php
// https://github.com/usmanhalalit/140-chars-router
// The router code
class R{
function a($r,callable $c){$this->r[$r]=$c;}
function e(){$s=$_SERVER;$i='PATH_INFO';$p=isset($s[$i])?$s[$i]:'/';$this->r[$p]();}
}
//


// Create 
$router = new R();

// Add a route with a callback function
$router->a('/a', 'callbackFunction');

// Add a route with a closure
$router->a('/b', function(){
    echo 'Hello B';
});

// Add homepage route
$router->a('/', function(){
	echo 'Hello World';
});

// Add route with class method
$router->a('/c', [new Foo, 'bar']);
// Add multiple slashed route with class method
$router->a('/c/d', [new Foo, 'bar']);

// Execute the route
$router->e();

// Callback handlers
function callbackFunction(){
	echo 'Hello A';
}

class Foo{
	function bar(){
		echo 'Hello Bar';
	}
}