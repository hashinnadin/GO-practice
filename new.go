package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter A number")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
