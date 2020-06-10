// 2.获取匿名信息

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

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "OK", 12}, title: "private_title"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	//获取到的是User,相对于Manager来说是匿名的，Anonymous:true
	//reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10b4580), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	//获取到的是User里面的Id字段，相对于User来说是非匿名的，Anonymous:false
	//reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x10a5000), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
}
