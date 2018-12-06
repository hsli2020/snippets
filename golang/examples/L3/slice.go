package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)
	
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//~ s[3] = "d"  
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))
	 //~ panic: runtime error: index out of range
	s = append(s, "d")
	s = append(s, "g", "f")
	fmt.Println("apd", s)
	
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c", c)
	
	t := []string{"g", "h", "i"}
	fmt.Println("x", t)
	
	twoD := make([][]int, 3)
	for i :=0; i < 3; i++ {
		inner := i + 1
		twoD[i] = make([]int, inner)
		for j := 0; j < inner; j ++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}