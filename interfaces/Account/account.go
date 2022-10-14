package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name           string
	DocumentNumber string
	Saldo          float32
}
