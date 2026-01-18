package main

import (
	"errors"
	"fmt"
)

// Task 1.9: Nil Pointer Check
// Напиши функцію SafeIncrement(num *int) яка збільшує значення на 1, але спочатку перевіряє чи pointer не nil.

type Person struct {
	Name string
	Age  int
}

func SafeIncrement(num *int) error {
	if num == nil {
		return errors.New("Is not correct value")
	}

	*num++
	return nil
}

func main() {

	count := 12

	SafeIncrement(&count)

	var count2 *int
	if err := SafeIncrement(count2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result", count)
}
