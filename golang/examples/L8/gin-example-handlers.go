package auth

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes routes for the App
func InitializeRoutes(r *gin.RouterGroup) {
	r.POST("/register", handleRegistration)
	r.POST("/login", handleLogin)
	r.POST("/logout", handleLogout)
	r.PUT("/user", handleUpdateUser)
	r.GET("/user", handleGetUser)
	r.POST("/password/change", handlePasswordChange)
}

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peterzernia/project/models"
	"github.com/peterzernia/project/utils"
)

func handleOptions(c *gin.Context) {
	c.Header("Allow", "POST, PUT, GET, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
}

func handleRegistration(c *gin.Context) {
	var auth models.Auth
	c.ShouldBindJSON(&auth)
	db := utils.GetDB()

	if auth.Password1 != auth.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords do not match",
		})
		return
	}

	err := utils.ValidatePassword(auth.Password1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	password := utils.HashAndSalt([]byte(auth.Password1))
	token, _ := utils.GenerateRandomString(32)

	user := models.User{
		Email:    auth.Email,
		Username: auth.Username,
		Password: password,
		Token:    token,
	}

	if err := db.Create(&user).Error; err != nil {
		message := utils.ParseUserDBError(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"email":      user.Email,
		"token":      user.Token,
		"username":   user.Username,
	})
}

func handleLogin(c *gin.Context) {
	var auth models.Auth
	var user models.User
	c.ShouldBindJSON(&auth)
	db := utils.GetDB()

	db.Where("username = ?", auth.Username).First(&user)

	err := utils.ComparePasswords(user.Password, auth.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	if user.Token == "" {
		token, _ := utils.GenerateRandomString(32)
		user.Token = token
		db.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"email":      user.Email,
		"token":      user.Token,
		"username":   user.Username,
	})
}

func handleLogout(c *gin.Context) {
	var user models.User
	db := utils.GetDB()

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	db.Where("token = ?", token).First(&user)

	if user.Username == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user.Token = ""
	db.Save(&user)
	c.Status(http.StatusOK)
	return
}

func handlePasswordChange(c *gin.Context) {
	var auth models.Auth
	var user models.User
	c.ShouldBindJSON(&auth)
	db := utils.GetDB()

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	db.Where("token = ?", token).First(&user)

	err := utils.ComparePasswords(user.Password, auth.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	if auth.Password1 != auth.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords do not match",
		})
		return
	}

	err = utils.ValidatePassword(auth.Password1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	password := utils.HashAndSalt([]byte(auth.Password1))
	user.Password = password
	db.Save(&user)

	c.Status(http.StatusOK)
}

func handleUpdateUser(c *gin.Context) {
	var auth models.Auth
	var user models.User
	c.ShouldBindJSON(&auth)
	db := utils.GetDB()

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	db.Where("token = ?", token).First(&user)

	if auth.Username != "" {
		user.Username = auth.Username
	}
	if auth.Email != "" {
		user.Email = auth.Email
	}

	if err := db.Save(&user).Error; err != nil {
		message := utils.ParseUserDBError(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"email":      user.Email,
		"token":      user.Token,
		"username":   user.Username,
	})
}

func handleGetUser(c *gin.Context) {
	var user models.User
	db := utils.GetDB()

	token := c.GetHeader("Authorization")

	err := db.Where("token = ?", token).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invailid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"email":      user.Email,
		"token":      user.Token,
		"username":   user.Username,
	})
}
