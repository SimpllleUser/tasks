package main

import "fmt"

// Task 1.6: Reverse Slice
// Напиши функцію Reverse(numbers []int) яка перевертає slice на місці (in-place).

func reverse(list []int) []int {
	result := []int{}

	for i := len(list); i > 0; i-- {
		result = append(result, list[i-1])
	}

	return result

}

func main() {
	fmt.Printf("Result %v \n", reverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
