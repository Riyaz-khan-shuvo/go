package main

import "fmt"

func main() {
	var marks int

	fmt.Print("Enter Your Marks : ")
	fmt.Scanf("%d", &marks)
	if marks > 100 || marks < 0 {
		fmt.Println("You Enter Invalid Number !!!")
	} else if marks > 79 {
		fmt.Println("A+")
	} else if marks > 69 {
		fmt.Println("A")
	} else if marks > 59 {
		fmt.Println("A-")
	} else if marks > 49 {
		fmt.Println("B")
	} else if marks > 39 {
		fmt.Println("C")
	} else if marks >= 33 {
		fmt.Println("D")
	} else if marks < 33 {
		fmt.Println("You Fail !!!")
	}

	// by using switch case

	marks = marks / 10
	fmt.Println(marks)

	if marks > 10 || marks < 0 {
		fmt.Println("You Enter Invalid Number !!!")
	} else {
		switch marks {
		case 0:
			fmt.Println("F")

		case 1:
			fmt.Println("F")
		case 2:
			fmt.Println("F")
		case 3:
			fmt.Println("F")
			fallthrough
		case 4:
			fmt.Println("D")
		case 5:
			fmt.Println("C")
		case 6:
			fmt.Println("B")
		case 7:
			fmt.Println("A-")
		case 8:
			fmt.Println("A")
		case 9, 10:
			fmt.Println("A+")
		}
	}

	// fmt.Println("Your marks is : ", marks)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
