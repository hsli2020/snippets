package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/peterzernia/project/auth"
	"github.com/peterzernia/project/models"
	"github.com/peterzernia/project/utils"
)

func main() {
	db := utils.InitDB()
	db.AutoMigrate(&models.User{})
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api/v1")
	auth.InitializeRoutes(api.Group("/auth"))
	api.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Catch all routes for React-Router
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./client/build/index.html")
	})

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}

package utils

import (
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt hashes and salts a password
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compares a hash with a password
func ComparePasswords(hashedPwd string, plainPwd string) error {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err
}

// ValidatePassword validates password
func ValidatePassword(pwd string) error {
	if len(pwd) < 6 {
		return errors.New("Password must be at least 6 characters")
	}
	return nil
}

// ParseUserDBError parses the error returned from Postgres
// when a unique kep exists
func ParseUserDBError(err error) string {
	var message string

	if strings.HasSuffix(err.Error(), "username_key\"") {
		message = "A user with that username already exists"
	} else if strings.HasSuffix(err.Error(), "email_key\"") {
		message = "A user registered with that email already exists"
	} else {
		message = "Oops! Something went wrong"
	}

	return message
}

package utils

import (
	"crypto/rand"
)

// GenerateRandomBytes securely generates random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a base64 encoded securely generated random string.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
