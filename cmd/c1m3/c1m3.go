package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var collection = make([]int, 3)
	var userInput string

	for count := 0; true; count++{
		fmt.Print("Enter integer (or X to quit): ")

		_, err := fmt.Scanln(&userInput)
		if err != nil || strings.ToLower(userInput) == "x"  {
			break
		}

		userInteger, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Wrong input")
			count--
			continue
		}

		// For initial slice size, replace 0 values.
		// If user tries to add more - append to original slice
		if count < 3 {
			insertIndex := sort.SearchInts(collection, 0)
			collection[insertIndex] = userInteger
		} else {
			collection = append(collection, userInteger)
		}

		sort.Ints(collection)
		fmt.Println(collection)
	}
}
