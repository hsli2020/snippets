/* 简明的数据库应用的组织方式
 *   example/main.go
 *   example/model/dao/init.go    => /database/ ?
 *   example/model/dao/user.go
 *   example/model/table/user.go  => /entity/ ?
 */

// ########## example/main.go
package main

import (
	"fmt"
	"time"

	"example/model/dao"
	"example/model/table"
)

func main() {
	users, _ := dao.GetAllUsers()
	for _, user := range users {
		fmt.Println(user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	}

	// User not exist
	err := dao.UpdateUsernameById("CHANGED", 1)
	if err != nil {
		fmt.Println(err)
	}

	// User not exist
	err = dao.DeleteUserById(2)
	if err != nil {
		fmt.Println(err)
	}
}

func seed() {
	user := table.User{
		//Id:        123,
		Username:  "USERNAME",
		Password:  "PASSWORD",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dao.CreateUser(&user)
}

// ########## example/model/dao/init.go
package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"example/model/table"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func init() {
	_DB = initDB()
	_DB.AutoMigrate(&table.User{})
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_web?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
}

// ########## example/model/dao/user.go
package dao

import "example/model/table"

func CreateUser(user *table.User) (err error) {
	err = DB().Create(user).Error
	return
}

func GetUserById(userId int64) (user *table.User, err error) {
	user = new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error
	return
}

func GetAllUsers() (users []*table.User, err error) {
	err = DB().Find(&users).Error
	return
}

func UpdateUsernameById(username string, userId int64) (err error) {
	user := new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error
	if err != nil {
		return
	}

	user.Username = username
	err = DB().Save(user).Error
	return
}

func DeleteUserById(userId int64) (err error) {
	user := new(table.User)
	err = DB().Where("id = ?", userId).First(user).Error
	if err != nil {
		return
	}

	err = DB().Delete(user).Error
	return
}

// ########## example/model/table/user.go
package table // entity ?

import "time"

type User struct {
	Id        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (m *User) TableName() string {
	return "users"
}
