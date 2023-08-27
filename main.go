package main

import (
	"coursera/pkg/module1"
	"coursera/pkg/module2"
	"coursera/pkg/module3"
	"coursera/pkg/module4"
	"fmt"
)

var moduleNumber int
var modules = [6]string{
	"1. Module 1: Hello, world!!!",
	"2. Module 2: Truncation enters a floating point number and prints the integer, truncated version of the floating point number that was entered.",
	"3. Module 2: Findian reads a string from the user, prints whether or not the string contains the characters a, i, and n, in order, with any number of characters between the three letters.",
	"4. Module 3: Slice reads integers from the user until the user enters the character X instead of an integer. Then prints the integers in the reverse order that they were read.",
	"5. Module 4: Makejson reads a name and an address from the user, stores them in a map, and then converts the map to JSON and prints it.",
	"0. Exit",
}

func main() {
	for {
		fmt.Println("What homework are you want to check? (any not numeric symbol to exit)")
		fmt.Println("========================================")

		for _, module := range modules {
			fmt.Println(module)
		}
		_, err := fmt.Scan(&moduleNumber)

		if nil != err || 0 == moduleNumber {
			fmt.Println("Exit...")

			return
		}

		switch moduleNumber {
		case 1:
			module1.HelloWorld()
		case 2:
			module2.Trunc()
		case 3:
			module2.Fundian()
		case 4:
			module3.Slice()
		case 5:
			module4.MakeJson()
		}

		fmt.Println("----------------------------------------")
	}
}
