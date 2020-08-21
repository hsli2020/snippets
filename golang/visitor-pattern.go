package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

// 这段代码的目的就是想解耦 数据结构和 算法
// 使用 Strategy 模式也是可以完成的，而且会比较干净。
// 但是在有些情况下，多个Visitor是来访问一个数据结构的不同部分，
// 这种情况下，数据结构有点像一个数据库，而各个Visitor会成为一个个小应用。

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}

	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
