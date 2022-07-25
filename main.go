package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var pages = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Rosa", Preco: 39.90, Quantidade: 10},
		{Nome: "Notebook", Descricao: "Gamer", Preco: 2049.90, Quantidade: 4},
		{Nome: "Violão RX213", Descricao: "Tampo Maciço em Abeto", Preco: 1599.90, Quantidade: 5},
		{Nome: "Encordoamento", Descricao: "Tensão Normal", Preco: 24.90, Quantidade: 20},
	}
	pages.ExecuteTemplate(w, "index", produtos)
}
