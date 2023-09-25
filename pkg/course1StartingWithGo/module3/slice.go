package module3

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func Slice() {
    intSlice := make([]int, 3)

    for {
        fmt.Print("Enter an integer (X to quit): ")
        var input string
        _, err := fmt.Scanln(&input)

        if err != nil {
            fmt.Println("Something went wrong:", err)

            break
        }

        input = strings.ToLower(input)
        if "x" == input {
            fmt.Println("Exit...")

            break
        }

        number, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Please, enter a valid integer number and try again.")

            continue
        }

        intSlice = append(intSlice, number)
        sort.Ints(intSlice)
        fmt.Println("The sorted slice is:", intSlice)
    }
}
