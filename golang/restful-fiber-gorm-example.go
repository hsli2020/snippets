// How to Build REST API using Go Fiber and Gorm ORM 

// go get github.com/gofiber/fiber/v2
// go get gorm.io/gorm
// go get gorm.io/driver/mysql

// @/main.go
package main

import (
    "log"

    "github.com/FranciscoMendes10866/gorm/config"
    "github.com/FranciscoMendes10866/gorm/handlers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.Connect()

    app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)

    log.Fatal(app.Listen(":3000"))
}

// @/entities/dog.go
package entities

import "gorm.io/gorm"

type Dog struct {
    gorm.Model
    Name      string `json:"name"`
    Breed     string `json:"breed"`
    Age       int    `json:"age"`
    IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}

// @/config/database.go
package config

import (
    "github.com/FranciscoMendes10866/gorm/entities"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&entities.Dog{})

    return nil
}

// @/handlers/dog.go
package handlers

import (
    "github.com/FranciscoMendes10866/gorm/config"
    "github.com/FranciscoMendes10866/gorm/entities"
    "github.com/gofiber/fiber/v2"
)

func AddDog(c *fiber.Ctx) error {
    dog := new(entities.Dog)

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Create(&dog)
    return c.Status(201).JSON(dog)
}

func GetDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog entities.Dog

    result := config.Database.Find(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&dog)
}

func UpdateDog(c *fiber.Ctx) error {
    dog := new(entities.Dog)
    id := c.Params("id")

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Where("id = ?", id).Updates(&dog)
    return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog entities.Dog

    result := config.Database.Delete(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}
