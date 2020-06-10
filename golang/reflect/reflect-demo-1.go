// 1.获取基础信息

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello")
}

func Info(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Typeof:", t.Name())

	//获取字段信息
	v := reflect.ValueOf(i)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v \n", f.Name, f.Type, val)
	}
	//获取方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s :%v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{Id: 1, Name: "Evan", Age: 18}
	Info(u)
}

//Typeof: User
//Id : int = 1  Name : string = Evan   Age : int = 18
