<?php

include __DIR__ . '/../public/init.php';

$db = Phalcon\Di::getDefault()->get('db');

$order = new AmazonOrder('bte-amazon-ca');
#$order->setOrderId('701-8728845-2735459'); // Canceled
$order->setOrderId('701-9942073-6661866'); // Shipped
$order->fetchOrder();

//print_r($order->getData()); die;

saveOrder($db, $order);

function yesNo($value)
{
    if ($value == 'false') return 'N';
    return 'Y';
}

function dtime($value)
{
    # "2016-09-22T06:59:59Z" => "2016-09-22 06:59:59"
    return str_replace(['T', 'Z'], [' ', ''], $value);
}

function saveOrder($db, $order)
{
    $data = $order->getData();

    if ($data['OrderStatus'] == 'Canceled') {
        return;
    }

    if (!isset($data['OrderTotal'])) {
        $data['OrderTotal']['Amount'] = '0.00';
        $data['OrderTotal']['CurrencyCode'] = 'CAD';
    }
    if (!isset($data['PaymentMethod'])) {
        $data['PaymentMethod'] = '';
    }
    if (!isset($data['BuyerName'])) {
        $data['BuyerName'] = '';
    }
    if (!isset($data['BuyerEmail'])) {
        $data['BuyerEmail'] = '';
    }
    if (!isset($data['ShippedByAmazonTFM'])) {
        $data['ShippedByAmazonTFM'] = 'false';
    }
    if (!isset($data['EarliestDeliveryDate'])) {
        $data['EarliestDeliveryDate'] = null;
    }
    if (!isset($data['LatestDeliveryDate'])) {
        $data['LatestDeliveryDate'] = null;
    }

    try {
        $db->insertAsDict('amazon_order', [
            'OrderId' => $data['AmazonOrderId'],
            'PurchaseDate' => dtime($data['PurchaseDate']),
            'LastUpdateDate' => dtime($data['LastUpdateDate']),
            'OrderStatus' => $data['OrderStatus'],
            'FulfillmentChannel' => $data['FulfillmentChannel'],
            'SalesChannel' => $data['SalesChannel'],
            'ShipServiceLevel' => $data['ShipServiceLevel'],
            'CurrencyCode' => $data['OrderTotal']['CurrencyCode'],
            'OrderTotalAmount' => $data['OrderTotal']['Amount'],
            'NumberOfItemsShipped' => $data['NumberOfItemsShipped'],
            'NumberOfItemsUnshipped' => $data['NumberOfItemsUnshipped'],
            'PaymentMethod' => $data['PaymentMethod'],
            'BuyerName' => $data['BuyerName'],
            'BuyerEmail' => $data['BuyerEmail'],
            'ShipmentServiceLevelCategory' => $data['ShipmentServiceLevelCategory'],
            'ShippedByAmazonTFM' => yesNo($data['ShippedByAmazonTFM']),
            'OrderType' => $data['OrderType'],
            'EarliestShipDate' => dtime($data['EarliestShipDate']),
            'LatestShipDate' => dtime($data['LatestShipDate']),
            'EarliestDeliveryDate' => dtime($data['EarliestDeliveryDate']),
            'LatestDeliveryDate' => dtime($data['LatestDeliveryDate']),
            'IsBusinessOrder' => yesNo($data['IsBusinessOrder']),
            'IsPrime' => yesNo($data['IsPrime']),
            'IsPremiumOrder' => yesNo($data['IsPremiumOrder']),
        ]);
    } catch (Exception $e) {
        echo $e->getMessage(), EOL;
    }

    saveOrderItem($db, $order);
    saveShippingAddress($db, $order);
}

function saveOrderItem($db, $order)
{
    $items = $order->fetchItems();
    $item = $items->getItems(0);

    try {
        $db->insertAsDict('amazon_order_item', [
            'OrderId' => $order->getAmazonOrderId(),
            'ASIN' => $item['ASIN'],
            'SellerSKU' => $item['SellerSKU'],
            'OrderItemId' => $item['OrderItemId'],
            'Title' => $item['Title'],
            'QuantityOrdered' => $item['QuantityOrdered'],
            'QuantityShipped' => $item['QuantityShipped'],
            'CurrencyCode' => $item['ItemPrice']['CurrencyCode'],
            'ItemPrice' => $item['ItemPrice']['Amount'],
            'ShippingPrice' => $item['ShippingPrice']['Amount'],
            'GiftWrapPrice' => $item['GiftWrapPrice']['Amount'],
            'ItemTax' => $item['ItemTax']['Amount'],
            'ShippingTax' => $item['ShippingTax']['Amount'],
            'GiftWrapTax' => $item['GiftWrapTax']['Amount'],
            'ShippingDiscount' => $item['ShippingDiscount']['Amount'],
            'PromotionDiscount' => $item['PromotionDiscount']['Amount'],
            'ConditionId' => $item['ConditionId'],
            'ConditionSubtypeId' => $item['ConditionSubtypeId'],
            'ConditionNote' => $item['ConditionNote'],
        ]);
    } catch (Exception $e) {
        echo $e->getMessage(), EOL;
    }
}

function saveShippingAddress($db, $order)
{
    $address = $order->getShippingAddress();

    try {
        $db->insertAsDict('amazon_order_shipping_address', [
            'OrderId' => $order->getAmazonOrderId(),
            'Name' => $address['Name'],
            'AddressLine1' => $address['AddressLine1'],
            'AddressLine2' => $address['AddressLine2'],
            'AddressLine3' => $address['AddressLine3'],
            'City' => $address['City'],
            'County' => $address['County'],
            'District' => $address['District'],
            'StateOrRegion' => $address['StateOrRegion'],
            'PostalCode' => $address['PostalCode'],
            'CountryCode' => $address['CountryCode'],
            'Phone' => $address['Phone'],
        ]);
    } catch (Exception $e) {
        echo $e->getMessage(), EOL;
    }
}
