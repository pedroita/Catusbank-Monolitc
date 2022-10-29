package controller

import (
	"fmt"
	"io"
	"net/http"

	"cactusbank.com.br/cactusbank/src/entities"
	// account "cactusbank.com.br/cactusbank/src/models/DAO"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Response-Type", "json")
	method := r.Method
	if method != "POST" {
		io.WriteString(w, `This method API is not support to create an account.`)
		return
	}
	fmt.Println(r)
	// account.Create()
}

func GetAccount(id int) *entities.Account {
	return nil
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{
		"nome": "teste"
		}`)
}

func RemoveAccount(id int) {
}

func UpdateAccount(account entities.Account) {

}

func Transferir(target entities.Account, ammount int) {

}
