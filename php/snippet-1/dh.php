<?php

$xmldoc=<<<EOS
<?xml version="1.0" encoding="UTF-8" ?>
<XMLRESPONSE>
<ITEM>
   <PARTNUM>52T</PARTNUM>
   <UNITPRICE>6.63</UNITPRICE>
   <BRANCHQTY>
       <BRANCH>Toronto</BRANCH>
       <QTY>84</QTY>
       <INSTOCKDATE />
   </BRANCHQTY>
   <TOTALQTY>84</TOTALQTY>
</ITEM>
<ITEM>
   <PARTNUM>123ABC</PARTNUM>
   <MESSAGE>Invalid Item Number</MESSAGE>
</ITEM>
<STATUS>success</STATUS>
</XMLRESPONSE>
EOS;

$xml = simplexml_load_string($xmldoc);

$result = new stdClass();
$result->items = [];

$result->status = strval($xml->STATUS);
if ($result->status == 'success') {
    $result->status = 'STATUS_OK';
}

foreach ($xml->ITEM as $xitem) {
    if (empty($xitem->UNITPRICE))
        $xitem->UNITPRICE = 99999;

    $item = new stdClass();

    $item->sku   = 'DH-'. strval($xitem->PARTNUM);
    $item->price = strval($xitem->UNITPRICE);

    if ($xitem->BRANCHQTY) {
        $item->avail = [
            [
                'branch' => strval($xitem->BRANCHQTY->BRANCH),
                'qty'    => strval($xitem->BRANCHQTY->QTY),
            ]
        ];
        $item->instockDate = strval($xitem->BRANCHQTY->INSTOCKDATE);
    };

    if ($xitem->TOTALQTY) {
        $item->totalQty = strval($xitem->TOTALQTY);
    }

    if ($xitem->MESSAGE) {
        $item->status = strval($xitem->MESSAGE);
    }

    $result->items[] = $item;
}

print_r($result);

