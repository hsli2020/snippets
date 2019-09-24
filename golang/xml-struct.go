package main

import (
	"encoding/xml"
	"fmt"
)

type OrderItem struct {
	LineNumber int    `xml:"lineNumber,attr"`
	SKU        string `xml:"SKU"`
	UnitPrice  string `xml:"UnitPrice"`
	OrderQty   string `xml:"OrderQuantity"`
}

type OrderRequest struct {
	XMLName xml.Name `xml:"SynnexB2B"`

	//Credential struct {
	//	UserID   string `xml:"UserID"`
	//	Password string `xml:"Password"`
	//} `xml:"Credential"`
	UserID   string `xml:"Credential>UserID"`
	Password string `xml:"Credential>Password"`

	Detail struct {
		Shipment struct {
			ShipFrom string `xml:"ShipFrom"`

			ShipTo struct {
				City    string `xml:"City"`
				State   string `xml:"State"`
				ZipCode string `xml:"ZipCode"`
				Country string `xml:"Country"`
			} `xml:"ShipTo"`

			//ShipMethod struct {
			//	Code string `xml:"Code"`
			//} `xml:"ShipMethod"`
			ShipMethod string `xml:"ShipMethod>Code"`
		} `xml:"Shipment"`

		//Items struct {
		//	Item []OrderItem `xml:"Item"`
		//} `xml:"Items"`
		Items []OrderItem `xml:"Items>Item"`
	} `xml:"OrderRequest"`
}

func main() {
	var x OrderRequest

	//x.Credential.UserID = "Username"
	//x.Credential.Password = "Password"
	x.UserID = "Username"
	x.Password = "Password"

	x.Detail.Shipment.ShipFrom = "Warehouse"
	x.Detail.Shipment.ShipTo.City = "City"
	x.Detail.Shipment.ShipTo.State = "State"
	x.Detail.Shipment.ShipTo.ZipCode = "ZipCode"
	x.Detail.Shipment.ShipTo.Country = "Country"

	//x.Detail.Shipment.ShipMethod.Code = "ShipMehod"
	x.Detail.Shipment.ShipMethod = "ShipMehod"

	items := make([]OrderItem, 2)

	items[0].LineNumber = 1
	items[0].SKU = "SKU-111"
	items[0].UnitPrice = "111.11"
	items[0].OrderQty = "11"

	items[1].LineNumber = 2
	items[1].SKU = "SKU-222"
	items[1].UnitPrice = "222.22"
	items[1].OrderQty = "22"

	//x.Detail.Items.Item = items
	x.Detail.Items = items

	output, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	request := xml.Header + string(output)
	fmt.Println(request)
}
