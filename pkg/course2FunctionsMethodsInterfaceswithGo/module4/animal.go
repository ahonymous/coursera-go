package module4

import (
    "bufio"
    "fmt"
    "os"
)

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct{}

func (c Cow) Eat() {
    fmt.Println("grass")
}
func (c Cow) Move() {
    fmt.Println("walk")
}
func (c Cow) Speak() {
    fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
    fmt.Println("worms")
}
func (b Bird) Move() {
    fmt.Println("fly")
}
func (b Bird) Speak() {
    fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
    fmt.Println("mice")
}

func (s Snake) Move() {
    fmt.Println("slither")
}

func (s Snake) Speak() {
    fmt.Println("hsss")
}

func Start() {
    animals := make(map[string]Animal)

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Commands list: `newanimal` or `query`.")
    fmt.Println("Available animal types: `cow`, `bird`, `snake`.")
    fmt.Println("Available actions: `eat`, `move`, `speak`.")
    fmt.Println("Enter `X` or `x` to exit.")
    fmt.Println("----------------------------------------")
    fmt.Println("Enter a command (<command> <name> <animal-type or action>): ")
    fmt.Print("> ")
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println("text: ", text)
        if text == "X" || text == "x" {
            fmt.Println("Exit...")

            return
        }

        var command, name, animalType string
        _, err := fmt.Sscan(text, &command, &name, &animalType)

        if err != nil {
            fmt.Println("Cannot read command: ", err)
            continue
        }

        switch command {
        case "newanimal":
            var animal Animal
            switch animalType {
            case "cow":
                animal = Cow{}
            case "bird":
                animal = Bird{}
            case "snake":
                animal = Snake{}
            default:
                fmt.Println("Invalid animal type. Please choose 'cow', 'bird', or 'snake'.")
                continue
            }
            animals[name] = animal
            fmt.Println("Created it!")
        case "query":
            animal, exists := animals[name]
            if !exists {
                fmt.Println("Animal not found.")
                continue
            }
            switch animalType {
            case "eat":
                animal.Eat()
            case "move":
                animal.Move()
            case "speak":
                animal.Speak()
            default:
                fmt.Println("Invalid information requested. Please choose 'eat', 'move', or 'speak'.")
            }
        default:
            fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
        }
        fmt.Print("> ")
    }
}
