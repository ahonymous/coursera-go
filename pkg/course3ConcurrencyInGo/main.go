package course3ConcurrencyInGo

import (
	"fmt"

	"coursera/pkg/course3ConcurrencyInGo/module2"
)

var moduleNumber int
var modules = [2]string{
	"1. Race condition in goroutines",
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
			module2.Start()
		}

		fmt.Println("----------------------------------------")
	}
}
