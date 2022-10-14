package main

import (
	account "cactusbank.com.br/cactusbank/interfaces/Account"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

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
	}
	println(db)

	conta_test1 := account.Account{Name: "Josafa Dieb", DocumentNumber: "412.523.541-98", Saldo: 452.45}

	db.AutoMigrate(&account.Account{}) // cria a tabela
	db.Create(&conta_test1)            // inserir dados na tabela

	//para mais informacoes
	//https://gorm.io/docs/create.html
}
