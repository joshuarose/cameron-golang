package main

import (
	"errors"
	"log"
)

// main is the entry point for the application. Every Go program must have a main package with a main function.
func main() {
	// Call the add function with two integer arguments.
	result, err := add(1, 2)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the result to the console.
	println(result)
}

// add is a function that adds two integers and returns the result.
func add(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("x or y is less than 0")
	}
	return x + y, nil
}
