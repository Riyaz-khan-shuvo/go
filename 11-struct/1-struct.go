package main

import "fmt"

// struct declaration

type Book struct {
	title  string
	author string
	ISBN   string
	price  float32
	page   int
}

// composite data
// array
// slice
// map
// struct -> user define data types / composite data types / concreate data types
// a struct is a collection of data fields or properties.
// struct literal

func main() {
	fmt.Println("Hello I am working")

	// set the value of struct and get the value

	var b1 Book
	b1.title = "An Introduction To Programming in GO"
	b1.author = "CAREB DOXSEY"
	b1.ISBN = "9781478355823"
	b1.price = 10.55
	b1.page = 155

	fmt.Println(b1)
	fmt.Println("This is Title : ", b1.title)

	student := struct {
		name string
		age  int
	}{
		name: "Riyaz",
		age:  23,
	}
	fmt.Println(student)

}
