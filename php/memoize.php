<?php

function memoize($target) {
    static $memo = new WeakMap;

    return new class ($target, $memo) {
        function __construct(
            protected $target,
            protected &$memo,
        ) {}

        function __call($method, $params)
        {
            $this->memo[$this->target] ??= [];

            $signature = $method . crc32(json_encode($params));

            return $this->memo[$this->target][$signature]
               ??= $this->target->$method(...$params);
        }
    };
}