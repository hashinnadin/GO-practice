// package main

// import "fmt"

// func main() {
// 	arr := [...]int{1, 2, 3, 9, 4, 5, 6}
// 	slice := arr[:]
// 	slice = append(slice, 7)
// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

// package main

// import "fmt"

// func main() {
// 	students := make(map[int]string)

// 	students[1] = "HAshin"
// 	students[2] = "shaju"
// 	students[3] = "anu"
// 	students[4] = "mnu"
// 	students[5] = "snu"
// 	students[6] = "lnu"
// 	students[7] = "Hanan"

// 	delete(students, 2)

// 	for id, name := range students {
// 		fmt.Println(id, name)
// 	}
// }

package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 6, 5}
	max := nums[0]
	min := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println("MAx :", max)
	fmt.Println("Min :", min)
}
