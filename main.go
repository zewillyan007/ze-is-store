package main

import (
	"net/http"
	"text/template"
	"ze-is-store/db"
)

var pages = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProduts := db.GetAllProduts()
	pages.ExecuteTemplate(w, "index", allProduts)
}
