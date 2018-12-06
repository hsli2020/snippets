<?php
/**
 * Twittee: A Dependency Injection Container in a Tweet
 * http://twittee.org/
 * 
 * Twittee is the smallest, and still useful, Dependency Injection Container in PHP;
 * it is also probably one of the first public software to use the newest anonymous 
 * functions support of PHP 5.3.
 * 
 * Packed in less than 140 characters, it fits in a tweet.
 * 
 * Despite its size, Twittee is a full-featured Dependency Injection Container with
 * support for object definitions, object injection and parameters.
 * 
 * Published in 2009 by Fabien Potencier, Twittee is in the Public Domain. Tweet me
 * if you find a bug! 
 */
class Container {
    protected $s=array();
    function __set($k, $c) { $this->s[$k]=$c; }
    function __get($k) { return $this->s[$k]($this); }
}

