https://github.com/MrBessrour/golang-CRUD-API
// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//making an instance of the type DB from the gorm package
var db *gorm.DB = nil
var err error

func main() {
	//establishing connection with mysql database 'CRUD'
	dsn := "root:@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}
	//database migration if not exist or if there is any modification made in the model 'Post'
	db.AutoMigrate(&Post{})

	server := gin.Default()

	//set up the different routes
	server.GET("/posts", Posts)
	server.GET("/posts/:id", Show)
	server.POST("/posts", Store)
	server.PATCH("/posts/:id", Update)
	server.DELETE("/posts/:id", Delete)

	//start the server and listen on the port 8000
	server.Run(":8000")
}

// type.go
package main

import (
	"gorm.io/gorm"
)

//this model represent a database table
type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status"`
}

// api.go
package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//select query with limit and offset
func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    posts,
	})
}

//showing a post with it's id passed in the URL with a GET request
func Show(c *gin.Context) {
	post := getById(c)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    post,
	})
}

//storing a mew post to the db with a POST request with
func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	post.Status = "Active"
	db.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    post,
	})
}

//delete a post by it's id
func Delete(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		return
	}
	db.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"messege": "deleted successfuly",
		"data":    "",
	})
}

//update a post with a Ptach request , the id sent in the URL
func Update(c *gin.Context) {
	oldpost := getById(c)
	var newpost Post
	if err := c.ShouldBindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	oldpost.Title = newpost.Title
	oldpost.Des = newpost.Des
	if newpost.Status != "" {
		oldpost.Status = newpost.Status
	}

	db.Save(&oldpost)

	c.JSON(http.StatusOK, gin.H{
		"messege": "Post has been updated",
		"data":    oldpost,
	})
}

func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"messege": "post not found",
			"data":    "",
		})
	}
	return post
}