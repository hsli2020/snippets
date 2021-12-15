<?php

$api = new SellingPartnerApi\Api\ReportsApi($config);
$reportType = ReportType::GET_MERCHANT_LISTINGS_ALL_DATA;
$spec = new SellingPartnerApi\Model\Reports\CreateReportSpecification();
$spec->setReportType($reportType['name']);
$spec->setMarketplaceIds(['ATVPDKIKX0DER']);

try {
    $result = $api->createReport($spec);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->createReport: ', $e->getMessage(), PHP_EOL;
}

//*
SellingPartnerApi\Model\Reports\CreateReportResponse Object
(
    [container:protected] => Array
    (
        [payload] => SellingPartnerApi\Model\Reports\CreateReportResult Object
        (
            [container:protected] => Array
            (
                [report_id] => 1384903018838
            )
        )
        [errors] =>
    )
)
//*/

# --------------------------------------------------------------------------------------

$api= new SellingPartnerApi\Api\ReportsApi($config);
$report_id = '1384903018838';

try {
    $result = $api->getReport($report_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->getReport: ', $e->getMessage(), PHP_EOL;
}

//*
SellingPartnerApi\Model\Reports\GetReportResponse Object
(
    [container:protected] => Array
    (
        [payload] => SellingPartnerApi\Model\Reports\Report Object
        (
            [container:protected] => Array
            (
                [marketplace_ids] => Array
                (
                    [0] => ATVPDKIKX0DER
                )

                [report_id] => 1384903018838
                [report_type] => GET_MERCHANT_LISTINGS_ALL_DATA
                [data_start_time] => DateTime Object
                (
                    [date] => 2021-07-30 20:31:06.000000
                    [timezone_type] => 1
                    [timezone] => +00:00
                )

                [data_end_time] => DateTime Object
                (
                    [date] => 2021-07-30 20:31:06.000000
                    [timezone_type] => 1
                    [timezone] => +00:00
                )

                [report_schedule_id] =>
                [created_time] => DateTime Object
                (
                    [date] => 2021-07-30 20:31:06.000000
                    [timezone_type] => 1
                    [timezone] => +00:00
                )

                [processing_status] => DONE
                [processing_start_time] => DateTime Object
                (
                    [date] => 2021-07-30 20:31:12.000000
                    [timezone_type] => 1
                    [timezone] => +00:00
                )

                [processing_end_time] => DateTime Object
                (
                    [date] => 2021-07-30 20:31:50.000000
                    [timezone_type] => 1
                    [timezone] => +00:00
                )

                [report_document_id] => amzn1.spdoc.1.3.d820dfb8-c770-429a-8fe8-e17ad14c63e7.T3T1O5V414OKWT.47700
            )
        )
        [errors] =>
    )
)
//*/

# --------------------------------------------------------------------------------------

$documentId = 'amzn1.spdoc.1.3.d820dfb8-c770-429a-8fe8-e17ad14c63e7.T3T1O5V414OKWT.47700';
$reportType = ReportType::GET_MERCHANT_LISTINGS_ALL_DATA;

$reportsApi = new SellingPartnerApi\Api\ReportsApi($config);
$reportDocumentInfo = $reportsApi->getReportDocument($documentId, $reportType['name']);

$doc = new SellingPartnerApi\Document($reportDocumentInfo->getPayload(), $reportType);
$contents = $doc->download();  // The raw report text

file_put_contents('amazon-us-full-listing.csv', $contents);
