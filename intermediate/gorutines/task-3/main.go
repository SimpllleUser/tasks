package main

import (
	"fmt"
	"time"
)

// Task 2.3: Timeout Pattern
// Створи goroutine яка виконує “важку операцію” (sleep 3 секунди).
// Якщо операція не завершилась за 1 секунду - поверни timeout error. Використай select і time.After

func main() {

	timeDuration := 6
	done := make(chan bool)
	progres := make(chan int)

	go func() {
		for i := 1; i <= timeDuration; i++ {
			time.Sleep(1 * time.Second)
			progres <- i
		}
		done <- true

	}()

	timeOut := time.After(5 * time.Second)

	for {
		select {
		case sec := <-progres:
			fmt.Println("Progress time", sec)
		case <-done:
			fmt.Println("Done operation")
			return

		case <-timeOut:
			fmt.Println("To long time operation")
			return
		}
	}
}
