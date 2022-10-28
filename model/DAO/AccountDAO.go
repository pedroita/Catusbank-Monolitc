package dao

import (
	"cactusbank.com.br/cactusbank/entities"
	"cactusbank.com.br/cactusbank/model"
)

func FindAll() []entities.Account {
	db := model.Instance()
	accounts := []entities.Account{}
	db.Find(&accounts)
	return accounts
}

func Find(id int) *entities.Account {
	account := entities.Account{}

	return &account
}

func Delete(id int) {
	// iniciar com begin tran

}

func Update(t entities.Account) int {
	// iniciar com begin tran
	return 0
}

func Create(t entities.Account) entities.Account {
	return t
}
