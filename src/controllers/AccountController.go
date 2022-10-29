package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"cactusbank.com.br/cactusbank/src/entities"
	account "cactusbank.com.br/cactusbank/src/models/DAO"
	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var newAccount entities.Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nameSize := len(newAccount.GetName())
	documentNumberSize := len(newAccount.GetDocumentNumber())

	if nameSize <= 0 || nameSize > 30 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "O nome da conta precisa ter no mínimo 4 caracteres e no máximo 30!"
		}`)
		return
	}
	if documentNumberSize != 14 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "O Documento informado é inválido!"
		}`)
		return
	}
	account.Create(&newAccount)
	io.WriteString(w, newAccount.GetName()+" Foi cadastrado com sucesso!")
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "Identificador da conta inválido!"
		}`)
		return
	}
	acc := account.Find(int(id))
	if acc.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "A conta informada não existe!"
		}`)
		return
	}
	data, _ := json.Marshal(acc)
	io.WriteString(w, string(data))
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []string

	for _, v := range account.FindAll() {
		data, _ := json.Marshal(v)
		accounts = append(accounts, string(data))
	}
	io.WriteString(w, string(strings.Join(accounts, "")))
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "Identificador da conta inválido!"
		}`)
		return
	}
	acc := account.Find(id)
	if acc.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "A conta informada não existe!"
		}`)
		return
	}
	account.Delete(id)
	io.WriteString(w, `{
		"error": true,
		"message": "A conta informada foi deletada com sucesso!!"
	}`)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "Identificador da conta inválido!"
		}`)
		return
	}
	acc := account.Find(id)
	if acc.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "A conta informada não existe!"
		}`)
		return
	}

	err2 := json.NewDecoder(r.Body).Decode(&acc)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nameSize := len(acc.GetName())
	documentNumberSize := len(acc.GetDocumentNumber())

	if nameSize <= 0 || nameSize > 30 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "O nome da conta precisa ter no mínimo 4 caracteres e no máximo 30!"
		}`)
		return
	}

	if documentNumberSize != 14 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{
			"error": true,
			"message": "O Documento informado é inválido!"
		}`)
		return
	}
	account.Update(acc)

	io.WriteString(w, `{
		"error": true,
		"message": "A conta informada foi atualizada com sucesso!!"
	}`)
}
