package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Hello I am Working")

	// more than array.
	// slice

	// array declaration

	var student = [3]string{"Riyaz", "Shuvo", "Ismail"}

	fmt.Println(student)

	// the other way to declare an array

	secondWay := [...]string{"Riyaz", "Khan", "Shuvo"}

	fmt.Println(secondWay)

	// slice array

	sliceArray := student[0:2]

	fmt.Println(sliceArray)

	// fruits := []string{}

	// or

	var fruits []string

	fruits = append(fruits, "apple", "banana")
	fmt.Println(fruits)

	fmt.Printf("%T \n", fruits)
	fmt.Printf("%T\n", secondWay)

	a := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(a)

	b := reflect.TypeOf(secondWay).Kind().String()
	fmt.Println(b)

}
