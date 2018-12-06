<?php
/**
 * 简单的PHP MYSQL PDO操作类 mysql.class.php, 功能简单粗糙，实用！
 * http://aowana.sinapp.com
 * 日期：2015-04-01
 */
class mysql
{
    public $pdo = null;
    public $results = null;
 
    public function find($sql, $array=array())
    {
        $ok = $this->process($sql, $array);
        if ($ok)
        {
            $this->results->setFetchMode(PDO::FETCH_ASSOC);
            $data = $this->results->fetch();
            return $data;
        }

        return false;
    }
 
    public function findAll($sql, $array=array())
    {
        $ok = $this->process($sql, $array);
        if ($ok)
        {
            $this->results->setFetchMode(PDO::FETCH_ASSOC);
            $data = $this->results->fetchAll();
            return $data;
        }

        return array();
    }
 
    public function update($sql, $array=array())
    {
        $ok = $this->process($sql, $array);
        if ($ok === false)
            return -1;//执行出错返回-1
        else if ($ok)
            return $this->results->rowCount();
        else
            return 0;
    }
 
    public function insert($sql, $array=array())
    {
        $ok = $this->process($sql, $array);
        if ($ok)
        {
            $id = $this->pdo->lastInsertId();
            $id = $id ? $id : 1;
            return $id;
        }

        return false;
    }
 
    public function delete($sql, $array=array())
    {
        $ok = $this->process($sql, $array);

        if ($ok === false)
            return -1; // 执行出错返回-1
        else if ($ok)
            return $this->results->rowCount();
        else
            return 0;
    }
 
    public function query($sql, $array=array())
    {
        return $this->process($sql, $array);
    }
 
    private function process($sql, $array)
    {
        if (is_null($this->pdo))
            $this->connect();

        $this->results = $this->pdo->prepare($sql);

        // print_r($this->pdo->errorInfo());
        // print_r($this->results->errorInfo());
 
        return $this->results->execute($array);
    }
 
    private function connect()
    {
        try
        {
            $this->pdo = new PDO(
                'mysql:host=' . MYSQL_HOST . 
                ';port=' . MYSQL_PORT . 
                ';dbname=' . MYSQL_DATABASE . 
                ';charset=utf8', 
                MYSQL_USERNAME, MYSQL_PASSWORD);
        }
        catch (PDOException $error)
        {
            $html = $error->getMessage();
            //SAE Mail to Master
            http503();
        }
    }
}
