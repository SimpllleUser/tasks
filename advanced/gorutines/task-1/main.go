package main

import (
	"fmt"
	"math/rand"
)

// Task 3.1: Pipeline Pattern
// Створи pipeline з трьох stages:
// 1.	Generator: генерує числа 1-100
// 2.	Square: множить на себе
// 3.	Filter: пропускає тільки числа > 50
// Кожен stage - окрема goroutine з channels між ними.

// func generator(n int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for i := 1; i <= n; i++ {
// 			out <- i
// 		}
// 	}()
// 	return out
// }

func Generator(length int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for i := 0; i < length; i++ {
			rand := rand.Intn(100)
			res <- rand
		}
	}()

	return res
}

func Square(number <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for num := range number {
			res <- num * num
		}
	}()

	return res
}

func Filter(number <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for num := range number {
			if num > 50 {
				res <- num
			}
		}
	}()

	return res
}

func main() {

	nums := Generator(10)
	square := Square(nums)
	results := Filter(square)

	for res := range results {
		fmt.Println("Result ", res)
	}
}
