package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func checkString(input string) bool {
	var result bool

	reg := regexp.MustCompile(`(?i)^i.*a.*n$`)
	result = reg.MatchString(input)

	return result
}

func main() {
	var userInput string
	var scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("Enter string: ")
	scanner.Scan()
	userInput = scanner.Text()

	if checkString(userInput) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
