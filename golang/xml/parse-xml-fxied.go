package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type datevalue struct {
	Date  customTime `xml:"date"`
	Value float32    `xml:"value"`
}

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02 15:04:05"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

type pv struct {
	XMLName    xml.Name  `xml:"series"`
	Unit       string    `xml:"unit"`
	Datevalues datevalue `xml:"values>dateValue"`
}

func main() {
	contents := `<series>
                   <timeUnit>DAY</timeUnit>
                   <unit>Wh</unit><measuredBy>INVERTER</measuredBy>
                   <values><dateValue>
                        <date>2015-11-04 00:00:00</date>
                        <value>5935.405</value>
                   </dateValue></values>
                </series>`

	m := &pv{}

	err := xml.Unmarshal([]byte(contents), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %f %s\n", m.Unit, m.Datevalues.Value, m.Datevalues.Date)
}
