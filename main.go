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

// package main

// import "fmt"

// func main() {
// 	students := []string{}
// 	students = append(students, "Hashin")
// 	students = append(students, "Salman")
// 	fmt.Println(students)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 55, 5, 5, 3, 7}

// 	for _, num := range arr {
// 		fmt.Println(num)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	students := map[int]string{
// 		1: "HAshin",
// 		2: "rahul",
// 		3: "mahul",
// 	}
// 	for i, name := range students {
// 		fmt.Println("Roll :", i, "Name :", name)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6}

// 	// fmt.Println(len(slice))
// 	// fmt.Println(cap(slice))
// 	slice = append(slice, 40, 50)
// 	// fmt.Println(slice)

// 	max := slice[0]

// 	for _, num := range slice {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	fmt.Println(max)
// }

// package main

// import "fmt"

// func main() {
// 	students := map[int]string{
// 		1: "Hashin",
// 		2: "Hanan",
// 		3: "Shahin",
// 		4: "Manhar",
// 		5: "Nabeel",
// 	}

// 	students[8] = "Halo"
// 	students[2] = "Shanu"

// 	name := students[1]
// 	fmt.Println(name)

// 	value, key := students[4]
// 	delete(students, 3)

// 	if key {
// 		fmt.Println("Found :", value)
// 	} else {
// 		fmt.Println("Not found the value")
// 	}

// 	for _, value := range students {
// 		fmt.Println(value)
// 	}

// 	fmt.Println(students)
// }

// package main

// import "fmt"

// func main() {
// 	type Student struct {
// 		Name string
// 		Age  int
// 		Mark int
// 	}

// 	s1 := Student{
// 		Name: "Hashin",
// 		Age:  21,
// 		Mark: 100,
// 	}

// 	s2 := Student{
// 		Name: "salman",
// 		Age:  12,
// 		Mark: 95,
// 	}
// 	fmt.Println(s2)
// 	fmt.Println(s1.Name)
// 	fmt.Println(s1.Age)
// 	fmt.Println(s1.Mark)
// }

// package main

// import "fmt"

// func main() {
// 	var Car struct {
// 		id   int
// 		name string
// 	}

// 	Car.id = 1
// 	Car.name = "BMW"
// 	fmt.Println(Car)
// 	fmt.Println(Car.id)
// }

// package main

// import "fmt"

// func main() {
// 	type Menu struct {
// 		name  string
// 		price map[string]float32
// 	}

// 	snacks := []Menu{
// 		{
// 			name: "Coffee",
// 			price: map[string]float32{
// 				"regular": 23.65,
// 				"Large":   34.5,
// 			},
// 		},
// 		{
// 			name: "Tea",
// 			price: map[string]float32{
// 				"Regular": 12.5,
// 				"Large":   20.5,
// 			},
// 		},
// 	}
// 	fmt.Println(snacks)
// }

// package main

// import "fmt"

// type Student struct {
// 	Name string
// 	Age  int
// 	Mark int
// }

// func (s Student) Greet() {
// 	fmt.Println("Hello my name is", s.Name)
// 	fmt.Println("My age is", s.Age)
// }

// func main() {

// 	persons := []Student{
// 		{
// 			Name: "Hashin",
// 			Age:  21,
// 		},
// 		{
// 			Name: "Rahul",
// 			Age:  22,
// 		},
// 	}

// 	for _, p := range persons {
// 		p.Greet()
// 	}
// }

// struct simple example

// package main

// import "fmt"

// type Student struct {
// 	name string
// 	age  int
// }

// func (s Student) Hello() {
// 	fmt.Println("My name is ", s.name)
// 	fmt.Println("My age is ", s.age)
// }
// func main() {
// 	s1 := Student{
// 		name: "Hashin",
// 		age:  21,
// 	}
// 	s1.Hello()
// }

// package main

// import "fmt"

// type Driver interface {
// 	Drive()
// }

// type Person1 struct {
// 	Name string
// }

// func (s Person1) Drive() {
// 	fmt.Println(s.Name, "You can Drive car")
// }

// type persons2 struct {
// 	Name string
// }

// func (s persons2) Drive() {
// 	fmt.Println(s.Name, "You can Drive Car")
// }

// func useDriver(d Driver) {
// 	d.Drive()
// }
// func main() {

// 	s := Person1{Name: "Hashin"}
// 	p := persons2{Name: "Salman"}

// 	useDriver(s)
// 	useDriver(p)
// }

// package main

// import "fmt"

// type Student struct {
// 	Name string
// 	Age  int
// }

// func updateAge(s *Student) {
// 	s.Age = 30
// }

// func main() {
// 	stu := Student{
// 		Name: "Hashin",
// 		Age:  21,
// 	}

// 	updateAge(&stu)
// 	fmt.Println(stu.Age)
// 	fmt.Println(stu)
// }

// package main

// import "fmt"

// func change(x *int) {
// 	*x = 20
// }

// func main() {
// 	a := 10
// 	change(&a)
// 	fmt.Println(a)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	fmt.Println("Hello")
// }

// func main() {
// 	go sayHello()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main done")
// }

// package main

// import (
// 	"fmt"
// )

// func printNum(n int) {
// 	fmt.Println(n)
// }
// func main() {
// 	for i := 1; i <= 5; i++ {
// 		go printNum(i)
// 	}
// }

package main

import "fmt"

func updateAge(n *int) {
	*n = 35
}

func main() {
	age := 21
	updateAge(&age)

	fmt.Println(age)
}
