package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Printf("eats %s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("makes %s sound\n", a.noise)
}

var Database = make(map[string]Animal, 3)

func main() {
	InitDb()

	fmt.Printf("Enter animal name and action.\nTo quit application enter Q\n")
	for {
		userInput := getInputFromUser(">")
		if userInput == "q" {
			fmt.Println("Bye")
			break
		}

		request := parseUserInput(userInput)
		if request == nil {
			fmt.Printf("Error parsing user input\n")
			continue
		}

		animal, ok := Database[request["name"]]
		if !ok {
			fmt.Printf("Animal %s not found in a database\n", request["name"])
			continue
		}

		fmt.Printf("%s ", strings.Title(request["name"]))
		switch request["info"] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("does nothing\n")
		}
	}
}

func InitDb() {
	Database["cow"] = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	Database["bird"] = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	Database["snake"] = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func parseUserInput(input string) map[string]string {
	result := make(map[string]string, 1)

	splitInput := strings.Fields(input)
	if len(splitInput) != 2 {
		return nil
	}
	result["name"] = splitInput[0]
	result["info"] = splitInput[1]

	return result
}

func getInputFromUser(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s ", prompt)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while getting user input: %v\n", err)
	}

	userInput := scanner.Text()
	userInput = strings.ToLower(userInput)

	return userInput
}
