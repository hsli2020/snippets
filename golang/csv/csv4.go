package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	csvfile, err := os.Open("somecsvfile.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	}
}
