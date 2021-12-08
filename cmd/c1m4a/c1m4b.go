package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func getUserInput(prompt string) string {
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s: ", prompt)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while getting user input: %v\n", err)
	}
	userInput = scanner.Text()

	return userInput
}

func main() {
	contact := make(map[string]string)

	contact["name"] = getUserInput("Enter name")
	contact["address"] = getUserInput("Enter address")

	contactJSON, err := json.Marshal(contact)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error converting to JSON: %v\n", err)
		return
	}

	fmt.Printf("JSON object:\n%v\n", string(contactJSON))
}
