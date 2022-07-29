package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	filename := "order-%Y%M%D-%h%m%s.txt"

	convertedDate := convertDateFormat(filename)
	fmt.Println("Converted date:", convertedDate)
}

func convertDateFormat(filename string) string {
	// Get the current date
	currentDate := time.Now()

	// Create a replacer to replace multiple strings in one go
	replacer := strings.NewReplacer(
		"%Y", currentDate.Format("2006"),
		"%M", currentDate.Format("01"),
		"%D", currentDate.Format("02"),

		"%h", currentDate.Format("15"),
		"%m", currentDate.Format("04"),
		"%s", currentDate.Format("05"),
	)

	// Replace the date format specifiers in the filename with the current date
	formattedFilename := replacer.Replace(filename)

	return formattedFilename
}
