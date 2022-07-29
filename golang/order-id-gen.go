package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateOrderNumber() string {
	now := time.Now()
	year := now.Year() - 2010
	ytod := now.YearDay()
	hms := now.Format("050415")

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100) + 1

	if year >= 10 {
		year = year - 10 + 'A'
	} else {
		year = year + '0'
	}

	orderID := fmt.Sprintf("%c%03d%s%02d", year, ytod, hms, random)

	return orderID
}

func main() {
	orderNumber := generateOrderNumber()
	fmt.Println(orderNumber)
}
