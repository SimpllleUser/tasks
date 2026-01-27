package main

import (
	"fmt"
	"time"
)

// Task 3.2: Rate Limiter
// Реалізуй rate limiter який дозволяє максимум 5 операцій в секунду.
// Є 20 tasks - обробляй їх з rate limiting.

// func main() {
// 	tasks := 20
// 	tTicker := time.NewTicker(time.Duration(200) * time.Millisecond)

// 	time := 0
// 	sec := 0

// 	for i := 0; i < tasks; i++ {
// 		<-tTicker.C

// 		time += 200

// 		if time%1000 == 0 {
// 			sec += 1
// 		}

// 		fmt.Println("Queue:", sec)
// 	}
// }

func worker(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Processing job:", job)
	}
}
func main() {

	amountOfTasks := 20
	ticker := time.NewTicker(time.Second)
	jobs := make(chan int)
	operationPerSecond := 5

	for i := 0; i < operationPerSecond; i++ {
		go worker(jobs)
	}

	for i := 0; i < amountOfTasks; i++ {
		if i%operationPerSecond == 0 && i != 0 {
			println("Waiting for next second...")
			<-ticker.C
		}
		jobs <- i

	}

	close(jobs)
	time.Sleep(time.Duration(operationPerSecond) * time.Second)

}
