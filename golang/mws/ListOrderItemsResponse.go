type ListOrderItemsResponse struct {
	XMLName xml.Name `xml:"ListOrderItemsResponse"`
	Xmlns   string   `xml:"xmlns,attr"`

	ListOrderItemsResult struct {
		AmazonOrderId string `xml:"AmazonOrderId"`
		OrderItems struct {
			OrderItem []struct {
				QuantityOrdered   string `xml:"QuantityOrdered"`
				Title             string `xml:"Title"`

				PromotionDiscount struct {
					Amount       string `xml:"Amount"`
					CurrencyCode string `xml:"CurrencyCode"`
				} `xml:"PromotionDiscount"`

				ConditionId    string `xml:"ConditionId"`
				IsGift         string `xml:"IsGift"`
				ASIN           string `xml:"ASIN"`
				SellerSKU      string `xml:"SellerSKU"`
				ConditionNote  string `xml:"ConditionNote"`
				OrderItemId    string `xml:"OrderItemId"`
				IsTransparency string `xml:"IsTransparency"`

				ProductInfo    struct {
					NumberOfItems string `xml:"NumberOfItems"`
				} `xml:"ProductInfo"`

				QuantityShipped    string `xml:"QuantityShipped"`
				ConditionSubtypeId string `xml:"ConditionSubtypeId"`

				ItemPrice          struct {
					Amount       string `xml:"Amount"`
					CurrencyCode string `xml:"CurrencyCode"`
				} `xml:"ItemPrice"`

				ItemTax struct {
					Amount       string `xml:"Amount"`
					CurrencyCode string `xml:"CurrencyCode"`
				} `xml:"ItemTax"`

				PromotionDiscountTax struct {
					Amount       string `xml:"Amount"`
					CurrencyCode string `xml:"CurrencyCode"`
				} `xml:"PromotionDiscountTax"`
			} `xml:"OrderItem"`
		} `xml:"OrderItems"`
	} `xml:"ListOrderItemsResult"`
	ResponseMetadata struct {
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
} 
