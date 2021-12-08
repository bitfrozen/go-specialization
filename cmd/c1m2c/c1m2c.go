package main

import (
	"fmt"
	"math"
)

func main() {
	var formattedInput float64

	fmt.Print("Enter floating point number: ")
	n, err := fmt.Scanf("%f", &formattedInput)

	if n > 0 {
		fmt.Printf("%.0f", math.Trunc(formattedInput))
	} else {
		fmt.Printf("Wrong input. Error: %v", err)
	}
}
