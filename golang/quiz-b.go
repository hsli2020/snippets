package main

type T int

func (t T) Foo() T {
  print("foo")
  //println("foo")
  return t
}

func (t T) Bar() T {
  print("bar")
  //println("bar")
  return t
}

func main() {
  var t T
  defer t.Foo().Bar()
  t.Bar().Foo()
  //println("---")
}

// try to change print to println
// then add println("---") to end of main
// 
// foo
// bar
// foo
// ---
// bar
