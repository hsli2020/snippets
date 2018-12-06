<?php

$xml =<<<EOS
<?xml version="1.0" encoding="UTF-8"?>
<ns2:PartnerFeedResponse xmlns:ns2="http://walmart.com/">
<ns2:feedId>1c349f8f-aec0-411f-8454-ead47d12946f</ns2:feedId>
<ns2:feedStatus>PROCESSED</ns2:feedStatus>
<ns2:ingestionErrors />
<ns2:itemsReceived>11</ns2:itemsReceived>
<ns2:itemsSucceeded>11</ns2:itemsSucceeded>
<ns2:itemsFailed>0</ns2:itemsFailed>
<ns2:itemsProcessing>0</ns2:itemsProcessing>
<ns2:offset>0</ns2:offset>
<ns2:limit>0</ns2:limit>
<ns2:itemDetails>
<ns2:itemIngestionStatus>
<ns2:martId>0</ns2:martId>
<ns2:sku>sku1</ns2:sku>
<ns2:index>8</ns2:index>
<ns2:ingestionStatus>SUCCESS</ns2:ingestionStatus>
<ns2:ingestionErrors />
</ns2:itemIngestionStatus>
<ns2:itemIngestionStatus>
<ns2:martId>0</ns2:martId>
<ns2:sku>sku2</ns2:sku>
<ns2:index>6</ns2:index>
<ns2:ingestionStatus>SUCCESS</ns2:ingestionStatus>
<ns2:ingestionErrors />
</ns2:itemIngestionStatus>
<ns2:itemIngestionStatus>
<ns2:martId>0</ns2:martId>
<ns2:sku>sku3</ns2:sku>
<ns2:index>9</ns2:index>
<ns2:ingestionStatus>SUCCESS</ns2:ingestionStatus>
<ns2:ingestionErrors />
</ns2:itemIngestionStatus>
</ns2:itemDetails>
</ns2:PartnerFeedResponse>
EOS;

$doc = simplexml_load_string($xml, null, 0, 'ns2', true);
print_r($doc);

