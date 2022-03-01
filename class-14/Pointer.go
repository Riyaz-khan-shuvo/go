package main

import "fmt"

func main() {
	var x int

	var y *int

	fmt.Println("X value is : ", x)
	fmt.Println("X memory address is : ", &x)
	fmt.Println("Y value is : ", y)
	fmt.Println("Y memory address is : ", &y)
}
