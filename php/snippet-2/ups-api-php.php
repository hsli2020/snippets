https://webcollage.wordpress.com/2011/05/13/ups-tracking-with-php/

UPS label print with PHP
========================

Code to print UPS label.

For generating UPS label you just need to pass shipping digest which you will get once you get ship confirmation response.

// SHIP ACCEPT REQUEST
$xmlRequest1='<?xml version=”1.0″ encoding=”ISO-8859-1″?>
<AccessRequest>
<AccessLicenseNumber>ACCESS LICENCE NUMBER</AccessLicenseNumber>
<UserId>UPS USERNAME</UserId>
<Password>UPS PASSWORD</Password>
</AccessRequest>
<?xml version=”1.0″ encoding=”ISO-8859-1″?>
<ShipmentAcceptRequest>
<Request>
<TransactionReference>
<CustomerContext>Customer Comment</CustomerContext>
</TransactionReference>
<RequestAction>ShipAccept</RequestAction>
<RequestOption>1</RequestOption>
</Request>
<ShipmentDigest>SHIPMENT DIGEST</ShipmentDigest>
</ShipmentAcceptRequest>
‘;

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, “https://wwwcie.ups.com/ups.app/xml/ShipAccept&#8221;);
// uncomment the next line if you get curl error 60: error setting certificate verify locations
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
// uncommenting the next line is most likely not necessary in case of error 60
// curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlRequest1);
curl_setopt($ch, CURLOPT_TIMEOUT, 3600);

//if ($this->logfile) {
//   error_log(“UPS REQUEST: ” . $xmlRequest . “\n”, 3, $this->logfile);
//}
$xmlResponse = curl_exec ($ch); // SHIP ACCEPT RESPONSE
//echo curl_errno($ch);

$xml = $xmlResponse;

preg_match_all( “/\<ShipmentAcceptResponse\>(.*?)\<\/ShipmentAcceptResponse\>/s”,
$xml, $bookblocks );

foreach( $bookblocks[1] as $block )
{
preg_match_all( “/\<GraphicImage\>(.*?)\<\/GraphicImage\>/”,
$block, $author ); // GET LABEL

preg_match_all( “/\<TrackingNumber\>(.*?)\<\/TrackingNumber\>/”,
$block, $tracking ); // GET TRACKING NUMBER
//echo( $author[1][0].”\n” );
}

echo ‘<img src=”data:image/gif;base64,’. $author[1][0]. ‘”/>’;






UPS address verification system(XAV)
====================================

Code to check about address validation for the UPS shipping.

Below code can be use to check whether customer shipping address is valid according to the UPS/USPS database or not.

//STREET LEVEL ADDRESS VARIFICATION REQUEST

$xmlRequest1='<?xml version=”1.0″?>
<AccessRequest xml:lang=”en-US”>
<AccessLicenseNumber>ACCESS LICENCE NUMBER</AccessLicenseNumber>
<UserId>UPS USERNAME</UserId>
<Password>UPS PASSWORD</Password>
</AccessRequest>
<?xml version=”1.0″?>
<AddressValidationRequest xml:lang=”en-US”>
<Request>
<TransactionReference>
<CustomerContext>Your Test Case Summary Description</CustomerContext>
<XpciVersion>1.0</XpciVersion>
</TransactionReference>
<RequestAction>XAV</RequestAction>
<RequestOption>3</RequestOption>
</Request>

<AddressKeyFormat>
<AddressLine>AIRWAY ROAD SUITE 7</AddressLine>
<PoliticalDivision2>SAN DIEGO</PoliticalDivision2>
<PoliticalDivision1>CA</PoliticalDivision1>
<PostcodePrimaryLow>92154</PostcodePrimaryLow>
<CountryCode>US</CountryCode>
</AddressKeyFormat>
</AddressValidationRequest>’;

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, “https://wwwcie.ups.com/ups.app/xml/XAV&#8221;);
// uncomment the next line if you get curl error 60: error setting certificate verify locations
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
// uncommenting the next line is most likely not necessary in case of error 60
// curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlRequest1);
curl_setopt($ch, CURLOPT_TIMEOUT, 3600);

//if ($this->logfile) {
//   error_log(“UPS REQUEST: ” . $xmlRequest . “\n”, 3, $this->logfile);
//}
echo $xmlResponse = curl_exec ($ch); // SHIP CONFIRMATION RESPONSE
$xml = $xmlResponse;

preg_match_all( “/\<AddressValidationResponse\>(.*?)\<\/AddressValidationResponse\>/s”,
$xml, $upsRes );

foreach( $upsRes[1] as $res )
{
preg_match_all( “/\<ResponseStatusCode\>(.*?)\<\/ResponseStatusCode\>/”,
$res, $response ); // XAV CODE

preg_match_all( “/\<ResponseStatusDescription\>(.*?)\<\/ResponseStatusDescription\>/”,
$res, $responseMsg ); // XAV MESSAGE

preg_match_all( “/\<AddressLine\>(.*?)\<\/AddressLine\>/”,
$res, $address ); // Possible Address Line

preg_match_all( “/\<Region\>(.*?)\<\/Region\>/”,
$res, $Region ); // Possible region

preg_match_all( “/\<PoliticalDivision2\>(.*?)\<\/PoliticalDivision2\>/”,
$res, $city );

preg_match_all( “/\<PoliticalDivision1\>(.*?)\<\/PoliticalDivision1\>/”,
$res, $state );

preg_match_all( “/\<PostcodePrimaryLow\>(.*?)\<\/PostcodePrimaryLow\>/”,
$res, $postCode );

preg_match_all( “/\<PostcodeExtendedLow\>(.*?)\<\/PostcodeExtendedLow\>/”,
$res, $extenCode );

}
for($i=0;$i<count($address[1]);$i++)
{
echo $address[1][$i];
echo “<br>”;
echo $Region[1][$i];
echo “<br>”;
echo $city[1][$i];
echo “<br>”;
echo $state[1][$i];
echo “<br>”;
echo $postCode[1][$i];
echo “<br>”;
echo $extenCode[1][$i];
echo “<br>”;
echo “<hr>”;
}





UPS shipping confirmation code in php
=====================================
10 May

Here is the code to generate UPS digest which will use to generate UPS label<Br /><BR />

Replace credential detail with the correct one and use with your customer address details<br /><Br />
// SHIP CONFIRMATION REQUEST

$xmlRequest1='<?xml version=”1.0″?>
<AccessRequest xml:lang=”en-US”>
<AccessLicenseNumber>ACCESS LICENCE NUMBER</AccessLicenseNumber>
<UserId>UPS USERNAME</UserId>
<Password>UPS PASSWORD</Password>
</AccessRequest>
<?xml version=”1.0″?>
<ShipmentConfirmRequest xml:lang=”en-US”>
<Request>
<TransactionReference>
<CustomerContext>Customer Comment</CustomerContext>
<XpciVersion/>
</TransactionReference>
<RequestAction>ShipConfirm</RequestAction>
<RequestOption>validate</RequestOption>
</Request>
<LabelSpecification>
<LabelPrintMethod>
<Code>GIF</Code>
<Description>gif file</Description>
</LabelPrintMethod>
<HTTPUserAgent>Mozilla/4.5</HTTPUserAgent>
<LabelImageFormat>
<Code>GIF</Code>
<Description>gif</Description>
</LabelImageFormat>
</LabelSpecification>
<Shipment>
<RateInformation>
<NegotiatedRatesIndicator/>
</RateInformation>
<Description/>
<Shipper>
<Name>TEST</Name>
<PhoneNumber>111-111-1111</PhoneNumber>
<ShipperNumber>SHIPPER NUMBER</ShipperNumber>
<TaxIdentificationNumber>1234567890</TaxIdentificationNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<PostcodeExtendedLow></PostcodeExtendedLow>
<CountryCode>US</CountryCode>
</Address>
</Shipper>
<ShipTo>
<CompanyName>Yats</CompanyName>
<AttentionName>Yats</AttentionName>
<PhoneNumber>123.456.7890</PhoneNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<CountryCode>US</CountryCode>
</Address>
</ShipTo>
<ShipFrom>
<CompanyName>Ship From Company Name</CompanyName>
<AttentionName>Ship From Attn Name</AttentionName>
<PhoneNumber>1234567890</PhoneNumber>
<TaxIdentificationNumber>1234567877</TaxIdentificationNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<CountryCode>US</CountryCode>
</Address>
</ShipFrom>
<PaymentInformation>
<Prepaid>
<BillShipper>
<AccountNumber>SHIPPER NUMBER</AccountNumber>
</BillShipper>
</Prepaid>
</PaymentInformation>
<Service>
<Code>02</Code>
<Description>2nd Day Air</Description>
</Service>
<Package>
<PackagingType>
<Code>02</Code>
<Description>Customer Supplied</Description>
</PackagingType>
<Description>Package Description</Description>
<ReferenceNumber>
<Code>00</Code>
<Value>Package</Value>
</ReferenceNumber>
<PackageWeight>
<UnitOfMeasurement/>
<Weight>60.0</Weight>
</PackageWeight>
<LargePackageIndicator/>
<AdditionalHandling>0</AdditionalHandling>
</Package>
</Shipment>
</ShipmentConfirmRequest>
‘;
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, “https://wwwcie.ups.com/ups.app/xml/ShipConfirm&#8221;);
// uncomment the next line if you get curl error 60: error setting certificate verify locations
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
// uncommenting the next line is most likely not necessary in case of error 60
// curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlRequest1);
curl_setopt($ch, CURLOPT_TIMEOUT, 3600);

//if ($this->logfile) {
//   error_log(“UPS REQUEST: ” . $xmlRequest . “\n”, 3, $this->logfile);
//}
echo $xmlResponse = curl_exec ($ch); // SHIP CONFORMATION RESPONSE
//echo curl_errno($ch);

$xml = $xmlResponse;

preg_match_all( “/\<ShipmentConfirmResponse\>(.*?)\<\/ShipmentConfirmResponse\>/s”,
$xml, $bookblocks );

foreach( $bookblocks[1] as $block )
{
preg_match_all( “/\<ShipmentDigest\>(.*?)\<\/ShipmentDigest\>/”,
$block, $author ); // SHIPPING DIGEST

//echo( $author[1][0].”\n” );
}

<br><br />
Enjoy!!!
Rate this:




PHP code to get UPS shipping rates
==================================

I have prepared this code to become helpfull for the developer who wants to integrate UPS in their customer website.

You just needs to make appropriate changes in credential detail and you are with the shipping rates.

//GET RATE FROM UPS API

$xmlRequest1='<?xml version=”1.0″?>
<AccessRequest xml:lang=”en-US”>
<AccessLicenseNumber>ACCESSLICENCENUMBER</AccessLicenseNumber>
<UserId>USERID</UserId>
<Password>UPS PASSWORD</Password>
</AccessRequest>
<?xml version=”1.0″?>
<RatingServiceSelectionRequest xml:lang=”en-US”>
<Request>
<TransactionReference>
<CustomerContext>Rating and Service</CustomerContext>
<XpciVersion>1.0</XpciVersion>
</TransactionReference>
<RequestAction>Rate</RequestAction>
<RequestOption>Rate</RequestOption>
</Request>
<PickupType>
<Code>07</Code>
<Description>Rate</Description>
</PickupType>
<Shipment>
<Description>Rate Description</Description>
<Shipper>
<Name>TEST</Name>
<PhoneNumber>888-748-7446</PhoneNumber>
<ShipperNumber>SHIPPER NUMBER</ShipperNumber>
<TaxIdentificationNumber>1234567877</TaxIdentificationNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<PostcodeExtendedLow></PostcodeExtendedLow>
<CountryCode>US</CountryCode>
</Address>
</Shipper>
<ShipTo>
<CompanyName>Yats</CompanyName>
<AttentionName>Yats</AttentionName>
<PhoneNumber>866.345.7638</PhoneNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<CountryCode>US</CountryCode>
</Address>
</ShipTo>
<ShipFrom>
<CompanyName>Ship From Company Name</CompanyName>
<AttentionName>Ship From Attn Name</AttentionName>
<PhoneNumber>1234567890</PhoneNumber>
<TaxIdentificationNumber>1234567877</TaxIdentificationNumber>
<Address>
<AddressLine1>AIRWAY ROAD SUITE 7</AddressLine1>
<City>SAN DIEGO</City>
<StateProvinceCode>CA</StateProvinceCode>
<PostalCode>92154</PostalCode>
<CountryCode>US</CountryCode>
</Address>
</ShipFrom>
<Service>
<Code>02</Code>
</Service>
<PaymentInformation>
<Prepaid>
<BillShipper>
<AccountNumber>Ship Number</AccountNumber>
</BillShipper>
</Prepaid>
</PaymentInformation>
<Package>
<PackagingType>
<Code>00</Code>
<Description>Customer Supplied</Description>
</PackagingType>
<Dimensions>
<UnitOfMeasurement>
<Code>IN</Code>
</UnitOfMeasurement>
<Length>30</Length>
<Width>34</Width>
<Height>34</Height>
</Dimensions>
<Description>Rate</Description>
<PackageWeight>
<UnitOfMeasurement>
<Code>LBS</Code>
</UnitOfMeasurement>
<Weight>150</Weight>
</PackageWeight>
</Package>
<ShipmentServiceOptions>
<OnCallAir>
<Schedule>
<PickupDay>02</PickupDay>
<Method>02</Method>
</Schedule>
</OnCallAir>
</ShipmentServiceOptions>
</Shipment>
</RatingServiceSelectionRequest>’;

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, “https://wwwcie.ups.com/ups.app/xml/Rate&#8221;);
// uncomment the next line if you get curl error 60: error setting certificate verify locations
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
// uncommenting the next line is most likely not necessary in case of error 60
// curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlRequest1);
curl_setopt($ch, CURLOPT_TIMEOUT, 3600);
echo $xmlResponse = curl_exec ($ch); // SHIP RATE RESPONSE



UPS tracking with php
=====================
13 May

Code to track shipped order.

Below code will help you to track your shipped order with UPS.

// UPS SHIP ORDER TRACKING
$xmlRequest1='<?xml version=”1.0″?>
<AccessRequest xml:lang=”en-US”>
<AccessLicenseNumber>ACCESS LICENCE NUMBER</AccessLicenseNumber>
<UserId>UPS USERNAME</UserId>
<Password>UPS PASSWORD</Password>
</AccessRequest>
<?xml version=”1.0″?>
<TrackRequest xml:lang=”en-US”>
<Request>
<TransactionReference>
<CustomerContext>Your Test Case Summary
Description</CustomerContext>
<XpciVersion>1.0</XpciVersion>
</TransactionReference>
<RequestAction>Track</RequestAction>
<RequestOption>activity</RequestOption>
</Request>
<TrackingNumber>W23WSDFFE23443</TrackingNumber>
</TrackRequest>’;

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, “https://wwwcie.ups.com/ups.app/xml/Track&#8221;);
// uncomment the next line if you get curl error 60: error setting certificate verify locations
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
// uncommenting the next line is most likely not necessary in case of error 60
// curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $xmlRequest1);
curl_setopt($ch, CURLOPT_TIMEOUT, 3600);

//if ($this->logfile) {
//   error_log(“UPS REQUEST: ” . $xmlRequest . “\n”, 3, $this->logfile);
//}
echo $xmlResponse = curl_exec ($ch); // TRACKING RESPONSE




