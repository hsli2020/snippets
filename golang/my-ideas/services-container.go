package main // test only

// myapp/service/services.go
//package service

// put all services together (ServiceBucket/ServiceContiner)
type serviceContainer struct {
	Orders *OrderService
	Gifts  *GiftService
	Users  *UserService
	Auth   *AuthService
}

var container *serviceContainer

func initContainer() {
	// repos, db ?
	container = &serviceContainer{
		Orders: &OrderService{},
		Gifts:  &GiftService{},
		Users:  &UserService{},
		Auth:   &AuthService{},
	}
}

func Get() *serviceContainer {
	if container == nil {
		initContainer()
	}
	return container
}

// this is how to visit services
func main() {
	//services := service.Get()
	services := Get()
	services.Orders.Hello()
	services.Gifts.Hello()
	services.Users.Hello()
	services.Auth.Hello()
}

// define services in seperated files

// myapp/service/orders.go
type OrderService struct{}
func (OrderService) Hello() { println("OrderService") }

// myapp/service/gifts.go
type GiftService struct{}
func (GiftService) Hello() { println("GiftService") }

// myapp/service/users.go
type UserService struct{}
func (UserService) Hello() { println("UserService") }

// myapp/service/auth.go
type AuthService struct{}
func (AuthService) Hello() { println("AuthService") }
