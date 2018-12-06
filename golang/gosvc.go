package main

import (
    "fmt"
)

type Application struct {
	skuService	 *SkuService
	orderService *OrderService
}

func NewApp() *Application {
	app := &Application{}

	service := Service{app: app}

	app.skuService   = &SkuService{Service: service}
	app.orderService = &OrderService{Service: service}

	return app;
}

type Service struct {
	// db *DB
	// redis *Redis
	app *Application
}

type SkuService struct {
	Service
}

func (s SkuService) FindSku(sku string) {
	fmt.Println(sku + " not found")
}

func (s SkuService) FindOrder(orderId string) {
	s.app.orderService.FindOrder(orderId)
}

type OrderService struct {
	Service
}

func (s OrderService) FindOrder(orderId string) {
	fmt.Println(orderId + " not found")
}

func (s OrderService) FindSku(sku string) {
	s.app.skuService.FindSku(sku)
}

func main() {
	app := NewApp()
	app.skuService.FindSku("SKU-11111")
	app.orderService.FindOrder("111-2222222-3333333")

	app.orderService.FindSku("SKU-22222")
	app.skuService.FindOrder("222-3333333-4444444")
}
