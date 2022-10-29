package main

import (
	"net/http"

	controller "cactusbank.com.br/cactusbank/src/controllers"
)

//	func getRoot(w http.ResponseWriter, r *http.Request) {
//		fmt.Printf("got / request\n")
//		io.WriteString(w, "This is my website!\n")
//	}

func main() {

	// START MIGRATIONS
	// db := model.Instance()
	// db.AutoMigrate(&entities.Account{})
	// db.AutoMigrate(&entities.Transaction{})
	// db.AutoMigrate(&entities.Pix{})
	// PORT = 8080
	// END MIGRATIONS
	http.HandleFunc("/account/create", controller.CreateAccount)
	http.ListenAndServe(":8080", nil)
	// 1 PIX VAI TER 1 Transacao e na transacao tera 2 contas o que envia (from) e o que vai receber (TO)
	// CRIAR CRUD (AccountDAO, PixDAO, TransactionDAO)
	// CRIAR METODOS DO NO CONTROLLER PARA SEREM CHAMADOS VIA REQUISICAO HTTP (CriarConta, Transferir, Imprimir)
	// CRIAR OS METODOS HTTP E ROTAS USANDO ALGUMA BIBLIOTECA
	// conta1 := entities.Account{Name: "Josafa", DocumentNumber: "452.475.452-95", Saldo: 452}
	// account.Create(&conta1)
}
