// package main

// import "fmt"

// func main() {
// 	var n int

// 	fmt.Print("Enter number of terms: ")
// 	fmt.Scanln(&n)

// 	a := 0
// 	b := 1

// 	fmt.Println("Fibonacci Series:")

// 	for i := 1; i <= n; i++ {
// 		fmt.Print(a)
// 		next := a + b
// 		a = b
// 		b = next
// 	}
// }

// package main

// import (
// 	dup "arrayslice/Dup"
// 	"fmt"
// )

// func main() {
// 	var nums int

// 	fmt.Println("===== USER CRUD MENU =====")
// 	fmt.Println("1. Create User")
// 	fmt.Println("2. Read Users")
// 	fmt.Println("3. Update User")
// 	fmt.Println("4. Delete User")
// 	fmt.Println("5. Exit")

// 	fmt.Print("Enter your choice: ")
// 	fmt.Scanln(&nums)

// 	switch nums {
// 	case 1:
// 		fmt.Println("user created")
// 	case 2:
// 		fmt.Println("Data receved")
// 	case 3:
// 		fmt.Println("Edited data")
// 	case 4:
// 		fmt.Println("deleted data")
// 	case 5:
// 		fmt.Println("exit")
// 	default:
// 		fmt.Println("Exited")
// 	}

// 	dup.Duplicate()
// }

// package main

// import "fmt"

// func main() {
// 	my := map[string]int{
// 		"hashin": 6,
// 		"nabeel": 10,
// 		"manhar": 15,
// 	}

//		val, ok := my["manhar"]
//		if ok {
//			fmt.Println(val)
//		}
//	}

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 2, 3, 4, 4, 5}

// 	unique := []int{}
// 	seen := make(map[int]bool)

// 	for _, v := range nums {
// 		if !seen[v] {
// 			seen[v] = true
// 			unique = append(unique, v)
// 		}
// 	}

// 	fmt.Println("Without duplicates:", unique)
// }

// package main

// import "fmt"

// func main() {
// 	grades := make(map[string]string)

// 	grades["rahul"] = "A"
// 	grades["Hashin"] = "A+"

// 	for name, grade := range grades {
// 		fmt.Println(name, ":", grade)
// 	}
// }

package main

import "fmt"

func main() {
	students := []string{}
	students = append(students, "Hashin")
	students = append(students, "Salman")
	fmt.Println(students)
}
