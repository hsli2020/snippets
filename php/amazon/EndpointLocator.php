<?php
/**
 * All Rights Reserved
 * @copyright Copyright (C) 2021 Web-Com Group
 */

namespace Webcom\Amazon\Rest;

/**
 * Class EndpointLocator
 * @author magik092
 */
class EndpointLocator
{
    /**
     * Endpoint details
     * @var string[][]
     */
    public static array $endpoints = [
        "DE" => [
            "id"                               => "A1PA6795UKMFR9",
            "name"                             => "Germany",
            "countryCode"                      => "DE",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.de",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.de",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "ES" => [
            "id"                               => "A1RKKUPIHCS9HS",
            "name"                             => "Spain",
            "countryCode"                      => "ES",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.es",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.es",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "FR" => [
            "id"                               => "A13V1IB3VIYZZH",
            "name"                             => "France",
            "countryCode"                      => "FR",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.fr",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.fr",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "IT" => [
            "id"                               => "APJ6JRA9NG5V4",
            "name"                             => "Italy",
            "countryCode"                      => "IT",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.it",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.it",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "UK" => [
            "id"                               => "A1F83G8C2ARO7P",
            "name"                             => "United Kingdom",
            "countryCode"                      => "UK",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.co.uk",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.co.uk",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "TR" => [
            "id"                               => "A33AVAJ2PDY3EV",
            "name"                             => "Turkey",
            "countryCode"                      => "TR",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.com.tr",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral-europe.amazon.com",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "EG" => [
            "id"                               => "ARBP9OOSHTCHU",
            "name"                             => "Egypt",
            "countryCode"                      => "EG",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.com",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral-europe.amazon.com",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "SA" => [
            "id"                               => "A17E79C6D8DWNP",
            "name"                             => "Saudi Arabia",
            "countryCode"                      => "SA",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.sa",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.sa",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "NL" => [
            "id"                               => "A1805IZSGTT6HS",
            "name"                             => "Netherlands",
            "countryCode"                      => "NL",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.nl",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.nl",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "SE" => [
            "id"                               => "A2NODRKZP88ZB9",
            "name"                             => "Sweden",
            "countryCode"                      => "SE",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.se",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.se",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "PL" => [
            "id"                               => "A1C3SOZRARQ6R3",
            "name"                             => "Poland",
            "countryCode"                      => "PL",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.pl",
            "mwsEndpoint"                      => "mws-eu.amazonservices.com",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "http://sell.amazon.pl",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "CA" => [
            "id"                               => "A2EUQ1WTGCTBG2",
            "name"                             => "Canada",
            "countryCode"                      => "CA",
            "region"                           => "us-east-1",
            "endpoint"                         => "amazon.ca",
            "mwsEndpoint"                      => "mws.amazonservices.ca",
            "imagesEndpoint"                   => "images-na.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sell.amazon.ca",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-na.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-na.amazon.com",
        ],
        "MX" => [
            "id"                               => "A1AM78C64UM0Y8",
            "name"                             => "Mexico",
            "countryCode"                      => "MX",
            "region"                           => "us-east-1",
            "endpoint"                         => "amazon.com.mx",
            "mwsEndpoint"                      => "mws.amazonservices.com.mx",
            "imagesEndpoint"                   => "images-na.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.com.mx",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-na.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-na.amazon.com",
        ],
        "US" => [
            "id"                               => "ATVPDKIKX0DER",
            "name"                             => "United States",
            "countryCode"                      => "US",
            "region"                           => "us-east-1",
            "endpoint"                         => "amazon.com",
            "mwsEndpoint"                      => "mws.amazonservices.com",
            "imagesEndpoint"                   => "images-na.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.com",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-na.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-na.amazon.com",
        ],
        "BR" => [
            "id"                               => "A2Q3Y263D00KWC",
            "name"                             => "Brazil",
            "countryCode"                      => "BR",
            "region"                           => "us-east-1",
            "endpoint"                         => "amazon.com.br",
            "mwsEndpoint"                      => "mws.amazonservices.com",
            "imagesEndpoint"                   => "images-na.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.com.br",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-na.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-na.amazon.com",
        ],
        "AE" => [
            "id"                               => "A2VIGQ35RCS4UG",
            "name"                             => "United Arab Emirates",
            "countryCode"                      => "AE",
            "region"                           => "us-west-2",
            "endpoint"                         => "amazon.ae",
            "mwsEndpoint"                      => "mws.amazonservices.ae",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.ae/",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-fe.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-fe.amazon.com",
        ],
        "IN" => [
            "id"                               => "A21TJRUUN4KGV",
            "name"                             => "India",
            "countryCode"                      => "IN",
            "region"                           => "eu-west-1",
            "endpoint"                         => "amazon.in",
            "mwsEndpoint"                      => "mws.amazonservices.in",
            "imagesEndpoint"                   => "images-eu.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://mws.amazonservices.in",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-eu.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-eu.amazon.com",
        ],
        "AU" => [
            "id"                               => "A39IBJ37TRP1C6",
            "name"                             => "Australia",
            "countryCode"                      => "AU",
            "region"                           => "us-west-2",
            "endpoint"                         => "amazon.com.au",
            "mwsEndpoint"                      => "mws.amazonservices.com.au",
            "imagesEndpoint"                   => "images-fe.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.com.au",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-fe.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-fe.amazon.com",
        ],
        "JP" => [
            "id"                               => "A1VC38T7YXB528",
            "name"                             => "Japan",
            "countryCode"                      => "JP",
            "region"                           => "us-west-2",
            "endpoint"                         => "amazon.co.jp",
            "mwsEndpoint"                      => "mws.amazonservices.jp",
            "imagesEndpoint"                   => "images-fe.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.co.jp",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-fe.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-fe.amazon.com",
        ],
        "SG" => [
            "id"                               => "A19VAU5U5O7RUS",
            "name"                             => "Singapore",
            "countryCode"                      => "SG",
            "region"                           => "us-west-2",
            "endpoint"                         => "amazon.sg",
            "mwsEndpoint"                      => "mws-fe.amazonservices.com",
            "imagesEndpoint"                   => "images-fe.ssl-images-amazon.com",
            "sellerCentralEndpoint"            => "https://sellercentral.amazon.sg",
            "sellingPartnerApiEndpoint"        => "https://sellingpartnerapi-fe.amazon.com",
            "sellingPartnerApiSandboxEndpoint" => "https://sandbox.sellingpartnerapi-fe.amazon.com",
        ],
    ];

    /**
     * @param string $marketplaceId
     * @return Endpoint
     * @throws ApiException
     */
    public static function resolveByMarketplaceId(string $marketplaceId): Endpoint
    {
        foreach (self::$endpoints as $endpoint) {
            if ($endpoint['id'] === $marketplaceId) {
                return self::resolveByCountryCode($endpoint['countryCode']);
            }
        }

        throw new ApiException('No endpoint found for marketplace ' . $marketplaceId);
    }

    /**
     * @param string $countryCode
     * @return Endpoint
     * @throws ApiException
     */
    public static function resolveByCountryCode(string $countryCode): Endpoint
    {
        if (isset(self::$endpoints[$countryCode])) {
            return new Endpoint(
                self::$endpoints[$countryCode]['id'],
                self::$endpoints[$countryCode]['name'],
                self::$endpoints[$countryCode]['countryCode'],
                self::$endpoints[$countryCode]['region'],
                self::$endpoints[$countryCode]['endpoint'],
                self::$endpoints[$countryCode]['mwsEndpoint'],
                self::$endpoints[$countryCode]['imagesEndpoint'],
                self::$endpoints[$countryCode]['sellerCentralEndpoint'],
                self::$endpoints[$countryCode]['sellingPartnerApiEndpoint'],
                self::$endpoints[$countryCode]['sellingPartnerApiSandboxEndpoint'],
            );
        }

        throw new ApiException('No endpoint found for country code ' . $countryCode);
    }
}
