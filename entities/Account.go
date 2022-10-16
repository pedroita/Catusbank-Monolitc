package entities

import (
	"cactusbank.com.br/cactusbank/model"
	"gorm.io/gorm"
)

type Account struct {
	model.Repository[Account]
	gorm.Model
	Id             int
	Name           string
	DocumentNumber string
	Saldo          float32
}
