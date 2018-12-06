<?php

namespace MyApp\Http;

class Curl
{
    /** @var resource cURL handle */
    private $ch;

    /** @var mixed The response */
    private $response = false;

    /**
     * @param string $url
     * @param array  $options
     */
    public function __construct($url, array $options = array())
    {
        $this->ch = curl_init($url);

        foreach ($options as $key => $val) {
            curl_setopt($this->ch, $key, $val);
        }

        curl_setopt($this->ch, CURLOPT_RETURNTRANSFER, true);
    }

    /**
     * Get the response
     * @return string
     * @throws \RuntimeException On cURL error
     */
    public function getResponse()
    {
         if ($this->response) {
             return $this->response;
         }

        $response = curl_exec($this->ch);
        $error    = curl_error($this->ch);
        $errno    = curl_errno($this->ch);

        if (is_resource($this->ch)) {
            curl_close($this->ch);
        }

        if (0 !== $errno) {
            throw new \RuntimeException($error, $errno);
        }

        return $this->response = $response;
    }

    /**
     * Let echo out the response
     * @return string
     */
    public function __toString()
    {
        return $this->getResponse();
    }
}

// usage
$curl = new \MyApp\Http\Curl('http://www.domain.com', array(
    CURLOPT_POSTFIELDS => array('username' => 'user1')
));

try {
    echo $curl;
} catch (\RuntimeException $ex) {
    die(sprintf('Http error %s with code %d', $ex->getMessage(), $ex->getCode()));
}
