package main

import (
	"fmt"
	"math/rand/v2"
)

// Task 1.3: Simple Channel
// Створи goroutine яка генерує числа від 1 до 5 і відправляє в channel. Main читає і виводить ці числа.

func main() {

	ch := make(chan int)

	go func() {

		ch <- rand.IntN(5) + 1
	}()

	fmt.Println("Random number", <-ch)

}
