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

// package main
// import (
//     "fmt"
//     "time"
// )
// type Email struct {
//     To      string
//     Subject string
//     Body    string
// }
// // Email worker - відправляє emails
// func emailWorker(id int, emails <-chan Email, done chan<- bool) {
//     for email := range emails {
//         fmt.Printf("[Worker %d] Sending email to %s...\n", id, email.To)

//         // Імітація відправки email (API call)
//         time.Sleep(500 * time.Millisecond)

//         fmt.Printf("[Worker %d] ✅ Sent: %s\n", id, email.Subject)
//     }

//     // Worker закінчив
//     done <- true
// }
// func main() {
//     // Channel для emails
//     emails := make(chan Email, 10)

//     // Channel для сигналу завершення
//     done := make(chan bool)

//     // Запускаємо 3 workers
//     for i := 1; i <= 3; i++ {
//         go emailWorker(i, emails, done)
//     }

//     // Відправляємо emails в channel
//     emailList := []Email{
//         {To: "user1@example.com", Subject: "Welcome!", Body: "..."},
//         {To: "user2@example.com", Subject: "Invoice", Body: "..."},
//         {To: "user3@example.com", Subject: "Newsletter", Body: "..."},
//         {To: "user4@example.com", Subject: "Reset password", Body: "..."},
//         {To: "user5@example.com", Subject: "Notification", Body: "..."},
//     }

//     for _, email := range emailList {
//         emails <- email
//     }

//     close(emails) // Більше emails не буде

//     // Чекаємо всіх workers
//     for i := 1; i <= 3; i++ {
//         <-done
//     }

// }
