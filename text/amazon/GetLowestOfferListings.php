<?php

$db = new \Phalcon\Db\Adapter\Pdo\Mysql(
    array(
        "host"     => "localhost",
        "username" => "root",
        "password" => "",
        "dbname"   => "bte",
    )
);

$text = file_get_contents('GetLowestOfferListings.xml');
$xml = simplexml_load_string($text);

$result      = $xml->GetLowestOfferListingsForSKUResult;
$product     = $result->Product;
$identifiers = $product->Identifiers;
$listings    = $product->LowestOfferListings->LowestOfferListing;

foreach ($listings as $listing) {
    $db->insertAsDict('Amazon_Lowest_Offer_Listings', [
        // Identifiers
        'MarketplaceId' => $identifiers->MarketplaceASIN->MarketplaceId,
        'SellerId'      => $identifiers->SKUIdentifier->SellerId,
        'SellerSKU'     => $identifiers->SKUIdentifier->SellerSKU,
        'ASIN'          => $identifiers->MarketplaceASIN->ASIN,

        'AllOfferListingsConsidered' => $result->AllOfferListingsConsidered,

        // Qualifiers
        'ItemCondition'                => $listing->Qualifiers->ItemCondition,
        'ItemSubcondition'             => $listing->Qualifiers->ItemSubcondition,
        'FulfillmentChannel'           => $listing->Qualifiers->FulfillmentChannel,
        'ShipsDomestically'            => $listing->Qualifiers->ShipsDomestically,
        'ShippingTime'                 => $listing->Qualifiers->ShippingTime->Max,
        'SellerPositiveFeedbackRating' => $listing->Qualifiers->SellerPositiveFeedbackRating,

        'NumberOfOfferListingsConsidered' => $listing->NumberOfOfferListingsConsidered,
        'SellerFeedbackCount'             => $listing->SellerFeedbackCount,
        'MultipleOffersAtLowestPrice'     => $listing->MultipleOffersAtLowestPrice,

        // Price
        'LandedPrice'  => $listing->Price->LandedPrice->Amount,
        'ListingPrice' => $listing->Price->ListingPrice->Amount,
        'Shipping'     => $listing->Price->Shipping->Amount,
        'CurrencyCode' => $listing->Price->LandedPrice->CurrencyCode,
    ]);
}
