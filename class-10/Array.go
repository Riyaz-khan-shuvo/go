package main

import "fmt"

func main() {
	//array: array is a special type of variable that stores multiple value in one single value;

	// primitive data type : int , float32,string, bool
	// primitive alias : rune byte

	fmt.Printf("Hello World!!!")

	var Student [3]string
	fmt.Println(Student)
	fmt.Println("The length of this array is : ", len(Student))

	// fmt.Println("Enter The size of the array : ")
	// var n int
	// fmt.Scan(&n)

	// var Students [3]string

	// array declaration in long way

	// Students := [3]string{}

	// array declaration in short way or array string literal

	// var i int
	// for i = 0; i < 3; i++ {
	// 	fmt.Scanf("%s", &Students[i])
	// }
	// for i = 0; i < 3; i++ {
	// 	fmt.Printf("Student[%d] : %s\n", i, Students[i])
	// }

	userName := [...]string{"Riyaz", "Khan", "Shuvo", "Shahid", "Khan"}

	fmt.Print(userName)

}
