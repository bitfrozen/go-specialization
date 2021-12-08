package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Name struct {
	fname [20]rune
	lname [20]rune
}

func main() {
	var nameCollection []Name

	// get filename
	var userInput string
	fmt.Println("Provide path to file containing names:")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error getting user input: %v", err)
		return
	}

	// open file
	absoluteFilePath, _ := filepath.Abs(userInput)
	userFile, err := os.Open(absoluteFilePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error opening file: %v", err)
		return
	}
	defer userFile.Close()

	// parse file into Name type slice
	fileScanner := bufio.NewScanner(userFile)
	for fileScanner.Scan() {
		fileLine := fileScanner.Text()

		splitLine := strings.Fields(fileLine)
		if len(splitLine) != 2 {
			_, _ = fmt.Fprintf(os.Stderr, "Invalid data format at line: %v. Skipping line\n", fileLine)
			continue
		}

		var firstName [20]rune
		copy(firstName[:], []rune(splitLine[0]))
		var lastName [20]rune
		copy(lastName[:], []rune(splitLine[1]))

		nameCollection = append(nameCollection, Name{fname: firstName, lname: lastName})
	}

	// print slice
	for _, name := range nameCollection {
		fmt.Println(string(name.fname[:]), string(name.lname[:]))
	}
}
