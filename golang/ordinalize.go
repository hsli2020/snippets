package main

import (
	"fmt"
	"math"
	"strconv"
)

func Ordinal(num int) string {

	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}

	// math.Abs() is to convert negative number to positive
	floatNum := math.Abs(float64(num))
	positiveNum := int(floatNum)

	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return "th"
	}

	return ordinalDictionary[positiveNum]

}

func Ordinalize(num int) string {

	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}

	// math.Abs() is to convert negative number to positive
	floatNum := math.Abs(float64(num))
	positiveNum := int(floatNum)

	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return strconv.Itoa(num) + "th"
	}

	return strconv.Itoa(num) + ordinalDictionary[positiveNum]

}

func main() {
	// oridinaL tests
	fmt.Println("1 : ", Ordinal(1))
	fmt.Println("2 : ", Ordinal(2))
	fmt.Println("3 : ", Ordinal(3))
	fmt.Println("4 : ", Ordinal(4))
	fmt.Println("5 : ", Ordinal(5))
	fmt.Println("6 : ", Ordinal(6))
	fmt.Println("7 : ", Ordinal(7))
	fmt.Println("8 : ", Ordinal(8))
	fmt.Println("9 : ", Ordinal(9))
	fmt.Println("102 : ", Ordinal(102))
	fmt.Println("-99 : ", Ordinal(-99))
	fmt.Println("-1021 : ", Ordinal(-1021))

	// oridinaLIZE tests
	fmt.Println("1 : ", Ordinalize(1))
	fmt.Println("2 : ", Ordinalize(2))
	fmt.Println("3 : ", Ordinalize(3))
	fmt.Println("4 : ", Ordinalize(4))
	fmt.Println("5 : ", Ordinalize(5))
	fmt.Println("6 : ", Ordinalize(6))
	fmt.Println("7 : ", Ordinalize(7))
	fmt.Println("8 : ", Ordinalize(8))
	fmt.Println("9 : ", Ordinalize(9))
	fmt.Println("102 : ", Ordinalize(102))
	fmt.Println("-99 : ", Ordinalize(-99))
	fmt.Println("-1021 : ", Ordinalize(-1021))
}
