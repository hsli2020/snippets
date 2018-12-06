<?php

//include('.\adodb5\adodb.inc.php'); 

$dbname = "E:\\DataCodes.accdb";
$dbname = "z:\\Purchasing\\General Purchase.accdb";
$dbname = "C:\\Users\BTE\Desktop\\General Purchase.accdb";

if (!file_exists($dbname)) {
    die("Could not find database file.");
}

$dsn = "odbc:Driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=$dbname;";

#try {
#    $db = new PDO($dsn);
#    $sql = "SELECT * FROM tblEquipClass";
#
#    $stmt = $db->query($sql);
#    while ($row = $stmt->fetch()) {
#        echo $row['EquipmentClass'] . "\n";
#    }
#}
#catch(PDOException $e){
#    echo $e->getMessage();
#}

$db = new PDO($dsn);
#$db->exec("UPDATE tblEquipClass SET EquipmentClass='Paste and Copy' WHERE Class=1");

#$db->exec("UPDATE Newegg SET [Stock Status]='Paste/Copy' WHERE ID=1");

$sql = "INSERT INTO Newegg ([ID], [Work Date], [Channel], [PO #], [Xpress?], [Stock Status], [Qty], [Supplier], [Supplier SKU], [Mfr #], [Supplier #], [Remarks], [RelatedSKU], [Dimension]) VALUES (NULL, '2016-11-23', 'Newegg', 'PO#12345', 0, 'In Stock', 2, 'Synnex', 'SKU-433221', 'Mfr-455667', '223344', '', '', '')";

$sql = "INSERT INTO Newegg (
            [Work Date], 
            [Channel], 
            [PO #], 
            [Xpress], 
            [Stock Status], 
            [Qty], 
            [Supplier], 
            [Supplier SKU], 
            [Mfr #], 
            [Supplier #], 
            [Remarks], 
            [RelatedSKU], 
            [Dimension]
        )
        VALUES (
            '2016-11-23', 
            'Newegg', 
            'PO#12345', 
            0, 
            'In Stock', 
            2, 
            'Synnex', 
            'SKU-433221', 
            'Mfr-455667', 
            '223344', 
            '', 
            '', 
            ''
        )";

#$sql = "INSERT INTO Newegg (
#        Channel, PO #, Xpress?, Stock Status, 
#        Qty, Supplier, Supplier SKU, Mfr #, Supplier #)
#    VALUES('Newegg', 'PO#12345', 0, 'In Stock', 2, 'Synnex', '433221', '455667')";

$ret = $db->exec($sql);
if ($ret) {
    echo "------\n";
} else {
    print_r($db->errorInfo());
}

$sql = "SELECT * FROM Newegg WHERE [PO #]='207017859'";
$sql = "SELECT * FROM Newegg WHERE ID>3390";

$stmt = $db->query($sql);

while ($row = $stmt->fetch()) {
    echo $row['ID'], ' ';
    echo $row['Work Date'], ' ';
    echo $row['Channel'], ' ';
    echo $row['PO #'], ' ';
    echo $row['Xpress'], ' ';
    echo $row['Stock Status'], ' ';
    echo $row['Qty'], ' ';
    echo $row['Supplier'], ' ';
    echo $row['Supplier SKU'], ' ';
    echo $row['Mfr #'], ' ';
    echo $row['Supplier #'], ' ';
    echo PHP_EOL;
}

#-------------------------------------------------------------------------------

$rakuten_order_info = 'a:6:{s:14:"component_mode";s:3:"buy";s:14:"payment_method";s:0:"";s:16:"channel_order_id";i:78565861;s:17:"channel_final_fee";i:0;s:12:"transactions";a:0:{}s:6:"tax_id";N;}';
$amazon_order_info = 'a:6:{s:14:"component_mode";s:6:"amazon";s:14:"payment_method";s:0:"";s:16:"channel_order_id";s:19:"102-4433114-6945028";s:17:"channel_final_fee";i:0;s:12:"transactions";a:0:{}s:6:"tax_id";N;}';
$ebay_order_info = 'a:6:{s:14:"component_mode";s:4:"ebay";s:14:"payment_method";s:6:"PayPal";s:16:"channel_order_id";s:26:"231616183399-1214063743013";s:17:"channel_final_fee";d:12.960000000000001;s:12:"transactions";a:1:{i:0;a:4:{s:14:"transaction_id";s:17:"8H496057D7441160T";s:3:"sum";d:350.68000000000001;s:3:"fee";d:9.0700000000000003;s:16:"transaction_date";s:19:"2015-08-11 18:16:51";}}s:6:"tax_id";s:0:"";}';

#var_export(unserialize($rakuten_order_info));
#var_export(unserialize($amazon_order_info));
#var_export(unserialize($ebay_order_info));

/*
$rakuten_order_info = array (
  'component_mode' => 'buy',
  'payment_method' => '',
  'channel_order_id' => 78565861,
  'channel_final_fee' => 0,
  'transactions' => array (),
  'tax_id' => NULL,
)

$amazon_order_info = array (
  'component_mode' => 'amazon',
  'payment_method' => '',
  'channel_order_id' => '102-4433114-6945028',
  'channel_final_fee' => 0,
  'transactions' => array (),
  'tax_id' => NULL,
)

$ebay_order_info = array (
  'component_mode' => 'ebay',
  'payment_method' => 'PayPal',
  'channel_order_id' => '231616183399-1214063743013',
  'channel_final_fee' => 12.960000000000001,
  'transactions' => array (
    0 => array (
      'transaction_id' => '8H496057D7441160T',
      'sum' => 350.68000000000001,
      'fee' => 9.0700000000000003,
      'transaction_date' => '2015-08-11 18:16:51',
    ),
  ),
  'tax_id' => '',
)
//*/

function get_order_id($string) // current code, not good
{
    $pos = strpos ($string,'channel_order_id') + 24;//start pos of order id
    $new_string = substr ($string,$pos);//string start from 1st number of order id to the end of string
    //echo $new_string.'<br>';
    $pos1 = strpos ($new_string,'"'); //pos of "
    //echo $pos1.'<br>';
    $order_id = substr ($new_string,0 ,$pos1 );

    if (stristr ($string, 'amazon')) $channel = 'Amazon';
    else if (stristr ($string, 'ebay') ) $channel = 'eBay';
    else if (stristr ($string, 'buy') ) {
        $channel = 'Rakuten';
        $pos2 = strpos ($string,'channel_order_id') + 20;//start pos of order id
        $new_string = substr ($string,$pos2);
        $order_id = substr ($new_string,0,8);
    }
    else $channel = 'website';

    return array ($channel, $order_id);
}

#var_export(get_order_id($rakuten_order_info));
#var_export(get_order_id($amazon_order_info));
#var_export(get_order_id($ebay_order_info));

#array('Rakuten', '78565861')
#array('Amazon', '102-4433114-6945028')
#array('eBay', '231616183399-1214063743013')

function get_order_id_($string) // new code, right way
{
    $order = unserialize($string);

    $order_id = $order['channel_order_id'];
    $channel = $order['component_mode'];

    $map = [
        'amazon' => 'Amazon',
        'buy'    => 'Rakuten',
        'ebay'   => 'eBay',
    ];

    $channel = isset($map[$channel]) ? $map[$channel] : 'website';

    return array ($channel, $order_id);
}
var_export(get_order_id_($rakuten_order_info));
var_export(get_order_id_($amazon_order_info));
var_export(get_order_id_($ebay_order_info));

#-------------------------------------------------------------------------------

function genInsertSql($table, $data)
{
    $columns = array_keys($data);
    $columnList = '`' . implode('`, `', $columns) . '`';

    $query = "INSERT INTO `$table` ($columnList) VALUES\n";

    $values = array();

    foreach($data as &$val) {
        $val = addslashes($val);
    }
    $values[] = "\t('" . implode("', '", $data). "')";

    $update = implode(",\n",
        array_map(function($name) {
            return "\t`$name`=VALUES(`$name`)";
        }, $columns)
    );

    return $query . implode(",\n", $values) . "\nON DUPLICATE KEY UPDATE\n$update";
}

$data = [ 'sku' => 'SKU-111', 'price' => 11.11, 'qty' => 11 ];
echo genInsertSql('test', $data), PHP_EOL, PHP_EOL;

function genBatchInsertSql($table, $columns, $data)
{
    $columnList = '`' . implode('`, `', $columns) . '`';

    $query = "INSERT INTO `$table` ($columnList) VALUES\n";

    $values = array();

    foreach($data as $row) {
        foreach($row as &$val) {
            $val = addslashes($val);
        }
        $values[] = "\t('" . implode("', '", $row). "')";
    }

    $update = implode(",\n",
        array_map(function($name) {
            return "\t`$name`=VALUES(`$name`)";
        }, $columns)
    );

    return $query . implode(",\n", $values) . "\nON DUPLICATE KEY UPDATE\n$update";
}

$columns = [ 'sku', 'price', 'qty' ];
$data = [
    [ 'SKU-111', 11.11, 11 ],
    [ 'SKU-222', 22.22, 22 ],
    [ 'SKU-333', 33.33, 33 ],
];

echo genBatchInsertSql('test', $columns, $data), PHP_EOL;

#-------------------------------------------------------------------------------

function getRedis()
{
    static $redis = null;

    if (!$redis) {
        $redis = new \Redis();
        $redis->connect('127.0.0.1');
    }

    return $redis;
}

function loadSkuList()
{
    $redis = getRedis();

    $timestamp = $redis->get('master_sku_list:timestamp');
    if (time() - $timestamp < 12*3600) {
        return;
    }

    $skulist = fopen('w:/data/master_sku_list.csv', 'r');
    $names = fgetcsv($skulist); // skip first line

    while (($result = fgetcsv($skulist)) !== false) {
        for ($i = 0; $i < 8 ; $i++) {
            $pn = $result[ $i * 3 + 2 ];
            if ($pn != '') {
                $redis->set($pn, json_encode($result));
            }
        }
    }

    $redis->set('master_sku_list:timestamp', time());

    fclose($skulist);
}

function getMasterSku($sku)
{
    $names = array(
        'SKU',
        'recommended_pn',
        'syn_pn', 'syn_cost', 'syn_qty',
        'td_pn', 'td_cost', 'td_qty',
        'ing_pn', 'ing_cost', 'ing_qty',
        'dh_pn', 'dh_cost', 'dh_qty',
        'asi_pn', 'asi_cost', 'asi_qty',
        'tak_pn', 'tak_cost', 'tak_qty',
        'ep_pn', 'ep_cost', 'ep_qty',
        'BTE_PN', 'BTE_cost', 'BTE_qty',
        'Manufacturer',
        'UPC',
        'MPN',
        'MAP_USD',
        'MAP_CAD',
        'Width',
        'Length',
        'Depth',
        'Weight',
        'ca_ebay_blocked',
        'us_ebay_blocked',
        'ca_newegg_blocked',
        'us_newegg_blocked',
        'us_amazon_blocked',
        'ca_amazon_blocked',
        'uk_amazon_blocked',
        'jp_amazon_blocked',
        'mx_amazon_blocked',
        'note',
        'name',
        'best_cost',
        'overall_qty'
    );

    $redis = getRedis();

    $value = json_decode($redis->get($sku));

    if (count($value) == count($names)) {
        return array_combine($names, $value);
    }

    return $value;
}

loadSkuList();
$info = getMasterSku('ING-4599CC');
print_r($info);

#-------------------------------------------------------------------------------
//include('.\adodb5\adodb.inc.php'); 

$dbname = "C:/Users/BTE/Desktop/General Purchase.accdb";
$neweggOrderReport = "E:/BTE/orders/newegg/neweggcanada_master_orders.csv";

$dsn = "odbc:Driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=$dbname;";

$db = new PDO($dsn);
//$db->exec("DELETE FROM Newegg WHERE ID>3397");

$today = date('Y-m-d');
$channel = 'NeweggCA';
$stockStatus = ' ';
$lastOrderNo = ' ';

$orderFile = fopen($neweggOrderReport, 'r');
$title = fgetcsv($orderFile);

while (($fields = fgetcsv($orderFile)) != false) {

    $orderNo    = $fields[0];
    $date       = date('Y-m-d', strtotime($fields[1]));
    $sku        = $fields[16];
    $qty        = $fields[26];
    $shipMethod = $fields[15];

    // check if the order already exported
    $sql = "SELECT * FROM Newegg WHERE [PO #]='$orderNo'";
    $result = $db->query($sql)->fetch();
    if ($result && $orderNo != $lastOrderNo) {
        echo "Skip $orderNo\n";
        continue;
    }

    if ($date != $today) {
        //continue;
    }

    $supplier = getSupplier($sku);
    $sku = getSku($sku);
    $xpress = isExpress($shipMethod);
    $mfrpn = getMfrPartNum($sku);

    #$data = array_combine($title, $fields);
    #print_r($data); break;

    echo "Importing $orderNo\n";

    $sql = "INSERT INTO Newegg (
                [Work Date], 
                [Channel], 
                [PO #], 
                [Xpress], 
                [Stock Status], 
                [Qty], 
                [Supplier], 
                [Supplier SKU], 
                [Mfr #], 
                [Supplier #], 
                [Remarks], 
                [RelatedSKU], 
                [Dimension]
            )
            VALUES (
                '$date',
                '$channel',
                '$orderNo',
                $xpress,
                '$stockStatus',
                '$qty',
                '$supplier',
                '$sku',
                '$mfrpn',
                '',
                '',
                '',
                '' 
            )";

    $ret = $db->exec($sql);

    if (!$ret) {
        print_r($db->errorInfo());
    }

    $lastOrderNo = $orderNo;
}

fclose($orderFile);

function getSku($sku)
{
    $parts = explode('-', $sku);
    array_shift($parts);
    return implode('-', $parts);
}

function getSupplier($sku)
{
    $names = [
        'AS'  => 'ASI',
        'SYN' => 'Synnex',
        'ING' => 'Ingram Micro',
        'EP'  => 'Eprom',
        'TD'  => 'Techdata',
        'TAK' => 'Taknology',
        'SP'  => 'Supercom',
    ];

    $parts = explode('-', $sku);
    $prefix = $parts[0];

    return isset($names[$prefix]) ? $names[$prefix] : $prefix;
}

function isExpress($shipMethod)
{
    # Expedited Shipping (3-5 business days)
    # One-Day Shipping(Next day)
    # Standard Shipping (5-7 business days)
    # Two-Day Shipping(2 business days)

    if (strpos($shipMethod, 'Standard Shipping') !== false) {
        return 0;
    }

    return 1;
}

function getMfrPartNum($sku)
{
    return '';
}

#-------------------------------------------------------------------------------
function getRedis()
{
    static $redis = null;

    if (!$redis) {
        $redis = new \Redis();
        $redis->connect('127.0.0.1');
    }

    //$redis->set('foo', json_encode($data));
    //$value = $redis->get('foo');
    //print_r(json_decode($value));

    return $redis;
}

function loadSkuList() // faster
{
    $redis = getRedis();

    $skulist = fopen('w:/data/selling_list.csv', 'r');
    fgetcsv($skulist); // skip first line

    while (($result = fgetcsv($skulist)) !== false) {

        $combined_string = '';
        $master_sku = strtoupper($result[0]);

        for ($i = 4; $i < 12 ; $i++) {

            if ($result[$i] != '') { //pn
                $PN = strtoupper($result[$i]);

                $length = $result[16]; // length
                $width  = $result[17]; // width
                $depth  = $result[18]; // depth
                $weight = $result[19]; // weight

                $price_array[0] = $result[2];  // cost for given pn
                $price_array[1] = $result[3];  // qty
                $price_array[2] = $result[12]; // note
                $price_array[3] = $result[15]; // mpn
                $price_array[4] = $result[1];  // recommended sku
                $price_array[5] = "$length x $width x $depth (inch); $weight (lbs)"; // package info
                $price_array[6] = $result[0];  // master sku

                //echo $PN, PHP_EOL;
                $redis->set($PN, json_encode($price_array));

                $combined_string = $combined_string . $result [$i]. ' | ';
            }
        }

        //echo $master_sku, ' ', $combined_string, PHP_EOL;
        $redis->set("master:$master_sku", $combined_string);
    }

    fclose($skulist);
}

function getSku()
{
    $skus = [
        'ING-92645Z',
        'SYN-4946765',
        'AS-89719',
        'SYN-5537007',
        'ING-64523S',
    ];

    $redis = getRedis();

    foreach ($skus as $sku) {
        $value = $redis->get($sku);
        print_r(json_decode($value));
    }
}

function getMasterSku()
{
    $skus = [
        'ODO-00AJ400-9190',
        'ODO-00AM066-8069',
        'ODO-00YJ199-1998',
        'ODO-00FK930-6232',
        'ODO-00FK936-6263',
        'ODO-00KA498-6294',
    ];

    $redis = getRedis();

    foreach ($skus as $sku) {
        $value = $redis->get("master:$sku");
        print_r($value);
        echo PHP_EOL;
    }
}

//loadSkuList();
//getSku();
getMasterSku();

#-------------------------------------------------------------------------------

use PHPUnit\Framework\TestCase;

class XmlTest extends TestCase
{
    public function testXml()
    {
        $xmlstr=<<<EOS
<XML_FreightEstimate_Submit>
  <Header>
    <UserName>YourUserID</UserName>
    <Password>YourPassword</Password>
  </Header>
  <Detail>
    <PostalCode>33611</PostalCode>
    <LineInfo>
      <AssignedID>001</AssignedID>
      <WhseCode2>A1</WhseCode2>
      <ShipViaCode2>23</ShipViaCode2>
      <RefIDQual>VP</RefIDQual>
      <RefID>ELISA</RefID>
      <Quantity>1</Quantity>
    </LineInfo>
  </Detail>
</XML_FreightEstimate_Submit>
EOS;

        $xml = simplexml_load_string($xmlstr, 'SimpleXMLElement', LIBXML_NOCDATA);

        $this->assertEquals($xml->getName(), 'XML_FreightEstimate_Submit');
        $this->assertEquals($xml->Header->UserName, 'YourUserID');
        $this->assertEquals($xml->Header->Password, 'YourPassword');
        $this->assertEquals($xml->Detail->PostalCode, '33611');
        $this->assertEquals($xml->Detail->LineInfo->AssignedID, '001');
        $this->assertEquals($xml->Detail->LineInfo->WhseCode2, 'A1');
        $this->assertEquals($xml->Detail->LineInfo->ShipViaCode2, '23');
        $this->assertEquals($xml->Detail->LineInfo->RefIDQual, 'VP');
        $this->assertEquals($xml->Detail->LineInfo->RefID, 'ELISA');
        $this->assertEquals($xml->Detail->LineInfo->Quantity, '1');
    }
}

#-------------------------------------------------------------------------------

$xml_single = '<?xml version="1.0" encoding="UTF-8"?>
<resultset>
    <row>
        <name>Happy</name>
        <age>20</age>
    </row>
</resultset>';
 
$xml_multi = '<?xml version="1.0" encoding="UTF-8"?>
<resultset>
    <row>
        <name><first>Happy</first><last>Year</last></name>
        <age>20</age>
    </row>
    <row>
        <name><first>Harry</first><last>Day</last></name>
        <age>25</age>
    </row>
</resultset>';
 
$simple_single = simplexml_load_string($xml_single);
$simple_multi = simplexml_load_string($xml_multi);

#print_r($simple_single);
print_r($simple_multi);
 
$single_array = xml2array($simple_single);
$multi_array = xml2array($simple_multi);
 
#rint_r($single_array);
print_r($multi_array);
 
function isNumericArray($var) {
    return is_array($var) && count(array_filter(array_keys($var), 'is_string')) == 0;
}

function xml2array($xml)
{
    $arr = array();
 
    foreach ($xml->children() as $r)
    {
        print_r($r);
        if (isNumericArray($r))
        {
            foreach ($r->children() as $c) {
               #$arr[$r->getName()][] = xml2array($c);
                $arr[$r->getName()][] = $c;
            }
        }
        else
        {
            $arr[$r->getName()] = $r;
           #$arr[$r->getName()] = strval($r);
        }
    }
    return $arr;
}

function xml2arr($xmlObject, $out = array())
{
    foreach ((array)$xmlObject as $index => $node)
    {
        $out[$index] = (is_object($node)) ? xml2array($node) : $node;
    }
    return $out;
}

#-------------------------------------------------------------------------------
# https://gist.github.com/Xeoncross/9401853

// Ignore errors
libxml_use_internal_errors(true) AND libxml_clear_errors();

// http://stackoverflow.com/q/10237238/99923
// http://stackoverflow.com/q/12034235/99923
// http://stackoverflow.com/q/8218230/99923

// original input (unknown encoding)
$html = 'hi</b><p>سلام<div>の家庭に、9 ☆ Mike Dubé';
#print $html . PHP_EOL;

#$doc = new DOMDocument();
#$doc->preserveWhiteSpace = false;
#$doc->loadHTML($html);
#print $doc->saveHTML($doc->documentElement) . PHP_EOL . PHP_EOL;
#
#$doc = new DOMDocument('1.0', 'UTF-8');
#$doc->loadHTML($html);
#$doc->encoding = 'utf-8';
#print $doc->saveHTML($doc->documentElement) . PHP_EOL . PHP_EOL;

$doc = new DOMDocument();
$doc->loadHTML('<?xml encoding="utf-8"?>' . $html);
$doc->encoding = 'utf-8';
print $doc->saveHTML($doc->documentElement) . PHP_EOL . PHP_EOL;

$doc = new DOMDocument('1.0', 'UTF-8');
$doc->loadHTML(mb_convert_encoding($html, 'HTML-ENTITIES', 'UTF-8'));
print $doc->saveHTML($doc->documentElement) . PHP_EOL . PHP_EOL;
/*
// Benchmark
print "Testing XML encoding spec" . PHP_EOL;
$time = microtime(TRUE);
for ($i=0; $i < 10000; $i++) { 
	$doc = new DOMDocument();
	$doc->loadHTML('<?xml encoding="utf-8"?>' . $html);
	foreach ($doc->childNodes as $item)
    	if ($item->nodeType == XML_PI_NODE)
        	$doc->removeChild($item); // remove hack
    
	$doc->encoding = 'utf-8';
	$doc->saveHTML();
	unset($doc);
}
print (microtime(TRUE) - $time) . " seconds" . PHP_EOL . PHP_EOL;

print "Testing mb_convert_encoding" . PHP_EOL;
$time = microtime(TRUE);
for ($i=0; $i < 10000; $i++) { 
	$doc = new DOMDocument();
	$doc->loadHTML(mb_convert_encoding($html, 'HTML-ENTITIES', 'UTF-8'));
	$doc->saveHTML();
	unset($doc);
}
print (microtime(TRUE) - $time) . " seconds" . PHP_EOL . PHP_EOL;
*/
#-------------------------------------------------------------------------------

function getCurrencyExchangeRate($from, $to)
{
    $url = "http://finance.yahoo.com/d/quotes.csv?f=l1d1t1&s={$from}{$to}=X";
    $handle = fopen($url, 'r');

    $exchangeRate = 1.25;

    if ($handle) {
        $result = fgetcsv($handle);

        if (isset($result[0])) {
            $exchangeRate = $result[0]-0.02;
        }

        fclose($handle);
    }

    return $exchangeRate;
}

$from = 'USD';
$to = 'CAD';

echo getCurrencyExchangeRate($from, $to), "\n";
echo getCurrencyExchangeRate('USD', 'CNY'), "\n";
echo getCurrencyExchangeRate('CAD', 'CNY'), "\n";

#-------------------------------------------------------------------------------
const EOL = PHP_EOL;

function pr($d) { var_export($d); echo EOL; }

class Arr
{
    protected $data = [];

    public function set($name, $value)
    {
        $keys = explode('.', $name);

        $data = &$this->data;
        foreach ($keys as $key) {
            if (isset($data[$key])) {
                $data = &$data[$key];
            } else {
                $data[$key] = [];
                $data = &$data[$key];
            }
        }
        $data = $value;
    }

    public function dump()
    {
        print_r($this->data);
    }

    public function toXml()
    {
        $xml = '<?xml version="1.0" encoding="UTF-8"?>'."\n";
        return $xml.self::array2xml('Request', $this->data);
    }

    public static function array2Xml($nodeName, $input, $indentLevel=0)
    {
        if (is_numeric($nodeName)) {
            throw new Exception("cannot parse into xml. remainder :". print_r($input, true));
        }

        $indent = str_repeat('  ', $indentLevel);

        if (!(is_array($input) || is_object($input))) {
            $input = htmlspecialchars($input);
            return "$indent<$nodeName>$input</$nodeName>\n";
        }
        else if (self::isNumericIndex($input)) {
            $newNode = '';
            foreach ($input as $key => $value) {
                $newNode .= self::array2Xml($nodeName, $value, $indentLevel + 1);
            }
            return $newNode;
        }
        else {
            $newNode="$indent<$nodeName>\n";
            foreach ($input as $key => $value) {
                if (is_numeric($key)) {
                    $key = substr($nodeName, 0, -1);
                }
                $newNode .= self::array2Xml($key, $value, $indentLevel + 1);
            }
            $newNode .= "$indent</$nodeName>\n";
            return $newNode;
        }
    }

    public static function isNumericIndex(array $var)
    {
        return array_keys($var) === range(0, sizeof($var) - 1);
    }

    public function flatten() 
    {
        $iter = new \RecursiveIteratorIterator(new \RecursiveArrayIterator($this->data));

        $result = array();
        foreach ($iter as $leafValue) {
            $keys = array();
            foreach (range(0, $iter->getDepth()) as $depth) {
                $keys[] = $iter->getSubIterator($depth)->key();
            }
            $result[join('.', $keys)] = $leafValue;
        }

        return $result;
    }
}

$a = new Arr;
$a->set('Head.Credential.Username', 'USERNAME');
$a->set('Head.Credential.Password', 'PASSWORD');
$a->set('Body.SkuList.Sku', 'SKU-111');
//$a->set('Boday.SkuList.Sku', 'SKU-222');
//$a->dump();
//pr($a->flatten());
echo $a->toXml();

#-------------------------------------------------------------------------------

class XmlNode
{
    public $name;
    public $value;
    public $attrs = [];

    public function __construct($name, $value='')
    {
        $this->name = $name;
        $this->value = $value;
    }

    public function setName($name)
    {
        $this->name = $name;
    }

    public function setValue($value)
    {
        $this->value = $value;
    }

    public function setAttr($name, $value)
    {
        $this->attrs[$name] = $value;
    }

    public function __toString()
    {
        $name = $this->name;

        $attrs = [];
        foreach ($this->attrs as $key => $val) {
            $attrs[] = "$key=\"$val\""; 
        }

        $attrs = implode(' ', $attrs);

        $openTag = "<$name>";
        $closeTag = "</$name>";

        if ($attrs) {
            $openTag = "<$name $attrs>";
        }

        if ($this->value instanceof XmlNode) {
            $openTag .= "\n";
            $value = (string)$this->value;
        } else {
            $value = htmlspecialchars($this->value);
        }

        if (strlen($value) == 0) {
            $openTag = trim($openTag, '>');
            $closeTag = "/>";
        }

        return "$openTag$value$closeTag\n";
    }
}

$x = new XmlNode('Company');
$x->setAttr('Lang', 'en');
$x->setAttr('Country', 'US');

$w = new XmlNode('Website');
$w->setAttr('URL', 'http://www.dandh.ca');

$x->setValue($w);

echo $x;
#-------------------------------------------------------------------------------

# RewriteRule ^(.+)\.(\d+)\.(js|css|png|jpg|gif)$ $1.$3 [L]

function cachedAssetUrl($url)
{
    $path = 'c:/xampp/htdocs/btenew/public';

    $filePath = $path . '/' .  ltrim($url, '/');

    if (!file_exists($filePath)) {
        throw new LogicException(
            'Unable to locate the asset "' . $url . '" in the "' . $path . '" directory.'
        );
    }

    $lastUpdated = filemtime($filePath);
    $pathInfo = pathinfo($url);

    $dirname = $pathInfo['dirname'];

    if ($dirname === '.') {
        $dirname = '';
    } elseif ($dirname === '/') {
        $dirname = '/';
    } else {
        $dirname = $dirname . '/';
    }

    $filename = $pathInfo['filename'];
    $extension = $pathInfo['extension'];

    return "$dirname$filename.$lastUpdated.$extension";
}

echo cachedAssetUrl('/assets/css/style.css');
#-------------------------------------------------------------------------------

#-------------------------------------------------------------------------------

