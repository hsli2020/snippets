<?php

class Request
{
    protected $parameters;
    protected $sanitize_commands;

    public function __construct()
    {
        $this->sanitize_commands = array(
            'clean'     => function ($var) { return htmlentities($var, ENT_QUOTES); },
            'boolean'   => function ($var) { return ((bool) $var); },
            'int'       => function ($var) { return intval($var); },
            'raw'       => function ($var) { return $var; },
            'striphtml' => function ($var) { return htmlentities(strip_tags($var), ENT_QUOTES); },
            'urlencode' => function ($var) { return urlencode($var); },
        );
    }

    public function getParameter($name, $sanitize = null)
    {
        if (!array_key_exists($name, $this->parameters))
            return false;

        $result = $this->parameters[$name];
        if ($sanitize !== null) {
            // Allow sanitization of inputs when the parameter is being fetched.
            if (array_key_exists($sanitize, $this->sanitize_commands)) {
                $result = $this->sanitize_commands[$sanitize]($result);
            }
        }

        return $result;
    }
}
