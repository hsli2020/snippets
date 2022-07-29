package main

import "fmt"

func main() {
	// Declaration of variables
	var (
		day = 0
		month = 0
		zodiac = ""
	)

	// Input data
	fmt.Scan(&day)
	fmt.Scan(&month)

	// Conditions
	if  day >= 21 && month == 1 || day <= 18 && month == 2 {
		zodiac = "Acuario"
	} else if day >= 19 && month == 2 || day <= 20 && month == 3 {
		zodiac = "Piscis"
	} else if day >= 21 && month == 3 || day <= 20 && month == 4 {
		zodiac = "Aries"
	} else if day >= 21 && month == 4 || day <= 20 && month == 5 {
		zodiac = "Tauro"
	} else if day >= 21 && month == 5 || day <= 21 && month == 6 {
		zodiac = "Géminis"
	} else if day >= 22 && month == 6 || day <= 22 && month == 7 {
		zodiac = "Cáncer"
	} else if day >= 23 && month == 7 || day <= 22 && month == 8 {
		zodiac = "Leo"
	} else if day >= 23 && month == 8 || day <= 22 && month == 9 {
		zodiac = "Virgo"
	} else if day >= 23 && month == 9 || day <= 22 && month == 10 {
		zodiac = "Libra"
	} else if day >= 23 && month == 10 || day <= 22 && month == 11 {
		zodiac = "Escorpio"
	} else if day >= 23 && month == 11 || day <= 21 && month == 12 {
		zodiac = "Sagitario"
	} else { // day >= 22 && month == 12 || day <= 20 && month == 1
		zodiac = "Capricornio"
	}

	// Print result
	println(zodiac)
}