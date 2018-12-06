http://www.oschina.net/question/1245337_2149527

GO new一个struct和反射生成struct疑问

package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

type Test struct {
    X int
    Y string
}

func main() {
    fmt.Println("hello world!")
    test1()
    fmt.Println("===========================")
    test2()
}

func test1() {
    a := Test{}
    fmt.Printf("a: %v %T \n", a, a)
    fmt.Println(a)
    err := json.Unmarshal([]byte(`{"X":1,"Y":"x"}`), &a)
    checkError(err)
    fmt.Printf("a: %v %T \n", a, a)
}

func test2() {
    m := make(map[string]reflect.Type)
    m["test"] = reflect.TypeOf(Test{})
    a := reflect.New(m["test"]).Elem().Interface()
    fmt.Printf("a: %v %T \n", a, a)
    fmt.Println(a)
    err := json.Unmarshal([]byte(`{"X":1,"Y":"x"}`), &a)
    checkError(err)
    fmt.Printf("a: %v %T \n", a, a)
}

结果如下：

a: {0 } main.Test
{0 }
a: {1 x} main.Test
===========================
a: {0 } main.Test
{0 }
a: map[X:1 Y:x] map[string]interface {}

求解为何这两种方式的结果会不同？

已在 stackoverflow 找到解答，详见：
http://stackoverflow.com/questions/34880504/golang-create-struct-in-different-way
