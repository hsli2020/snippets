<?php

const EOL = PHP_EOL;

function pr($d) { var_export($d); echo EOL; }

//$xml = [ 'fields' => [ ],

class XmlNode
{
    protected $name;
    protected $value;

    public function __construct($name, $value='')
    {
        $this->name = $name;
        $this->value = $value;
    }

    public function setValue($value)
    {
        $this->value = $value;
    }

    public function __toString()
    {
        if ($this->value instanceof XmlNode) {
            return "<{$this->name}>\n{$this->value}</{$this->name}>\n";
        } else {
            return "<{$this->name}>{$this->value}</{$this->name}>\n";
        }
    }
}

$x1 = new XmlNode('SynnexB2B', '...');
$x2 = new XmlNode('OrderResponse', '...');
$x1->setValue($x2);
echo $x1;

/*
<?xml version="1.0" encoding="UTF-8"?>
<SynnexB2B>
    <OrderResponse>
        <CustomerNumber>1150897</CustomerNumber>
        <PONumber>702-4768232-1989057</PONumber>
        <Code>rejected</Code>
        <Reason>DUPPLICATED INSERT.. </Reason>
        <ResponseDateTime>2016-09-01T15:44:47</ResponseDateTime>
        <ResponseElapsedTime>2.146s</ResponseElapsedTime>
        <Items>
            <Item lineNumber="1">
                <SKU>5537035</SKU>
                <OrderQuantity>1</OrderQuantity>
                <Code>rejected</Code>
                <Reason>This PO# already exists in our system for your account, therefore we are rejecting this request to prevent a duplicate shipment. Please call your SYNNEX sales rep if any questions.</Reason>
                <OrderNumber>19236858</OrderNumber>
                <OrderType>99</OrderType>
                <ShipFromWarehouse></ShipFromWarehouse>
                <SynnexInternalReference>DUPREJECT,R-PODUPE---DUPPOKILL,R-PODUPE</SynnexInternalReference>
            </Item>
        </Items>
    </OrderResponse>
</SynnexB2B>
*/
