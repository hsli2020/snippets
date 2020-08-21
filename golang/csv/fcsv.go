package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	filename := "items.csv"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", filename, err)
		return
	}
	defer file.Close()

	fcsv := csv.NewWriter(file)
	defer fcsv.Flush()

	// Write CSV header
	fcsv.Write([]string{"Name", "Price", "URL", "Image URL"})

    for (...) {
		fcsv.Write([]string{
            "Name",
            "Price",
            "URL",
            "Img",
		})
    }
}
