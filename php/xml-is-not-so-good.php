<?php

// This is an issue I encountered when integrating walmart api,
// If response is json format, the issue never happen.

// I get some orders through walmart api, the response looks like
$res1 = <<<EOS
<Response>
  <Order>
    <OrderNum>ORDER-1</OrderNum>
    <Item>ITEM-1</Item>
    <Qty>1</Qty>
  </Order>
  <Order>
    <OrderNum>ORDER-2</OrderNum>
    <Item>ITEM-2</Item>
    <Qty>2</Qty>
  </Order>
  <Order>
    <OrderNum>ORDER-3</OrderNum>
    <Item>ITEM-3</Item>
    <Qty>3</Qty>
  </Order>
</Response>
EOS;

$xml = simplexml_load_string($res1);
print_r($xml);

// This is no problem, everything is fine
/*
SimpleXMLElement Object
(
    [Order] => Array
    (
        [0] => SimpleXMLElement Object
        (
            [OrderNum] => ORDER-1
            [Item] => ITEM-1
            [Qty] => 1
        )

        [1] => SimpleXMLElement Object
        (
            [OrderNum] => ORDER-2
            [Item] => ITEM-2
            [Qty] => 2
        )

        [2] => SimpleXMLElement Object
        (
            [OrderNum] => ORDER-3
            [Item] => ITEM-3
            [Qty] => 3
        )
    )
)
*/

// In last api call, only one order returned from walmart
// the response looks like
$res2 = <<<EOS
<Response>
  <Order>
    <OrderNum>ORDER-ONLY-1</OrderNum>
    <Item>ITEM-ONLY-1</Item>
    <Qty>1</Qty>
  </Order>
</Response>
EOS;

$xml = simplexml_load_string($res2);
print_r($xml);

// This is really bad
/*
SimpleXMLElement Object
(
    [Order] => SimpleXMLElement Object
    (
        [OrderNum] => ORDER-ONLY-1
        [Item] => ITEM-ONLY-1
        [Qty] => 1
    )
)
*/

// It should be

/*
SimpleXMLElement Object
(
    [Order] => Array
    (
        [0] => SimpleXMLElement Object
        (
            [OrderNum] => ORDER-ONLY-1
            [Item] => ITEM-ONLY-1
            [Qty] => 1
        )
    )
)
*/

// To fix it, I wrote some hacking code
//
// if (empty($xml->Order[0])) {
//   $order = $xml->Order;
//   $xml->Order = [];
//   $xml->Order[] = $order;
// }
