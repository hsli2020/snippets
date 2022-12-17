/**
 * EVM From Scratch
 * Go template
 *
 * To work on EVM From Scratch in Go:
 *
 * - Install Golang: https://golang.org/doc/install
 * - Go to the `go` directory: `cd go`
 * - Edit `evm.go` (this file!), see TODO below
 * - Run `go run evm.go` to run the tests
 */

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
)

type code struct {
	Bin string
	Asm string
}

type expect struct {
	Stack   []string
	Success bool
	Return  string
}

type TestCase struct {
	Name   string
	Hint   string
	Code   code
	Expect expect
}

func evm(code []byte) (bool, []big.Int) {
	var stack []big.Int
	pc := 0

	for pc < len(code) {
		op := code[pc]
		pc++
		
		if op == 0x00 {
		}
		// TODO: Implement the EVM here!
	}


	return true, stack
}

func main() {
	content, err := ioutil.ReadFile("../evm.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []TestCase
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json.Unmarshal(): ", err)
	}

	for index, test := range payload {
		fmt.Printf("Test #%v of %v: %v\n", index + 1, len(payload), test.Name)

		bin, err := hex.DecodeString(test.Code.Bin)
		if err != nil {
			log.Fatal("Error during hex.DecodeString(): ", err)
		}

		var expectedStack []big.Int
		for _, s := range test.Expect.Stack {
			i, ok := new(big.Int).SetString(s, 0)
			if !ok {
				log.Fatal("Error during big.Int.SetString(): ", err)
			}
			expectedStack = append(expectedStack, *i)
		}

		// Note: as the test cases get more complex, you'll need to modify this
		// to pass down more arguments to the evm function and return more than
		// just the stack.
		success, stack := evm(bin)

		match := len(stack) == len(expectedStack)
		if match {
			for i, s := range stack {
				match = match && (s.Cmp(&expectedStack[i]) == 0)
			}
		}
		match = match && (success == test.Expect.Success)

		if !match {
			fmt.Printf("Instructions: \n%v\n", test.Code.Asm)
			fmt.Printf("Expected: success=%v, stack=%v\n", test.Expect.Success, toStrings(expectedStack))
			fmt.Printf("Got:      success=%v, stack=%v\n\n", success, toStrings(stack))
			fmt.Printf("Hint: %v\n\n", test.Hint)
			fmt.Printf("Progress: %v/%v\n\n", index, len(payload))
			log.Fatal("Stack mismatch")
		}
	}
}

func toStrings(stack []big.Int) []string {
	var strings []string
	for _, s := range stack {
		strings = append(strings, s.String())
	}
	return strings
}
