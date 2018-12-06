<?php

class RedisSessionHandler implements SessionHandlerInterface
{
    public $ttl = 1800; // 30 minutes default
    protected $db;
    protected $prefix;
 
    public function __construct(PredisClient $db, $prefix = 'PHPSESSID:')
    {
        $this->db = $db;
        $this->prefix = $prefix;
    }
 
    public function open($savePath, $sessionName)
    {
        // No action necessary because connection is injected
        // in constructor and arguments are not applicable.
    }
 
    public function close()
    {
        $this->db = null;
        unset($this->db);
    }
 
    public function read($id)
    {
        $id = $this->prefix . $id;
        $sessData = $this->db->get($id);
        $this->db->expire($id, $this->ttl);
        return $sessData;
    }
 
    public function write($id, $data)
    {
        $id = $this->prefix . $id;
        $this->db->set($id, $data);
        $this->db->expire($id, $this->ttl);
    }
 
    public function destroy($id)
    {
        $this->db->del($this->prefix . $id);
    }
 
    public function gc($maxLifetime)
    {
        // no action necessary because using EXPIRE
    }
}

// Redis as a PHP Session Handler (20 June 2013)
//
// extension=redis.so
// 
// session.save_handler = redis
// session.save_path = "tcp://localhost:6379"
//
// or directly in your page:
// 
// ini_set('session.save_handler', 'redis');
// ini_set('session.save_path',    'tcp://127.0.0.1:6379');
