package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Account_TO_ID   int
	Account_FROM_ID int
	To              Account `gorm:"foreignKey:Account_TO_ID;references:ID"`
	From            Account `gorm:"foreignKey:Account_FROM_ID;references:ID"`
}

func (transaction Transaction) GetId() int {
	return int(transaction.ID)
}
func (transaction Transaction) GetTo() *Account {
	return &transaction.To
}

func (transaction *Transaction) SetTo(account Account) {
	transaction.To = account
}

func (transaction Transaction) GetFrom() *Account {
	return &transaction.From
}

func (transaction *Transaction) SetFrom(account Account) {
	transaction.To = account
}
