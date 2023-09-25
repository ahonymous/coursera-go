package module1

import (
    "fmt"
    "strconv"
    "strings"
)

func Sort() {
    fmt.Println("Please enter 10 numbers: ")
    var numbers = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i := 0; i < 10; i++ {
        fmt.Print("Enter number: ")

        for {
            var rawNumber string
            _, err := fmt.Scan(&rawNumber)

            if nil == err {
                rawNumber = strings.TrimSpace(rawNumber)
                numbers[i], err = strconv.Atoi(rawNumber)
                if nil == err {
                    break
                }
            }

            fmt.Print("Enter valid integer number:")
        }
    }

    BubbleSort(numbers)
}

func BubbleSort(sli []int) {
    fmt.Println("Start bubble sort for slice: ", sli)
    fmt.Println("----------------------------------------")

    for i := 0; i < len(sli); i++ {
        for j := i + 1; j < len(sli); j++ {
            if sli[i] > sli[j] {
                Swap(&sli[i], &sli[j])
            }
        }
    }
    fmt.Println("End bubble sort for slice: ", sli)
}

func Swap(a *int, b *int) {
    *a, *b = *b, *a
}
