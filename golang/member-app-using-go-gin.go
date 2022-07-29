// https://github.com/ai285063/member-app-using-go-gin
// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	// router := gin.Default()

	ConnectMysql()
	ConnectRedis()

	g1 := router.Group("/users").Use(middleware)
	g1.GET("/", GetUserList)
	g1.POST("/register", Register)
	g1.PUT("/:id", PutUser)
	g1.DELETE("/:id", DeleteUser)

	g2 := router.Group("/viewcount")

	g2.GET("/", GetViewCount)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}

func middleware(c *gin.Context) {
	c.Next()
	AddViewCount()
}

// Mysql.go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	// SERVER   = "127.0.0.1"
	// docker-compose 裡面有自己的 dns，api 如果在 docker 裡面  不能用127.0.0.1
	SERVER   = "mysql"
	PORT     = 3306
	DATABASE = "practice"
)

func ConnectMysql() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("MySQL connection failed: " + err.Error())
	} else {
		log.Println("MySQL connected.")
	}

	if err := MysqlDB.AutoMigrate(&User{}); err != nil {
		panic("MySql create table failed: " + err.Error())
	}
}

// redis.go
package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Ctx = context.Background()

const ViewCount = "viewcount"

func ConnectRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		// Addr: "0.0.0.0:7001",
		// docker-compose 裡面有自己的 dns，api 如果在 docker 裡面  不能用127.0.0.1
		Addr: "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check if redis db is connected
	_, err := RedisDB.Ping(Ctx).Result()
	if err == nil {
		log.Println("Redis connected.")
	} else {
		panic("Redis connection failed: " + err.Error())
	}

	// set viewcount as initial 0
	err = RedisDB.Set(Ctx, ViewCount, 0, 0).Err()
	if err != nil {
		panic("Redis cannot initialize viewcount: " + err.Error())
	}
}

// controller.go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64
	Account  string
	Email    string
	Password string
}

func GetUserList(c *gin.Context) {
	var users []User
	if err := MysqlDB.Table("users").Find(&users).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "success",
			"Users":   users,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
}

func Register(c *gin.Context) {
	var user User
	user.Account = c.Request.FormValue("account")
	user.Password = c.Request.FormValue("password")
	user.Email = c.Request.FormValue("email")

	if err := MysqlDB.Table("users").Create(&user).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "Successfully registered.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}

func PutUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := findUser(int(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User doesn't exist",
		})
		return
	}

	account := c.Request.FormValue("account")
	password := c.Request.FormValue("password")
	email := c.Request.FormValue("email")

	if err := MysqlDB.Table("users").
		Where("ID = ?", id).
		Updates(User{
			Account:  account,
			Email:    email,
			Password: password,
		}).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "User data updated.",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := findUser(int(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User doesn't exist",
		})
		return
	}

	if err := MysqlDB.Table("users").Delete(User{}, id).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "User deleted.",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetViewCount(c *gin.Context) {
	val, err := RedisDB.Get(Ctx, ViewCount).Result()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"View count": val,
	})
}

func AddViewCount() {
	RedisDB.Incr(Ctx, ViewCount)
}

func findUser(id int) error {
	var user User
	err := MysqlDB.Table("users").Where("ID = ?", id).First(&user).Error
	return err
}
