package entities

import (
	"cactusbank.com.br/cactusbank/model"
	"gorm.io/gorm"
)

type Transaction struct {
	model.Repository[Transaction]
	gorm.Model
	To   int `gorm:"foreignkey:id"`
	From int `gorm:"foreignkey:id"`
}
