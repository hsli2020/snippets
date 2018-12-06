<?php

const EOL = PHP_EOL;

function pr($d) { var_export($d); echo EOL; }

$xml=<<<EOL
<?xml version="1.0" encoding="utf-8" ?>
<GetOrdersRequest xmlns="urn:ebay:apis:eBLBaseComponents">
<Order>Seller</Order>
<Status>All</Status>
<Detail>ReturnSummary</Detail>
</GetOrdersRequest>
EOL;

#$d = simplexml_load_string($xml);
#var_dump($d->NotFound); // OK
#var_dump($d->NotFound->EvenMore); // OK
#var_dump($d->NotFound->EvenMore->TooMuch); // warnning

//In YYMMDD format.  Example: July 15, 2002 is represented as 020715
#list($y, $m, $d) = str_split('020715', 2);
#echo "20$y-$m-$d";

#$dt = new DateTime();
#echo $dt->format(DateTime::ISO8601), EOL;
#echo $dt->format(DateTime::W3C), EOL;
#echo $dt->format(DATE_ATOM), EOL;

$fn = '20161028110806208_BTE_COMPUTER_856.xml';
$fn = '20161028_000442571_BTE_COMPUTER_810.xml';

$datetime = substr(str_replace('_', '', $fn), 0, 14);

#if ($fn[8] == '_') {
#    $date = substr($fn, 0, 8);
#    $time = substr($fn, 9, 6);
#} else {
#    $date = substr($fn, 0, 8);
#    $time = substr($fn, 8, 6);
#}
#echo date('Y-m-d H:i:s', strtotime($date.$time));

#echo strtotime($datetime), EOL;
#echo date('Y-m-d H:i:s', strtotime($datetime));

$array = [ 'foo' => 1, 'bar' => 2, 'baz' => 3, 'egg' => 5, 'cow' => 6 ];

#$deletion = [ 'a', 'b', 'c', 'd', 'e' ];
#
#for ($i=0; $i<count($deletion); $i++) {
#    $sql = "delete From [BTE_Records].[dbo].[Chitchat] where Tracking_Number='$deletion[$i]'";
#    echo $sql, EOL;
#}

//$orderId = '115-8641817-4801849';
//$d = file_get_contents("http://192.168.0.12/order/get?id=$orderId");
//$json = json_decode($d);
//print_r($json);
/*
$id = '111-5075550-5781819';
$id = '9400110200829252056732';
$d = file_get_contents("http://192.168.0.12/query/shippingeasy?id=$id");
$json = json_decode($d);
print_r($json);

if ($json->status == 'OK') {
    $SKU = $json->data->SKU;
    $Price = $json->data->OrderTotal;
    $weight = $json->data->WeightOZ/16;
    $description = $json->data->ItemName;
    if ($weight<1){
        $class='1st Class';
        $handling='$0.65';
    }
    else if($weight>=1){
        $class='Priority';
        $handling='$1.00';
    }
    //$Item = getItemClass($description);
    echo $SKU, EOL;
    echo $Price, EOL;
    echo $weight, EOL;
    echo $description, EOL;
}
*/

//echo date('Y-m-d 00:00:00', strtotime('-1 days')), EOL;
//echo date('Y-m-d H:i:s', strtotime('now')), EOL;
//echo date('Y-m-d H:i:s', strtotime('10 minutes')), EOL;
//echo date('Y-m-d H:i:s', strtotime('1 days')), EOL;

//echo date('Y-m-d H:i:s', filemtime('e:/tr.php')), EOL;
//echo date('Y-m-d H:i:s', strtotime('10 minutes', filemtime('e:/tr.php')));
//if (time() < strtotime('10 minutes', filemtime('e:/tr.php'))) echo "file is not old";

class Foo {
    protected $name = __CLASS__;
    public function show() { echo "Name of this class is '$this->name'\n"; }
}
$foo = new Foo();
$foo->show();


