package main

 import (
     "encoding/csv"
     "fmt"
     "os"
 )

 type TestRecord struct {
     Email string
     Date  string
 }

 func main() {

     csvfile, err := os.Open("somecsvfile.csv")

     if err != nil {
             fmt.Println(err)
             os.Exit(1)
     }

     defer csvfile.Close()

     reader := csv.NewReader(csvfile)

     reader.FieldsPerRecord = -1

     rawCSVdata, err := reader.ReadAll()

     if err != nil {
         fmt.Println(err)
         os.Exit(1)
     }

     // sanity check, display to standard output
     for _, each := range rawCSVdata {
         fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
     }

     // now, safe to move raw CSV data to struct

     var oneRecord TestRecord

     var allRecords []TestRecord

     for _, each := range rawCSVdata {
         oneRecord.Email = each[0]
         oneRecord.Date = each[1]
         allRecords = append(allRecords, oneRecord)
     }

     // second sanity check, dump out allRecords and see if 
     // individual record can be accessible
     fmt.Println(allRecords)
     fmt.Println(allRecords[2].Email)
     fmt.Println(allRecords[2].Date)
}