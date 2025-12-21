package Bank

import (
	"errors"
)

type BankAccount struct {
	ID      uint64
	Name    string
	Balance float64
}

func newBankAccount(id uint64, name string, balance float64) *BankAccount {
	return &BankAccount{
		ID:      id,
		Name:    name,
		Balance: balance,
	}
}

func (account *BankAccount) GetBalance() float64 {
	return account.Balance
}

func (account *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be greater than zero")
	}
	account.Balance += amount
	return nil
}

func (account *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be greater than zero")
	}

	if amount > account.Balance {
		return errors.New("Insufficient funds")
	}

	account.Balance -= amount
	return nil
}
