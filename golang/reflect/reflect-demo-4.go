// 4.*通过反射动态调用方法

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

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", My Name is ", u.Name)
}
func main() {
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u) //通过valueOf获取u对象
	mv := v.MethodByName("Hello")

	//创建slice参数
	args := []reflect.Value{reflect.ValueOf("Evan")}
	mv.Call(args)
}
