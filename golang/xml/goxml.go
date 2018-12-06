package main

import (
    "encoding/xml"
    "fmt"
    "os"
    "io"
)

func main() {

    type Staff struct {
        XMLName   xml.Name `xml:"staff"`
        ID        int      `xml:"id"`
        FirstName string   `xml:"firstname"`
        LastName  string   `xml:"lastname"`
        UserName  string   `xml:"username"`
    }

    type Company struct {
        XMLName xml.Name `xml:"company"`
        Staffs  []Staff  `xml:"staffs>staff"`
    }

    v := &Company{}

    // add two staff details
    v.Staffs = append(v.Staffs, Staff{ID: 103, FirstName: "Adam", LastName: "Ng", UserName: "adamng"})
    v.Staffs = append(v.Staffs, Staff{ID: 108, FirstName: "Jennifer", LastName: "Loh", UserName: "jenniferloh"})

    filename := "newstaffs.xml"
    file, _ := os.Create(filename)

    xmlWriter := io.Writer(file)

    enc := xml.NewEncoder(xmlWriter)
    enc.Indent("  ", "    ")
    if err := enc.Encode(v); err != nil {
        fmt.Printf("error: %v\n", err)
    }
}
