package main

import (
	account "cactusbank.com.br/cactusbank/interfaces/Account"
	pix "cactusbank.com.br/cactusbank/interfaces/Pix"
	transaction "cactusbank.com.br/cactusbank/interfaces/Transaction"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func loadMigrations(db *gorm.DB) {
	db.AutoMigrate(&account.Account{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&pix.Pix{})
}

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

	// simulando uma transacao
	// contas
	loadMigrations(db) // cria a tabela
	conta_to := account.Account{Name: "Josafa Dieb 1", DocumentNumber: "412.523.541-98", Saldo: 452.45}
	db.Create(&conta_to)
	conta_from := account.Account{Name: "Josafa Dieb 2", DocumentNumber: "412.523.541-99", Saldo: 452.45}
	db.Create(&conta_from)

	transaction := transaction.Transaction{To: int(conta_to.ID), From: int(conta_from.ID)}
	db.Create(&transaction)
	pix_test := pix.Pix{Transaction_Id: int(transaction.ID), Message: "Ola mundo", Value: 45.2}
	db.Create(&pix_test)

	//para mais informacoes
	//https://gorm.io/docs/create.html
}
