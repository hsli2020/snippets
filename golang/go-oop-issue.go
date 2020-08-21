package main

import "fmt"
import "math"

type ShapeInterface interface {
	Area() float64
	GetName() string
	PrintArea()
}

// 标准形状，它的面积为0.0
type Shape struct {
	name string
}

func (s *Shape) Area() float64 {
	return 0.0
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", s.name, s.Area())
}

// 矩形 : 重新定义了Area方法
type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h
}

// 圆形  : 重新定义 Area 和PrintArea 方法
type Circle struct {
	Shape
	r float64
}

func (c *Circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func (c *Circle) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", c.GetName(), c.Area())
}

func main() {
	s := Shape{name: "Shape"}
	c := Circle{Shape: Shape{name: "Circle"}, r: 10}
	r := Rectangle{Shape: Shape{name: "Rectangle"}, w: 5, h: 4}

	listshape := []ShapeInterface{&s, &c, &r}

	for _, si := range listshape {
		si.PrintArea() //!! 猜猜哪个Area()方法会被调用 !!
	}

	// Go的OOP与其它语言的OOP的不同之处

	// 对于Rectangle类型来说si.PrintArea()将调用Shape.PrintArea()
	// 因为没有为Rectangle类型定义PrintArea()方法（没有接受者是*Rectangle的PrintArea()方法），
	// 而Shape.PrintArea()方法的实现调用的是Shape.Area()
	// 而不是Rectangle.Area()-如前面所讨论的，Shape不知道Rectangle的存在。

	// r.PrintArea() // This is what goes wrong

	fmt.Println("\nFIX\n")
	for _, si := range listshape {
		PrintArea(si)
	}
}

// This is how to fix
func PrintArea(s ShapeInterface) {
	fmt.Printf("%s : Area %v\r\n", s.GetName(), s.Area())
}

// output:
// Shape : Area 0
// Circle : Area 314.1592653589793
// Rectangle : Area 0                 <==== ???
