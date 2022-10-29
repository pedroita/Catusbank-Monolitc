package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name           string
	DocumentNumber string
	Saldo          float32
}

func (account Account) GetId() int {
	return int(account.ID)
}
func (account Account) GetName() string {
	return account.Name
}

func (account *Account) SetName(name string) {
	account.Name = name
}

func (account Account) GetDocumentNumber() string {
	return account.DocumentNumber
}

func (account *Account) SetDocumentNumber(documentNumber string) {
	account.DocumentNumber = documentNumber
}

func (account Account) GetSaldo() float32 {
	return float32(account.Saldo)
}

func (account *Account) SetSaldo(saldo float32) {
	account.Saldo = saldo
}
