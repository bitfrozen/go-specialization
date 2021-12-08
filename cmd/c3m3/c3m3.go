package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func printAndSort(slice *[]int, wg *sync.WaitGroup) {
	fmt.Printf("Will sort this slice: %v\n", *slice)
	sort.Ints(*slice)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	userNumbers := getSliceOfIntFromUser()
	quarterSlices := splitIntoQuarters(userNumbers)

	wg.Add(4)
	go printAndSort(&quarterSlices[0], &wg)
	go printAndSort(&quarterSlices[1], &wg)
	go printAndSort(&quarterSlices[2], &wg)
	go printAndSort(&quarterSlices[3], &wg)
	wg.Wait()

	firstHalf := mergeSorted(quarterSlices[0], quarterSlices[1])
	secondHalf := mergeSorted(quarterSlices[2], quarterSlices[3])
	whole := mergeSorted(firstHalf, secondHalf)

	fmt.Printf("Sorted array: %v\n", whole)
}

func splitIntoQuarters(slice []int) [4][]int {
	var result [4][]int
	size := len(slice)

	for i := 0; i < size; i++ {
		targetArray := i % 4
		result[targetArray] = append(result[targetArray], slice[i])
	}

	return result
}

func mergeSorted(a, b []int) []int {
	aLen := len(a)
	bLen := len(b)
	result := make([]int, aLen+bLen)
	var i, j, k int

	for i, j, k = 0, 0, 0; i < aLen && j < bLen; k++ {
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
	}
	for i < aLen {
		result[k] = a[i]
		i++
		k++
	}
	for j < bLen {
		result[k] = b[j]
		j++
		k++
	}

	return result
}

func getSliceOfIntFromUser() []int {
	userInput := getInputFromUser()
	userSlice := parseInputFromUser(userInput)

	return userSlice
}

func getInputFromUser() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type in a sequence of integers (separate using spaces): ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while getting user input: %v\n", err)
	}
	userInput := scanner.Text()

	return userInput
}

func parseInputFromUser(input string) []int {
	splitInput := strings.Fields(input)
	var result = make([]int, len(splitInput))

	for i, s := range splitInput {
		// If input is wrong, Atoi will return 0 and error.
		// 0 in case of error satisfies our requirements, so error state can be ignored.
		result[i], _ = strconv.Atoi(s)
	}

	return result
}
