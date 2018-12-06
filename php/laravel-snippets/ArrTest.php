<?php

class Arr
{
    public static function accessible($value)
    {
        return is_array($value) || $value instanceof ArrayAccess;
    }

    public static function dot($array, $prepend = '')
    {
        $results = [];

        foreach ($array as $key => $value) {
            if (is_array($value)) {
                $results = array_merge($results, static::dot($value, $prepend.$key.'.'));
            } else {
                $results[$prepend.$key] = $value;
            }
        }

        return $results;
    }

    public static function exists($array, $key)
    {
        if ($array instanceof ArrayAccess) {
            return $array->offsetExists($key);
        }

        return array_key_exists($key, $array);
    }

    public static function flatten($array, $depth = INF)
    {
        $result = [];

        foreach ($array as $item) {

            if (is_array($item)) {
                if ($depth === 1) {
                    $result = array_merge($result, $item);
                    continue;
                }

                $result = array_merge($result, static::flatten($item, $depth - 1));
                continue;
            }

            $result[] = $item;
        }

        return $result;
    }

    public static function get($array, $key, $default = null)
    {
        if (! $array) {
            return value($default);
        }

        if (is_null($key)) {
            return $array;
        }

        if (static::exists($array, $key)) {
            return $array[$key];
        }

        foreach (explode('.', $key) as $segment) {
            if (static::accessible($array) && static::exists($array, $segment)) {
                $array = $array[$segment];
            } else {
                return value($default);
            }
        }

        return $array;
    }

    public static function has($array, $key)
    {
        if (! $array) {
            return false;
        }

        if (is_null($key)) {
            return false;
        }

        if (static::exists($array, $key)) {
            return true;
        }

        foreach (explode('.', $key) as $segment) {
            if (static::accessible($array) && static::exists($array, $segment)) {
                $array = $array[$segment];
            } else {
                return false;
            }
        }

        return true;
    }

    public static function isAssoc(array $array)
    {
        $keys = array_keys($array);

        return array_keys($keys) !== $keys;
    }

    public static function only($array, $keys)
    {
        return array_intersect_key($array, array_flip((array) $keys));
    }

    public static function set(&$array, $key, $value)
    {
        if (is_null($key)) {
            return $array = $value;
        }

        $keys = explode('.', $key);

        while (count($keys) > 1) {
            $key = array_shift($keys);

            // If the key doesn't exist at this depth, we will just create an empty array
            // to hold the next value, allowing us to create the arrays to hold final
            // values at the correct depth. Then we'll keep digging into the array.
            if (! isset($array[$key]) || ! is_array($array[$key])) {
                $array[$key] = [];
            }

            $array = &$array[$key];
        }

        $array[array_shift($keys)] = $value;

        return $array;
    }
}

const EOL = (PHP_SAPI == 'cli') ? PHP_EOL : '<br/>';

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


$arr = [
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

#vd(Arr::get($arr, 'db.master.host'));
#Arr::set($arr, 'db.master.host', 'blabla-value');
#vd(Arr::get($arr, 'db.master.host'));

vd(Arr::dot($arr));
#vd(Arr::flatten($arr));

#$ar = [ 'a' => 11, 'b' => 22, 'c' => 33, 'd' => 44 ];
#vd(Arr::only($ar, ['b', 'c']));

