
type AuthorizationService interface {
	GetAuthorizationCode(params *GetAuthorizationCodeParams)
        (*GetAuthorizationCodeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type CatalogService interface {
	GetCatalogItem(params *GetCatalogItemParams)
        (*GetCatalogItemOK, error)

	ListCatalogCategories(params *ListCatalogCategoriesParams)
        (*ListCatalogCategoriesOK, error)

	ListCatalogItems(params *ListCatalogItemsParams)
        (*ListCatalogItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type FbaInboundService interface {
	ConfirmPreorder(params *ConfirmPreorderParams)
        (*ConfirmPreorderOK, error)

	ConfirmTransport(params *ConfirmTransportParams)
        (*ConfirmTransportOK, error)

	CreateInboundShipment(params *CreateInboundShipmentParams)
        (*CreateInboundShipmentOK, error)

	CreateInboundShipmentPlan(params *CreateInboundShipmentPlanParams)
        (*CreateInboundShipmentPlanOK, error)

	EstimateTransport(params *EstimateTransportParams)
        (*EstimateTransportOK, error)

	GetBillOfLading(params *GetBillOfLadingParams)
        (*GetBillOfLadingOK, error)

	GetInboundGuidance(params *GetInboundGuidanceParams)
        (*GetInboundGuidanceOK, error)

	GetLabels(params *GetLabelsParams)
        (*GetLabelsOK, error)

	GetPreorderInfo(params *GetPreorderInfoParams)
        (*GetPreorderInfoOK, error)

	GetPrepInstructions(params *GetPrepInstructionsParams)
        (*GetPrepInstructionsOK, error)

	GetShipmentItems(params *GetShipmentItemsParams)
        (*GetShipmentItemsOK, error)

	GetShipmentItemsByShipmentID(params *GetShipmentItemsByShipmentIDParams)
        (*GetShipmentItemsByShipmentIDOK, error)

	GetShipments(params *GetShipmentsParams)
        (*GetShipmentsOK, error)

	GetTransportDetails(params *GetTransportDetailsParams)
        (*GetTransportDetailsOK, error)

	PutTransportDetails(params *PutTransportDetailsParams)
        (*PutTransportDetailsOK, error)

	UpdateInboundShipment(params *UpdateInboundShipmentParams)
        (*UpdateInboundShipmentOK, error)

	VoidTransport(params *VoidTransportParams)
        (*VoidTransportOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type FbaOutboundService interface {
	CancelFulfillmentOrder(params *CancelFulfillmentOrderParams)
        (*CancelFulfillmentOrderOK, error)

	CreateFulfillmentOrder(params *CreateFulfillmentOrderParams)
        (*CreateFulfillmentOrderOK, error)

	CreateFulfillmentReturn(params *CreateFulfillmentReturnParams)
        (*CreateFulfillmentReturnOK, error)

	GetFeatureInventory(params *GetFeatureInventoryParams)
        (*GetFeatureInventoryOK, error)

	GetFeatureSKU(params *GetFeatureSKUParams)
        (*GetFeatureSKUOK, error)

	GetFeatures(params *GetFeaturesParams)
        (*GetFeaturesOK, error)

	GetFulfillmentOrder(params *GetFulfillmentOrderParams)
        (*GetFulfillmentOrderOK, error)

	GetFulfillmentPreview(params *GetFulfillmentPreviewParams)
        (*GetFulfillmentPreviewOK, error)

	GetPackageTrackingDetails(params *GetPackageTrackingDetailsParams)
        (*GetPackageTrackingDetailsOK, error)

	ListAllFulfillmentOrders(params *ListAllFulfillmentOrdersParams)
        (*ListAllFulfillmentOrdersOK, error)

	ListReturnReasonCodes(params *ListReturnReasonCodesParams)
        (*ListReturnReasonCodesOK, error)

	UpdateFulfillmentOrder(params *UpdateFulfillmentOrderParams)
        (*UpdateFulfillmentOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type FeedService interface {
	CancelFeed(params *CancelFeedParams)
        (*CancelFeedOK, error)

	CreateFeed(params *CreateFeedParams)
        (*CreateFeedAccepted, error)

	CreateFeedDocument(params *CreateFeedDocumentParams)
        (*CreateFeedDocumentCreated, error)

	GetFeed(params *GetFeedParams)
        (*GetFeedOK, error)

	GetFeedDocument(params *GetFeedDocumentParams)
        (*GetFeedDocumentOK, error)

	GetFeeds(params *GetFeedsParams)
        (*GetFeedsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type FeeService interface {
	GetMyFeesEstimateForASIN(params *GetMyFeesEstimateForASINParams)
        (*GetMyFeesEstimateForASINOK, error)

	GetMyFeesEstimateForSKU(params *GetMyFeesEstimateForSKUParams)
        (*GetMyFeesEstimateForSKUOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type MerchantFulfillmentService interface {
	CancelShipment(params *CancelShipmentParams)
        (*CancelShipmentOK, error)

	CancelShipmentOld(params *CancelShipmentOldParams)
        (*CancelShipmentOldOK, error)

	CreateShipment(params *CreateShipmentParams)
        (*CreateShipmentOK, error)

	GetAdditionalSellerInputs(params *GetAdditionalSellerInputsParams)
        (*GetAdditionalSellerInputsOK, error)

	GetAdditionalSellerInputsOld(params *GetAdditionalSellerInputsOldParams)
        (*GetAdditionalSellerInputsOldOK, error)

	GetEligibleShipmentServices(params *GetEligibleShipmentServicesParams)
        (*GetEligibleShipmentServicesOK, error)

	GetEligibleShipmentServicesOld(params *GetEligibleShipmentServicesOldParams)
        (*GetEligibleShipmentServicesOldOK, error)

	GetShipment(params *GetShipmentParams)
        (*GetShipmentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type MessagingService interface {
	CreateAmazonMotors(params *CreateAmazonMotorsParams)
        (*CreateAmazonMotorsCreated, error)

	CreateWarranty(params *CreateWarrantyParams)
        (*CreateWarrantyCreated, error)

	GetAttributes(params *GetAttributesParams)
        (*GetAttributesOK, error)

	ConfirmCustomizationDetails(params *ConfirmCustomizationDetailsParams)
        (*ConfirmCustomizationDetailsCreated, error)

	CreateConfirmDeliveryDetails(params *CreateConfirmDeliveryDetailsParams)
        (*CreateConfirmDeliveryDetailsCreated, error)

	CreateConfirmOrderDetails(params *CreateConfirmOrderDetailsParams)
        (*CreateConfirmOrderDetailsCreated, error)

	CreateConfirmServiceDetails(params *CreateConfirmServiceDetailsParams)
        (*CreateConfirmServiceDetailsCreated, error)

	CreateDigitalAccessKey(params *CreateDigitalAccessKeyParams)
        (*CreateDigitalAccessKeyCreated, error)

	CreateLegalDisclosure(params *CreateLegalDisclosureParams)
        (*CreateLegalDisclosureCreated, error)

	CreateNegativeFeedbackRemoval(params *CreateNegativeFeedbackRemovalParams)
        (*CreateNegativeFeedbackRemovalCreated, error)

	CreateUnexpectedProblem(params *CreateUnexpectedProblemParams)
        (*CreateUnexpectedProblemCreated, error)

	GetMessagingActionsForOrder(params *GetMessagingActionsForOrderParams)
        (*GetMessagingActionsForOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type NotificationsService interface {
	CreateDestination(params *CreateDestinationParams)
        (*CreateDestinationOK, error)

	CreateSubscription(params *CreateSubscriptionParams)
        (*CreateSubscriptionOK, error)

	DeleteDestination(params *DeleteDestinationParams)
        (*DeleteDestinationOK, error)

	DeleteSubscriptionByID(params *DeleteSubscriptionByIDParams)
        (*DeleteSubscriptionByIDOK, error)

	GetDestination(params *GetDestinationParams)
        (*GetDestinationOK, error)

	GetDestinations(params *GetDestinationsParams)
        (*GetDestinationsOK, error)

	GetSubscription(params *GetSubscriptionParams)
        (*GetSubscriptionOK, error)

	GetSubscriptionByID(params *GetSubscriptionByIDParams)
        (*GetSubscriptionByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type OperationService interface {
	ListFinancialEventGroups(params *ListFinancialEventGroupsParams)
        (*ListFinancialEventGroupsOK, error)

	ListFinancialEvents(params *ListFinancialEventsParams)
        (*ListFinancialEventsOK, error)

	ListFinancialEventsByGroupID(params *ListFinancialEventsByGroupIDParams)
        (*ListFinancialEventsByGroupIDOK, error)

	ListFinancialEventsByOrderID(params *ListFinancialEventsByOrderIDParams)
        (*ListFinancialEventsByOrderIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type OrderService interface {
	GetOrder(params *GetOrderParams)
        (*GetOrderOK, error)

	GetOrderAddress(params *GetOrderAddressParams)
        (*GetOrderAddressOK, error)

	GetOrderBuyerInfo(params *GetOrderBuyerInfoParams)
        (*GetOrderBuyerInfoOK, error)

	GetOrderItems(params *GetOrderItemsParams)
        (*GetOrderItemsOK, error)

	GetOrderItemsBuyerInfo(params *GetOrderItemsBuyerInfoParams)
        (*GetOrderItemsBuyerInfoOK, error)

	GetOrders(params *GetOrdersParams)
        (*GetOrdersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type ProductPricingService interface {
	GetCompetitivePricing(params *GetCompetitivePricingParams)
        (*GetCompetitivePricingOK, error)

	GetItemOffers(params *GetItemOffersParams)
        (*GetItemOffersOK, error)

	GetListingOffers(params *GetListingOffersParams)
        (*GetListingOffersOK, error)

	GetPricing(params *GetPricingParams)
        (*GetPricingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type ReportService interface {
	CancelReport(params *CancelReportParams)
        (*CancelReportOK, error)

	CancelReportSchedule(params *CancelReportScheduleParams)
        (*CancelReportScheduleOK, error)

	CreateReport(params *CreateReportParams)
        (*CreateReportAccepted, error)

	CreateReportSchedule(params *CreateReportScheduleParams)
        (*CreateReportScheduleCreated, error)

	GetReport(params *GetReportParams)
        (*GetReportOK, error)

	GetReportDocument(params *GetReportDocumentParams)
        (*GetReportDocumentOK, error)

	GetReportSchedule(params *GetReportScheduleParams)
        (*GetReportScheduleOK, error)

	GetReportSchedules(params *GetReportSchedulesParams)
        (*GetReportSchedulesOK, error)

	GetReports(params *GetReportsParams)
        (*GetReportsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type SaleService interface {
	GetOrderMetrics(params *GetOrderMetricsParams)
        (*GetOrderMetricsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type SellerService interface {
	GetMarketplaceParticipations(params *GetMarketplaceParticipationsParams)
        (*GetMarketplaceParticipationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type ShippingService interface {
	CancelShipment(params *CancelShipmentParams)
        (*CancelShipmentOK, error)

	CreateShipment(params *CreateShipmentParams)
        (*CreateShipmentOK, error)

	GetAccount(params *GetAccountParams)
        (*GetAccountOK, error)

	GetRates(params *GetRatesParams)
        (*GetRatesOK, error)

	GetShipment(params *GetShipmentParams)
        (*GetShipmentOK, error)

	GetTrackingInformation(params *GetTrackingInformationParams)
        (*GetTrackingInformationOK, error)

	PurchaseLabels(params *PurchaseLabelsParams)
        (*PurchaseLabelsOK, error)

	PurchaseShipment(params *PurchaseShipmentParams)
        (*PurchaseShipmentOK, error)

	RetrieveShippingLabel(params *RetrieveShippingLabelParams)
        (*RetrieveShippingLabelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type ServiceJobService interface {
	AddAppointmentForServiceJobByServiceJobID(
            params *AddAppointmentForServiceJobByServiceJobIDParams)
        (*AddAppointmentForServiceJobByServiceJobIDOK, error)

	CancelServiceJobByServiceJobID(params *CancelServiceJobByServiceJobIDParams)
        (*CancelServiceJobByServiceJobIDOK, error)

	CompleteServiceJobByServiceJobID(params *CompleteServiceJobByServiceJobIDParams)
        (*CompleteServiceJobByServiceJobIDOK, error)

	GetServiceJobByServiceJobID(params *GetServiceJobByServiceJobIDParams)
        (*GetServiceJobByServiceJobIDOK, error)

	GetServiceJobs(params *GetServiceJobsParams) (*GetServiceJobsOK, error)

	RescheduleAppointmentForServiceJobByServiceJobID(
            params *RescheduleAppointmentForServiceJobByServiceJobIDParams)
        (*RescheduleAppointmentForServiceJobByServiceJobIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type SmallAndLightService interface {
	DeleteSmallAndLightEnrollmentBySellerSKU(
            params *DeleteSmallAndLightEnrollmentBySellerSKUParams)
        (*DeleteSmallAndLightEnrollmentBySellerSKUNoContent, error)

	GetSmallAndLightEligibilityBySellerSKU(
            params *GetSmallAndLightEligibilityBySellerSKUParams)
        (*GetSmallAndLightEligibilityBySellerSKUOK, error)

	GetSmallAndLightEnrollmentBySellerSKU(params *GetSmallAndLightEnrollmentBySellerSKUParams)
        (*GetSmallAndLightEnrollmentBySellerSKUOK, error)

	GetSmallAndLightFeePreview(params *GetSmallAndLightFeePreviewParams)
        (*GetSmallAndLightFeePreviewOK, error)

	PutSmallAndLightEnrollmentBySellerSKU(params *PutSmallAndLightEnrollmentBySellerSKUParams)
        (*PutSmallAndLightEnrollmentBySellerSKUOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type SolicitationsService interface {
	CreateProductReviewAndSellerFeedbackSolicitation(
            params *CreateProductReviewAndSellerFeedbackSolicitationParams)
        (*CreateProductReviewAndSellerFeedbackSolicitationCreated, error)

	GetSolicitationActionsForOrder(params *GetSolicitationActionsForOrderParams)
        (*GetSolicitationActionsForOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

type UploadService interface {
	CreateUploadDestinationForResource(params *CreateUploadDestinationForResourceParams)
        (*CreateUploadDestinationForResourceCreated, error)

	SetTransport(transport runtime.ClientTransport)
}
