package account

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Name     string
	Password string
	Balance  float64
}

func NewBankAccount(name, password string, balance float64) *BankAccount {
	return &BankAccount{
		Name:     name,
		Password: password,
		Balance:  balance,
	}
}

func (b *BankAccount) Info(password string) (string, error) {
	if password != b.Password {
		return "", errors.New("Invalid Password Please try it again.")
	}

	return fmt.Sprintf("Name: %s\nPassword: %s\nBalance: %.2f", b.Name, b.Password, b.Balance), nil
}
