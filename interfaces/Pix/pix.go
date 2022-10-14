package pix

import (
	"gorm.io/gorm"
)

type Pix struct {
	gorm.Model
	Transaction_Id int `gorm:"foreignkey:id"`
	Message        string
	Value          float32
}
