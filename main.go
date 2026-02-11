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

// package main

// import "fmt"

// func updateAge(n *int) {
// 	*n = 35
// }

// func main() {
// 	age := 21
// 	updateAge(&age)

// 	fmt.Println(age)
// }

// package main

// import "fmt"

// func Greet(c chan string, d chan bool) {
// 	name := <-c
// 	fmt.Println("Hello ", name)
// 	d <- true

// }

// func main() {
// 	c := make(chan string)
// 	done := make(chan bool)

// 	go Greet(c, done)

// 	c <- "Hashin"
// 	<-done
// 	close(c)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		ch1 <- "Hello"
// 	}()

// 	go func() {
// 		ch2 <- "Hy"
// 	}()

// 	select {
// 	case value1 := <-ch1:
// 		time.Sleep(1 * time.Second)

// 		fmt.Println("Receved", value1)

// 	case msg2 := <-ch2:
// 		time.Sleep(1 * time.Second)

// 		fmt.Println("Receved :", msg2)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var count = 0
// var mu sync.Mutex

// func increment() {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// }

// func main() {
// 	var vg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		vg.Add(1)
// 		go func() {
// 			increment()
// 			vg.Done()
// 		}()

// 		// fmt.Println(i)
// 	}
// 	vg.Wait()
// 	fmt.Println(count)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printeven(vg *sync.WaitGroup) {
// 	defer vg.Done()
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println("Even", i)
// 		}
// 	}
// }

// func printOdd(vg *sync.WaitGroup) {
// 	defer vg.Done()
// 	for i := 1; i <= 10; i++ {
// 		if i%2 != 0 {
// 			fmt.Println("Odd :", i)
// 		}
// 	}
// }

// func main() {
// 	var vg sync.WaitGroup
// 	vg.Add(2)

// 	go printeven(&vg)
// 	go printOdd(&vg)
// 	vg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func task1(wg *sync.WaitGroup) {

// 	fmt.Println("Hello")
// 	fmt.Println("Task A Completed")
// 	wg.Done()
// 	wg.Done()
// }

// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	fmt.Println("hyy")
// 	go task1(&wg)

// 	wg.Wait()
// 	fmt.Println("Main exit")
// }

// package main

// import "fmt"

// func main() {
// 	marks := 33

// 	switch {
// 	case marks >= 90:
// 		fmt.Println("A Grade")

// 	case marks >= 75:
// 		fmt.Println("B Grade")
// 	case marks >= 45:
// 		fmt.Println("c grade")
// 	default:
// 		fmt.Println("Failed")
// 	}
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	num, num2 := 0, 2
// 	err := errors.New("Invalid input")

// 	if num == 0 {
// 		fmt.Println("Error :", err)
// 	} else {
// 		fmt.Println("Result :", num%num2)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var choice int

// 	fmt.Println("///Menu////")
// 	fmt.Println("1. Say Hello")
// 	fmt.Println("2. Add number")
// 	fmt.Println("3. Say")
// 	fmt.Println("4. hyy")
// 	fmt.Print("Enter A choice")
// 	fmt.Scanln(&choice)

// 	switch choice {
// 	case 1:
// 		fmt.Println("Hello How are you")
// 	case 2:
// 		var a, b int
// 		fmt.Print("Enter First Number  :")
// 		fmt.Scanln(&a)
// 		fmt.Print("Enter Second Number :")
// 		fmt.Scanln(&b)
// 		fmt.Println("Result :", a+b)
// 	case 3:
// 		fmt.Println("Good Morning")
// 	case 4:
// 		fmt.Println("Exit the program")
// 	}

// }

// package main

// import "fmt"

// func main() {
// 	arr := make([]int, 0, 3)
// 	fmt.Println(cap(arr))
// 	arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 89, 9, 9, 2, 3, 4, 5)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func max[T int | float64](a T, b T) T {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}

// }
// func main() {
// 	fmt.Println(max(4, 2))
// 	fmt.Println(max(2.4, 4.4))
// }

// package main

// import "fmt"

// func printSlice[T any](items []T) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }
// func main() {
// 	printSlice([]int{1, 2, 3, 4, 55, 6, 7})
// 	printSlice([]string{"hy", "Hello"})
// }

package main

import "fmt"

type Box[T any] struct {
	Value T
}

func main() {
	b1 := Box[int]{Value: 23}
	b2 := Box[string]{Value: "Hyy"}

	fmt.Println(b1, b2)
}
