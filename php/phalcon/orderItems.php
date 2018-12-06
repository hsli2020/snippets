<?php

use Phalcon\Di;

use Phalcon\Mvc\Model;
use Phalcon\Mvc\Model\Manager as ModelsManager;
use Phalcon\Mvc\Model\Metadata\Memory as MetaData;

const EOL = PHP_EOL;

function pr($var) { var_export($var); echo EOL; }

$di = new Di();

$di->set('db', function () {
    $config = [
        "host"     => "127.0.0.1",
        "username" => "root",
        "password" => "",
        "dbname"   => "test",
        "options"  => [ PDO::ATTR_CASE => PDO::CASE_LOWER ]
    ];

    $db = new \Phalcon\Db\Adapter\Pdo\Mysql($config);

    return $db;
});

$di->set("modelsManager",  new ModelsManager());
$di->set("modelsMetadata", new MetaData());

class Orders extends Model
{
    public $order_id;
    public $channel;

    public function initialize()
    {
        $this->hasMany("order_id", "OrderItems", "order_id", ['alias' => 'items']);
        $this->hasOne("order_id", "OrderShippingAddress", "order_id", 
                ['alias' => 'shippingAddress']);
    }
}

class OrderItems extends Model
{
    public $order_id;
    public $sku;
    public $qty;
    public $price;

    public function initialize()
    {
        $this->belongsTo('order_id', 'Orders', 'order_id', ['alias' => 'order']);
    }
}

class OrderShippingAddress extends Model
{
    public $order_id;
    public $name;
    public $address;

    public function initialize()
    {
        $this->belongsTo('order_id', 'Orders', 'order_id', ['alias' => 'order']);
    }
}

########################################

$orders = Orders::find();

foreach ($orders as $order) {

    $shippingAddress = $order->shippingAddress;

    foreach ($order->items as $item) {
        echo $order->order_id, ': ';

        echo $item->sku,   ' ';
        echo $item->qty,   ' ';
        echo $item->price, ' ';

        echo $shippingAddress->name,    ' ';
        echo $shippingAddress->address, ' ';

        echo EOL;
    }
}
echo EOL;

$items = OrderItems::find();
foreach ($items as $item) {
    echo $item->sku,   ' ';
    echo $item->qty,   ' ';
    echo $item->price, ': ';
    echo $item->order->order_id, ' ';
    echo EOL;
}
echo EOL;

$addresses = OrderShippingAddress::find();
foreach ($addresses as $address) {
    $order = $address->order;
    $items = $order->items;

    foreach ($items as $item) {
        echo $address->address, ' ';
        echo $address->name,    ': ';

        echo $order->order_id,  ' ';

        echo $item->sku,   ' ';
        echo $item->qty,   ' ';
        echo $item->price, ' ';

        echo EOL;
    }
}
