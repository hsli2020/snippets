// 3.修改struct的值(通过反射修改对象的值)

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

func Set(i interface{}) {
	v := reflect.ValueOf(i)

	//ptr 指针类型
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		// 如果是指针，则获取其所指向的元素
		v = v.Elem()
	}

	//测试设置下Name
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("It's Not Valid")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("Evan")
	}
}

func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}
