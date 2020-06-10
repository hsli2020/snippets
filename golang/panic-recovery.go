package main

import "fmt"

func main() {
	x := 20
	y := 0
	printAllOperations(x, y)
	fmt.Println("Exiting main without any issues")
}

func printAllOperations(x int, y int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
			fmt.Println("Proceeding to alternative flow skipping division.")
			printOperationsSkipDivide(x, y)
		}
	}()

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf("sum=%v, subtract=%v, multiply=%v, divide=%v \n", sum, subtract, multiply, divide)
}

func printOperationsSkipDivide(x int, y int) {
	sum, subtract, multiply := x+y, x-y, x*y
	fmt.Printf("sum=%v, subtract=%v, multiply=%v \n", sum, subtract, multiply)
}
