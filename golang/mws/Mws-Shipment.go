package shipment

type Shipment struct {
	ShipmentId        string
	AmazonOrderId     string
	SellerOrderId     string
	ItemList          []Item
	ShipFromAddress   Address
	ShipToAddress     Address
	PackageDimensions PackageDimensions
	Weight            Weight
	Insurance         CurrencyAmount
	ShippingService   ShippingService
	Label             Label
	Status            string
	TrackingId        string
	CreatedDate       string
	LastUpdatedDate   string
}

type PackageDimensions struct {
	Length string
	Width  string
	Height string
	Unit   string

	PredefinedPackageDimensions string
}

type SellerInputDefinition struct {
	IsRequired          bool
	DataType            string
	Constraints         string
	InputDisplayText    string
	InputTarget         string
	StoredValue         string
	RestrictedSetValues string
}

type AdditionalInputs struct {
	AdditionalInputFieldName string
	SellerInputDefinition    SellerInputDefinition
}

type Address struct {
	Name                string
	AddressLine1        string
	AddressLine2        string
	AddressLine3        string
	DistrictOrCounty    string
	Email               string
	City                string
	StateOrProvinceCode string
	PostalCode          string
	CountryCode         string
	Phone               string
}

type FileContents struct {
	Contents string
	FileType string
	Checksum string
}

type CurrencyAmount struct {
	CurrencyCode string
	Amount       string
}

type Weight struct {
	Value string
	Unit  string
}

type TransparencyCodeList struct {
	TransparencyCode []string
}

const (
	Status_Purchased      = "Purchased"
	Status_RefundPending  = "RefundPending"
	Status_RefundRejected = "RefundRejected"
	Status_RefundApplied  = "RefundApplied"
)
