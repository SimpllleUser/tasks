package main

import (
	"fmt"
	"sync"
)

// Task 1.2: Counter Race
// Є змінна counter := 0. Запусти 1000 goroutines, кожна робить counter++. Виправ race condition використовуючи Mutex. Перевір з go run -race

func makeCount(count *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*count++
	mu.Unlock()
}

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup

	count := 0
	length := 1000

	wg.Add(length)
	for i := 0; i < length; i++ {
		go makeCount(&count, &mu, &wg)
	}

	wg.Wait()

	fmt.Println("Count: ", count)
}
