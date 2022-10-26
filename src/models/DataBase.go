package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Instance() *gorm.DB {
	db := database()
	return db
}
func database() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:Smoke123@@tcp(127.0.0.1:3306)/trabalho_sd?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
		return nil
	}
	return db
}
