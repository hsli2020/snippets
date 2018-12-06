package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	lats := make([]string, len(lines)-1)
	lngs := make([]string, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}
		lats[i-1] = line[0]
		lngs[i-1] = line[1]
	}

	for i, _ := range lats {
		fmt.Println(lats[i], lngs[i])
	}
}

var csvFile = strings.NewReader(`LatitudeMeasure,LongitudeMeasure,GeographicalLocation_Id,GeoTemporalMeasurement_Id
32.53,-116.99,0,0
32.53,-116.99,1,1
32.53,-116.99,2,2
32.53,-116.99,3,3
32.53,-116.99,4,4
32.53,-116.99,5,5
32.53,-116.99,6,6
32.53,-116.99,7,7
32.53,-116.99,8,8
32.53,-116.99,9,9
32.53,-116.99,10,10
32.53,-116.99,11,11
32.53,-116.99,12,12
32.53,-116.99,13,13
32.53,-116.99,14,14
32.53,-116.99,15,15
32.53,-116.99,16,16
32.53,-116.99,17,17
32.53,-116.99,18,18
32.53,-116.99,19,19`)
