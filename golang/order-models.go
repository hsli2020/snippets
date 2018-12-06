package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street,
	City,
	State,
	Zip string
}

type Customer struct {
	FirstName,
	LastName,
	Email,
	Phone string
	Addresses Address
}

type Order struct {
	Id int
	Customer
	PlacedOn   time.Time
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	Product
	Quantity int
}

type Product struct {
	SKU, // Code,
	Name,
	UPC,
	Condition,
	Description string
	UnitPrice float64
}
