<?php
#
# https://www.codefixup.com/how-to-use-amazon-mws-api-to-add-products-via-submit-feed/
#
$xml_string='<?xml version="1.0" encoding="iso-8859-1"?>
<AmazonEnvelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:noNamespaceSchemaLocation="amzn-envelope.xsd">
  <Header>
    <DocumentVersion>1.01</DocumentVersion>
    <MerchantIdentifier>M_EXAMPLE_123456</MerchantIdentifier>
  </Header>
  <MessageType>Product</MessageType>
  <PurgeAndReplace>false</PurgeAndReplace>
  <Message>
    <MessageID>1</MessageID>
    <OperationType>Update</OperationType>
    <Product>
      <SKU>56789</SKU>
      <StandardProductID>
        <Type>UPC</Type>
        <Value>463563647487</Value>
      </StandardProductID>
      <ProductTaxCode>A_GEN_NOTAX</ProductTaxCode>
      <DescriptionData>
        <Title>Example Product Title</Title>
        <Brand>Example Product Brand</Brand>
        <Description>This is an example product description.</Description>
        <BulletPoint>Example Bullet Point 1</BulletPoint>
        <BulletPoint>Example Bullet Point 2</BulletPoint>
        <MSRP currency="USD">25.19</MSRP>
        <Manufacturer>Example Product Manufacturer</Manufacturer>
        <ItemType>example-item-type</ItemType>
      </DescriptionData>
      <ProductData>
        <Health>
          <ProductType>
            <HealthMisc>
              <Ingredients>Example Ingredients</Ingredients>
              <Directions>Example Directions</Directions>
            </HealthMisc>
          </ProductType>
        </Health>
      </ProductData>
    </Product>
  </Message>
</AmazonEnvelope>'; 
 
$var = base64_encode(md5(trim($xml_string), true));
 
$param = array();
$param['AWSAccessKeyId']   = 'AKI************UPA'; 
$param['Action']           = 'SubmitFeed';  
$param['FeedType']         = '_POST_PRODUCT_DATA_';  
$param['MWSAuthToken']     = 'amzn.mws.**************-4565d8b3f16c';  
$param['MarketplaceId.Id'] = '************';    
$param['SellerId']         = '************'; 
$param['ContentMD5Value']  = $var; 
$param['FeedContent']      = trim($xml_string); 
$param['SignatureMethod']  = 'HmacSHA256'; 
$param['SignatureVersion'] = '2'; 
$param['Timestamp'] = gmdate("Y-m-d\TH:i:s.\\0\\0\\0\\Z", time()); 
$param['PurgeAndReplace']  = 'false';  
$param['Version']          = '2009-01-01';   
 
$secret = 'jFtS4y*****************gXUUgY';
 
$url = array();
foreach ($param as $key => $val) {
    $key = str_replace("%7E", "~", rawurlencode($key));
    $val = str_replace("%7E", "~", rawurlencode($val));
    $url[] = "{$key}={$val}";
}
 
sort($url);
 
$arr   = implode('&', $url);
 
$sign  = 'POST' . "\n";
$sign .= 'mws.amazonservices.com' . "\n";
$sign .= '/Feeds/2009-01-01' . "\n";
$sign .= $arr;
 
$signature = hash_hmac("sha256", $sign, $secret, true);
$signature = urlencode(base64_encode($signature));
 
$link  = "https://mws.amazonservices.com/Feeds/2009-01-01?";
$link .= $arr . "&Signature=" . $signature;
echo($link); //for debugging - you can paste this into a browser and see if it loads.
 
$headers=array(
    'Content-Type: application/xml',
    'Content-Length: ' . strlen($xml_string),
    'Content-MD5: ' . $var
);
 
$ch = curl_init($link);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST"); 
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xml_string);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE); 
$response = curl_exec($ch);
$info = curl_getinfo($ch);
curl_close($ch);
?>
