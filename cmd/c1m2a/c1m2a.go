package main

import "fmt"

type Bridge struct {
	length int
	height int
}

func main() {
	var crossings [3]Bridge

	crossings[0] = Bridge{4, 4}

	for i, v := range crossings {
		fmt.Printf("index %v, value %v\n", i, v)
	}
}