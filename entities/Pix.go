package entities

import (
	"cactusbank.com.br/cactusbank/model"
	"gorm.io/gorm"
)

type Pix struct {
	model.Repository[Pix]
	gorm.Model
	Transaction_Id int `gorm:"foreignkey:id"`
	Message        string
	Value          float32
}
