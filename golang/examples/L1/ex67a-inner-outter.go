package main

import "fmt"

type I interface {
	m1()
	m2()
	callParentM1()
	callParentM2()
}

type Inner struct {
}

func (s *Inner) callParentM1() {
  s.m1()
}

func (s *Inner) callParentM2() {
  s.m2()
}

func (s *Inner) m1() {
	fmt.Println("Inner m1 entered")
	s.m2()
}

func (s *Inner) m2() {
	fmt.Println("Inner m2 entered")
}

type Outter struct {
	I
}

func (s *Outter) m2() {
	fmt.Println("Outter m2 entered")
}

func main() {
	var i I = &Outter{&Inner{}}
	i.m1()
	i.m2()
	i.callParentM2()
}
