package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"ze-is-store/db"
)

var pages = template.Must(template.ParseGlob("../templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProduts := db.GetAllProduts()
	pages.ExecuteTemplate(w, "Index", allProduts)
}

func New(w http.ResponseWriter, r *http.Request) {
	pages.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvert, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		db.CreateNewProduts(nome, descricao, precoConvert, quantidadeConvert)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	db.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
