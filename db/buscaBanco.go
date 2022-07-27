package db

import (
	"ze-is-store/model"
)

func GetAllProduts() []model.Produto {
	db := ConectionMySql()
	selectAllProduts, err := db.Query("select * from produtos p")
	if err != nil {
		panic(err.Error())
	}
	p := model.Produto{}
	produtos := []model.Produto{}

	for selectAllProduts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectAllProduts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
