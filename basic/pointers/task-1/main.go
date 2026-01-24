package main

// Task 1.7: Swap Two Numbers
// Напиши функцію Swap(a, b *int) яка міняє місцями два числа.

func swap(first *int, second *int) {

	temp := *second
	*second = *first
	*first = temp
}

func main() {
	first := 1
	second := 2

	swap(&first, &second)

	println("Result ", first, second)
}
