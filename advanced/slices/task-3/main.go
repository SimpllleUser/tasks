package main

import "fmt"

// Є slice чисел і розмір вікна. Знайди максимум в кожному вікні.

// Input: [1,3,-1,-3,5,3,6,7], window=3
// Output: [3,3,5,5,6,7]

func getLargest(list []int) int {
	if len(list) == 0 {
		return 0
	}
	result := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > result {
			result = list[i]
		}
	}

	return result

}

func getSlicedList(start, end int, list []int) []int {
	return list[start:end]
}

func main() {

	window := 3
	results := []int{}
	step := 0
	list := []int{1, 3, -1, -3, 5, 3, 6, 7}

	for step+window <= len(list) {
		slicedList := list[step : step+window]
		results = append(results, getLargest(slicedList))
		step += 1
	}

	fmt.Printf("Result : %v", results)

}
