package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	displaceFn := func(t float64) float64 {
		displace := 0.5*a*math.Pow(t, 2) + vo*t + so

		return displace
	}

	return displaceFn
}

func main() {
	initialValues := getInitialData()
	fn := GenDisplaceFn(initialValues["a"], initialValues["vo"], initialValues["so"])

	userTimeInput := getInputFromUser("Enter time")
	userTime, _ := strconv.ParseFloat(userTimeInput, 64)

	fmt.Printf("Displacement after provided time: %f", fn(userTime))
}

func getInitialData() map[string]float64 {
	userValues := make(map[string]float64, 3)
	var userInput string

	userInput = getInputFromUser("Enter acceleration")
	userValues["a"], _ = strconv.ParseFloat(userInput, 64)

	userInput = getInputFromUser("Enter initial velocity")
	userValues["vo"], _ = strconv.ParseFloat(userInput, 64)

	userInput = getInputFromUser("Enter initial displacement")
	userValues["so"], _ = strconv.ParseFloat(userInput, 64)

	return userValues
}

func getInputFromUser(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s: ", prompt)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error while getting user input: %v\n", err)
	}
	userInput := scanner.Text()

	return userInput
}
