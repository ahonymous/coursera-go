package final

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

var names []Name

func ReadNames() {
	fmt.Println("Final: Reads a file containing a list of names, and then prints the content of the file to the screen.")
	fmt.Println("----------------------------------------")

	file := getFileWithNames()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Close file with error: ", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("Unable to read file: ", err)
	}

	for scanner.Scan() {
		nameParts := strings.Split(scanner.Text(), " ")
		name := Name{
			fname: nameParts[0],
			lname: nameParts[1],
		}
		names = append(names, name)
	}

	for _, name := range names {
		fmt.Printf("Firstname is \"%s\" and lastname is \"%s\".\n", name.fname, name.lname)
	}
}

func getFileWithNames() *os.File {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome! Please enter the path to text file with names: ")
		scanner.Scan()
		filePath := scanner.Text()

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Cannot open file: "+filePath, err)
			continue
		}

		return file
	}
}
