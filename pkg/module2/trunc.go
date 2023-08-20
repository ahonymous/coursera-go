package module2

import "fmt"

func Trunc() {
	fmt.Println("TRUNCATION FLOAT NUMBERS")
	truncatedNumber := getNumber()
	fmt.Println(fmt.Sprintf("Truncated numbrer is: %d", truncatedNumber))
}

func getNumber() int64 {
	var number float64
	fmt.Println("Enter any float number:")
	_, err := fmt.Scanln(&number)

	if err != nil {
		fmt.Println(err)
	}

	return int64(number)
}
