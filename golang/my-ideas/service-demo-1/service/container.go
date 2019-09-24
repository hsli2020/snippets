package service

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
