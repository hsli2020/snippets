type GetOrderResponse struct {
	XMLName xml.Name `xml:"GetOrderResponse"`
	Xmlns   string   `xml:"xmlns,attr"`

	GetOrderResult struct {
		Orders struct {
			Order struct {
				LatestShipDate         string `xml:"LatestShipDate"`
				OrderType              string `xml:"OrderType"`
				PurchaseDate           string `xml:"PurchaseDate"`
				BuyerEmail             string `xml:"BuyerEmail"`
				AmazonOrderId          string `xml:"AmazonOrderId"`
				LastUpdateDate         string `xml:"LastUpdateDate"`
				IsReplacementOrder     string `xml:"IsReplacementOrder"`
				NumberOfItemsShipped   string `xml:"NumberOfItemsShipped"`
				ShipServiceLevel       string `xml:"ShipServiceLevel"`
				OrderStatus            string `xml:"OrderStatus"`
				SalesChannel           string `xml:"SalesChannel"`
				ShippedByAmazonTFM     string `xml:"ShippedByAmazonTFM"`
				IsBusinessOrder        string `xml:"IsBusinessOrder"`
				NumberOfItemsUnshipped string `xml:"NumberOfItemsUnshipped"`
				LatestDeliveryDate     string `xml:"LatestDeliveryDate"`

				PaymentMethodDetails   struct {
					PaymentMethodDetail string `xml:"PaymentMethodDetail"`
				} `xml:"PaymentMethodDetails"`

				IsGlobalExpressEnabled string `xml:"IsGlobalExpressEnabled"`
				IsSoldByAB             string `xml:"IsSoldByAB"`
				BuyerName              string `xml:"BuyerName"`
				EarliestDeliveryDate   string `xml:"EarliestDeliveryDate"`

				OrderTotal struct {
					Amount       string `xml:"Amount"`
					CurrencyCode string `xml:"CurrencyCode"`
				} `xml:"OrderTotal"`

				IsPremiumOrder     string `xml:"IsPremiumOrder"`
				EarliestShipDate   string `xml:"EarliestShipDate"`
				MarketplaceId      string `xml:"MarketplaceId"`
				FulfillmentChannel string `xml:"FulfillmentChannel"`
				PaymentMethod      string `xml:"PaymentMethod"`

				ShippingAddress struct {
					City                         string `xml:"City"`
					AddressType                  string `xml:"AddressType"`
					PostalCode                   string `xml:"PostalCode"`
					IsAddressSharingConfidential string `xml:"isAddressSharingConfidential"`
					StateOrRegion                string `xml:"StateOrRegion"`
					Phone                        string `xml:"Phone"`
					CountryCode                  string `xml:"CountryCode"`
					Name                         string `xml:"Name"`
					AddressLine1                 string `xml:"AddressLine1"`
					AddressLine2                 string `xml:"AddressLine2"`
					AddressLine3                 string `xml:"AddressLine3"`
				} `xml:"ShippingAddress"`

				IsPrime                      string `xml:"IsPrime"`
				ShipmentServiceLevelCategory string `xml:"ShipmentServiceLevelCategory"`
			} `xml:"Order"`
		} `xml:"Orders"`
	} `xml:"GetOrderResult"`
	ResponseMetadata struct {
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}
