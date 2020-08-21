<?php	# https://github.com/crazyfactory/php-curl

namespace CrazyFactory\Curl;

class Curl
{
    /**
     * @param string          $url
     * @param string[]|null   $postFields
     * @param string[]        $options
     * @param string[]|null   $curlInfo
     *
     * @return string|null
     */
    public function post($url, array $postFields = array(), array $options = array(), array &$curlInfo = array())
    {
        $options[CURLOPT_URL] = $url;
        $options[CURLOPT_POST] = 1;

        if (!empty($postFields)) {
            $options[CURLOPT_POSTFIELDS] = $postFields;
        }

        return $this->call($options, $curlInfo);
    }

    /**
     * @param string        $url
     * @param string[]|null $getFields
     * @param string[]      $options
     * @param string[]|null $curlInfo
     *
     * @return string|null
     */
    public function get($url, array $getFields = array(), array $options = array(), array &$curlInfo = array())
    {
        $options[CURLOPT_URL] = empty($getFields)
            ? $url
            : $url . (strpos($url, '?') === false ? '?' : '&') . http_build_query($getFields);

        return $this->call($options, $curlInfo);
    }

    /**
     * @param string[]      $options
     * @param string[]|null $curlInfo
     *
     * @return string|null
     * @throws Exception
     */
    public function call(array $options, array &$curlInfo = array())
    {
        $curlOptions = $this->makeOptions($options);

        $ch = curl_init();
        curl_setopt_array($ch, $curlOptions);

        $result   = curl_exec($ch);
        $error    = curl_error($ch);
        $curlInfo = curl_getinfo($ch);

        if (!empty($error)) {
            throw new Exception($error, $curlOptions, $curlInfo);
        }

        return $curlInfo['http_code'] == 204 || empty($result)
            ? null
            : $result;
    }

    /**
     * @return array
     */
    public static function getDefaultOptions()
    {
        return array(
            CURLOPT_ENCODING => '',
            CURLOPT_FAILONERROR => 1,
            CURLOPT_RETURNTRANSFER => 1,
            CURLOPT_SSL_VERIFYPEER => 0,
            CURLOPT_TIMEOUT => 10,
            CURLOPT_FOLLOWLOCATION => ini_get('open_basedir')
                ? 0
                : 1,
        );
    }

    /**
     * @param string[] $options
     *
     * @return array
     */
    public function makeOptions(array $options = array())
    {
        // merge in default options
        $options += static::getDefaultOptions();

        // Allow redirects
        // only if open_basedir is falsy, as system curl would deny it anyway to protect local files
        if (!ini_get('open_basedir')) {
            $options[CURLOPT_FOLLOWLOCATION] = 1;
        }

        // We support cookies as a key-value-array
        if (!empty($options[CURLOPT_COOKIE]) && is_array($options[CURLOPT_COOKIE])) {
            $options[CURLOPT_COOKIE] = http_build_query($options[CURLOPT_COOKIE], '', '; ');
        }

        // We support post fields as a key-value-array
        if (!empty($options[CURLOPT_POSTFIELDS])) {
            if (is_array($options[CURLOPT_POSTFIELDS])) {
                $options[CURLOPT_POSTFIELDS] = http_build_query($options[CURLOPT_POSTFIELDS], '', '&');
            }
        }

        return $options;
    }
}
