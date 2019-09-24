package main

import (
	"./service"
)

func main() {
	// Method 1: use ServiceContainer
	services := service.Get()
	services.Orders.Hello()
	services.Gifts.Hello()
	services.Users.Hello()
	services.Auth.Hello()

	// Method 2: use package
	service.Orders.Hello()
	service.Gifts.Hello()
	service.Users.Hello()
	service.Auth.Hello()
}
