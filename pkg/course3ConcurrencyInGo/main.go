package course3ConcurrencyInGo

import (
    "coursera/pkg/course2FunctionsMethodsInterfaceswithGo/module3"
    "coursera/pkg/course3ConcurrencyInGo/module4"
    "fmt"

    "coursera/pkg/course3ConcurrencyInGo/module2"
)

var moduleNumber int
var modules = [4]string{
    "1. Race condition in goroutines", // module2
    "2. Goroutine sort",               // module3
    "3. Dining philosophers",          // module4
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
        case 2:
            module3.Start()
        case 3:
            module4.Start()
        }

        fmt.Println("----------------------------------------")
    }
}
