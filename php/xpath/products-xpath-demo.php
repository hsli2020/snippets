<?php

$xmldoc = <<<XML
<?xml version="1.0" encoding="UTF-8"?>
<PRODUCTS>
  <PRODUCT category="software">
    <SKU>soft1234</SKU>
    <SUB_CATEGORY>Business Analysis</SUB_CATEGORY>
    <NAME>Widget Reporting</NAME>
    <PRICE>4500</PRICE>
  </PRODUCT>
  <PRODUCT category="software">
    <SKU>soft5678</SKU>
    <SUB_CATEGORY>Business Analysis</SUB_CATEGORY>
    <NAME>Pro Reporting</NAME>
    <PRICE>2300</PRICE>
  </PRODUCT>
  <PRODUCT category="storage">
    <SKU>soft2233</SKU>
    <SUB_CATEGORY>Tape Systems</SUB_CATEGORY>
    <NAME>Tapes Abound</NAME>
    <PRICE>2300</PRICE>
  </PRODUCT>
  <PRODUCT category="storage">
    <SKU>soft2233</SKU>
    <SUB_CATEGORY>Disk Systems</SUB_CATEGORY>
    <NAME>Widget100 Series</NAME>
    <PRICE>6500</PRICE>
  </PRODUCT>
</PRODUCTS>
XML;

$xml = simplexml_load_string($xmldoc);

//$products = $xml->xpath("/PRODUCTS");
//print_r($products);

//$products = $xml->xpath("/PRODUCTS/PRODUCT/NAME");
//print_r($products);

//$products = $xml->xpath("/PRODUCTS/PRODUCT[SKU='soft5678']/NAME");
//$products = $xml->xpath("/PRODUCTS/PRODUCT[SKU='soft2233']/NAME");
//print_r($products);

//$products = $xml->xpath("/PRODUCTS/PRODUCT[@category='software' and PRICE > 2500]"); 
//print_r($products);

//$doc = new DOMDocument;
//$doc->loadXML($xmldoc);
//$xpath = new DOMXPath($doc);
//$products = $xpath->query("/PRODUCTS/PRODUCT[SKU='soft5678']/NAME");
//  
//foreach ($products as $product) {
//   print($product->nodeValue);
//}
