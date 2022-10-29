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
	db := model.Instance()
	account := entities.Account{}
	db.Find(&account, `id = ?`, id)
	return &account
}

func Delete(id int) {
	db := model.Instance()
	account := Find(id)
	tx := db.Begin()
	tx.Delete(&account)
	tx.Commit()

}

func Update(t *entities.Account) {
	// iniciar com begin tran
	db := model.Instance()
	tx := db.Begin()
	tx.Model(&t).Updates(entities.Account{Name: t.GetName(), DocumentNumber: t.GetDocumentNumber(), Saldo: t.GetSaldo()})
	tx.Commit()
}

func Create(t *entities.Account) {
	db := model.Instance()
	tx := db.Begin()
	tx.Create(&t)
	tx.Commit()
}
