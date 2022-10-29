package account

import (
	"cactusbank.com.br/cactusbank/src/entities"
	model "cactusbank.com.br/cactusbank/src/models"
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

func Create(t *entities.Account) {
	db := model.Instance()
	db.Begin()
	db.Create(&t)
	db.Commit()
}
