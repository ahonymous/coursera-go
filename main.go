package main

import (
    "fmt"

    "coursera/pkg/course1StartingWithGo"
    "coursera/pkg/course2FunctionsMethodsInterfaceswithGo"
    "coursera/pkg/course3ConcurrencyInGo"
)

var courseNumber int
var courses = [4]string{
    "1. Course 1: Starting with Go",
    "2. Course 2: Functions, Methods, and Interfaces in Go",
    "3. Course 3: Concurrency in Go",
    "0. Exit",
}

func main() {
    for {
        fmt.Println("What course are you want to check? (any not numeric symbol to exit)")
        fmt.Println("========================================")

        for _, course := range courses {
            fmt.Println(course)
        }
        _, err := fmt.Scan(&courseNumber)

        if nil != err || 0 == courseNumber {
            fmt.Println("Exit...")

            return
        }

        switch courseNumber {
        case 1:
            course1StartingWithGo.Start()
        case 2:
            course2FunctionsMethodsInterfaceswithGo.Start()
        case 3:
            course3ConcurrencyInGo.Start()
        }

        fmt.Println("========================================")
    }
}
