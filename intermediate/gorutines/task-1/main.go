package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Task 2.1: Parallel File Processing
// Є slice з іменами файлів []string{"file1.txt", "file2.txt", ...}.
// Створи worker pool (3 workers) який “обробляє” файли паралельно (можна просто sleep). Виведи який worker обробив який файл.

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		randomTime := rand.IntN(5)
		fmt.Println("Work with file with ID:", id, job)
		time.Sleep(time.Duration(randomTime) * time.Second)
	}
}

func main() {

	var wg sync.WaitGroup
	jobs := make(chan string)

	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt"}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i+1, jobs, &wg)
	}

	for _, file := range files {
		jobs <- file
	}

	close(jobs)

	wg.Wait()

}
