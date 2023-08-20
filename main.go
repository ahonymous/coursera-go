package main

import (
	"coursera/pkg/module1"
	"coursera/pkg/module2"
	"fmt"
)

var modules []string
var moduleNumber int

func main() {

	// Add modules to the slice
	modules = append(modules, "1. Module 1: Hello, world!!!")
	modules = append(modules, "2. Module 2: Truncation enters a floating point number and prints the integer, truncated version of the floating point number that was entered.")
	modules = append(modules, "3. Module 2: Findian reads a string from the user, prints whether or not the string contains the characters a, i, and n, in order, with any number of characters between the three letters.")

	fmt.Println("What homework are you want to check?")

	for _, module := range modules {
		fmt.Println(module)
	}
	_, err := fmt.Scan(&moduleNumber)

	if err != nil {
		fmt.Println(err)

		return
	}

	switch moduleNumber {
	case 1:
		module1.HelloWorld()
	case 2:
		module2.Trunc()
	case 3:
		module2.Fundian()
	}
}
