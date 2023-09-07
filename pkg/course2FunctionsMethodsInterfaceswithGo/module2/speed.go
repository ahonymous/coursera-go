package module2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
}

func Speed() {
	var a, vo, so, t float64

	accelerationValue(&a)
	velocityValue(&vo)
	displacementValue(&so)

	fn := GenDisplaceFn(a, vo, so)

	timeValue(&t)

	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, displacement)
}

func inputValue(message string, value *float64) {
	fmt.Print(message)
	for {
		var rawValue string
		_, err := fmt.Scan(&rawValue)
		if err == nil {
			rawValue = strings.TrimSpace(rawValue)
			*value, err = strconv.ParseFloat(rawValue, 64)
			if err == nil {
				break
			}
		}
		fmt.Println("Please, enter a valid value and try again.")
	}
}

func accelerationValue(a *float64) {
	inputValue("Enter acceleration (a): ", a)
}

func velocityValue(vo *float64) {
	inputValue("Enter initial velocity (vo): ", vo)
}

func displacementValue(so *float64) {
	inputValue("Enter initial displacement (so): ", so)
}

func timeValue(t *float64) {
	inputValue("Enter time (t): ", t)
}
