<?php

function pipe($thing) {
    $fluent = new class ($thing) {
        public $calls = 0;
        public $ticks = 0;
        public $onTick;
        public $kill = false;
        public $fluent;

        public function __construct(public $thing) {}

        public function setFluentRef(&$fluent)
        {
            $this->fluent = $fluent;
        }

        public function &__invoke(...$params) {
            if (empty($params)) return $this->thing;

            $this->calls++;

            $before = [];
            $through = null;
            $after = [];

            foreach ($params as $key => $param) {
                if (! $through) {
                    if (is_callable($param)) $through = $param;
                    else $before[$key] = $param;
                } else {
                    $after[$key] = $param;
                }
            }

            $params = array_merge($before, [$this->thing], $after);

            $this->thing = $through(...$params);

            $this->ticks = 0;

            declare(ticks=1);

            return $this->fluent;
        }
    };

    $fluent->setFluentRef($fluent);

    $onTick = function () use (&$fluent) {
        if (is_string($fluent)) return;

        if ($fluent->ticks === 1 && $fluent->calls === 0) {
            $fluent->fluent = $fluent->thing;
        }

        $fluent->ticks++;
        $fluent->calls = 0;
    };

    register_tick_function($onTick);

    return $fluent;
}