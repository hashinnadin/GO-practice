package main

import "fmt"

func aain() {
	var a, b int
	var op string

	fmt.Print("Enter First Num: ")
	fmt.Scanln(&a)

	fmt.Print("Enter Second Num: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operator (+ - * /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println("Result:", a+b)

	case "-":
		fmt.Println("Result:", a-b)

	case "*":
		fmt.Println("Result:", a*b)

	case "/":
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Error: division by zero")
		}

	default:
		fmt.Println("Invalid operator")
	}
}
