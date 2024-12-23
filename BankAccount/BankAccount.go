package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Println("Deposit", amount, b.Balance)
	}
	fmt.Println("Deposit can`t be negative.")
}
func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdraw amount must be positive!")
		return
	}
	if amount > b.Balance {
		fmt.Println("Insufficient balance!")
		return
	}
	b.Balance -= amount
	fmt.Printf("Withdrew: ", amount, b.Balance)
}

func (b *BankAccount) GetBalance() {
	fmt.Println("Balance now: ", b.Balance)
}
func (b *BankAccount) Transaction(account *BankAccount, transacitions []float64) {
	for _, amount := range transacitions {
		if amount > 0 {
			account.Deposit(amount)
		}
		account.Withdraw(-amount)
	}
}
func main() {
	account := BankAccount{
		AccountNumber: "777124",
		HolderName:    "Askat Narinbetov",
		Balance:       0.0,
	}
	var choice int
	var amount float64

	for {
		fmt.Println("\n ---Bank Account Menu ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
