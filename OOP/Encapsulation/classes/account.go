package classes

import (
	"errors"
)

var (
	errorValueLessThanZero = errors.New("Баланс должен быть положительным") 
	errorBalanceLessThanZero = errors.New("Общий баланс стал отрицательным")
)

type Account struct {
	 balance (float64)
	 owner (string)
}

func NewAccount(owner string) *Account {
	newAccount := Account{
		owner: owner,
		balance: 0,
	}
	return &newAccount
}

func (account *Account) SetBalance(value float64) error {
	if value < 0 {
		return errorValueLessThanZero
	}
	account.balance = value
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.balance
}

func (account *Account) Deposit(value float64) error {
	if value < 0 {
		return errorValueLessThanZero
	}
	account.balance += value
	return nil
}

func (account *Account) Withdraw(value float64) error {
	if value < 0 {
		return errorValueLessThanZero
	}
	if result := account.balance - value; result < 0 {
		return errorBalanceLessThanZero
	} else  {
		account.balance = result
	}
	return nil
}