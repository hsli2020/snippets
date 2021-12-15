package main

import "fmt"

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}

func executePanic() {
	defer recoveryFunction()	// Case #1
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely") // Never Go Here
}

func main() {
	// defer recoveryFunction()	// Case #2
	executePanic()
	fmt.Println("Main block is executed completely...")
}

/*
Case #1
This is Panic Situation
This is recovery function...
Main block is executed completely...

Case #2
This is Panic Situation
This is recovery function...
*/