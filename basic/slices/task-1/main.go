package main

// Task 1.4: Sum of Slice
// Напиши функцію Sum(numbers []int) int яка повертає суму всіх елементів slice.

func getSum(list []int) int {
	result := 0

	for _, num := range list {
		result += num
	}

	return result

}

func main() {
	println("Result ", getSum([]int{1, 2, 3, 4, 5, 7, 8, 9, 10}))
}
