package main

import "fmt"

// Task 2.8: Modify Slice via Pointer
// Напиши функцію
// AppendValue(slice *[]int, value int)
// яка додає значення до оригінального slice через pointer

func AppendValue(slice *[]int, value int) {

	*slice = append(*slice, value)
}

func main() {
	val := 6
	list := []int{1, 2, 3, 4, 5}

	AppendValue(&list, val)
	fmt.Println("Result %v", list)
}
