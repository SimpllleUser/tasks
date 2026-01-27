package main

import (
	"context"
	"fmt"
	"time"
)

// Task 3.3: Context Cancellation
// Створи 3 workers які працюють в нескінченному циклі.
//  Використай context.WithCancel щоб зупинити всіх workers після 5 секунд.

func worker(ctx context.Context, workerID int, sec *int) {
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d sec:%d stopping\n", workerID, *sec)
			return
		default:
			*sec += 1
			fmt.Printf("Worker %d sec: %d working\n", workerID, *sec)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	sec := []int{0, 0, 0}
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go worker(ctx, i+1, &sec[i])
	}

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
