<?php
//Functions from Athena:
/**
 * Get url or send POST data
 *
 * @param string $url
 * @param array  $param['Header']
 *               $param['Post']
 * @return array $return['ok'] 1  - success, (0,-1) - fail
 *               $return['body']  - response
 *               $return['error'] - error, if "ok" is not 1
 *               $return['head']  - http header
 */
function fetchURL($url, $param) // curlGet/curlPost
{
    $return = array();

    $ch = curl_init();

    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_TIMEOUT, 0);
    curl_setopt($ch, CURLOPT_FORBID_REUSE, 1);
    curl_setopt($ch, CURLOPT_FRESH_CONNECT, 1);
    curl_setopt($ch, CURLOPT_HEADER, 1);
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);

    if (!empty($param)) {
        if (!empty($param['Header'])) {
            curl_setopt($ch, CURLOPT_HTTPHEADER, $param['Header']);
        }
        if (!empty($param['Post'])) {
            curl_setopt($ch, CURLOPT_POSTFIELDS, $param['Post']);
        }
    }

    $data = curl_exec($ch);

    if (curl_errno($ch)) {
        $return['ok'] = -1;
        $return['error'] = curl_error($ch);
        return $return;
    }

    if (is_numeric(strpos($data, 'HTTP/1.1 100 Continue'))) {
        $data = str_replace('HTTP/1.1 100 Continue', '', $data);
    }

    $data = preg_split("/\r\n\r\n/",$data, 2, PREG_SPLIT_NO_EMPTY);

    if (!empty($data)) {
        $return['head'] = (isset($data[0]) ? $data[0] : null);
        $return['body'] = (isset($data[1]) ? $data[1] : null);
    } else {
        $return['head'] = null;
        $return['body'] = null;
    }

    $matches = array();
    $data = preg_match("/HTTP\/[0-9.]+ ([0-9]+) (.+)\r\n/",$return['head'], $matches);

    if (!empty($matches)) {
        $return['code'] = $matches[1];
        $return['answer'] = $matches[2];
    }

    $data = preg_match("/meta http-equiv=.refresh. +content=.[0-9]*;url=([^'\"]*)/i",
            $return['body'], $matches);

    if (!empty($matches)) {
        $return['location'] = $matches[1];
        $return['code'] = '301';
    }

    if ($return['code'] == '200' || $return['code'] == '302') {
        $return['ok'] = 1;
    } else {
        $return['error'] = (($return['answer'] and $return['answer'] != 'OK')
                         ? $return['answer'] : 'Something wrong!');
        $return['ok'] = 0;
    }

    foreach (preg_split('/\n/', $return['head'], -1, PREG_SPLIT_NO_EMPTY) as $value) {
        $data = preg_split('/:/', $value, 2, PREG_SPLIT_NO_EMPTY);
        if (is_array($data) and isset($data['1'])) {
            $return['headarray'][$data['0']] = trim($data['1']);
        }
    }

    curl_close($ch);

    return $return;
}
// End Functions from Athena
