package main

import "fmt"

type orderFacade struct {
    userService UserService
    productService ProductService
    paymentService PaymentService
    notificationService NotificationService
}

func (o *orderFacade) placeOrder(userId string, productId string) {
    fmt.Println("[Facade] Starting order placement")

    userValid := o.userService.isUserValid(userId)
    productAvailable := o.productService.productAvailable(productId)

    if userValid && productAvailable {
        o.productService.assignProductToUser(productId, userId)
        o.paymentService.makePayment(userId, productId)
        o.notificationService.notifyDealer(productId)
    }
}

type UserService struct {}

func (u *UserService) isUserValid(userId string) bool {
    fmt.Println("[UserService] validating the user: ", userId);
    return true    // Complex logic for checking validity
}

type ProductService struct {}

func (p *ProductService) productAvailable(productId string) bool {
    fmt.Println("[ProductService] checking availability of product: ", productId)
    return true    // Complex logic for checking availability
}

func (p *ProductService) assignProductToUser(productId string, userId string) {
    fmt.Printf("[ProductService] assigning product %s to user %s\n", productId, userId)
    // complex logic for product assignment
}

type PaymentService struct {}

func (p *PaymentService) makePayment(userId string, productId string) {
    fmt.Printf("[PaymentService] charging user %s for product %s\n", userId, productId)
    // complex logic for making payment
}

type NotificationService struct {}

func (n *NotificationService) notifyDealer(productId string) {
    fmt.Printf("[NotificationService] notifying dealer about sale of product %s\n", productId)
    // complex notification logic
}

func main() {
    orderModule := &orderFacade{
        userService:  UserService{},
        productService:  ProductService{},
        paymentService:  PaymentService{},
        notificationService:  NotificationService{},
    }

    userId := "test-user-id"
    productId := "test-product-id"

    orderModule.placeOrder(userId, productId)
}