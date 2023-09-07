package module3

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func Start() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print("Enter an animal {cow, bird, snake} and an action {eat, move, speak} (e.g., 'cow eat'): ")
		var animalName, action string
		_, err := fmt.Scan(&animalName, &action)

		if nil != err {
			fmt.Println("Invalid input. Please try again.", err)
			continue
		}

		var animal Animal

		switch animalName {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("Invalid animal name. Please choose 'cow', 'bird', or 'snake'.")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid action. Please choose 'eat', 'move', or 'speak'.")
		}
	}
}
