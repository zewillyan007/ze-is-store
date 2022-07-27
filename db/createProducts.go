package db

import "fmt"

func CreateNewProduts(nome, descricao string, preco float64, quantidade int) {
	db := ConectionMySql()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)")

	fmt.Println(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}
	_, err2 := insertData.Exec(nome, descricao, preco, quantidade)

	if err2 != nil {
		fmt.Printf("%#v", err2)
	}

	defer db.Close()
}
