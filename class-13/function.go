package main

import (
	"fmt"
)

// methode - 1

// we have to initialise variable type and the return type

func add(a int, b int) int {
	r := a + b
	return r
}

// methode - 2

//  the the variables type same in a function we can use only one time the variable type

func sub(a, b int) int {
	result := a - b
	return result
}

// methode - 3

//  we can initialise the return variable type co creating the function

func multi(a, b int) (r int) {
	r = a * b
	return r
}

// methode 4

func divi(a, b float32) (r float32) {
	r = a / b
	return
}

// multiple return value

func rectangle(l, base int) (area, parameter int) {
	parameter = 2 * (l + base)
	area = l * base
	// return area , parameter
	// or
	return
}

// pointer

// func update(a *int, b *string) {
// 	*a = *a + 5 // this is called defrencing pointer address
// 	*b = *b + " Hossain"
// 	return
// }

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

//anonymous function

func main() {

	// A function is a group of statements that perform a particular task.

	// or

	// A function is a group of statements that can be used repeatedly in a program.

	fmt.Println("Hello World!!!")

	result := add(5, 2)
	fmt.Println(result)
	subtraction := sub(9, 5)

	fmt.Println("Subtraction is : ", subtraction)

	multiplication := multi(5, 5)
	fmt.Println("Multiplication is : ", multiplication)

	division := divi(5, 6)
	fmt.Println(division)

	// multiple return value

	a, p := rectangle(10, 5)

	fmt.Println("Area is		: ", a, "\nParameter is	: ", p)

	//
	number := 25
	name := "Riyaz"

	update(&number, &name)
	fmt.Println(number, name)

	//anonymous function => a function without name is called anonymous function

	x := func(a, b int) (c int) {
		c = a * b
		return
	}

	fmt.Println("Multiplication	: ", x(5, 5))

}
