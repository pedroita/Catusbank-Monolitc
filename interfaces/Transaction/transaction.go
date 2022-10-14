package transaction

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	To   int `gorm:"foreignkey:id"`
	From int `gorm:"foreignkey:id"`
}
