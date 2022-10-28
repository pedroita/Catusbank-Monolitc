package main

import (
	"cactusbank.com.br/cactusbank/entities"
	"cactusbank.com.br/cactusbank/model"
)

func main() {

	// START MIGRATIONS
	db := model.Instance()
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Transaction{})
	db.AutoMigrate(&entities.Pix{})
	// END MIGRATIONS

	// 1 PIX VAI TER 1 Transacao e na transacao tera 2 contas o que envia (from) e o que vai receber (TO)
	// CRIAR CRUD (AccountDAO, PixDAO, TransactionDAO)
	// CRIAR METODOS DO NO CONTROLLER PARA SEREM CHAMADOS VIA REQUISICAO HTTP (CriarConta, Transferir, Imprimir)
	// CRIAR OS METODOS HTTP E ROTAS USANDO ALGUMA BIBLIOTECA

}
