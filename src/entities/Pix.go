package entities

import (
	"gorm.io/gorm"
)

type Pix struct {
	gorm.Model
	Message        string
	Value          float32
	Transaction_ID int
	Transaction    Transaction `gorm:"foreignKey:Transaction_ID;references:ID"`
}

func (pix Pix) GetId() int {
	return int(pix.ID)
}
func (pix Pix) GetMessage() string {
	return pix.Message
}

func (pix Pix) GetValue() float32 {
	return float32(pix.Value)
}
