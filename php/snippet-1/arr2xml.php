<?php
/*
    public static function array2xml($array, $xml = false)
    {
        if ($xml === false){
            $xml = new SimpleXMLElement('<root/>');
        }

        foreach($array as $key => $value){
            if (is_array($value)){
                self::array2xml($value, $xml->addChild($key));
            }else{
                $xml->addChild($key, $value);
            }
        }

        return $xml->asXML();
    }

    public function toXml()
    {
        return self::array2xml($this->users_array, new SimpleXMLElement('<Response/>'));
    }

    ##

    public static function array_to_xml($data, $root = null)
    {
        $xml = new SimpleXMLElement($root ? '<' . $root . '/>' : '<root/>');
        array_walk_recursive($data, function($value, $key) use ($xml) {
            $xml->addChild($key, $value);
        });
        return $xml->asXML();
    }

    public function toXml()
    {
        return self::array_to_xml($this->users_array, 'Response');
    }

function array_to_xml( $data, &$xml_data ) {
    foreach( $data as $key => $value ) {
        if( is_array($value) ) {
            if( is_numeric($key) ){
                $key = 'item'.$key; //dealing with <0/>..<n/> issues
            }
            $subnode = $xml_data->addChild($key);
            array_to_xml($value, $subnode);
        } else {
            $xml_data->addChild("$key",htmlspecialchars("$value"));
        }
     }
}

// initializing or creating array
$data = include 'users_array.php';

// creating object of SimpleXMLElement
$xml = new SimpleXMLElement('<data/>');

// function call to convert array to xml
array_to_xml($data,$xml);

//saving generated xml file; 
$result = $xml->asXML();
echo $result;
*/

/**
* Converts an array to XML
*
* @param array $array
* @param SimpleXMLElement $xml
* @param string $child_name
*
* @return SimpleXMLElement $xml
*/
/*
function arrayToXML($array, SimpleXMLElement $xml, $child_name)
{
    foreach ($array as $k => $v) {
        if(is_array($v)) {
            (is_int($k)) ? arrayToXML($v, $xml->addChild($child_name), $v) : arrayToXML($v, $xml->addChild(strtolower($k)), $child_name);
        } else {
            (is_int($k)) ? $xml->addChild($child_name, $v) : $xml->addChild(strtolower($k), $v);
        }
    }

    return $xml->asXML();
}

$array = include "users_array.php";

echo arrayToXML($array, new SimpleXMLElement('<root/>'), 'user');
*/

/* BEST

$array = include "users_array.php";

echo arrayToXml("response",$array);

function arrayToXml($thisNodeName, $input)
{
    if (is_numeric($thisNodeName)) {
        throw new Exception("cannot parse into xml. remainder :".print_r($input,true));
    }

    if (!(is_array($input) || is_object($input))) {
        return "<$thisNodeName>$input</$thisNodeName>\n";
    }
    else {
        $newNode = "<$thisNodeName>\n";
        foreach ($input as $key => $value){
            if (is_numeric($key)) {
               #$key = substr($thisNodeName, 0, strlen($thisNodeName)-1);
                $key = substr($thisNodeName, 0, -1);
            }
            $newNode .= arrayToXml($key, $value);
        }
        $newNode .= "</$thisNodeName>\n";
        return $newNode;
    }
}
*/

$users_array = include 'users_array.php';

class Arr
{
    public static function toXml($array, $xml = false, $root = null)
    {
        if ($xml === false){
            $xml = new SimpleXMLElement($root ? "<$root/>" : '<root/>');
        }

        foreach($array as $key => $value){
            if (is_array($value)) {
                self::toXml($value, $xml->addChild($key));
            } else {
                $xml->addChild($key, $value);
            }
        }

        return $xml->asXML();
    }
}

echo Arr::toXml($users_array, false, 'Request');
