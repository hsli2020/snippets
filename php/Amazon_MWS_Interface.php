<?php
/**
 * Marketplace Web Service PHP5 Library - Generated: Thu May 07 13:07:36 PDT 2009
 */

/**
 * The Amazon Marketplace Web Service contain APIs for inventory and order management.
 */
interface MarketplaceWebService_Interface
{
    /**
     * Get Report
     *
     * The GetReport operation returns the contents of a report. Reports can potentially
     * be very large (>100MB) which is why we only return one report at a time, and in a
     * streaming fashion.
     *
     * @see http://docs.amazonwebservices.com/${docPath}GetReport.html
     *
     * @param mixed $request array of parameters for 
     *    MarketplaceWebService_Model_GetReportRequest request
     * or 
     *    MarketplaceWebService_Model_GetReportRequest object itself
     *
     * @see MarketplaceWebService_Model_GetReportRequest
     *
     * @return MarketplaceWebService_Model_GetReportResponse
     *
     * @throws MarketplaceWebService_Exception
     */
    public function getReport($request);

    /**
     * Get Report Schedule Count
     *
     * returns the number of report schedules
     *
     * @param MarketplaceWebService_Model_GetReportScheduleCountRequest
     * @return MarketplaceWebService_Model_GetReportScheduleCountResponse
     */
    public function getReportScheduleCount($request);

    /**
     * Get Report Request List By Next Token
     *
     * retrieve the next batch of list items and if there are more items to retrieve
     *
     * @param MarketplaceWebService_Model_GetReportRequestListByNextTokenRequestreques
     * @return MarketplaceWebService_Model_GetReportRequestListByNextTokenResponse
     */
    public function getReportRequestListByNextToken($request);

    /**
     * Update Report Acknowledgements
     *
     * The UpdateReportAcknowledgements operation updates the acknowledged status 
     * of one or more reports.
     *
     * @param MarketplaceWebService_Model_UpdateReportAcknowledgementsRequest
     * @return MarketplaceWebService_Model_UpdateReportAcknowledgementsResponse
     */
    public function updateReportAcknowledgements($request);

    /**
     * Submit Feed
     *
     * Uploads a file for processing together with the necessary
     * metadata to process the file, such as which type of feed it is.
     * PurgeAndReplace if true means that your existing e.g. inventory is
     * wiped out and replace with the contents of this feed - use with
     * caution (the default is false).
     *
     * @param MarketplaceWebService_Model_SubmitFeedRequest
     * @return MarketplaceWebService_Model_SubmitFeedResponse
     */
    public function submitFeed($request);

    /**
     * Get Report Count
     *
     * returns a count of reports matching your criteria;
     * by default, the number of reports generated in the last 90 days,
     * regardless of acknowledgement status
     *
     * @param MarketplaceWebService_Model_GetReportCountRequest request
     * @return MarketplaceWebService_Model_GetReportCountResponse
     */
    public function getReportCount($request);

    /**
     * Get Feed Submission List By Next Token
     *
     * retrieve the next batch of list items and if there are more items to retrieve
     *
     * @param MarketplaceWebService_Model_GetFeedSubmissionListByNextTokenRequest request
     * @return MarketplaceWebService_Model_GetFeedSubmissionListByNextTokenResponse
     */
    public function getFeedSubmissionListByNextToken($request);

    /**
     * Cancel Feed Submissions
     *
     * cancels feed submissions - by default all of the submissions of the
     * last 30 days that have not started processing
     *
     * @param MarketplaceWebService_Model_CancelFeedSubmissionsRequest
     * @return MarketplaceWebService_Model_CancelFeedSubmissionsResponse
     */
    public function cancelFeedSubmissions($request);

    /**
     * Request Report
     *
     * requests the generation of a report
     *
     * @param MarketplaceWebService_Model_RequestReportRequest
     * @return MarketplaceWebService_Model_RequestReportResponse
     */
    public function requestReport($request);

    /**
     * Get Feed Submission Count
     *
     * returns the number of feeds matching all of the specified criteria
     *
     * @param MarketplaceWebService_Model_GetFeedSubmissionCountRequest
     * @return MarketplaceWebService_Model_GetFeedSubmissionCountResponse
     */
    public function getFeedSubmissionCount($request);

    /**
     * Cancel Report Requests
     *
     * cancels report requests that have not yet started processing,
     * by default all those within the last 90 days
     *
     * @param MarketplaceWebService_Model_CancelReportRequestsRequest
     * @return MarketplaceWebService_Model_CancelReportRequestsResponse
     */
    public function cancelReportRequests($request);

    /**
     * Get Report List
     *
     * returns a list of reports; by default the most recent ten reports,
     * regardless of their acknowledgement status
     *
     * @param MarketplaceWebService_Model_GetReportListRequest
     * @return MarketplaceWebService_Model_GetReportListResponse
     */
    public function getReportList($request);

    /**
     * Get Feed Submission Result
     *
     * retrieves the feed processing report
     *
     * @param MarketplaceWebService_Model_GetFeedSubmissionResultRequest
     * @return MarketplaceWebService_Model_GetFeedSubmissionResultResponse
     */
    public function getFeedSubmissionResult($request);

    /**
     * Get Feed Submission List
     *
     * returns a list of feed submission identifiers and their associated metadata
     *
     * @param MarketplaceWebService_Model_GetFeedSubmissionListRequest
     * @return MarketplaceWebService_Model_GetFeedSubmissionListResponse
     */
    public function getFeedSubmissionList($request);

    /**
     * Get Report Request List
     *
     * returns a list of report requests ids and their associated metadata
     *
     * @param MarketplaceWebService_Model_GetReportRequestListRequest
     * @return MarketplaceWebService_Model_GetReportRequestListResponse
     */
    public function getReportRequestList($request);

    /**
     * Get Report Schedule List By Next Token
     *
     * retrieve the next batch of list items and if there are more items to retrieve
     *
     * @param MarketplaceWebService_Model_GetReportScheduleListByNextTokenRequest
     * @return MarketplaceWebService_Model_GetReportScheduleListByNextTokenResponse
     */
    public function getReportScheduleListByNextToken($request);

    /**
     * Get Report List By Next Token
     *
     * retrieve the next batch of list items and if there are more items to retrieve
     *
     * @param MarketplaceWebService_Model_GetReportListByNextTokenRequest
     * @return MarketplaceWebService_Model_GetReportListByNextTokenResponse
     */
    public function getReportListByNextToken($request);

    /**
     * Manage Report Schedule
     *
     * Creates, updates, or deletes a report schedule
     * for a given report type, such as order reports in particular.
     *
     * @param MarketplaceWebService_Model_ManageReportScheduleRequest
     * @return MarketplaceWebService_Model_ManageReportScheduleResponse
     */
    public function manageReportSchedule($request);

    /**
     * Get Report Request Count
     *
     * returns a count of report requests; by default all the report
     * requests in the last 90 days
     *
     * @param MarketplaceWebService_Model_GetReportRequestCountRequest
     * @return MarketplaceWebService_Model_GetReportRequestCountResponse
     */
    public function getReportRequestCount($request);

    /**
     * Get Report Schedule List
     * returns the list of report schedules
     *
     * @param MarketplaceWebService_Model_GetReportScheduleListRequest
     * @return MarketplaceWebService_Model_GetReportScheduleListResponse
     */
    public function getReportScheduleList($request);
}
