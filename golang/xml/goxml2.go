package main

import (
    "encoding/xml"
    "fmt"
    "io"
    "os"
)

// if variable name starts with small cap
// for example : "tape" instead of "Tape"
// the final value will not appear in XML

type Thing struct {
    XMLName xml.Name `xml:"thing"`
    Tape    string   `xml:"tape"`
    // Tape string `xml:", innerxml"` -- keep it simple!
}

type Staff struct {
    XMLName   xml.Name `xml:"staff"`
    ID        int      `xml:"id"`
    FirstName string   `xml:"firstname"`
    LastName  string   `xml:"lastname"`
    UserName  string   `xml:"username"`
    TapeBrand Thing
}

type Company struct {
    XMLName xml.Name `xml:"company"`
    Staffs  StaffArray
}

type StaffArray struct {
    Staffs []Staff
}

func (s *StaffArray) AddStaff(sID int, sFName string, sLName string, sUName string, brandName string) {
    daThing := Thing{Tape: brandName}
    staffRecord := Staff{ID: sID, FirstName: sFName, LastName: sLName, UserName: sUName, TapeBrand: daThing}

    s.Staffs = append(s.Staffs, staffRecord)
}

func main() {

    v := Company{}

    // put a for loop here to add more data
    // this example will just add 2 rows of data.

    v.Staffs.AddStaff(103, "Adam", "Ng", "adamng", "scotch")
    v.Staffs.AddStaff(103, "Jennifer", "Loh", "jenniferloh", "sellotape")

    // sanity check - display on screen
    xmlString, err := xml.MarshalIndent(v, "", "    ")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%s \n", string(xmlString))

    // everything ok now, write to file.
    filename := "newstaffs2.xml"
    file, _ := os.Create(filename)

    xmlWriter := io.Writer(file)

    enc := xml.NewEncoder(xmlWriter)
    enc.Indent("  ", "    ")
    if err := enc.Encode(v); err != nil {
        fmt.Printf("error: %v\n", err)
    }
}
