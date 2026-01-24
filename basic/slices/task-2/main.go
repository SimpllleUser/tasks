package main

import "fmt"

// Task 1.5: Filter Even Numbers
// Напиши функцію FilterEven(numbers []int) []int яка повертає новий slice тільки з парними числами.

func filterEvent(list []int) []int {
	result := []int{}

	for _, num := range list {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result

}

func main() {
	fmt.Printf("Result %v \n", filterEvent([]int{1, 2, 3, 4, 5, 7, 8, 9, 10}))
}
