package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

func (a Cow) Eat()   { fmt.Println("eats grass") }
func (a Cow) Move()  { fmt.Println("walks") }
func (a Cow) Speak() { fmt.Println("moos") }

func (a Bird) Eat()   { fmt.Println("eats worms") }
func (a Bird) Move()  { fmt.Println("flies") }
func (a Bird) Speak() { fmt.Println("peeps") }

func (a Snake) Eat()   { fmt.Println("eats mice") }
func (a Snake) Move()  { fmt.Println("slithers") }
func (a Snake) Speak() { fmt.Println("hisses") }

var Database = make(map[string]Animal, 3)

type UserCommand struct {
	command string
	name    string
	field   string
}

func main() {
	printInstructions()

	for {
		userInput := getInputFromUser(">")
		if userInput == "q" {
			fmt.Println("Bye")
			break
		}

		request, err := parseUserInput(userInput)
		if err != nil {
			fmt.Println("Error parsing user input")
			continue
		}

		switch request.command {
		case "newanimal":
			createAnimal(request.name, request.field)
		case "query":
			queryAnimal(request.name, request.field)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func createAnimal(name string, kind string) {
	var animal Animal
	switch kind {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Println("Invalid animal type")
		return
	}
	addAnimalToDatabase(name, animal)
	fmt.Println("Created it!")
}

func queryAnimal(name string, info string) {
	animal := getAnimalFromDatabase(name)
	if animal == nil {
		fmt.Printf("Animal with name %s not found in the database\n", name)
		return
	}

	fmt.Printf("%s ", strings.Title(name))
	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("does nothing")
	}
}

func getAnimalFromDatabase(name string) Animal {
	animal, ok := Database[name]
	if !ok {
		return nil
	}

	return animal
}

func addAnimalToDatabase(name string, animal Animal) {
	Database[name] = animal
}

func parseUserInput(input string) (*UserCommand, error) {
	result := UserCommand{}

	splitInput := strings.Fields(input)
	if len(splitInput) != 3 {
		return nil, errors.New("not enough information")
	}

	result.command = splitInput[0]
	result.name = splitInput[1]
	result.field = splitInput[2]

	return &result, nil
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

func printInstructions() {
	instructions := "To create new animal enter: newanimal NAME TYPE\n" +
		"\tTYPE has to be one of: cow, bird, snake\n" +
		"To request information about animal enter: query NAME INFORMATION\n" +
		"\tINFORMATION has to be one of: eat, move, speak\n" +
		"To quit application enter Q\n"
	fmt.Print(instructions)
}
