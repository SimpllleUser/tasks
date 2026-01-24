package main

import (
	"errors"
	"fmt"
)

// Task 2.9: Method with Pointer Receiver

// type BankAccount struct {
// Balance float64
// }

// Реалізуй методи: Deposit(amount float64), Withdraw(amount float64) error. Використай pointer receivers.

type BankAccount struct {
	Balance float64
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount > ba.Balance {
		return errors.New("To small balance")
	}

	ba.Balance -= amount
	return nil
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount

}

func main() {

	bankAccount := BankAccount{
		Balance: float64(100.12),
	}

	bankAccount.Deposit(float64(12.0))

	err := bankAccount.Withdraw(13.0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%.2f", bankAccount.Balance)

}
