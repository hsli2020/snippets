package main

import (
    "encoding/xml"
    "fmt"
)

type datevalue struct {
    Date  int     `xml:"date"`
    Value float32 `xml:"value"`
}

type pv struct {
    XMLName    xml.Name  `xml:"series"`
    Unit       string    `xml:"unit"`
    datevalues datevalue `xml:"values>dateValue"`
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
    xml.Unmarshal([]byte(contents), &m)
    fmt.Printf("%s %f %d\n", m.Unit, m.datevalues.Value, m.datevalues.Date)
}

// https://stackoverflow.com/questions/33557401/unmarshal-nested-xml-with-go