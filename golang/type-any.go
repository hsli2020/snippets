package main

import "fmt"

func main() {
	cells := []interface{}{"x", 124, true}

	for i, v := range cells {
		fmt.Printf("%d: %v ", i, v)

		switch v.(type) {
		case int:
			fmt.Printf("int: %d\n", v.(int))
		case string:
			fmt.Printf("string: %s\n", v.(string))
		case bool:
			fmt.Printf("bool: %t\n", v.(bool))
		}
	}
}
