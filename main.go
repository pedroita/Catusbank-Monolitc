package main

import "cactusbank.com.br/cactusbank/entities"

func main() {
	// db := model.Instance()
	// db.AutoMigrate(&model.Account{})
	// db.AutoMigrate(&model.Transaction{})
	// db.AutoMigrate(&model.Pix{})
	account := entities.Account{}
	print(account.FindAll())
}
