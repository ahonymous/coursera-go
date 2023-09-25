package module2

import (
    "fmt"
    "strings"
)

func Fundian() {
    var input string
    fmt.Print("Enter a string: ")
    _, err := fmt.Scanln(&input)

    if err != nil {
        fmt.Println(err)
    }
    lowerInput := strings.ToLower(input)

    foundA := false
    foundI := false

    for _, char := range lowerInput {
        if char == 'a' && !foundA {
            foundA = true
        } else if char == 'i' && foundA && !foundI {
            foundI = true
        } else if char == 'n' && foundA && foundI {
            fmt.Println("Found 'a', 'i', and 'n' in order!")
            return
        }
    }

    fmt.Println("Did not find 'a', 'i', and 'n' in order.")
}
