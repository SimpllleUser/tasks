package main

import "fmt"

// Task 1.8: Update Struct

// type Person struct {
// Name string
// Age  int
// }

// Напиши функцію UpdateAge(p *Person, newAge int) яка змінює вік.

type Person struct {
	Name string
	Age  int
}

func changeName(p *Person, name string) {
	p.Name = name
}

func changeAge(p *Person, age int) {
	p.Age = age
}

func main() {

	person := Person{Name: "Alex", Age: 18}

	changeName(&person, "Tim")
	changeName(&person, "Tim")

	fmt.Printf("Result %v", person)
}
