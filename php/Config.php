<?php

class Config
{
    protected static $data;

    public static function load($filename, $section = null)
    {
        $pathInfo = pathinfo($filename);

        if (!$section) {
            $section = $pathInfo['filename'];
        }

        $fileExt = strtolower($pathInfo['extension']);

        switch ($fileExt) {
        case 'ini':
            self::$data[$section] = parse_ini_file($filename, true);
            break;
        case 'php':
            self::$data[$section] = include $filename;
            break;
        case 'json':
            $json = file_get_contents($filename);
            self::$data[$section] = json_decode($json, true);
            break;
        case 'xml':
        case 'yml':
        case 'yaml':
        default:
            throw new Exception("Unsupported config file: " . $filename);
            break;
        }
        #print_r(self::$data);
    }

    public static function get($key)
    {
        $keys = explode('.', $key);
        $data = self::$data;

        while (null !== ($name = array_shift($keys))) {
            if (!isset($data[$name])) {
                return null;
            }

            $data = $data[$name];
        }

        return $data;
    }
}

const EOL = PHP_EOL;

function loadPhp()
{
    Config::load('db.php');
    Config::load('user.php');
}

function loadIni()
{
    Config::load('db.ini');
    Config::load('user.ini');
}

function loadJson()
{
    Config::load('db.json');
    Config::load('user.json');
}

function loadYml()
{
    Config::load('db.yml');
    Config::load('user.yml');
}

function showInfo()
{
    echo Config::get('db.master.host'), EOL;
    echo Config::get('db.master.port'), EOL, EOL;

    echo Config::get('db.slave.host'), EOL;
    echo Config::get('db.slave.port'), EOL, EOL;

#   print_r(Config::get('db.master'));
#   print_r(Config::get('db.slave'));

    echo Config::get('user.guest.username'), EOL;
    echo Config::get('user.guest.password'), EOL, EOL;

    echo Config::get('user.admin.username'), EOL;
    echo Config::get('user.admin.password'), EOL, EOL;

#   print_r(Config::get('user.guest'));
#   print_r(Config::get('user.admin'));
}

#loadPhp();
#loadIni();
#loadJson();
#loadYml();
showInfo();
