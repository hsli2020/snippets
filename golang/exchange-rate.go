package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ExchangeRate struct {
	XMLName xml.Name `xml:"Envelope"`
	Subject string   `xml:"subject"`

	Sender string `xml:"Sender>name"`

	Rates struct {
		Time string `xml:"time,attr"`
		Rate []struct {
			Currency string `xml:"currency,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"Cube"`
	} `xml:"Cube>Cube"`
}

var xmlresp = []byte(`<?xml version="1.0" encoding="UTF-8"?>
    <gesmes:Envelope
        xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01"
        xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
        <gesmes:subject>Reference rates</gesmes:subject>
        <gesmes:Sender>
            <gesmes:name>European Central Bank</gesmes:name>
        </gesmes:Sender>
        <Cube>
            <Cube time='2018-04-20'>
                <Cube currency='USD' rate='1.2309'/>
                <Cube currency='JPY' rate='132.41'/>
                <Cube currency='BGN' rate='1.9558'/>
                <Cube currency='CZK' rate='25.340'/>
                <Cube currency='DKK' rate='7.4477'/>
                <Cube currency='GBP' rate='0.87608'/>
                <Cube currency='HUF' rate='310.52'/>
                <Cube currency='PLN' rate='4.1677'/>
                <Cube currency='RON' rate='4.6586'/>
                <Cube currency='SEK' rate='10.3703'/>
                <Cube currency='CHF' rate='1.1970'/>
                <Cube currency='ISK' rate='123.30'/>
                <Cube currency='NOK' rate='9.6050'/>
                <Cube currency='HRK' rate='7.4110'/>
                <Cube currency='RUB' rate='75.7375'/>
                <Cube currency='TRY' rate='4.9803'/>
                <Cube currency='AUD' rate='1.5983'/>
                <Cube currency='BRL' rate='4.1892'/>
                <Cube currency='CAD' rate='1.5557'/>
                <Cube currency='CNY' rate='7.7449'/>
                <Cube currency='HKD' rate='9.6568'/>
                <Cube currency='IDR' rate='17142.74'/>
                <Cube currency='ILS' rate='4.3435'/>
                <Cube currency='INR' rate='81.3900'/>
                <Cube currency='KRW' rate='1316.26'/>
                <Cube currency='MXN' rate='22.7424'/>
                <Cube currency='MYR' rate='4.7924'/>
                <Cube currency='NZD' rate='1.7032'/>
                <Cube currency='PHP' rate='64.179'/>
                <Cube currency='SGD' rate='1.6172'/>
                <Cube currency='THB' rate='38.552'/>
                <Cube currency='ZAR' rate='14.8008'/>
            </Cube>
        </Cube>
    </gesmes:Envelope>
`)

func main() {
	var ex ExchangeRate

	//xml.Unmarshal(xmlresp, &ex)

	//fmt.Println(ex.Subject)
	//fmt.Println(ex.Sender)
	//fmt.Println(ex.Rates.Time)

	//for _, rate := range ex.Rates.Rate {
	//    fmt.Println(rate.Currency, rate.Rate)
	//}

	url := "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	xml.Unmarshal(body, &ex)

	fmt.Println(ex.Subject)
	fmt.Println(ex.Sender)
	fmt.Println(ex.Rates.Time)

	for _, rate := range ex.Rates.Rate {
		fmt.Println(rate.Currency, rate.Rate)
	}
}
