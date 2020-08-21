package main

import "fmt"

type Marketplace struct {
	Name     string
	ID       string
	Country  string
	Endpoint string
	Region   string
}

var marketplaceList = []Marketplace{
	Marketplace{
		Name:     "Canada", // "加拿大",
		ID:       "A2EUQ1WTGCTBG2",
		Country:  "CA",
		Endpoint: "https://sellingpartnerapi-na.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-na.amazon.com",
		Region:   "us-east-1",
	},

	Marketplace{
		Name:     "United States", // "美国",
		ID:       "ATVPDKIKX0DER",
		Country:  "US",
		Endpoint: "https://sellingpartnerapi-na.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-na.amazon.com",
		Region:   "us-east-1",
	},

	Marketplace{
		Name:     "Mexico", // "墨西哥",
		ID:       "A1AM78C64UM0Y8",
		Country:  "MX",
		Endpoint: "https://sellingpartnerapi-na.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-na.amazon.com",
		Region:   "us-east-1",
	},

	Marketplace{
		Name:     "Brazil", // "巴西",
		ID:       "A2Q3Y263D00KWC",
		Country:  "BR",
		Endpoint: "https://sellingpartnerapi-na.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-na.amazon.com",
		Region:   "us-east-1",
	},

	Marketplace{
		Name:     "Spain", // "西班牙",
		ID:       "A1RKKUPIHCS9HS",
		Country:  "ES",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "United Kingdom", // "英国",
		ID:       "A1F83G8C2ARO7P",
		Country:  "GB",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "France", // "法国",
		ID:       "A13V1IB3VIYZZH",
		Country:  "FR",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Netherlands", // "荷兰",
		ID:       "A1805IZSGTT6HS",
		Country:  "NL",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Germany", // "德国",
		ID:       "A1PA6795UKMFR9",
		Country:  "DE",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Italy", // "意大利",
		ID:       "APJ6JRA9NG5V4",
		Country:  "IT",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Sweden",
		ID:       "A2NODRKZP88ZB9",
		Country:  "SE",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Poland",
		ID:       "A1C3SOZRARQ6R3",
		Country:  "PL",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Turkey", // "土耳其",
		ID:       "A33AVAJ2PDY3EV",
		Country:  "TR",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "United Arab Emirates", // "阿拉伯联合酋长国",
		ID:       "A2VIGQ35RCS4UG",
		Country:  "AE",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "India", // "印度",
		ID:       "A21TJRUUN4KGV",
		Country:  "IN",
		Endpoint: "https://sellingpartnerapi-eu.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
	},

	Marketplace{
		Name:     "Singapore", // "新加坡",
		ID:       "A19VAU5U5O7RUS",
		Country:  "SG",
		Endpoint: "https://sellingpartnerapi-fe.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-fe.amazon.com",
		Region:   "us-west-2",
	},

	Marketplace{
		Name:     "Austrilia", // "澳大利亚",
		ID:       "A39IBJ37TRP1C6",
		Country:  "AU",
		Endpoint: "https://sellingpartnerapi-fe.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-fe.amazon.com",
		Region:   "us-west-2",
	},

	Marketplace{
		Name:     "Japan", // "日本",
		ID:       "A1VC38T7YXB528",
		Country:  "JP",
		Endpoint: "https://sellingpartnerapi-fe.amazon.com",
		Sandbox:  "https://sandbox.sellingpartnerapi-fe.amazon.com",
		Region:   "us-west-2",
	},
}

var marketplaces map[string]Marketplace

func init() {
	marketplaces = make(map[string]Marketplace)

	for _, marketplace := range marketplaceList {
		id := marketplace.ID
		country := marketplace.Country

		marketplaces[id] = marketplace
		marketplaces[country] = marketplace
	}
}

func main() {
	fmt.Printf("%#v\n", marketplaces)
}
