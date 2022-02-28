package main

import "fmt"

// methode - 1

// we have to initalize variable type and the return type

func add(a int, b int) int {
	r := a + b
	return r
}

// methode - 2

func sub(a, b int) int {
	result := a - b
	return result
}

func main() {

	// A function is a group of statements that perform a particular task.

	// or

	// A function is a group of statements that can be used repeatedly in a program.

	fmt.Println("Hello World!!!")

	result := add(5, 2)
	fmt.Println(result)
	subtraction := sub(9, 5)

	fmt.Println("Subtraction is : ", subtraction)

}
