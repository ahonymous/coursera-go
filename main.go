package main

import (
	"coursera/pkg/module1"
	"coursera/pkg/module2"
	"fmt"
)

var moduleNumber int
var modules = [4]string{
	"1. Module 1: Hello, world!!!",
	"2. Module 2: Truncation enters a floating point number and prints the integer, truncated version of the floating point number that was entered.",
	"3. Module 2: Findian reads a string from the user, prints whether or not the string contains the characters a, i, and n, in order, with any number of characters between the three letters.",
	"0. Exit",
}

func main() {
	for {
		fmt.Println("What homework are you want to check?")
		fmt.Println("========================================")

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
		case 0:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Wrong module number! Try again.")
		}

		fmt.Println("----------------------------------------")
	}
}
