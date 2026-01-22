package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Task 2.2: Race Condition Debug
// Дано код:

// var balance int

// for i := 0; i < 100; i++ {
// go func() {
// balance = balance + 10
// }()
// }

// Знайди і виправ race condition двома способами: Mutex і atomic.

//  ===== >  Mutex

// func count(val *int, mu *sync.Mutex, wg *sync.WaitGroup) {
// 	mu.Lock()
// 	*val += 10
// 	mu.Unlock()
// }

// func main() {

// 	var mu sync.Mutex

// 	var balance int
// 	var wg sync.WaitGroup

// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go count(&balance, &mu, &wg)
// 	}
// 	defer wg.Done()
// 	wg.Wait()

// 	fmt.Println("Hello!", balance)

// }

// ===== > Atomic

func count(at *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()
	at.Add(10)
}

func main() {
	var at atomic.Int32

	// var balance int
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go count(&at, &wg)
	}
	wg.Wait()

	fmt.Println("Hello!", at.Load())

}
