package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Printf("Enter Your Name and Age : ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Your Name is	: %s \nYour Age is	: %d", name, age)
}
