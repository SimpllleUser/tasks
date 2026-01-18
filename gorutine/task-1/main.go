package main

import (
	"fmt"
	"sync"
)

// Task 1.1: Hello Goroutines
// Створи 5 goroutines, кожна виводить “Hello from goroutine X” де X - номер. Використай WaitGroup щоб main дочекався всіх.

func helloWorld(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine", number)
}

func main() {

	var wg sync.WaitGroup
	length := 5

	wg.Add(length)
	for i := 0; i < length; i++ {
		go helloWorld(i+1, &wg)
	}

	wg.Wait()
}
