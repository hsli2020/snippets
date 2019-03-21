package main

import "fmt"

// map & nil

func show(usr map[string]string) {
	//if usr == nil { return } // not necessary

	for k, v := range usr { // for nil, no output
		fmt.Println(k, v)
	}

	fmt.Println("fname:", usr["fname"])
	fmt.Println("lname:", usr["lname"])
}

func main() {
	show(nil) // no error

	usr := map[string]string{
		"fname": "Some",
		"lname": "Wang",
	}
	show(usr)
}
