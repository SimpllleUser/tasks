package main

import (
	"fmt"
	"time"
)

// Task 3.2: Rate Limiter
// Реалізуй rate limiter який дозволяє максимум 5 операцій в секунду.
// Є 20 tasks - обробляй їх з rate limiting.

func main() {
	tasks := 20
	tTicker := time.NewTicker(time.Duration(200) * time.Millisecond)

	time := 0
	sec := 0

	for i := 0; i < tasks; i++ {
		<-tTicker.C

		time += 200

		if time%1000 == 0 {
			sec += 1
		}

		fmt.Println("Queue:", sec)
	}
}
