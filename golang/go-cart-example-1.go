package main // https://www.practical-go-lessons.com/chap-14-methods

import (
    "os/user"
    "time"

    "github.com/Rhymond/go-money"
)

type User struct {
    ID        string
    Firstname string
    Lastname  string
}

type Product struct {
    ID    string
    Name  string
    Price *money.Money
}

type Item struct {
    Product					// 嵌入
    Quantity uint8
}

type Cart struct {
    ID        	 string
    User					// 嵌入
    Items        []Item
    CurrencyCode string
    isLocked     bool
    CreatedAt 	 time.Time
    UpdatedAt 	 time.Time
    lockedAt  	 time.Time
}

func (c *Cart) TotalPrice() (*money.Money, error) {
    total := money.New(0, c.CurrencyCode)
    var err error
    for _, v := range c.Items {
        itemSubtotal := v.Product.Price.Multiply(int64(v.Quantity))
        total, err = total.Add(itemSubtotal)
        if err != nil {
            return nil, err
        }
    }
    return total, nil
}

func (c *Cart) Lock() error {
    if c.isLocked {
        return errors.New("cart is already locked")
    }
    c.isLocked = true
    c.lockedAt = time.Now()
    return nil
}
