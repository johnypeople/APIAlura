package models

import "alura/db"

//Produto é o tipo dos produtos cadastrados
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//BuscaTodosOsProdutos retorna a lista completa
func BuscaTodosOsProdutos() []Produto {

	db := db.Conectar()

	selectDeTodosOsProdutos, erro := db.Query("select * from produtos")
	if erro != nil {
		panic(erro.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if erro != nil {
			panic(erro.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos

}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Conectar()

	//Prepara o script de inserção para verificar se houve algum erro
	insereDadosNoBanco, erro := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?,?,?,?)")
	if erro != nil {
		panic(erro.Error())
	}

	//Aqui os valores são inseridos no banco
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.Conectar()
	deletarOProduto, erro := db.Prepare("delete from produtos where id=?")
	if erro != nil {
		panic(erro.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.Conectar()
	produtoDoBanco, erro := db.Query("select * from produtos where id=?", id)
	if erro != nil {
		panic(erro.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = produtoDoBanco.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if erro != nil {
			panic(erro.Error())
		}

		produtoParaAtualizar.ID = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()

	return produtoParaAtualizar
}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := db.Conectar()
	AtualizaProduto, erro := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id=?")
	if erro != nil {
		panic(erro.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
