package project

import (
	"fmt"
	"math/rand"
)

type BankAccount struct {
	AccountNumber int32
	HolderName    string
	Balance       float64
}

func (b *BankAccount) CkBalance() float64 {
	return b.Balance
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("deposit successful\n\n")
}
func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance > amount {
		b.Balance -= amount
		fmt.Printf("withdraw successful\n\n")

	} else {
		fmt.Println("insufficient balance : ", b.Balance)
	}

}

func CrAccount(acNumber int32, hlName string, crBalance float64) BankAccount {
	user := BankAccount{
		AccountNumber: acNumber,
		HolderName:    hlName,
		Balance:       crBalance,
	}
	return user
}

func accData(accNumber int32, accHolder string, amount float64) {
	john := CrAccount(accNumber, accHolder, amount)
	var userInp int

	for {
		fmt.Printf("\n\n1.Check Balance\n2.Deposit Balance\n3.Withdraw Balance\n4.Exit\n\n")
		fmt.Println("<-------------------------------->")
		fmt.Printf("Enter your option: ")
		fmt.Scanln(&userInp)
		if userInp == 1 {
			fmt.Println("your current balance is ", john.CkBalance())
		} else if userInp == 2 {
			var dAmount float64
			fmt.Printf("Enter your deposit amount: ")
			fmt.Scanln(&dAmount)
			john.Deposit(dAmount)
		} else if userInp == 3 {
			var wAmount float64
			fmt.Printf("Enter your withdraw amount: ")
			fmt.Scanln(&wAmount)
			john.Withdraw(wAmount)
		} else {
			fmt.Println("Thank you for use our bank!")
			break
		}
	}

}

func Grand() int64 {
	randomNumber := rand.Int63n(900000000000) + 100000000000
	return randomNumber
}

func BankProject() {
	accNumber := int32(Grand())
	var hlName string
	var amount float64
	fmt.Printf("Enter your holder name: ")
	fmt.Scanln(&hlName)
	fmt.Printf("Enter your deposit amount: ")
	fmt.Scanln(&amount)
	accData(accNumber, hlName, amount)
}
