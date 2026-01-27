package main

import "fmt"

// Task 3.5: Find All Permutations
// Напиши функцію яка знаходить всі перестановки slice.

// Input: [1, 2, 3]
// Output: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]

func replace(a, b int) (int, int) {
	return b, a
}

func getElementNotWithIndex(index int, list []int) []int {
	copyList := append([]int(nil), list...)
	return append(copyList[:index], copyList[index+1:]...)
}

func permutations(list []int) [][]int {

	if len(list) <= 1 {
		return [][]int{list}
	}
	results := [][]int{}

	for i := 0; i < len(list); i++ {
		current := list[i]
		rest := getElementNotWithIndex(i, list)

		subList := permutations(rest)
		fmt.Printf("subList : %v \n", subList)

		for _, perm := range subList {
			fmt.Printf("perm : %v \n", subList)
			newPerm := append([]int{current}, perm...)
			results = append(results, newPerm)
		}
	}
	return results
}

func main() {

	list := []int{10, 20, 30}

	fmt.Printf("Result : %v", permutations(list))

}
