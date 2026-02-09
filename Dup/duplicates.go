package main

import "fmt"

func Duplicate() {
	my := map[string]int{
		"hashin": 6,
		"nabeel": 10,
		"manhar": 15,
	}

	val, ok := my["manhar"]
	if ok {
		fmt.Println(val)
	}
}
