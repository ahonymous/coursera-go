package course2FunctionsMethodsInterfaceswithGo

import (
	"coursera/pkg/course2FunctionsMethodsInterfaceswithGo/module1"
	"coursera/pkg/course2FunctionsMethodsInterfaceswithGo/module2"
	"fmt"
)

var moduleNumber int
var modules = [3]string{
	"1. Bubble sort",
	"2. Speed calculation",
	"0. Back to main menu",
}

func Start() {
	for {
		fmt.Println("What homework are you want to check? (any not numeric symbol to exit)")
		fmt.Println("----------------------------------------")

		for _, module := range modules {
			fmt.Println(module)
		}
		_, err := fmt.Scan(&moduleNumber)

		if nil != err || 0 == moduleNumber {
			fmt.Println("Back...")

			return
		}

		switch moduleNumber {
		case 1:
			module1.Sort()
		case 2:
			module2.Speed()
		}

		fmt.Println("----------------------------------------")
	}
}
