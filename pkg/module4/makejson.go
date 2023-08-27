package module4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func MakeJson() {
	data := loadData()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Cannot marshal data for entered values", data)

		return
	}

	fmt.Println(string(jsonData))
}

func loadData() map[string]string {
	reader := bufio.NewReader(os.Stdin)
	mapData := make(map[string]string)

	for {
		fmt.Print("Enter a name: ")
		name, err := reader.ReadString('\n')
		if err == nil && name != "" && name != "\n" {
			mapData["name"] = strings.TrimSuffix(name, "\n")
			break
		}
		fmt.Println("Please, enter a valid name and try again.")
	}
	for {
		fmt.Print("Enter an address: ")
		address, err := reader.ReadString('\n')

		if err == nil && address != "" && address != "\n" {
			mapData["address"] = strings.TrimSuffix(address, "\n")

			break
		}
		fmt.Println("Please, enter a valid address and try again.")
	}

	return mapData
}
