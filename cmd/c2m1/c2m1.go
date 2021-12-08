package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	lastItemIdx := len(slice) - 1
	if i < 0 || i >= lastItemIdx {
		return
	}

	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	lastItemIdx := len(slice) - 1
	lastSwappedIdx := lastItemIdx

	for i := 0; i < len(slice); i++ {
		for j := 0; j < lastItemIdx; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				lastSwappedIdx = j + 1
			}
		}
		lastItemIdx = lastSwappedIdx
	}
}

func main() {
	userNumbers := getSliceOfIntFromUser()
	BubbleSort(userNumbers)

	for _, number := range userNumbers {
		fmt.Printf("%v ", number)
	}
}

func getSliceOfIntFromUser() []int {
	userInput := getInputFromUser()
	userSlice := parseInputFromUser(userInput, 10)

	return userSlice
}

func getInputFromUser() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type in a sequence of up to 10 integers (separate using spaces): ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while getting user input: %v\n", err)
	}
	userInput := scanner.Text()

	return userInput
}

func parseInputFromUser(input string, sizeLimit int) []int {
	finalSliceSize := sizeLimit
	splitInput := strings.Fields(input)

	if len(splitInput) < sizeLimit {
		finalSliceSize = len(splitInput)
	}
	var result = make([]int, finalSliceSize)

	for i, s := range splitInput {
		if i >= finalSliceSize {
			break
		}
		// If input is wrong, Atoi will return 0 and error.
		// 0 in case of error satisfies our requirements, so error state can be ignored.
		result[i], _ = strconv.Atoi(s)
	}

	return result
}
